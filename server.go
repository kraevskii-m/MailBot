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

func server() {
	http.HandleFunc("/send", postHandler)
	http.HandleFunc("/updates", getUpdates)
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

func postHandler(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var letter SendMailStruct
	err := decoder.Decode(&letter)
	if err != nil {
		panic(err)
	}
	mailSender(letter.to, letter.body, letter.subject)
}

func getUpdates(writer http.ResponseWriter, request *http.Request) {
	var formattedMessages [][]string
	for _, msg := range messagesBase {
		formattedMessages = append(formattedMessages, []string{msg.to, msg.body, msg.subject})
	}
	output, _ := json.Marshal(formattedMessages)
	writer.Write(output)
}
