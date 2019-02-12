package main

import (
	"os"

	"github.com/droxey/goslackit/slack"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	botToken := os.Getenv("BOT_OAUTH_ACCESS_TOKEN")
	slackClient := slack.CreateSlackClient(botToken)
	slack.RespondToEvents(slackClient)
}
