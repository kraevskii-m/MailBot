package main

import (
	"github.com/kraevskii-m/MailBot/server"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go server.Start()

	wg.Wait()
}
