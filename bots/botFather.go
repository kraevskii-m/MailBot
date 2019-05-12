package bots

import (
	"github.com/kraevskii-m/MailBot/botLib"
	"github.com/kraevskii-m/MailBot/data"
	"log"
	"time"
)

var BotFatherEmail = "fatherofbots@yandex.ru"

func BotFatherController(storage data.Storage) {
	token, _ := storage.AddBot("BotFather")
	bot := botLib.NewMailBot(token, "BotFather")
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
	token, err := storage.AddBot(message.Body)
	if err != nil {
		log.Print(err)
		botLib.NewMessage(BotFatherEmail, message.From, "Bot registering", "Choose another name! "+err.Error())
		return
	}

	botLib.NewMessage(BotFatherEmail, message.From, "Successful registration!", token)
}
