package messaging

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

type Service interface {
	SendMessage(to string, text string)
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) SendMessage(to string, text string) {
	posturl := "https://graph.facebook.com/v16.0/108319032187494/messages"

	message := fmt.Sprintf(`{
		"messaging_product": "whatsapp",
		"to": "%s",
		"type": "text",
		"text": {
			"body": "%s"
		}
	}`, to, text)

	body := []byte(message)

	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Bearer EAAKrLKufN4oBAEHfQ2Bt6VJYbl6ryehZBHZB1bMn1xTaFIHA7jz7F5YEKBU9MrbBaZBNVgI9iIuOpeiYKJXBZCiF8cWa2105fe15wCa1cFy9FwJdHylFaqFsrbOiJc6bVM8MfxWKG5gUZBu5HuG4PDaB6EUPuW9Q51HjRD4CEJTgSoNyYPrZBWx22eh2YEZCKZBLSCXzk0urbRlZCKEv1L4mc")

	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode == 200 {
		fmt.Println("Mensaje enviado")
	}
}
