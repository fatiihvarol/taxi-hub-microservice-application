package database

import (
	"context"
	"log"
	"time"

	"auth-service/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var UserCollection *mongo.Collection

func ConnectMongo() {
	uri := config.GetEnv("MONGO_URI")
	dbName := config.GetEnv("MONGO_DB")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("MongoDB Connection Error:", err)
	}

	Client = client
	UserCollection = client.Database(dbName).Collection("users")

	log.Println("MongoDB connected")
}
