package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
)

type args struct {
	Channel string `json:"channel"`
	Message string `json:"message"`
}

func main() {
	if os.Getenv("GPTSCRIPT_SLACK_TOKEN") == "" {
		logrus.Error("GPTSCRIPT_SLACK_TOKEN is not set")
		os.Exit(1)
	}

	slackClient := slack.New(os.Getenv("GPTSCRIPT_SLACK_TOKEN"))

	if len(os.Args) != 2 {
		logrus.Errorf("Usage: %s <JSON parameters>", os.Args[0])
		os.Exit(1)
	}

	var a args
	if err := json.Unmarshal([]byte(os.Args[1]), &a); err != nil {
		logrus.Errorf("failed to parse arguments: %v", err)
		os.Exit(1)
	}

	if a.Channel == "" {
		logrus.Error("channel is required")
		os.Exit(1)
	} else if a.Message == "" {
		logrus.Error("message is required")
		os.Exit(1)
	}

	if _, _, err := slackClient.PostMessage(a.Channel, slack.MsgOptionText(a.Message, false)); err != nil {
		logrus.Errorf("failed to post message: %v", err)
		os.Exit(1)
	}
	fmt.Println("message sent successfully")
}
