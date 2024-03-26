package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

type args struct {
	Channel string `json:"channel"`
	Message string `json:"message"`
}

func main() {
	if os.Getenv("GPTSCRIPT_SLACK_TOKEN") == "" {
		panic("GPTSCRIPT_SLACK_TOKEN is not set")
	}

	slackClient := slack.New(os.Getenv("GPTSCRIPT_SLACK_TOKEN"))

	if len(os.Args) != 2 {
		panic("Usage: " + os.Args[0] + " <JSON parameters>")
	}

	var a args
	if err := json.Unmarshal([]byte(os.Args[1]), &a); err != nil {
		panic(err)
	}

	if a.Channel == "" {
		panic("channel is required")
	} else if a.Message == "" {
		panic("message is required")
	}

	if _, _, err := slackClient.PostMessage(a.Channel, slack.MsgOptionText(a.Message, false)); err != nil {
		panic(err)
	}
	fmt.Println("message sent successfully")
}
