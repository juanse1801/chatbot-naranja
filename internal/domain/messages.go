package domain

type NewMessageReceived struct {
	Object string  `json:"object"`
	Entry  []Entry `json:"entry"`
}

type Entry struct {
	Id      string    `json:"id"`
	Changes []Changes `json:"changes"`
}

type Changes struct {
	Value Value  `json:"value"`
	Field string `json:"field"`
}

type Value struct {
	MessagingProduct string     `json:"messaging_product"`
	Metadata         Metadata   `json:"metadata"`
	Contacts         []Contacts `json:"contacts"`
	Messages         []Messages `json:"messages"`
}

type Metadata struct {
	DisplayPhoneNumber string `json:"display_phone_number"`
	PhoneNumberId      string `json:"phone_number_id"`
}

type Contacts struct {
	Profile Profile `json:"profile"`
	WaId    string  `json:"wa_id"`
}

type Profile struct {
	Name string `json:"name"`
}

type Messages struct {
	From      string `json:"from"`
	Id        string `json:"id"`
	TimeStamp string `json:"timestamp"`
	Text      Text   `json:"text"`
	Type      string `json:"type"`
}

type Text struct {
	Body string `json:"body"`
}
