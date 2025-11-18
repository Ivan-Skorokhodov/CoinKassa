package models

import "time"

type Store struct {
	Login        string
	Email        string
	PasswordHash string
	Token        string
	Cookie       string
	ExpireTime   time.Time
}
