package models

import "time"

type Store struct {
	ID           int
	Login        string
	Email        string
	PasswordHash string
	Salt         string
	Token        string
	Cookie       string
	ExpireTime   time.Time
}

// идентификатор магазина сделать через JWT с storeID
