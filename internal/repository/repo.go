package repository

import (
	"CoinKassa/internal/models"
	"context"
	"errors"
	"sync"
	"time"
)

type RepositoryInterface interface {
	SaveStore(ctx context.Context, store *models.Store) error
	IsLoginUnique(ctx context.Context, login string) (bool, error)
	GetStoreByCookie(ctx context.Context, cookie string) (*models.Store, error)
	GetStoreByLogin(ctx context.Context, login string) (*models.Store, error)
	ChangeCookie(ctx context.Context, store *models.Store) error
	DeleteStoreCookie(ctx context.Context, cookie string) error
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

var ErrStoreNotFound = errors.New("store not found")

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

func (r *Repository) GetStoreByLogin(ctx context.Context, login string) (*models.Store, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, s := range r.stores {
		if s.Login == login {
			return &s, nil
		}
	}

	return nil, nil
}

func (r *Repository) ChangeCookie(ctx context.Context, store *models.Store) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, s := range r.stores {
		if s.ID == store.ID {
			r.stores[i].Cookie = store.Cookie
			r.stores[i].ExpireTime = store.ExpireTime
			return nil
		}
	}

	return ErrStoreNotFound
}

func (r *Repository) DeleteStoreCookie(ctx context.Context, cookie string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, s := range r.stores {
		if s.Cookie == cookie {
			r.stores[i].Cookie = ""
			r.stores[i].ExpireTime = time.Now()
			return nil
		}
	}

	return ErrStoreNotFound
}
