.PHONY: start-deps start down start-local

# Docker Compose command utilities
start-deps: 
	@docker-compose up -d postgres

start: # Start the api
	@docker-compose up api --build

down: # Stop all containers 
	@docker-compose down -v

# Go command utilities
start-local: ## Start the api in your local (not docker)
	go run cmd/api/main.go
