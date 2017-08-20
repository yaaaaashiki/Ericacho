package facebook

import (
	"os"

	"golang.org/x/oauth2"
)

const (
	authorizeEndpoint = "https://www.facebook.com/dialog/oauth"
	tokenEndpoint     = "https://graph.facebook.com/oauth/access_token"
)

func GetConnect() *oauth2.Config {
	config := &oauth2.Config{
		ClientID:     os.Getenv("facebookClientID"),
		ClientSecret: os.Getenv("facebookClientSecret"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  authorizeEndpoint,
			TokenURL: tokenEndpoint,
		},
		Scopes:      []string{"email"},
		RedirectURL: "http://localhost:8080/facebook/callback",
	}

	return config
}
