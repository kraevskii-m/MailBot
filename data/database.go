package data

import "github.com/lib/pq"

type DataBase struct {
}

func (DataBase) AddBot(username string, password string) (string, error) {
	panic("implement me")
}

func (DataBase) GetBot(token string) (Bot, error) {
	panic("implement me")
}

func (DataBase) GetAllBots() []Bot {
	panic("implement me")
}

func (DataBase) GetMessages(bot Bot, offset int, limit int) []Message {
	panic("implement me")
}

func (DataBase) AddMessages(messages []Message, token string) {
	panic("implement me")
}
