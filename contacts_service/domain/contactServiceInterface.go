package domain

type ContactService interface {
	CreateContact(contact *Input) error
	GetAllContacts() (contacts *Output, error error)
	GetOneContactByID(id string) (contact *Input, error error)
	DeleteOneContactByID(id string) (error error)
}
