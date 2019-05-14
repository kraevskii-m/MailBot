package mailController

import (
	"crypto/tls"
	"fmt"
	"github.com/kraevskii-m/MailBot/data"
	"log"
	"net/smtp"
	"strings"
)

type Mail struct {
	senderId string
	toIds    []string
	subject  string
	body     string
}

type smtpServer struct {
	host string
	port string
}

func (s *smtpServer) serverName() string {
	return s.host + ":" + s.port
}

func (mail *Mail) buildMessage() string {
	message := ""
	message += fmt.Sprintf("From: %s\r\n", mail.senderId)
	if len(mail.toIds) > 0 {
		message += fmt.Sprintf("To: %s\r\n", strings.Join(mail.toIds, ";"))
	}

	message += fmt.Sprintf("Subject: %s\r\n", mail.subject)
	message += "\r\n" + mail.body

	return message
}

func Sender(bot data.Bot, recepient string, body string, subject string) error {
	mail := Mail{}
	mail.senderId = bot.Username + "@yandex.ru"
	mail.toIds = []string{recepient}
	mail.subject = subject
	mail.body = body

	messageBody := mail.buildMessage()

	smtpServer := smtpServer{host: "smtp.yandex.ru", port: "465"}

	log.Println(smtpServer.host)
	auth := smtp.PlainAuth("", mail.senderId, bot.Password, smtpServer.host)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer.host,
	}

	conn, err := tls.Dial("tcp", smtpServer.serverName(), tlsconfig)
	if err != nil {
		return err
	}

	client, err := smtp.NewClient(conn, smtpServer.host)
	if err != nil {
		return err
	}

	if err = client.Auth(auth); err != nil {
		return err
	}

	if err = client.Mail(mail.senderId); err != nil {
		return err
	}
	for _, k := range mail.toIds {
		if err = client.Rcpt(k); err != nil {
			return err
		}
	}

	w, err := client.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(messageBody))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	client.Quit()

	log.Println("Mail sent successfully " + string(mail.toIds[0]))
	return nil
}
