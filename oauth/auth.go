package oauth

import (
	"context"
	"encoding/json"
	"fmt"
	server "invoice_fetch/httpServer"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

func RunAuth() {
	codeCh := make(chan string)
	go server.Init(codeCh)
	googleCreds := EMAILogin()
	SaveToken(codeCh, googleCreds)
}

func EMAILogin() *oauth2.Config {
	var _ *oauth2.Config
	var _ = google.Endpoint

	_, err := os.Stat("token.json")
	fileExists := err == nil

	if fileExists {
		fmt.Println("Token file already exists, if you wish to login as someone else, delete it!")
		return nil
	}

	creds, errCreds := os.ReadFile("credentials.json")

	if errCreds != nil {
		return nil
	}

	scope := gmail.GmailReadonlyScope
	googleCreds, _ := google.ConfigFromJSON(creds, scope)
	googleCreds.RedirectURL = "http://localhost:8080/auth"

	url := googleCreds.AuthCodeURL("ahshit", oauth2.AccessTypeOffline, oauth2.ApprovalForce)

	fmt.Print(url)
	return googleCreds
}

func SaveToken(codeCh chan string, googleCreds *oauth2.Config) {
	code := <-codeCh

	ctx := context.Background()
	token, err := googleCreds.Exchange(ctx, code)

	if err != nil {
		return
	}

	bytes, _ := json.Marshal(token)
	os.Create("token.json")
	os.WriteFile("token.json", bytes, 0600)
}
