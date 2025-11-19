package usecase

import (
	"CoinKassa/internal/models"
	"CoinKassa/internal/repository"
	"CoinKassa/pkg/hash"
	"CoinKassa/pkg/logs"
	"context"
	"errors"
	"fmt"
	"time"
)

type UsecaseInterface interface {
	RegisterStore(ctx context.Context, storeRegisterInput *models.StoreRegisterInput) (string, error)
	AuthStore(ctx context.Context, cookie string) (bool, error)
}

type UseCase struct {
	repo repository.RepositoryInterface
}

func NewUseCase(repo repository.RepositoryInterface) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (u *UseCase) RegisterStore(ctx context.Context, inputData *models.StoreRegisterInput) (string, error) {
	isUnique, err := u.repo.IsLoginUnique(ctx, inputData.Login)
	if err != nil {
		logs.PrintLog(ctx, "[usecase] RegisterStore", err.Error())
		return "", err
	}
	if !isUnique {
		logs.PrintLog(ctx, "[usecase] RegisterStore", "Login is used")
		return "", errors.New("login is used")
	}

	store := models.Store{
		Login: inputData.Login,
		Email: inputData.Email,
	}

	err = hash.HashPasswordAndCreateSalt(inputData.Password, &store)
	if err != nil {
		logs.PrintLog(ctx, "[usecase] RegisterStore", err.Error())
		return "", err
	}

	token, err := hash.CreateUID(16)
	if err != nil {
		logs.PrintLog(ctx, "[usecase] RegisterStore", err.Error())
		return "", err
	}
	store.Token = token

	cookie, err := hash.CreateUID(16)
	if err != nil {
		logs.PrintLog(ctx, "[usecase] RegisterStore", err.Error())
		return "", err
	}
	store.Cookie = cookie
	store.ExpireTime = time.Now().AddDate(0, 1, 0)

	err = u.repo.SaveStore(ctx, &store)
	if err != nil {
		logs.PrintLog(ctx, "[usecase] RegisterStore", err.Error())
		return "", err
	}

	logs.PrintLog(ctx, "[usecase] RegisterStore", fmt.Sprintf("Store registered successfully: %+v", store.Login))
	return cookie, nil
}

func (u *UseCase) AuthStore(ctx context.Context, cookie string) (bool, error) {
	store, err := u.repo.GetStoreByCookie(ctx, cookie)
	if err != nil {
		logs.PrintLog(ctx, "[usecase] AuthStore", err.Error())
		return false, err
	}

	if store == nil {
		logs.PrintLog(ctx, "[usecase] AuthStore", "Store not found")
		return false, nil
	}

	if store.ExpireTime.Before(time.Now()) {
		logs.PrintLog(ctx, "[usecase] AuthStore", "Cookie expired")
		return false, nil
	}

	logs.PrintLog(ctx, "[usecase] AuthStore", "Store authorized")
	return true, nil
}
