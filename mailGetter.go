package main

import (
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"log"
	"time"
)

var messagesBase []*imap.Message

//var lastId uint32
//= messagesBase[len(messagesBase)-1].Uid

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

	check(c)

	log.Println("Done!")
}

func check(c *client.Client) {
	for {
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

		time.Sleep(time.Second * 5)
		from := uint32(len(messagesBase) + 1)
		to := mbox.Messages
		if from > to {
			continue
		}
		seqset := new(imap.SeqSet)
		seqset.AddRange(from, to)

		messages := make(chan *imap.Message)
		done = make(chan error, 1)
		go func() {
			done <- c.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope}, messages)
		}()

		for msg := range messages {
			messagesBase = append(messagesBase, msg)
			log.Println("* " + msg.Envelope.Subject)
		}

		if err := <-done; err != nil {
			log.Fatal(err)
		}
	}
}
