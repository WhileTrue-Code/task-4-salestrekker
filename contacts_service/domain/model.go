package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Input struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	FirstName string             `bson:"firstName" json:"firstName"`
	LastName  string             `bson:"lastName" json:"lastName"`
	Telephone string             `bson:"telephone" json:"telephone"`
	Deleted   bool               `bson:"deleted" json:"deleted"`
}

type Output struct {
	Results []Input `json:"results"`
}
