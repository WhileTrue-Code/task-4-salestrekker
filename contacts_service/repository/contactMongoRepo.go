package repository

import (
	"contacts_service/domain"
	"contacts_service/errors"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var collectionName = "contacts"

type ContactMongoRepo struct {
	dbClient     *mongo.Client
	dbCollection *mongo.Collection
}

func NewContactMongoRepo(dbClient *mongo.Client, dbName string) domain.ContactRepository {
	contactsCollection := dbClient.Database(dbName).Collection(collectionName)
	return &ContactMongoRepo{
		dbClient:     dbClient,
		dbCollection: contactsCollection,
	}
}

func (repo *ContactMongoRepo) CreateContact(contact *domain.Input) error {
	result, err := repo.dbCollection.InsertOne(context.TODO(), contact)
	if err != nil {
		log.Printf("Error in saving contact to database because of: %s", err)
		return err
	}
	log.Printf("Inserted id is: %s", result.InsertedID)
	return nil
}

func (repo *ContactMongoRepo) DoesContactExist(contact *domain.Input) error {
	mongoResult := repo.dbCollection.FindOne(context.TODO(), bson.M{"firstName": contact.FirstName,
		"lastName":  contact.LastName,
		"telephone": contact.Telephone})

	var result domain.Input
	err := mongoResult.Decode(&result)
	if err != nil {
		return nil
	}

	return fmt.Errorf(errors.ContactAlreadyExist)
}

func (repo *ContactMongoRepo) GetAllContacts() (contacts *domain.Output, error error) {
	//TODO implement me
	panic("implement me")
}

func (repo *ContactMongoRepo) GetOneContactByID(id primitive.ObjectID) (contact *domain.Input, error error) {
	//TODO implement me
	panic("implement me")
}

func (repo *ContactMongoRepo) DeleteOneContactByID(id primitive.ObjectID) (error error) {
	//TODO implement me
	panic("implement me")
}
