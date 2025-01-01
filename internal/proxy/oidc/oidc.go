package oidc

import "net/http"

type Tokens struct {
	AccessToken        string
	AccessTokenExpiry  int64
	RefreshToken       string
	RefreshTokenExpiry int64
	IDToken            string
	SessionID          string
}

func GetTokens(r *http.Request) (Tokens, error) {
	// TODO implement
	// handle token refresh
	return Tokens{
		AccessToken:        "ey...",
		AccessTokenExpiry:  0,
		RefreshToken:       "",
		RefreshTokenExpiry: 0,
		IDToken:            "",
	}, nil
}
