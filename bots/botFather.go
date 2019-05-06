package bots

import (
	"github.com/kraevskii-m/MailBot/botLib"
	"github.com/kraevskii-m/MailBot/data"
	"log"
	"time"
)

var url = "http://localhost:8080/updates"
var address = ""

func BotController() {
	bot := botLib.NewMailBot(data.TokenBotFather) //todo change
	for {
		time.Sleep(5 * time.Second)
		messages, err := bot.GetUpdates()
		if err != nil {
			log.Print(err)
		}
		for _, message := range messages {
			go register(message, data.MemoryStorage{}) //todo fix
		}
	}
}
func register(message botLib.Message, storage data.Storage) {
	err := storage.AddBot(message.Body)
	if err != nil {
		log.Print(err)
		botLib.NewMessage(address, message.From, message.Subject, "Произошла ошибка! "+err.Error())
		return
	}
}

//func getUpdates() {
//	req, err := http.NewRequest("GET", url, nil)
//	req.Header.Add("Authorization", tokenBotFather)
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		log.Println("Error on response.\n[ERRO] -", err)
//	}
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Print(err)
//		os.Exit(1)
//	}
//	var letters = map[string]map[string]string{}
//	if err := json.Unmarshal(body, &letters); err != nil {
//		panic(err)
//	}
//
//	for _, letter := range letters {
//
//		if !isValidAddress(letter["Address"]) {
//			log.Println("Wrong e-mailController: " + letter["Address"])
//			continue
//		}
//		if letter["Theme"] != "Register" {
//			sendWelcomeMessage(letter["Address"])
//		} else {
//			registerBot(letter["Address"], letter["Body"])
//		}
//	}
//}
//
//func registerBot(address string, name string) {
//	if !isValidAddress("name" + "@yandex.ru") {
//		sendWelcomeMessage(address)
//	}
//}
//
//func sendWelcomeMessage(address string) {
//
//}
//
//func isValidAddress(address string) bool {
//Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
//return Re.MatchString(address)
//}
