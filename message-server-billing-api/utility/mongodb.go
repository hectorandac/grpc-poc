package utility

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println("Connected to MongoDB")
	return client
}
