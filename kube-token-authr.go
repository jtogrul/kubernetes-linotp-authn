package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"kubernetes-linotp-auth/auth"
	"kubernetes-linotp-auth/models"
	"log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	ServerHost string `yaml:"serverHost"`
	ServerPort string `yaml:"serverPort"`
	ServerCert string `yaml:"serverCert"`
	ServerKey string `yaml:"serverKey"`
	LinotpUrl string `yaml:"linotpUrl"`
}

func handler(w http.ResponseWriter, r *http.Request) {

	var tr models.TokenReview
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&tr)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	//fmt.Printf(tr.Spec.Token)

	authn := auth.LinotpBase64Authn{}

	auth_success, username, uid :=  authn.Check(tr)

	fmt.Print("AUTH: ")
	fmt.Println(auth_success)

	var tr_res models.TokenReview
	tr_res.Kind = "TokenReview"
	tr_res.ApiVersion = "authentication.k8s.io/v1beta1"

	if (auth_success) {
		var tr_res_status models.TokenReviewStatus
		tr_res_status.Authenticated = true

		var tr_res_status_user models.TokenReviewStatusUser
		tr_res_status_user.Username = username
		tr_res_status_user.Uid = uid

		tr_res_status.User = tr_res_status_user

		tr_res.Status = tr_res_status

	} else {
		var tr_res_status models.TokenReviewStatus
		tr_res_status.Authenticated = false

		var tr_res_status_user models.TokenReviewStatusUser
		tr_res_status.User = tr_res_status_user

		tr_res.Status = tr_res_status
	}

	b, err := json.Marshal(tr_res)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	fmt.Fprintf(w, string(b))

}

func main() {

	configData, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println(string(configData))

	config := Config{}

	marshall_err := yaml.Unmarshal([]byte(configData), &config)
	if marshall_err != nil {
		log.Fatalf("error: %v", marshall_err)
	}

	http.HandleFunc("/", handler)
	http.ListenAndServeTLS(fmt.Sprintf("%s:%s", config.ServerHost, config.ServerPort), config.ServerCert, config.ServerKey, nil)
}
