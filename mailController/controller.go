package mailController

import (
	"github.com/kraevskii-m/MailBot/data"
)

func UpdateMailBox(bot data.Bot) {
	messages := checkMailBox(bot)
	data.Base.AddMessages(messages, bot.Token)
}

func SendMessage(message data.Message, bot data.Bot) error {
	return Sender(bot, message.To, message.Body, message.Subject)
}
