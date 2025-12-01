package repositories

import (
	"auth-service/models"
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
}
