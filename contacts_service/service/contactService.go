package service

import (
	"contacts_service/domain"
	"contacts_service/errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type ContactService struct {
	repository domain.ContactRepository
}

func NewContactService(repository domain.ContactRepository) *ContactService {
	return &ContactService{repository: repository}
}

func (service *ContactService) CreateContact(contact *domain.Input) error {
	contact.ID = primitive.NewObjectID()

	err := validateFields(contact)
	if err != nil {
		log.Println("Error in validating fields of contact object: %s", err)
		return err
	}

	err = service.repository.DoesContactExist(contact)
	if err != nil {
		log.Println("Error in getting contact existing information: %s", err)
		return err
	}

	return service.repository.CreateContact(contact)
}

func (service *ContactService) GetAllContacts() (contacts *domain.Output, error error) {
	return service.repository.GetAllContacts()
}

func (service *ContactService) GetOneContactByID(id string) (contact *domain.Input, error error) {
	primitiveID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Error in parsing hex string to primitive.ObjectID because of: %s", err)
		return nil, fmt.Errorf(errors.WrongIdFormatError)
	}

	return service.repository.GetOneContactByID(primitiveID)
}

func (service *ContactService) DeleteOneContactByID(id string) (error error) {
	//TODO implement me
	panic("implement me")
}

func validateFields(contact *domain.Input) error {
	if contact.FirstName == "" || contact.LastName == "" || contact.Telephone == "" {
		return fmt.Errorf(errors.EmptyFieldError)
	}

	return nil
}
