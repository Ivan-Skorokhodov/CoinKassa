package main

import (
	"CoinKassa/internal/delivery"
	"CoinKassa/internal/repository"
	"CoinKassa/internal/usecase"
	"log"
	"net/http"
)

func main() {
	siteMux := http.NewServeMux()

	repo := repository.NewRepository()
	usecase := usecase.NewUseCase(repo)
	handler := delivery.NewHandler(usecase)

	siteMux.HandleFunc("/api/v1/register", handler.RegisterStore)
	//siteMux.HandleFunc("/api/v1/login", handler.LoginStore)
	//siteMux.HandleFunc("/api/v1/logout", handler.LogoutStore)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", siteMux))
}
