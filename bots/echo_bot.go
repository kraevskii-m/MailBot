package bots

import (
	"github.com/kraevskii-m/MailBot/botLib"
	"log"
	"time"
)

var address = "echobot-mailbot@yandex.ru"

func EchoBotController(token string) {
	bot := botLib.NewMailBot(token, "EchoBot")
	for {
		time.Sleep(5 * time.Second)
		messages, err := bot.GetUpdates(0, 0) //todo use offset and limit
		if err != nil {
			log.Print(err)
		}
		for _, message := range messages {
			botLib.NewMessage(address, message.From, message.Subject, message.Body)
		}
	}
}
