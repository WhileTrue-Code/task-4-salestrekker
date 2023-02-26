package service

import (
	"contacts_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ContactService struct {
	repository domain.ContactRepository
}

func NewContactService(repository domain.ContactRepository) *ContactService {
	return &ContactService{repository: repository}
}

func (service *ContactService) CreateContact(contact *domain.Input) error {
	contact.ID = primitive.NewObjectID()
	return service.repository.CreateContact(contact)
}

func (service *ContactService) GetAllContacts() (contacts *domain.Output, error error) {
	//TODO implement me
	panic("implement me")
}

func (service *ContactService) GetOneContactByID(id string) (contact *domain.Input, error error) {
	//TODO implement me
	panic("implement me")
}

func (service *ContactService) DeleteOneContactByID(id string) (error error) {
	//TODO implement me
	panic("implement me")
}
