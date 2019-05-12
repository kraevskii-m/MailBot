package mailController

import "github.com/kraevskii-m/MailBot/data"

//todo refactor

func UpdateMailBox(token string) {
	panic("Not Implemented")
}

func SendMessage(message data.Message, bot data.Bot) error {
	panic("Not Implemented")
}

//func Get(token string, offset string, limit string) ([]byte, error) {
//	letters, err := data.GetLetters(token)
//	if err != nil {
//		return nil, err
//	}
//	if offset != "" {
//		num, err := strconv.Atoi(offset)
//		if err != nil {
//			return nil, err
//		}
//		letters = letters[num:]
//	}
//	if limit != "" {
//		num, err := strconv.Atoi(limit)
//		if err != nil {
//			return nil, err
//		}
//		letters = letters[len(letters)-num:] //TODO fix
//	}
//
//	var formattedLetters [][]string
//
//	for _, let := range letters {
//		formattedLetters = append(formattedLetters, []string{let.From, let.To, let.Subject, let.Body})
//	}
//	output, _ := json.Marshal(formattedLetters)
//	return output, nil
//}
//
//func MailSender(token string, body io.ReadCloser) {
//	decoder := json.NewDecoder(body)
//	var letters [][]string
//	err := decoder.Decode(&letters)
//	if err != nil {
//		panic(err)
//	}
//	for _, let := range letters {
//		sender(let[0], let[1], let[2], let[3])
//	}
//}
//
//func Initialize() {
//	data.NewBot("botfather")
//}
//
//func UpdatesController() {
//	Initialize()
//
//	for {
//		time.Sleep(1 * time.Second)
//		var base = data.BotStorage.Load()
//		var botBase []data.Bot
//		if base == nil {
//			continue
//		}
//
//		wg := &sync.WaitGroup{}
//
//		wg.Add(1) //TODO Check
//		for _, bot := range botBase {
//			go UpdateMailbox(bot)
//		}
//	}
//}
//
//func UpdateMailbox(bot data.Bot) {
//	var base = data.LetterStorage.Load()
//	var letterBase map[string][]data.Message
//	if base == nil {
//		data.LetterStorage.Store(letterBase)
//		return
//	}
//	letterBase = base.(map[string][]data.Message)
//	letterBase[bot.Token] = GetUpdatesForBot(bot.Token)
//	data.LetterStorage.Store(letterBase)
//}
