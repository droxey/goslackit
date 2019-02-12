package main

import (
	"os"

	"github.com/droxey/goslackit/slack"
	_ "github.com/joho/godotenv/autoload"
)

// main is our entrypoint, where the application initializes the Slackbot.
// DO NOT EDIT THIS FUNCTION. This is a fully complete implementation.
func main() {
	botToken := os.Getenv("BOT_OAUTH_ACCESS_TOKEN")
	slackClient := slack.CreateSlackClient(botToken)
	slack.RespondToEvents(slackClient)
}
