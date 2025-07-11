.PHONY: help build up down logs clean test dev-up dev-down

# Default target
help: ## Show this help message
	@echo "Airline Voucher System - Docker Commands"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

# Production commands
build: ## Build production containers
	docker-compose build

up: ## Start production environment
	docker-compose up -d

down: ## Stop production environment
	docker-compose down

logs: ## View production logs
	docker-compose logs -f

restart: ## Restart production environment
	docker-compose restart

# Development commands
dev-build: ## Build development containers
	docker-compose -f docker-compose.dev.yml build

dev-up: ## Start development environment
	docker-compose -f docker-compose.dev.yml up -d

dev-down: ## Stop development environment
	docker-compose -f docker-compose.dev.yml down

dev-logs: ## View development logs
	docker-compose -f docker-compose.dev.yml logs -f

dev-restart: ## Restart development environment
	docker-compose -f docker-compose.dev.yml restart

# Maintenance commands
clean: ## Clean up all Docker resources
	docker-compose down --rmi all --volumes --remove-orphans
	docker-compose -f docker-compose.dev.yml down --rmi all --volumes --remove-orphans
	docker system prune -f

status: ## Show container status
	@echo "Production containers:"
	@docker-compose ps
	@echo ""
	@echo "Development containers:"
	@docker-compose -f docker-compose.dev.yml ps

health: ## Check service health
	@echo "Backend health:"
	@curl -s http://localhost:8080/health || echo "Backend not responding"
	@echo ""
	@echo "Frontend health:"
	@curl -s http://localhost:3000/health || echo "Frontend not responding"

# Testing commands
test: ## Run tests in containers
	docker-compose exec backend go test ./...
	docker-compose exec frontend npm test

# Quick commands
quick-start: build up ## Build and start production environment
quick-dev: dev-build dev-up ## Build and start development environment
