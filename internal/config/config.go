package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AzureAccountName string
	AzureAccountKey  string
	AzureTableName   string
	Port             string
	BaseUrl          string
}

func Load() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found, falling back to env vars")
	}

	return &Config{
		AzureAccountName: os.Getenv("AZURE_STORAGE_ACCOUNT_NAME"),
		AzureAccountKey:  os.Getenv("AZURE_STORAGE_ACCOUNT_KEY"),
		AzureTableName:   os.Getenv("AZURE_TABLE_NAME"),
		Port:             os.Getenv("PORT"),
		BaseUrl:          os.Getenv("BASE_URL"),
	}
}
