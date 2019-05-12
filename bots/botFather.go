package bots

import (
	"github.com/kraevskii-m/MailBot/botLib"
	"github.com/kraevskii-m/MailBot/data"
	"log"
	"strings"
	"time"
)

var BotFatherEmail = "fatherofbots@yandex.ru"

func BotFatherController(storage data.Storage) {
	token, _ := storage.AddBot("fatherofbots", "lermonter07")
	bot := botLib.NewMailBot(token, "fatherofbots")
	for {
		time.Sleep(5 * time.Second)
		messages, err := bot.GetUpdates(0, 0)
		if err != nil {
			log.Print(err)
		}
		for _, message := range messages {
			register(message, storage)
		}
	}
}
func register(message botLib.Message, storage data.Storage) {
	userData := strings.Fields(message.Body)
	token, err := storage.AddBot(userData[0], userData[1])
	if err != nil {
		log.Print(err)
		botLib.NewMessage(BotFatherEmail, message.From, "Bot registering", "Choose another name! "+err.Error())
		return
	}

	botLib.NewMessage(BotFatherEmail, message.From, "Successful registration!", token)
}
