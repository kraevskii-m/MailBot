package data

type Bot struct {
	Token    string
	Username string
	Password string
}

var database map[string]Bot

func Initialize() {
	NewBot("botfather")
}

func NewBot(name string) Bot {
	output := Bot{}
	//TODO It's temporary solution
	if name == "botfather" {
		output.Token = "qwertyui"
		output.Username = "fatherofbots"
		output.Password = "lermonter07"
	}
	if name == "echobot" {
		output.Token = "asdfghjk"
		output.Username = "echobot-mailbot"
		output.Password = "lermonter07"
	}

	return output
}

func GetBot(name string) Bot {
	return Bot{}
}

func GetByToken(token string) Bot {
	return Bot{}
}

type Letter struct {
	From    string
	To      string
	Subject string
	Body    string
}

func GetLetters(token string) []Letter {

}
