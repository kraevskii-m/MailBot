package data

type Storage interface {
	addBot()
	getBot() Bot
	getMessages(bot Bot, offset int, limit int) []Message
}
