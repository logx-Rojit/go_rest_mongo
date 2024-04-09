package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDb() (*mongo.Client, error) {

	DB_URI := os.Getenv("DB_URI")
	if DB_URI == "" {
		log.Panic("Please provide mongodb uri for database connection")
	}
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(DB_URI))
	if err != nil {
		panic(err)
	}
	fmt.Println("Mongodb connected successfully")
	return client, nil
}

func Disconnect(client *mongo.Client) error {
	return client.Disconnect(context.Background())
}
