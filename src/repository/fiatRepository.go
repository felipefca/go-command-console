package repository

import (
	"console/src/config"
	"console/src/database"
	"console/src/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type fiat struct {
	client *mongo.Client
}

func GetFiats() ([]models.Fiat, error) {
	client, err := database.ConnectToMongoDB()
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	collection := client.Database(config.Database).Collection(config.FiatCollection)

	var fiat []models.Fiat

	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &fiat); err != nil {
		log.Fatal(err)
	}

	return fiat, nil
}
