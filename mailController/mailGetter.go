package mailController

import (
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
	"github.com/kraevskii-m/MailBot/data"
	"io"
	"io/ioutil"
	"log"
	"time"
)

func checkMailBox(bot data.Bot) []data.Message {
	c, err := client.DialTLS("imap.yandex.ru:993", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	if err := c.Login(bot.Username, bot.Password); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")

	defer c.Logout()

	for {
		mbox, err := c.Select("INBOX", false)
		if err != nil {
			log.Fatal(err)
		}
		if mbox.Messages == 0 {
			log.Println("No messages in mailbox")
		}
		time.Sleep(time.Second * 5)

		from := uint32(1)
		to := mbox.Messages
		if from > to {
			continue
		}
		seqset := new(imap.SeqSet)
		seqset.AddRange(from, to)

		section := &imap.BodySectionName{}
		items := []imap.FetchItem{section.FetchItem()}

		messages := make(chan *imap.Message)
		go func() {
			if err := c.Fetch(seqset, items, messages); err != nil {
				log.Fatal(err)
			}
		}()

		var output []data.Message

		for msg := range messages {
			if msg == nil {
				log.Fatal("Server didn't returned message")
			}

			r := msg.GetBody(section)
			if r == nil {
				log.Fatal("Server didn't returned message body")
			}

			mr, err := mail.CreateReader(r)
			if err != nil {
				log.Fatal(err)
			}

			for {
				p, err := mr.NextPart()
				if err == io.EOF {
					break
				} else if err != nil {
					log.Fatal(err)
				}

				header := mr.Header
				switch h := p.Header.(type) {
				case mail.TextHeader:
					b, _ := ioutil.ReadAll(p.Body)
					from, _ := header.AddressList("From")
					to, _ := header.AddressList("To")
					subject, _ := header.Subject()
					message := data.Message{From: from[0].Address, To: to[0].Address, Subject: subject, Body: string(b)}
					output = append(output, message) //todo fix
				case mail.AttachmentHeader:
					filename, _ := h.Filename()
					log.Println("Got attachment: %v", filename)
				}
			}
		}
	}
}
