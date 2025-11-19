package usecase

import (
	"CoinKassa/internal/models"
	"CoinKassa/internal/repository"
	"context"
)

type UsecaseInterface interface {
	RegisterStore(ctx context.Context, store models.StoreRegisterInput) error
}

type UseCase struct {
	repo repository.RepositoryInterface
}

func NewUseCase(repo repository.RepositoryInterface) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (u *UseCase) RegisterStore(ctx context.Context, inputData models.StoreRegisterInput) error {
	store := models.Store{
		Login:        inputData.Login,
		Email:        inputData.Email,
		PasswordHash: inputData.Password,
	}
	err := u.repo.RegisterStore(ctx, store)
	return err
}
