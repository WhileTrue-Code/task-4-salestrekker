package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ContactRepository interface {
	CreateContact(contact *Input) error
	GetAllContacts() (contacts *Output, error error)
	GetOneContactByID(id primitive.ObjectID) (contact *Input, error error)
	DeleteOneContactByID(id primitive.ObjectID) (error error)
}
