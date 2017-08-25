package facebook

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

const (
	authorizeEndpoint = "https://www.facebook.com/dialog/oauth"
	tokenEndpoint     = "https://graph.facebook.com/oauth/access_token"
)

func GetConnect() *oauth2.Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := &oauth2.Config{
		ClientID:     os.Getenv("FacebookClientID"),
		ClientSecret: os.Getenv("FacebookClientSecret"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  authorizeEndpoint,
			TokenURL: tokenEndpoint,
		},
		Scopes:      []string{"email"},
		RedirectURL: "http://localhost:8080/facebook/callback",
	}
	return config
}
