package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var UserCollection *mongo.Collection
var PlanCollection *mongo.Collection
var APIKeyCollection *mongo.Collection

func Connect(connectionString string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}

	Client = client
	UserCollection = client.Database("go-sail-auth").Collection("users")
	PlanCollection = client.Database("go-sail-auth").Collection("plans")
	APIKeyCollection = client.Database("go-sail-auth").Collection("apikeys")
	log.Println("Connected to Database!")
}
