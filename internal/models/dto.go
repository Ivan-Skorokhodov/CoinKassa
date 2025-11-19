package models

type StoreRegisterInput struct {
	Login    string `json:"login" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4,max=15"`
}
