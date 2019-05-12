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

type MemoryStorage struct {
	BotStorage    atomic.Value
	LetterStorage atomic.Value
}

func (m *MemoryStorage) AddMessages(messages []Message, token string) {
	var base = m.LetterStorage.Load()
	letterBase := make(map[string][]Message)
	if base != nil {
		letterBase = base.(map[string][]Message)
	}
	if val, ok := letterBase[token]; ok {
		letterBase[token] = append(val, messages...)
	} else {
		letterBase[token] = messages
	}
	m.LetterStorage.Store(letterBase)
}

func (m *MemoryStorage) GetMessages(bot Bot, offset int, limit int) []Message { //todo offset and limit
	var base = m.LetterStorage.Load()
	var letterBase map[string][]Message
	if base != nil {
		letterBase = base.(map[string][]Message)
	}
	return letterBase[bot.Token]
}

func (m *MemoryStorage) AddBot(username string, password string) (string, error) {
	botProfile := Bot{Token: GenerateToken(username), Username: username, Password: password}
	_, err := m.GetBot(botProfile.Token)
	if err == nil {
		return "", errors.New("BOT ALREADY EXIST")
	}
	var base = m.BotStorage.Load()
	var botBase []Bot
	if base != nil {
		botBase = base.([]Bot)
	}
	botBase = append(botBase, botProfile)
	m.BotStorage.Store(botBase)
	return botProfile.Token, nil
}

func (m *MemoryStorage) GetBot(token string) (Bot, error) {
	var base = m.BotStorage.Load()
	var botBase []Bot
	if base != nil {
		botBase = base.([]Bot)
	}
	for _, bot := range botBase {
		if bot.Token == token {
			return bot, nil
		}
	}
	return Bot{}, errors.New("CANNOT FIND BOT")
}

func (m *MemoryStorage) GetAllBots() []Bot {
	var base = m.BotStorage.Load()
	var botBase []Bot
	if base != nil {
		botBase = base.([]Bot)
	}
	return botBase
}
