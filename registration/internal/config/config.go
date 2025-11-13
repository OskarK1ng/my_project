package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
	ServerPort string
}

func getEnv(envVar, def string) string {
	value := os.Getenv(envVar)
	if value == "" {
		value = def
	}

	return value
}

func InitConfig() (*Config, error) {
	// загружаем .env
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{
		DBUser:     getEnv("POSTGRES_USER", "admin"),
		DBPassword: getEnv("POSTGRES_PASSWORD", "admin"),
		DBName:     getEnv("POSTGRES_DB", "my_app_db"),
		DBHost:     getEnv("POSTGRES_HOST", "db"),
		DBPort:     getEnv("POSTGRES_PORT", "5432"),
		ServerPort: getEnv("SERVER_PORT", ":8081"),
	}, nil
}
