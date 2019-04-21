package main

import (
	"bytes"
	"github.com/gorilla/mux"
	"github.com/kraevskii-m/MailBot/data"
	"github.com/kraevskii-m/MailBot/mailController"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	server()

	wg.Add(1)
	data.UpdatesController()

	wg.Wait()
}

func server() {
	r := mux.NewRouter()
	log.Println("Listening...")
	r.Handle("/", http.FileServer(http.Dir("./about/")))
	r.Handle("/bot{token}/getupdates", getUpdates).Methods("GET")
	r.Handle("/bot{token}/sendmessage", sendMessage).Methods("POST")
	r.Handle("/bot{token}/register", register).Methods("POST")

	http.ListenAndServe(":3000", r)
}

var register = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["token"] != data.GetBot("botfather").Token {
		http.Error(w, "You are not BotFather!", http.StatusBadRequest)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	name := bytes.NewBuffer(body).String()
	log.Println("Adding bot" + name)
	token := data.NewBot(name).Token
	w.Write([]byte(token))
})

var getUpdates = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := vars["token"]

	var offset string
	var limit string

	offsetQuery, ok := r.URL.Query()["offset"]
	if !ok && len(offsetQuery[0]) < 1 {
		offset = ""
	}
	limitQuery, ok := r.URL.Query()["limit"]
	if !ok || len(limitQuery[0]) < 1 {
		limit = ""
	}

	output, err := mailController.Get(token, offset, limit)
	if err != nil {
		http.Error(w, "No new messages", http.StatusBadRequest)
		return
	}
	w.Write(output)
})

var sendMessage = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := vars["token"]

	mailController.MailSender(token, r.Body)
})
