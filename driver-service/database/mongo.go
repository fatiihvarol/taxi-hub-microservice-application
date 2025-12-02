package database

import (
	"context"
	"log"
	"time"

	"driver-service/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var UserCollection *mongo.Collection
var DriverCollection *mongo.Collection

func ConnectMongo() {
	uri := config.GetEnv("MONGO_URI")
	dbName := config.GetEnv("MONGO_DATABASE") // env ile aynı isim
	driverCollectionName := config.GetEnv("MONGO_DRIVER_COLLECTION")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("MongoDB Connection Error:", err)
	}

	Client = client
	UserCollection = client.Database(dbName).Collection("users")
	DriverCollection = client.Database(dbName).Collection(driverCollectionName)

	log.Println("MongoDB connected")
}

func GetDriverCollection() *mongo.Collection {
	if Client == nil {
		log.Fatal("MongoDB client not initialized")
	}
	return DriverCollection
}

func GetCollection(dbName, colName string) *mongo.Collection {
	if Client == nil {
		log.Fatal("MongoDB client not initialized")
	}
	return Client.Database(dbName).Collection(colName)
}
