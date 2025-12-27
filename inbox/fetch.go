package inbox

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type Message struct {
	Title       string
	Contents    string
	From        string
	Attachments [][]byte
}

func decode(data string) string {
	b, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return ""
	}
	return string(b)
}

func getHeader(msg *gmail.Message, name string) string {
	for _, h := range msg.Payload.Headers {
		if h.Name == name {
			return h.Value
		}
	}
	return ""
}

func extractBody(msg *gmail.Message) string {
	if msg.Payload.Body.Data != "" {
		return decode(msg.Payload.Body.Data)
	}

	for _, part := range msg.Payload.Parts {
		if part.MimeType == "text/plain" && part.Body.Data != "" {
			return decode(part.Body.Data)
		}
	}
	return ""
}

func extractAttachments(msg *gmail.Message, svc *gmail.Service) [][]byte {
	var files [][]byte

	for _, part := range msg.Payload.Parts {

		if part.Filename == "" || part.Body.AttachmentId == "" {
			continue
		}

		att, err := svc.Users.Messages.Attachments.Get("me", msg.Id, part.Body.AttachmentId).Do()

		if err != nil {
			continue
		}

		data, err := base64.URLEncoding.DecodeString(att.Data)
		if err != nil {
			continue
		}

		files = append(files, data)

	}
	return files
}

func FetchMessages(keywords []string) ([]Message, error) {
	var messages []Message
	tokenBytes, err := os.ReadFile("token.json")
	if err != nil {
		fmt.Println("Couldn't find token.json file. Make sure you login first using auth option.")
		return nil, err
	}

	var token oauth2.Token
	json.Unmarshal(tokenBytes, &token)

	creds, errCreds := os.ReadFile("credentials.json")

	if errCreds != nil {
		return nil, errCreds
	}

	ctx := context.Background()
	config, _ := google.ConfigFromJSON(creds, gmail.GmailReadonlyScope)
	client := config.Client(ctx, &token)

	service, _ := gmail.NewService(ctx, option.WithHTTPClient(client))
	query := fmt.Sprintf(
		"(%s)",
		strings.Join(keywords, " OR "),
	)

	res, _ := service.Users.Messages.List("me").Q(query).MaxResults(50).Do()

	for _, msg := range res.Messages {
		fullMsg, _ := service.Users.Messages.Get("me", msg.Id).Format("full").Do()
		title := getHeader(fullMsg, "Subject")
		from := getHeader(fullMsg, "From")
		contents := extractBody(fullMsg)
		attachments := extractAttachments(fullMsg, service)

		m := Message{Title: title, Contents: contents, From: from, Attachments: attachments}

		messages = append(messages, m)
	}

	return messages, nil
}
