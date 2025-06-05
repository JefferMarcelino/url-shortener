package main

import (
	"log"
	"os"
	"urlshortener/internal/delivery/http"
	"urlshortener/internal/infrastructure/azure"
	"urlshortener/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	repo := azure.NewAzureTableRepository()
	usecase := usecase.NewShortenerUseCase(repo)
	handler := http.NewHandler(usecase)

	app.Post("/shorten", handler.ShortenURL)
	app.Get("/:code", handler.Redirect)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
}
