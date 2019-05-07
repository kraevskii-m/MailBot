package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() { //todo make cleanup!
	r := mux.NewRouter()
	log.Println("Listening...")
	r.Handle("/", http.FileServer(http.Dir("./about/")))
	r.Handle("/bot{token}/getupdates", getUpdates).Methods("GET")
	r.Handle("/bot{token}/sendmessage", sendMessage).Methods("POST")

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal("Start not starting!")
	}
}

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

	log.Print(token + offset + limit)
	//output, err := mailController.Get(token, offset, limit)
	//if err != nil {
	//	http.Error(w, "No new messages", http.StatusBadRequest)
	//	return
	//}
	//w.Write(output)
})

var sendMessage = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := vars["token"]
	log.Print(token)
	//mailController.MailSender(token, r.Body)
})
