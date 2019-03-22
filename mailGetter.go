package main

import (
	"log"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

func mailGetter() {

	log.Println("Connecting to server...")

	c, err := client.DialTLS("imap.yandex.ru:993", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	defer c.Logout()

	if err := c.Login("fatherofbots", "lermonter07"); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")

	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.List("", "*", mailboxes)
	}()

	if err := <-done; err != nil {
		log.Fatal(err)
	}

	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}

	from := uint32(1)
	to := mbox.Messages
	if mbox.Messages > 3 {
		from = mbox.Messages - 3
	}
	seqset := new(imap.SeqSet)
	seqset.AddRange(from, to)

	messages := make(chan *imap.Message, 10)
	done = make(chan error, 1)
	go func() {
		done <- c.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope}, messages)
	}()

	for msg := range messages {
		log.Println("* " + msg.Envelope.Subject)
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}

	log.Println("Done!")
}
