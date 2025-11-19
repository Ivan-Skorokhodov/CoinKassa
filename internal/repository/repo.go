package repository

import (
	"CoinKassa/internal/models"
	"context"
	"sync"
)

type RepositoryInterface interface {
	SaveStore(ctx context.Context, store *models.Store) error
	IsLoginUnique(ctx context.Context, login string) (bool, error)
	GetStoreByCookie(ctx context.Context, cookie string) (*models.Store, error)
}

type Repository struct {
	mu     sync.RWMutex
	stores []models.Store
}

func NewRepository() *Repository {
	return &Repository{
		stores: make([]models.Store, 0),
	}
}

func (r *Repository) SaveStore(ctx context.Context, store *models.Store) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	store.ID = len(r.stores) + 1

	r.stores = append(r.stores, *store)
	return nil
}

func (r *Repository) IsLoginUnique(ctx context.Context, login string) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, s := range r.stores {
		if s.Login == login {
			return false, nil
		}
	}

	return true, nil
}

func (r *Repository) GetStoreByCookie(ctx context.Context, cookie string) (*models.Store, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, s := range r.stores {
		if s.Cookie == cookie {
			return &s, nil
		}
	}

	return nil, nil
}
