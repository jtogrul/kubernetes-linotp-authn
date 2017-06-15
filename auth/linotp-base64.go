package auth

import (
	"kubernetes-linotp-auth/models"
	"fmt"
	"encoding/json"
	"encoding/base64"
	"strings"
	"net/http"
	"io/ioutil"
	"crypto/tls"
	"net/url"
	"log"
)

type LinotpResponseStatus struct {
	Status bool `json:"status"`
	Value bool `json:"value"`
}

type LinotpResponse struct {
	ID int64 `json:"id"`
	Version string `json:"version"`
	JsonRPC string `json:"jsonrpc"`
	Result LinotpResponseStatus `json:"result"`
}

type LinotpBase64Authn struct {
	linotpUrl string
}

func (a LinotpBase64Authn) Check(tr models.TokenReview) (bool, string, string) {
	// Parse username:password from base64 encoded token
	username, password := a.getCreds(tr.Spec.Token)

	fmt.Println()

	// disable SSL certificate checking
	trans := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: trans}

	api_url := "https://linotp.azercell.com/validate/check"

	// Add username and password to URL
	u, err := url.Parse(api_url)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Set("pass", password)
	q.Set("user", username)
	u.RawQuery = q.Encode()

	fmt.Println(u.String())

	// get response
	res, err := client.Get(u.String())
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(body))

	// Parse to struct
	s, err := a.parseResponse([]byte(body))

	return  s.Result.Value, username, username
}

func (a LinotpBase64Authn) getCreds(token string) (string, string) {
	decoded, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		fmt.Println("decode error:", err)
		//return nil, nil
	}
	fmt.Println(string(decoded))

	creds := strings.Split(string(decoded), ":")
	username, password := creds[0], creds[1]
	fmt.Println(username)
	fmt.Println(password)
	return username, password
}

func (a LinotpBase64Authn) parseResponse (body []byte) (*LinotpResponse, error) {
	var s = new(LinotpResponse)
	err := json.Unmarshal(body, &s)
	if(err != nil){
		fmt.Println("whoops:", err)
	}
	return s, err
}