package data

import (
	"sync/atomic"
)

type Bot struct {
	Token    string
	Username string
	Password string
}

type MemoryStorage struct {
	BotStorage    atomic.Value
	LetterStorage atomic.Value
}

func (MemoryStorage) AddBot(name string) (string, error) {
	panic("implement me")
}

func (MemoryStorage) GetBot(token string) (Bot, error) {
	panic("implement me")
}

func (MemoryStorage) GetAllBots() []Bot {
	panic("implement me")
}

func (MemoryStorage) GetMessages(bot Bot, offset string, limit string) []Message {
	panic("implement me")
}

// //todo obsolete
//func NewBot(name string) Bot {
//	output := Bot{}
//	if name == "botfather" {
//		output.Token = "qwertyui"
//		output.Username = "fatherofbots"
//		output.Password = "lermonter07"
//	}
//	if name == "echobot" {
//		output.Token = "asdfghjk"
//		output.Username = "echobot-mailbot"
//		output.Password = "lermonter07"
//	}
//
//	var base = BotStorage.Load()
//	var botBase []Bot
//	if base != nil {
//		botBase = base.([]Bot)
//	}
//	botBase = append(botBase, output)
//	BotStorage.Store(botBase)
//
//	var mailBase = LetterStorage.Load()
//	var letterBase map[string][]Message
//	if mailBase != nil {
//		letterBase = mailBase.(map[string][]Message)
//	}
//	letterBase[output.Token] = []Message{}
//	LetterStorage.Store(letterBase)
//
//	return output
//}
//
//func GetBot(name string) (Bot, error) {
//	var bots = BotStorage.Load().([]Bot)
//	for _, bot := range bots {
//		if bot.Username == name {
//			return bot, nil
//		}
//	}
//
//	return Bot{}, errors.New("There is no bot")
//}
//
//func GetByToken(token string) (Bot, error) {
//	var bots = BotStorage.Load().([]Bot)
//	for _, bot := range bots {
//		if bot.Token == token {
//			return bot, nil
//		}
//	}
//
//	return Bot{}, errors.New("There is no bot")
//}
//
//func GetLetters(token string) ([]Message, error) {
//	var base = LetterStorage.Load()
//	var letterBase map[string][]Message
//	if base == nil {
//		return nil, errors.New("Empty database")
//	}
//	letterBase = base.(map[string][]Message)
//	if val, ok := letterBase[token]; ok {
//		return val, nil
//	}
//
//	return nil, errors.New("Empty mailbox")
//}
