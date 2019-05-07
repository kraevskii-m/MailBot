package main

import (
	"github.com/kraevskii-m/MailBot/bots"
	"github.com/kraevskii-m/MailBot/server"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go server.Start()

	wg.Add(1)
	go bots.BotController()

	wg.Wait()
}
