package oidc

import "net/http"

type Tokens struct {
	AccessToken  string
	RefreshToken string
	IDToken      string
}

func GetTokens(r *http.Request) (Tokens, error) {
	// TODO implement
	// handle token refresh
	return Tokens{
		AccessToken:  "ey...",
		RefreshToken: "",
		IDToken:      "",
	}, nil
}
