package data

import (
	"errors"
	"sync/atomic"
	"time"
)

type Bot struct {
	Token    string
	Username string
	Password string
}

var database atomic.Value

func Initialize() {
	NewBot("botfather")
}

func NewBot(name string) Bot {
	output := Bot{}
	//TODO It's temporary solution
	if name == "botfather" {
		output.Token = "qwertyui"
		output.Username = "fatherofbots"
		output.Password = "lermonter07"
	}
	if name == "echobot" {
		output.Token = "asdfghjk"
		output.Username = "echobot-mailbot"
		output.Password = "lermonter07"
	}

	var base = database.Load()
	var botBase []Bot
	if base != nil {
		botBase = base.([]Bot)
	}
	botBase = append(botBase, output)
	database.Store(botBase)
	return output
}

func GetBot(name string) (Bot, error) {
	var bots = database.Load().([]Bot)
	for _, bot := range bots {
		if bot.Username == name {
			return bot, nil
		}
	}

	return Bot{}, errors.New("There is no bot")
}

func GetByToken(token string) (Bot, error) {
	var bots = database.Load().([]Bot)
	for _, bot := range bots {
		if bot.Token == token {
			return bot, nil
		}
	}

	return Bot{}, errors.New("There is no bot")
}

type Letter struct {
	From    string
	To      string
	Subject string
	Body    string
}

func GetLetters(token string) []Letter {
	return []Letter{}
}

func UpdatesController() {
	Initialize()

	for {
		time.Sleep(1 * time.Second)
		var base = database.Load()
		var botBase []Bot
		if base == nil {
			continue
		}

		for _, bot := range botBase {
			updateMailbox(bot)
		}
	}
}

func updateMailbox(bot Bot) {

}
