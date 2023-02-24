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

	fmt.Println(message)

	body := []byte(message)

	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Bearer EAAKrLKufN4oBAKlNPh1636CzoJUq5JnhMYH6z9BJl6R60a0JUnvlOEYSjIl4Qs7wwFGgPY7UWqY35na5eZAwoU0EZCIXcSUjquflBDJWznKaCqSyKmZCCDn9nH7ZADZAFWNGE9TpWLaliYCX7Q5JBaNnCntkKRDLxNQ9eYJTJxKVddgEr6teaBU2psrdUeYzplynggQ1gVszzIZCxX1ZAbo")

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

	fmt.Println(message)

	body := []byte(message)

	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Bearer EAAKrLKufN4oBAKlNPh1636CzoJUq5JnhMYH6z9BJl6R60a0JUnvlOEYSjIl4Qs7wwFGgPY7UWqY35na5eZAwoU0EZCIXcSUjquflBDJWznKaCqSyKmZCCDn9nH7ZADZAFWNGE9TpWLaliYCX7Q5JBaNnCntkKRDLxNQ9eYJTJxKVddgEr6teaBU2psrdUeYzplynggQ1gVszzIZCxX1ZAbo")

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

	fmt.Println(message)

	body := []byte(message)

	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Bearer EAAKrLKufN4oBAKlNPh1636CzoJUq5JnhMYH6z9BJl6R60a0JUnvlOEYSjIl4Qs7wwFGgPY7UWqY35na5eZAwoU0EZCIXcSUjquflBDJWznKaCqSyKmZCCDn9nH7ZADZAFWNGE9TpWLaliYCX7Q5JBaNnCntkKRDLxNQ9eYJTJxKVddgEr6teaBU2psrdUeYzplynggQ1gVszzIZCxX1ZAbo")

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
