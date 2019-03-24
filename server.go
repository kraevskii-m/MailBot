package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type SendMailStruct struct {
	to      string
	body    string
	subject string
}

var sentMailNumber = 0

func server() {
	http.HandleFunc("/send", sengLetters)
	http.HandleFunc("/updates", getUpdates)
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

func sengLetters(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var letters [][]string
	err := decoder.Decode(&letters)
	if err != nil {
		panic(err)
	}
	for _, let := range letters[sentMailNumber:] {
		mailSender(let[0], let[1], let[2])
	}
	sentMailNumber = len(letters)
}

func getUpdates(writer http.ResponseWriter, request *http.Request) {
	var formattedMessages [][]string
	var mbase = MessagesBase.Load().([]SendMailStruct)
	for _, msg := range mbase {
		formattedMessages = append(formattedMessages, []string{msg.to, msg.body, msg.subject})
	}
	output, _ := json.Marshal(formattedMessages)
	writer.Write(output)
}
