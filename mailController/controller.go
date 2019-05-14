package mailController

import (
	"fmt"
	"github.com/kraevskii-m/MailBot/bots"
	"github.com/kraevskii-m/MailBot/data"
	"log"
	"time"
)

func UpdateMailBox(bot data.Bot) {
	for {
		time.Sleep(5 * time.Second)
		messages, err := checkMailBox(bot)
		if err != nil {
			bots.RemoveBot(bot)
			break
		}
		data.Base.AddMessages(messages, bot.Token)
	}
}

func SendMessage(message data.Message, bot data.Bot) error {
	time.Sleep(5 * time.Second)
	return Sender(bot, message.To, message.Body, message.Subject)
}

func BotMonitor() {
	botIsMoitoring := make(map[data.Bot]bool)
	for {
		time.Sleep(5 * time.Second)
		registeredBots := data.Base.GetAllBots()
		for _, b := range registeredBots {
			if !botIsMoitoring[b] {
				log.Println(fmt.Sprintf("Bot %v is monitoring", b.Username))
				go UpdateMailBox(b)
				if b.Username == "echobot-mailbot" {
					log.Println("Echobot started!")
					go bots.EchoBotController(b.Token)
				}
				botIsMoitoring[b] = true
			}
		}
	}
}
