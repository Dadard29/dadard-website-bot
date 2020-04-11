package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func id(u tgbotapi.Update) string {
	return fmt.Sprintf("Your ID is `%d`", u.Message.From.ID)
}
