package mailing

import (
	"errors"
	"fmt"
	"net/smtp"
	"os"

	"github.com/juanse1801/chatbot-naranja/pkg/configs"
)

type Service interface {
	SendEmail(service string, zone string, data string) error
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) SendEmail(service string, zone string, data string) error {
	from := "juanse1801.dev@gmail.com"

	to := configs.EmailsConfig[service][zone]

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	subject := "Subject: This is the subject of the mail\n"
	message := []byte(subject + data)

	auth := smtp.PlainAuth("", from, os.Getenv("MAIL_PASS"), host)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New(err.Error())
	}

	return nil
}
