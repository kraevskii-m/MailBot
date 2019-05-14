package data

import (
	"crypto/md5"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type Message struct {
	From    string
	To      string
	Subject string
	Body    string
}

type Storage interface {
	AddBot(username string, password string) (string, error)
	GetBot(token string) (Bot, error)
	GetAllBots() []Bot
	GetMessages(bot Bot, offset int, limit int) []Message
	AddMessages(messages []Message, token string)
	RemoveBot(bot Bot)
}

func GenerateToken(name string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(name), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	hasher := md5.New()
	hasher.Write(hash)
	return hex.EncodeToString(hasher.Sum(nil))
}

var Base = MemoryStorage{}
