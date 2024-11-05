package data

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Collection *mongo.Collection

func Connect() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal((err))
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal((err))
	}

	fmt.Println("mongo connected")

	Collection = client.Database("Task_manager").Collection("tasks")

	fmt.Println(Collection)
}