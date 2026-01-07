package oauth

import (
	"context"
	"encoding/json"
	"errors"
	server "invoice_fetch/httpServer"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

type EmailLogin struct {
	googleCreds *oauth2.Config
	url         string
	err         error
}

func RunAuth(ctx context.Context) (url string, err error) {
	codeCh := make(chan string)
	go server.Init(codeCh)
	emailLogin := EMAILogin()
	if emailLogin.err != nil {
		return "", emailLogin.err
	}
	go SaveToken(ctx, codeCh, emailLogin.googleCreds)
	return emailLogin.url, nil
}

func EMAILogin() EmailLogin {
	var _ *oauth2.Config
	var _ = google.Endpoint

	_, err := os.Stat("token.json")
	fileExists := err == nil

	if fileExists {
		fileExistsErr := errors.New("Token file already exists, if you wish to login as someone else, delete it!")
		return EmailLogin{googleCreds: nil, url: "", err: fileExistsErr}
	}

	creds, errCreds := os.ReadFile("credentials.json")

	if errCreds != nil {
		return EmailLogin{googleCreds: nil, url: "", err: errCreds}
	}

	scope := gmail.GmailReadonlyScope
	googleCreds, _ := google.ConfigFromJSON(creds, scope)
	googleCreds.RedirectURL = "http://localhost:8080/auth"

	url := googleCreds.AuthCodeURL("ahshit", oauth2.AccessTypeOffline, oauth2.ApprovalForce)

	return EmailLogin{googleCreds, url, nil}
}

func SaveToken(ctx context.Context, codeCh chan string, googleCreds *oauth2.Config) {
	code := <-codeCh

	token, err := googleCreds.Exchange(ctx, code)

	if err != nil {
		return
	}

	bytes, _ := json.Marshal(token)
	os.Create("token.json")
	os.WriteFile("token.json", bytes, 0600)

	runtime.EventsEmit(ctx, "auth:success")
}
