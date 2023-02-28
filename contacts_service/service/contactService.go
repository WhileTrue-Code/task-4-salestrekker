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

	contact.Deleted = false
	exist, err := service.repository.DoesContactExist(contact)
	if err != nil && err.Error() == errors.ContactAlreadyExist {
		return err
	} else if err != nil && err.Error() == errors.DeletedContactMsg {
		err := service.repository.RecoverContact(exist)
		if err != nil {
			return err
		}
		return nil
	}

	return service.repository.CreateContact(contact)
}

func (service *ContactService) GetAllContacts() (contacts *domain.Output, error error) {
	return service.repository.GetAllContacts()
}

func (service *ContactService) GetOneContactByID(id string) (contact *domain.Input, error error) {
	primitiveID, err := getPrimitiveIDFromHex(id)
	if err != nil {
		log.Println("getPrimitiveIDFromHex parsing error got")
		return nil, err
	}

	return service.repository.GetOneContactByID(*primitiveID)
}

func (service *ContactService) DeleteOneContactByID(id string) (error error) {
	primitiveID, err := getPrimitiveIDFromHex(id)
	if err != nil {
		log.Println("getPrimitiveIDFromHex parsing error got")
		return err
	}

	return service.repository.DeleteOneContactByID(*primitiveID)
}

func validateFields(contact *domain.Input) error {
	if contact.FirstName == "" || contact.LastName == "" || contact.Telephone == "" {
		return fmt.Errorf(errors.EmptyFieldError)
	}

	return nil
}

func getPrimitiveIDFromHex(id string) (*primitive.ObjectID, error) {
	primitiveID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Error in parsing hex string to primitive.ObjectID because of: %s", err)
		return nil, fmt.Errorf(errors.WrongIdFormatError)
	}

	return &primitiveID, nil
}
