package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type InteractionModel struct {
	Id           primitive.ObjectID `json:"id,omitempty"`
	State        string             `json:"state"`
	ClientNumber string             `json:"client_number"`
	Entity       string
	Service      string
	Type         string
	Zone         string
}

type HistoryModel struct {
	Id           primitive.ObjectID `json:"id,omitempty"`
	ClientNumber string             `json:"client_number"`
	Entity       string
	Service      string
	Type         string
	Zone         string
	Data         string
	Date         string
}

type Files struct {
	Directory string
	Data      []byte
}
