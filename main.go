package main

import (
	"os"

	"github.com/toeflbank/translation-slackbot/internal/bot"
)

var (
	slackBotToken = os.Getenv("SLACK_BOT_TOKEN")
	slackAppToken = os.Getenv("SLACK_APP_TOKEN")
	translatorApiKey = os.Getenv("TRANSLATOR_API_KEY")
	translatorApiSecret = os.Getenv("TRANSLATOR_API_SECRET")
)

func main() {
	c, err := bot.New(slackBotToken, slackAppToken, translatorApiKey,translatorApiSecret)
	if err != nil {
		panic(err)
	}

	panic(c.Process())
}
