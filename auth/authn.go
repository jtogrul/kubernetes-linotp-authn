package auth

type Authn interface {
	Check() (bool, string, string)
}
