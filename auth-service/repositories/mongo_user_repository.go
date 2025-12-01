package repositories

import (
	"auth-service/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	Collection *mongo.Collection
}

func NewMongoUserRepository(collection *mongo.Collection) *MongoUserRepository {
	return &MongoUserRepository{Collection: collection}
}

func (r *MongoUserRepository) CreateUser(ctx context.Context, user models.User) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	_, err := r.Collection.InsertOne(ctx, user)
	return err
}

func (r *MongoUserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var user models.User
	err := r.Collection.FindOne(ctx, map[string]interface{}{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
