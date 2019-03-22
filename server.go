package main

import (
	"fmt"
	"log"
	"net/http"
)

func server() {
	http.HandleFunc("/", postHandler)
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

func postHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello")
}
