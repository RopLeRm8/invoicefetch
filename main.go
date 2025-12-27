package main

import (
	"fmt"
	"invoicefetch/inbox"
	"invoicefetch/oauth"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: invoicefetch <auth | fetch>")
		return
	}

	switch os.Args[1] {
	case "auth":
		oauth.RunAuth()
	case "fetch":
		messages, _ := inbox.FetchMessages([]string{"As long as the giveaway is live, every completed order gets you more. The more tickets you stack, the better your odds."})

		for _, msg := range messages {
			fmt.Print(msg.Title)
		}

	case "verify":
		email, err := oauth.VerifyIdentity()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("You are logged in! Your email is %s\n", *email)

	default:
		fmt.Println("Command not found")
	}
}
