package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	envConfigGoogleClientId     = "GOOGLE_CLIENT_ID"
	envConfigGoogleClientSecret = "GOOGLE_CLIENT_SECRET"
)

type GoogleHandler struct {
	ClientId     string
	ClientSecret string
}

func (g *GoogleHandler) GoogleOauthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     g.ClientId,
		ClientSecret: g.ClientSecret,
		RedirectURL:  "http://localhost:8000/auth/google/callback",
		Scopes: []string{
			"openid",
			"profile",
			"email",
		},
		Endpoint: google.Endpoint,
	}
}
