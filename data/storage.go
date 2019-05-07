package data

type Message struct {
	From    string
	To      string
	Subject string
	Body    string
}

type Storage interface {
	AddBot(name string) (string, error)
	GetBot(token string) Bot
	GetAllBots() []Bot
	GetMessages(bot Bot, offset int, limit int) []Message
}

func GenerateToken() string { // todo implement
	panic("Implement me")
}
