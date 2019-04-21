package mailController

import (
	"encoding/json"
	"github.com/kraevskii-m/MailBot/data"
	"io"
	"strconv"
)

func Get(token string, offset string, limit string) ([]byte, error) {
	letters := data.GetLetters(token)
	if offset != "" {
		num, err := strconv.Atoi(offset)
		if err != nil {
			return nil, err
		}
		letters = letters[num:]
	}
	if limit != "" {
		num, err := strconv.Atoi(limit)
		if err != nil {
			return nil, err
		}
		letters = letters[len(letters)-num:] //TODO fix
	}

	var formattedLetters [][]string

	for _, let := range letters {
		formattedLetters = append(formattedLetters, []string{let.From, let.To, let.Subject, let.Body})
	}
	output, _ := json.Marshal(formattedLetters)
	return output, nil
}

func MailSender(token string, body io.ReadCloser) {
	decoder := json.NewDecoder(body)
	var letters [][]string
	err := decoder.Decode(&letters)
	if err != nil {
		panic(err)
	}
	for _, let := range letters {
		sender(let[0], let[1], let[2], let[3])
	}
}
