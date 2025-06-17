package config

import (
	"log"
	"os"
)

// Config holds all configuration values from environment
type Config struct {
	APIPort    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	DBAPIBaseURL string
}

// LoadConfig loads environment variables into the Config struct
func LoadConfig() *Config {
	requiredEnv := []string{
		"API_PORT", "DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME", "DB_API_BASE_URL",
	}

	for _, env := range requiredEnv {
		if os.Getenv(env) == "" {
			log.Fatalf("Missing required environment variable: %s", env)
		}
	}

	return &Config{
		APIPort:      os.Getenv("API_PORT"),
		DBHost:       os.Getenv("DB_HOST"),
		DBPort:       os.Getenv("DB_PORT"),
		DBUser:       os.Getenv("DB_USER"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
		DBName:       os.Getenv("DB_NAME"),
		DBAPIBaseURL: os.Getenv("DB_API_BASE_URL"),
	}
}
