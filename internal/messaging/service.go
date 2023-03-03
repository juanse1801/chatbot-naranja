package messaging

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type messageError struct {
	Error error `json:"error"`
}

type error struct {
	Message string `json:"message"`
}

type Service interface {
	SendMessage(to string, text string)
	SendRecontactMessage(to string)
	SendNoContactMessage(to string)
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
	r.Header.Add("Authorization", "Bearer EAAKrLKufN4oBAPUr9Pq8W7F6GSuirm4LlCUNrrVkZBgRFxRZBYwj1kSK9sZBipPErwUgPYpnb23ZBmZBa7py9HhZAZCXs7UvEZCYeTTO7mYZBoPleLhlz2aX7wUB2wxkGjMqnvyJOXtBjdvIZC6fchkMYjENAgjZAj6SiyZCOooqSKc4xXJYEZBT4QAxKvjvsIwZC9m9je8S7C9b8zxkYzayZCLcbdV")

	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	msgerror := &messageError{}

	derr := json.NewDecoder(response.Body).Decode(msgerror)
	if derr != nil {
		log.Fatal(derr)
	}

	if response.StatusCode == 200 {
		fmt.Println("Mensaje enviado")
	} else {
		log.Println(response.StatusCode)
		log.Println(msgerror.Error.Message)
	}
}

func (s *service) SendRecontactMessage(to string) {
	posturl := "https://graph.facebook.com/v16.0/108319032187494/messages"

	message := fmt.Sprintf(`{
		"messaging_product": "whatsapp",
		"to": "%s",
		"type": "text",
		"text": {
			"body": "¿Sigues ahí?"
		}
	}`, to)

	body := []byte(message)

	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Bearer EAAKrLKufN4oBAPUr9Pq8W7F6GSuirm4LlCUNrrVkZBgRFxRZBYwj1kSK9sZBipPErwUgPYpnb23ZBmZBa7py9HhZAZCXs7UvEZCYeTTO7mYZBoPleLhlz2aX7wUB2wxkGjMqnvyJOXtBjdvIZC6fchkMYjENAgjZAj6SiyZCOooqSKc4xXJYEZBT4QAxKvjvsIwZC9m9je8S7C9b8zxkYzayZCLcbdV")

	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	msgerror := &messageError{}

	derr := json.NewDecoder(response.Body).Decode(msgerror)
	if derr != nil {
		log.Fatal(derr)
	}

	if response.StatusCode == 200 {
		fmt.Println("Mensaje enviado")
	} else {
		log.Println(response.StatusCode)
		log.Println(msgerror.Error.Message)
	}
}

func (s *service) SendNoContactMessage(to string) {
	posturl := "https://graph.facebook.com/v16.0/108319032187494/messages"

	message := fmt.Sprintf(`{
		"messaging_product": "whatsapp",
		"to": "%s",
		"type": "text",
		"text": {
			"body": "Muchas gracias por comunicarte con nosotros."
		}
	}`, to)

	body := []byte(message)

	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Bearer EAAKrLKufN4oBAPUr9Pq8W7F6GSuirm4LlCUNrrVkZBgRFxRZBYwj1kSK9sZBipPErwUgPYpnb23ZBmZBa7py9HhZAZCXs7UvEZCYeTTO7mYZBoPleLhlz2aX7wUB2wxkGjMqnvyJOXtBjdvIZC6fchkMYjENAgjZAj6SiyZCOooqSKc4xXJYEZBT4QAxKvjvsIwZC9m9je8S7C9b8zxkYzayZCLcbdV")

	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	msgerror := &messageError{}

	derr := json.NewDecoder(response.Body).Decode(msgerror)
	if derr != nil {
		log.Fatal(derr)
	}

	if response.StatusCode == 200 {
		fmt.Println("Mensaje enviado")
	} else {
		log.Println(response.StatusCode)
		log.Println(msgerror.Error.Message)
	}
}
