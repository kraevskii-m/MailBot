package main

import (
	"github.com/kraevskii-m/MailBot/bots"
	"github.com/kraevskii-m/MailBot/mailController"
	"github.com/kraevskii-m/MailBot/server"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go server.Router()

	wg.Add(1)
	go bots.BotFatherController()

	wg.Add(1)
	go mailController.BotMonitor()

	wg.Wait()
}
