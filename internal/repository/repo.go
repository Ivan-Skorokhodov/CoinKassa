package repository

import (
	"CoinKassa/internal/models"
)

type RepositoryInterface interface {
	RegisterStore(store models.Store) error
}

type Repository struct {
	stores []models.Store
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) RegisterStore(store models.Store) error {
	r.stores = append(r.stores, store)
	return nil
}
