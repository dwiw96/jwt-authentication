package database

import (
	"context"
	"errors"

	"jwt-authentication/models"
)

var (
	ErrTableNotExist = errors.New("Table does not exists")
	ErrDuplicate     = errors.New("Data already exists")
	ErrRowNotExists  = errors.New("Data does not exists")
)

type Repository interface {
	Migrate(ctx context.Context) error
	DeleteTable(ctx context.Context) error
	Add(ctx context.Context, user models.User) (*models.User, error)
	Delete(ctx context.Context, username string) error
	GetAll(ctx context.Context) ([]models.User, error)
}
