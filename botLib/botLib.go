package botLib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type MailBot struct {
	token string
}

type Message struct {
	from    string
	to      string
	subject string
	body    string
}

func NewMailBot(token string) *MailBot {
	bot := new(MailBot)
	bot.token = token
	return bot
}

func NewMessage(from string, to string, subject string, body string) *Message {
	message := new(Message)
	message.from = from
	message.to = to
	message.subject = subject
	message.body = body
	return message
}

func (bot *MailBot) SendMessage(message Message) error {
	address := fmt.Sprintf("https://localhost:3000/bot%v/getupdates", bot.token)
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

func (bot *MailBot) GetUpdates() ([]Message, error) {
	address := fmt.Sprintf("https://localhost:3000/bot%v/getupdates", bot.token)
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
