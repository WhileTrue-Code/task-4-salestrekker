package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func GetDatabaseClient(dbPort, dbUsername, dbPassword string) *mongo.Client {
	connectionUri := fmt.Sprintf("mongodb://%s:%s@mongo:%s", dbUsername, dbPassword, dbPort)
	clientOptions := options.Client().ApplyURI(connectionUri)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatalf("Error in connecting to database because of: %s", err)
	}
	log.Printf("Succesfully connected to database and got client.")

	return client
}
