package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const (
	commandStart = "start"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю такой команды =(")
	switch message.Command() {
	case commandStart:
		msg.Text = "Команда /start"
	}

	_, err := b.bot.Send(msg)
	if err != nil {
		log.Panicln(err)
	}
}
