package facebook

import (
	"fmt"
	"os"

	"golang.org/x/oauth2"
)

const (
	authorizeEndpoint = "https://www.facebook.com/dialog/oauth"
	tokenEndpoint     = "https://graph.facebook.com/oauth/access_token"
)

// GetConnect 接続を取得する
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

	fmt.Println(config)

	return config
}
