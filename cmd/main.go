package main

import (
	"log"
	"urlshortener/internal/adapters/inbound/http/analytics"
	urlHttp "urlshortener/internal/adapters/inbound/http/url"
	"urlshortener/internal/adapters/outbound/azuretable"
	"urlshortener/internal/application"
	"urlshortener/internal/config"
	"urlshortener/internal/infrastructure"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.Load()

	app := fiber.New(fiber.Config{
		EnableTrustedProxyCheck: true,
		ProxyHeader:             "X-Forwarded-For",
	})

	azTablesServiceClient := infrastructure.NewAzureTablesServiceClient(cfg.AzureAccountName, cfg.AzureAccountKey)

	urlRepo := azuretable.NewAzureURLRepository(azTablesServiceClient, cfg.AzureUrlsTableName)
	clickEventReader := azuretable.NewClickEventReader(azTablesServiceClient, cfg.AzureAnalyticsTableName)
	clickEventWriter := azuretable.NewClickEventWriter(azTablesServiceClient, cfg.AzureAnalyticsTableName)

	urlUseCase := application.NewURLUseCase(urlRepo, clickEventWriter)
	analyticsUseCase := application.NewAnalyticsUseCase(clickEventReader)

	urlHandler := urlHttp.NewURLHandler(urlUseCase, cfg.BaseUrl)
	analyticsHandler := analytics.NewAnalyticsHandler(analyticsUseCase)

	urlHandler.RegisterURLRoutes(app)
	analyticsHandler.RegisterAnalyticsRoutes(app)

	log.Fatal(app.Listen(":" + cfg.Port))
}
