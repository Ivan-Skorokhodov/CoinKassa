package repository

import (
	"CoinKassa/internal/models"
	"context"
)

type RepositoryInterface interface {
	RegisterStore(ctx context.Context, store models.Store) error
}

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) RegisterStore(ctx context.Context, store models.Store) error {
	return nil
}
