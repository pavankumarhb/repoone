package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Create Struct
// omitempty means that If we didn’t assign any value our struct’s field, don’t show this field after serialize process. Also, bson is relevant to mongo-driver. It works for creating filter query. We will see as soon as.
type Book struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string `json:"firstname,omitempty" bson:"firstname,omitempty"`        
	Title  string             `json:"title" bson:"title,omitempty"`
	Author *Author            `json:"author" bson:"author,omitempty"`
}

type Author struct {
	FirstName string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty" bson:"lastname,omitempty"`
}