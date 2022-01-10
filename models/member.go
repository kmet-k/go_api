package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Member struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Firstname string             `json:"firstname" bson:"firstname"`
	Lastname  string             `json:"lastname" bson:"lastname"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"password" bson:"password"`
}

type PutMember struct {
	Firstname string `json:"firstname" bson:"firstname"`
	Lastname  string `json:"lastname" bson:"lastname"`
	Username  string `json:"username" bson:"username"`
	Password  string `json:"password" bson:"password"`
}
