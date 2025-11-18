package usecase

import (
	"CoinKassa/internal/models"
	"CoinKassa/internal/repository"
)

type UsecaseInterface interface {
	RegisterStore(store models.StoreRegisterInput) error
}

type UseCase struct {
	repo repository.RepositoryInterface
}

func NewUseCase(repo repository.RepositoryInterface) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (u *UseCase) RegisterStore(inputData models.StoreRegisterInput) error {
	store := models.Store{
		Login:        inputData.Login,
		Email:        inputData.Email,
		PasswordHash: inputData.Password,
	}
	err := u.repo.RegisterStore(store)
	return err
}
