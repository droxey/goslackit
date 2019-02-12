package slack

import (
	"fmt"
	"strings"

	"github.com/nlopes/slack"
)

// TODO: Change @BOT_NAME to the same thing you entered when creating your Slack application.
const helpMessage = "type in '@BOT_NAME <command_arg_1> <command_arg_2> <lastname>'"

// CreateSlackClient sets up the slack RTM (real-timemessaging) client library,
// initiating the socket connection and returning the client.
// DO NOT MODIFY! This function is complete.
func CreateSlackClient(apiKey string, debug bool) *slack.RTM {
	println("API:", apiKey)
	api := slack.New(apiKey)

	rtm := api.NewRTM()
	go rtm.ManageConnection() // goroutine!
	return rtm
}

// RespondToEvents waits for messages on the slack client's incomingEvents
// channel, sending responses when it detects the bot has been tagged in a
// message with @<botTag>
func RespondToEvents(slackClient *slack.RTM) {
	for msg := range slackClient.IncomingEvents {
		fmt.Println("Event Received: ", msg.Type)
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			botTagString := fmt.Sprintf("<@%s> ", slackClient.GetInfo().User.ID)
			if !strings.Contains(ev.Msg.Text, botTagString) {
				continue
			}
			message := strings.Replace(ev.Msg.Text, botTagString, "", -1)

			// TODO: Make your bot do more than respond to a help command. See notes below.
			// Make changes below this line and add additional funcs to support your bot's functionality.
			// sendHelp is provided as a simple example. Your team may want to call a free external API
			// in a function called sendResponse that you'd create below the definition of sendHelp,
			// and call in this context to ensure execution when the bot receives an event.

			sendResponse(slackClient, message, ev.Channel)
			sendHelp(slackClient, message, ev.Channel)

		default:

		}
	}
}

// working sendHelp function for reference.
func sendHelp(slackClient *slack.RTM, message, slackChannel string) {
	if strings.ToLower(message) != "help" {
		return
	}
	slackClient.SendMessage(slackClient.NewOutgoingMessage(helpMessage, slackChannel))
}

func sendResponse(slackClient *slack.RTM, message, slackChannel string) {
	// TODO: Implement sendResponse.

	// TODO (STRETCH CHALLENGE): Write a goroutine that calls an external API
	// based on the message you received in this function.
}
