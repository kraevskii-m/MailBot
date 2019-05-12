package botLib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	From    string
	To      string
	Subject string
	Body    string
}

func NewMessage(from string, to string, subject string, body string) *Message {
	message := new(Message)
	message.From = from
	message.To = to
	message.Subject = subject
	message.Body = body
	return message
}

type MailBot struct {
	Token string
	Name  string
}

func NewMailBot(token string, name string) *MailBot {
	bot := new(MailBot)
	bot.Token = token
	bot.Name = name
	return bot
}

func (bot *MailBot) SendMessage(message Message) error {
	address := fmt.Sprintf("http://localhost:3000/bot%v/sendmessage", bot.Token)
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	r := bytes.NewReader(data)
	_, err = http.Post(address, "application/json", r)
	if err != nil {
		return err
	}
	return nil
}

func (bot *MailBot) GetUpdates(offset int, limit int) ([]Message, error) {
	address := fmt.Sprintf("http://localhost:3000/bot%v/getupdates?offset=%v&limit=%v", bot.Token, offset, limit)
	resp, err := http.Get(address)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var messages []Message
	err = decoder.Decode(&messages)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
