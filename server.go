package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type sendMailStruct struct {
	to      string
	body    string
	subject string
}

func server() {
	http.HandleFunc("/send", postHandler)
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

func postHandler(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var letter sendMailStruct
	err := decoder.Decode(&letter)
	if err != nil {
		panic(err)
	}
	mailSender(letter.to, letter.body, letter.subject)
}
