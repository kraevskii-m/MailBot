package bots

import (
	"github.com/kraevskii-m/MailBot/botLib"
	"github.com/kraevskii-m/MailBot/data"
	"log"
	"strings"
	"time"
)

var BotFatherEmail = "fatherofbots@yandex.ru"
var botFather *botLib.MailBot

func BotFatherController() {
	token, _ := data.Base.AddBot("fatherofbots", "lermonter07")
	botFather = botLib.NewMailBot(token, "fatherofbots")
	for {
		time.Sleep(5 * time.Second)
		messages, err := botFather.GetUpdates()
		if err != nil {
			log.Print(err)
		}
		for _, message := range messages {
			register(message)
		}
	}
}
func register(message botLib.Message) {
	if message.Subject != "Register" {
		log.Println("Wrong subj " + message.Subject)
		return
	}
	userData := strings.Fields(message.Body)
	if len(userData) < 2 {
		log.Println("Wrong data " + message.Body)
		return
	}
	login := strings.Replace(userData[0], "<div>", "", -1) //todo Change for better solution
	password := strings.Replace(userData[0], "</div>", "", -1)

	token, err := data.Base.AddBot(login, password)
	if err != nil {
		log.Print(err)
		log.Println("Bot already exist!")
		message := botLib.NewMessage(BotFatherEmail, message.From, "Bot registering", "Choose another name! "+err.Error())
		botFather.SendMessage(*message)
		return
	}

	log.Println("Bot registered! " + login)
	msg := botLib.NewMessage(BotFatherEmail, message.From, "Successful registration!", token)
	botFather.SendMessage(*msg)
}

func RemoveBot(bot data.Bot) {
	data.Base.RemoveBot(bot)
	msg := botLib.NewMessage(BotFatherEmail, bot.Username+"yandex.ru", "Can't log in!", "Wrong credentials, bot removed!")
	botFather.SendMessage(*msg)
}
