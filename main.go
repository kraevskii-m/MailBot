package main

import (
	"github.com/kraevskii-m/MailBot/bots"
	"github.com/kraevskii-m/MailBot/data"
	"github.com/kraevskii-m/MailBot/server"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go server.Router()

	wg.Add(1)
	go bots.BotFatherController(data.Base)

	wg.Wait()
}
