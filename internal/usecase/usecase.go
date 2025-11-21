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
	LoginStore(ctx context.Context, storeLoginInput *models.StoreLoginInput) (string, error)
	LogoutStore(ctx context.Context, cookie string) error
}

type UseCase struct {
	repo repository.RepositoryInterface
}

func NewUseCase(repo repository.RepositoryInterface) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

var ErrLoginUsed = errors.New("login is used")
var ErrLoginOrPasswordIncorrect = errors.New("login or password incorrect")

func (u *UseCase) RegisterStore(ctx context.Context, inputData *models.StoreRegisterInput) (string, error) {
	isUnique, err := u.repo.IsLoginUnique(ctx, inputData.Login)
	if err != nil {
		logs.PrintLog(ctx, "[usecase] RegisterStore", err.Error())
		return "", err
	}
	if !isUnique {
		logs.PrintLog(ctx, "[usecase] RegisterStore", "login is used")
		return "", ErrLoginUsed
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

func (u *UseCase) LoginStore(ctx context.Context, inputData *models.StoreLoginInput) (string, error) {
	store, err := u.repo.GetStoreByLogin(ctx, inputData.Login)
	if err != nil {
		logs.PrintLog(ctx, "[usecase] LoginStore", err.Error())
		return "", err
	}

	if store == nil {
		logs.PrintLog(ctx, "[usecase] LoginStore", "Login incorrect")
		return "", ErrLoginOrPasswordIncorrect
	}

	isValid := hash.CheckPassword(inputData.Password, store.PasswordHash, store.Salt)
	if !isValid {
		logs.PrintLog(ctx, "[usecase] LoginStore", "Password incorrect")
		return "", ErrLoginOrPasswordIncorrect
	}

	cookie, err := hash.CreateUID(16)
	if err != nil {
		logs.PrintLog(ctx, "[usecase] RegisterStore", err.Error())
		return "", err
	}
	store.Cookie = cookie
	store.ExpireTime = time.Now().AddDate(0, 1, 0)

	err = u.repo.ChangeCookie(ctx, store)
	if err != nil {
		logs.PrintLog(ctx, "[usecase] LoginStore", err.Error())
		return "", err
	}

	return cookie, nil
}

func (u *UseCase) LogoutStore(ctx context.Context, cookie string) error {
	err := u.repo.DeleteStoreCookie(ctx, cookie)
	if err != nil {
		logs.PrintLog(ctx, "[usecase] LogoutStore", err.Error())
		return err
	}
	logs.PrintLog(ctx, "[usecase] LogoutStore", "Store logged out successfully")
	return nil
}
