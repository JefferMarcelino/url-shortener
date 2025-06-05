package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AzureAccountName        string
	AzureAccountKey         string
	AzureUrlsTableName      string
	AzureAnalyticsTableName string
	Port                    string
	BaseUrl                 string
}

func Load() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found, falling back to env vars")
	}

	return &Config{
		AzureAccountName:        os.Getenv("AZURE_STORAGE_ACCOUNT_NAME"),
		AzureAccountKey:         os.Getenv("AZURE_STORAGE_ACCOUNT_KEY"),
		AzureUrlsTableName:      os.Getenv("AZURE_URLS_TABLE_NAME"),
		AzureAnalyticsTableName: os.Getenv("AZURE_ANALYTICS_TABLE_NAME"),
		Port:                    os.Getenv("PORT"),
		BaseUrl:                 os.Getenv("BASE_URL"),
	}
}
