package database

import (
	"context"

	"jwt-authentication/models"
)

type Repository interface {
	Migrate(ctx context.Context) error
	DeleteTable(ctx context.Context) error
	Add(ctx context.Context, user models.User) (*models.User, error)
	GetAll(ctx context.Context) ([]models.User, error)
}
