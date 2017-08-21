package mail

import (
	"log"
	"net/smtp"
	"os"
	"time"

	"github.com/VG-Tech-Dojo/treasure2017/mid/yaaaaashiki/crypto"
	"github.com/joho/godotenv"
)

const (
	SubTitle        = "Ericacho からの通知です!!!!!!"
	PerMinute       = 60
	CharacterLength = 512
)

type Mail struct {
	From     string
	Username string
	Password string
	To       string
	Sub      string
	Msg      string
}

func Env_load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func SendMail(toAddress string) {
	Env_load()
	m := &Mail{
		From:     os.Getenv("USERNAME"),
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
		To:       toAddress,
		Sub:      SubTitle,
		Msg:      crypto.Salt(CharacterLength),
	}

	if err := GmailSend(m); err != nil {
		log.Println(err)
	}
}
func (m *Mail) body() string {
	return "To: " + m.To + "\r\n" +
		"Subject: " + m.Sub + "\r\n\r\n" +
		m.Msg + "\r\n"
}

func GmailSend(m *Mail) error {
	smtpSvr := os.Getenv("HOSTSERVER")
	auth := smtp.PlainAuth("", m.Username, m.Password, os.Getenv("HOST"))
	if err := smtp.SendMail(smtpSvr, auth, m.From, []string{m.To}, []byte(m.body())); err != nil {
		return err
	}
	return nil
}

func InfiniteMail(toAddress string) {
	i := 0
	for {
		SendMail(toAddress)
		if i != 0 {
			time.Sleep(PerMinute * time.Second)
		}
		i++
	}
}
