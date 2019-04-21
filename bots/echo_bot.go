package bots

import (
	"log"
	"net/http"
)

var url = "http://localhost:3000"
var token = "asdfghjk"

func echoBot() {
	resp, err := http.Get(url + "/bot" + token + "/getupdates")
	if err != nil {
		log.Fatal(err)
	}
}
