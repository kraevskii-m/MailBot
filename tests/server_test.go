package tests

import (
	"github.com/kraevskii-m/MailBot/server"
	"log"
	"testing"
)

func TestUpdates(t *testing.T) {
	go server.Start()
	log.Print("It works!")
}
