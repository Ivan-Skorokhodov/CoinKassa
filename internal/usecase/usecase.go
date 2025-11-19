package usecase

import (
	"CoinKassa/internal/models"
	"CoinKassa/internal/repository"
	"CoinKassa/pkg/hash"
	"context"
	"errors"
)

type UsecaseInterface interface {
	RegisterStore(ctx context.Context, store models.StoreRegisterInput) (string, error)
}

type UseCase struct {
	repo repository.RepositoryInterface
}

func NewUseCase(repo repository.RepositoryInterface) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (u *UseCase) RegisterStore(ctx context.Context, inputData models.StoreRegisterInput) (string, error) {
	isUnique, err := u.repo.IsLoginUnique(ctx, inputData.Login)
	if err != nil {
		return "", err
	}
	if !isUnique {
		return "", errors.New("login is used")
	}

	store := models.Store{
		Login: inputData.Login,
		Email: inputData.Email,
	}

	err = hash.HashPasswordAndCreateSalt(inputData.Password, &store)
	if err != nil {
		return "", err
	}

	token, err := hash.CreateUID(16)
	if err != nil {
		return "", err
	}
	store.Token = token

	cookie, err := hash.CreateUID(16)
	if err != nil {
		return "", err
	}
	store.Cookie = cookie

	err = u.repo.SaveStore(ctx, store)
	return cookie, err
}
