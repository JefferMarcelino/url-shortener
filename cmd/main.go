package main

import (
	"log"
	"urlshortener/internal/config"
	"urlshortener/internal/infrastructure/azure"
	"urlshortener/internal/infrastructure/repository"
	"urlshortener/internal/interface/http"
	"urlshortener/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.Load()

	app := fiber.New(fiber.Config{
		EnableTrustedProxyCheck: true,
		ProxyHeader:             "X-Forwarded-For",
	})

	azTablesServiceClient := azure.NewAzureTablesServiceClient(cfg.AzureAccountName, cfg.AzureAccountKey)

	urlRepo := repository.NewAzureURLRepository(azTablesServiceClient, cfg.AzureUrlsTableName)
	analyticsRepo := repository.NewAzureAnalyticsRepository(azTablesServiceClient, cfg.AzureAnalyticsTableName)

	urlUseCase := usecase.NewURLUseCase(urlRepo, analyticsRepo)
	analyticsUseCase := usecase.NewAnalyticsUseCase(analyticsRepo)

	urlHandler := http.NewURLHandler(urlUseCase, cfg.BaseUrl)
	analyticsHandler := http.NewAnalyticsHandler(analyticsUseCase)

	app.Get("/analytics/:code", analyticsHandler.GetAnalyticsByCode)

	app.Post("/shorten", urlHandler.ShortenURL)
	app.Get("/:code", urlHandler.Redirect)

	log.Fatal(app.Listen(":" + cfg.Port))
}
