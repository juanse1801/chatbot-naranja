package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type InteractionModel struct {
	Id           primitive.ObjectID `json:"id,omitempty"`
	State        string             `json:"state"`
	ClientNumber string             `json:"client_number"`
	ClientMail   string             `json:"client_mail"`
}
