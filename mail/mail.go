package mail

import (
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

const (
	SubTitle = "Ericacho 通知です!!!!!!"
)

func Env_load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	Env_load()
	m := Mail{
		From:     os.Getenv("USERNAME"),
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
		To:       "yaaaaakishi@gmail.com",
		Sub:      SubTitle,
		Msg:      "本文insert",
	}

	if err := GmailSend(m); err != nil {
		log.Println(err)
	}
}

type Mail struct {
	From     string
	Username string
	Password string
	To       string
	Sub      string
	Msg      string
}

func (m Mail) body() string {
	return "To: " + m.To + "\r\n" +
		"Subject: " + m.Sub + "\r\n\r\n" +
		m.Msg + "\r\n"
}

func GmailSend(m Mail) error {
	smtpSvr := os.Getenv("HOSTSERVER")
	auth := smtp.PlainAuth("", m.Username, m.Password, os.Getenv("HOST"))
	if err := smtp.SendMail(smtpSvr, auth, m.From, []string{m.To}, []byte(m.body())); err != nil {
		return err
	}
	return nil
}
