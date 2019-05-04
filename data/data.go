package data

import (
	"errors"
	"sync/atomic"
)

type Bot struct {
	Token    string
	Username string
	Password string
}

var BotStorage atomic.Value
var LetterStorage atomic.Value

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

	var base = BotStorage.Load()
	var botBase []Bot
	if base != nil {
		botBase = base.([]Bot)
	}
	botBase = append(botBase, output)
	BotStorage.Store(botBase)

	var mailBase = LetterStorage.Load()
	var letterBase map[string][]Letter
	if mailBase != nil {
		letterBase = mailBase.(map[string][]Letter)
	}
	letterBase[output.Token] = []Letter{}
	LetterStorage.Store(letterBase)

	return output
}

func GetBot(name string) (Bot, error) {
	var bots = BotStorage.Load().([]Bot)
	for _, bot := range bots {
		if bot.Username == name {
			return bot, nil
		}
	}

	return Bot{}, errors.New("There is no bot")
}

func GetByToken(token string) (Bot, error) {
	var bots = BotStorage.Load().([]Bot)
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

func GetLetters(token string) ([]Letter, error) {
	var base = LetterStorage.Load()
	var letterBase map[string][]Letter
	if base == nil {
		return nil, errors.New("Empty database")
	}
	letterBase = base.(map[string][]Letter)
	if val, ok := letterBase[token]; ok {
		return val, nil
	}

	return nil, errors.New("Empty mailbox")
}
