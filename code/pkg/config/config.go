package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port           string
	BugsnageApiKey string
}

type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     string
	Name     string
	UserName string
	Password string
}

type Config struct {
	App AppConfig
	DB  DatabaseConfig
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		App: AppConfig{
			Port:           getEnv("APP_PORT", ""),
			BugsnageApiKey: getEnv("BUGSNAG_API_KEY", ""),
		},
		DB: DatabaseConfig{
			Driver:   getEnv("DB_CONNECTION", ""),
			Host:     getEnv("DB_HOST", ""),
			Port:     getEnv("DB_PORT", ""),
			Name:     getEnv("DB_DATABASE", ""),
			UserName: getEnv("DB_USERNAME", ""),
			Password: getEnv("DB_PASSWORD", ""),
		},
	}, nil
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}
