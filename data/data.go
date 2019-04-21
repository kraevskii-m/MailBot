package data

import (
	"errors"
	"github.com/kraevskii-m/MailBot/mailController"
	"sync"
	"sync/atomic"
	"time"
)

type Bot struct {
	Token    string
	Username string
	Password string
}

var botStorage atomic.Value
var letterStorage atomic.Value

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

	var base = botStorage.Load()
	var botBase []Bot
	if base != nil {
		botBase = base.([]Bot)
	}
	botBase = append(botBase, output)
	botStorage.Store(botBase)

	var mailBase = letterStorage.Load()
	var letterBase map[string][]Letter
	if mailBase != nil {
		letterBase = mailBase.(map[string][]Letter)
	}
	letterBase[output.Token] = []Letter{}
	letterStorage.Store(letterBase)

	return output
}

func GetBot(name string) (Bot, error) {
	var bots = botStorage.Load().([]Bot)
	for _, bot := range bots {
		if bot.Username == name {
			return bot, nil
		}
	}

	return Bot{}, errors.New("There is no bot")
}

func GetByToken(token string) (Bot, error) {
	var bots = botStorage.Load().([]Bot)
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
	var base = letterStorage.Load()
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

func UpdatesController() {
	Initialize()

	for {
		time.Sleep(1 * time.Second)
		var base = botStorage.Load()
		var botBase []Bot
		if base == nil {
			continue
		}

		wg := &sync.WaitGroup{}

		wg.Add(1) //TODO Check
		for _, bot := range botBase {
			go updateMailbox(bot)
		}
	}
}

func updateMailbox(bot Bot) {
	var base = letterStorage.Load()
	var letterBase map[string][]Letter
	if base == nil {
		letterStorage.Store(letterBase)
		return
	}
	letterBase = base.(map[string][]Letter)
	letterBase[bot.Token] = mailController.GetUpdatesForBot(bot.Token)
	letterStorage.Store(letterBase)
}
