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

type stocks struct {
	client *mongo.Client
}

func GetStocks() ([]models.Stock, error) {
	client, err := database.ConnectToMongoDB()
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	collection := client.Database(config.Database).Collection(config.StockCollection)

	var stock []models.Stock

	// Realizando a consulta
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.Background())

	// Iterando sobre os resultados
	if err := cursor.All(context.Background(), &stock); err != nil {
		log.Fatal(err)
	}

	return stock, nil
}
