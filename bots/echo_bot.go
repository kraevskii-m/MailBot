package bots

import (
	"log"
	"net/http"
)

var address = "" //todo add current BotFatherEmail

func echoBot() { //todo complete ECHOBOT
	resp, err := http.Get(url + "/bot" + token + "/getupdates")
	if err != nil {
		log.Fatal(err)
	}
}
