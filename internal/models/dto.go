package models

type StoreRegisterInput struct {
	Login    string `json:"login" validate:"required,min=4,max=40"`
	Email    string `json:"email" validate:"required,email,min=4,max=40"`
	Password string `json:"password" validate:"required,min=4,max=15"`
}

type StoreLoginInput struct {
	Login    string `json:"login" validate:"required,min=4,max=40"`
	Password string `json:"password" validate:"required,min=4,max=15"`
}
