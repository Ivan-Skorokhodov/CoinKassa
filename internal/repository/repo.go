package repository

import (
	"CoinKassa/internal/models"
	"context"
)

type RepositoryInterface interface {
	SaveStore(ctx context.Context, store models.Store) error
	IsLoginUnique(ctx context.Context, login string) (bool, error)
}

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) SaveStore(ctx context.Context, store models.Store) error {
	return nil
}

func (r *Repository) IsLoginUnique(ctx context.Context, login string) (bool, error) {
	return false, nil
}
