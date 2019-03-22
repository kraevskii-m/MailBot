package main

import (
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go mailGetter()

	wg.Add(1)
	go server()

	wg.Wait()
}
