package models

type TokenReview struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Spec       TokenReviewSpec `json:"spec"`
	Status     TokenReviewStatus `json:"status"`
}

type TokenReviewSpec struct {
	Token string `json:"token"`
}

type TokenReviewStatus struct {
	Authenticated bool `json:"authenticated"`
	User          TokenReviewStatusUser `json:"user"`
}

type TokenReviewStatusUser struct {
	Username string `json:"username"`
	Uid      string `json:"uid"`
}
