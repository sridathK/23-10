package middlewear

import "myserviceapp/auth"

type Middlewear struct {
	a *auth.Auth
}

func NewMiddleWear(a *auth.Auth) (Middlewear, error) {
	return Middlewear{a: a}, nil
}
