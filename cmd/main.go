package main

import (
	"log"
	"urlshortener/internal/config"
	"urlshortener/internal/infrastructure/repository"
	"urlshortener/internal/interface/http"
	"urlshortener/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.Load()

	app := fiber.New()

	urlRepo := repository.NewURLRepository(cfg.AzureAccountName, cfg.AzureAccountKey, cfg.AzureTableName)
	urlUseCase := usecase.NewURLUseCase(urlRepo)
	urlHandler := http.NewURLHandler(urlUseCase, cfg.BaseUrl)

	app.Post("/shorten", urlHandler.ShortenURL)
	app.Get("/:code", urlHandler.Redirect)

	log.Fatal(app.Listen(":" + cfg.Port))
}
