package data

type Message struct {
	From    string
	To      string
	Subject string
	Body    string
}

type Storage interface {
	AddBot(name string) (string, error)
	GetBot(token string) (Bot, error)
	GetAllBots() []Bot
	GetMessages(bot Bot, offset string, limit string) []Message
}

func GenerateToken() string { // todo implement
	panic("Implement me")
}

var Base = MemoryStorage{}
