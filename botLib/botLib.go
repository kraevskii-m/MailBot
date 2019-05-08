package botLib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	address := fmt.Sprintf("https://localhost:3000/bot%v/sendmessage", bot.Token)
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

func (bot *MailBot) GetUpdates(offset int, limit int) ([]Message, error) { //todo fix offset and limit
	address := fmt.Sprintf("https://localhost:3000/bot%v/getupdates", bot.Token)
	resp, err := http.Get(address)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var messages []Message
	err = json.Unmarshal(body, &messages)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return messages, nil
}
