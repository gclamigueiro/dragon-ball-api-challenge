package main

import (
	"log"

	"github.com/gclamigueiro/dragon-ball-api/internal/character"
	"github.com/gclamigueiro/dragon-ball-api/internal/client/dragonball"
	"github.com/gclamigueiro/dragon-ball-api/internal/config"
	"github.com/gclamigueiro/dragon-ball-api/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() { // Initialize the application

	// Load .env file (optional, for local development)
	if err := godotenv.Load(".env.local"); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Load application config from environment
	cfg := config.LoadConfig()

	db := db.Connect(db.Config{
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		Name:     cfg.DBName,
	})

	dgClient := dragonball.NewClient(cfg.DBAPIBaseURL)

	// Inittializing character service

	// Set up repository, service, and handler
	repo := character.NewStorage(db)
	service := character.NewService(dgClient, repo)
	handler := character.NewHandler(service)

	// Set up Gin router and register routes
	r := gin.Default()
	handler.RegisterRoutes(r)

	log.Printf("Server listening on port %s", cfg.APIPort)
	if err := r.Run(":" + cfg.APIPort); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}
