package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kraevskii-m/MailBot/data"
	"github.com/kraevskii-m/MailBot/mailController"
	"log"
	"net/http"
	"strconv"
)

func Router() {
	r := mux.NewRouter()
	log.Println("Listening...")
	r.Handle("/", http.FileServer(http.Dir("./about/")))
	r.Handle("/bot{token}/getupdates", getUpdates).Methods("GET")
	r.Handle("/bot{token}/sendmessage", sendMessage).Methods("POST")

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal("Router didn't start!")
	}
}

var getUpdates = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := vars["token"]

	var offset string
	var limit string

	offsetQuery, ok := r.URL.Query()["offset"]
	if !ok || len(offsetQuery[0]) < 1 { //todo test
		offset = ""
	} else {
		offset = offsetQuery[0]
	}

	limitQuery, ok := r.URL.Query()["limit"]
	if !ok || len(limitQuery[0]) < 1 {
		limit = ""
	} else {
		limit = limitQuery[0]
	}

	bot, err := data.Base.GetBot(token)
	if err != nil {
		http.Error(w, "Bot doesn't exist! Send message to BotFather!", http.StatusBadRequest)
		return
	}
	mailController.UpdateMailBox(bot)

	numOffset, err := strconv.Atoi(offset)
	if err != nil {
		http.Error(w, "Offset must be a number!", http.StatusBadRequest)
		return
	}
	numLimit, err := strconv.Atoi(limit)
	if err != nil {
		http.Error(w, "Limit must be a number!", http.StatusBadRequest)
		return
	}

	messages := data.Base.GetMessages(bot, numOffset, numLimit)
	result, _ := json.Marshal(messages)

	w.Write(result)
})

var sendMessage = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := vars["token"]

	bot, err := data.Base.GetBot(token)
	if err != nil {
		http.Error(w, "Bot doesn't exist! Send message to BotFather!", http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var message data.Message
	err = decoder.Decode(&message)
	if err != nil {
		http.Error(w, "Can't parse your request!", http.StatusBadRequest)
		return
	}

	err = mailController.SendMessage(message, bot)
	if err != nil {
		http.Error(w, "Can't send message!", http.StatusBadRequest)
		return
	}
	w.Write([]byte("Message sent!"))
})
