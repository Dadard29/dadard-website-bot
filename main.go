package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

var inputMapping = map[string]func(u tgbotapi.Update) string {
	"/id": id,

	"ok": ok,
}

func main() {
	botToken := os.Getenv("BOT_TOKEN")
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		if update.Message.From.IsBot { // ignoring bots
			continue
		}

		inputText := update.Message.Text
		responseText := "I see"
		if h, check := inputMapping[inputText]; check {
			responseText = h(update)
		} else {
			// not managed
			continue
		}


		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, responseText)
		msg.ParseMode = tgbotapi.ModeMarkdown
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
