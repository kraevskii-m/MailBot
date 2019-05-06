package data

type Storage interface {
	AddBot(name string) error
	GetBot(token string) Bot
	GetMessages(bot Bot, offset int, limit int) []Message
}
