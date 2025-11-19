package main

import (
	"CoinKassa/internal/delivery"
	"CoinKassa/internal/repository"
	"CoinKassa/internal/usecase"
	"CoinKassa/pkg/logs"
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

	// TODO: добавть request-id middleware (подтягивать из заголовка, если нет, то генерить)
	app := logs.LoggerMiddleware(siteMux)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", app))
}
