package oauth

import (
	"context"
	"encoding/json"
	"errors"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

func GetIdentity() (string, error) {
	tokenBytes, err := os.ReadFile("token.json")
	if err != nil {
		tokenErr := errors.New("Couldn't find token.json file. Make sure you login first using auth option.")
		return "", tokenErr
	}

	var token oauth2.Token
	json.Unmarshal(tokenBytes, &token)

	creds, errCreds := os.ReadFile("credentials.json")

	if errCreds != nil {
		return "", errCreds
	}

	ctx := context.Background()
	config, _ := google.ConfigFromJSON(creds, gmail.GmailReadonlyScope)
	client := config.Client(ctx, &token)

	service, _ := gmail.NewService(ctx, option.WithHTTPClient(client))
	profile, _ := service.Users.GetProfile("me").Do()

	return profile.EmailAddress, nil

}
