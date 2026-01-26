include .env

MIGRATIONS_DIR = ./internal/db/migrations
SERVER_BINARY = bin/api

# PHONY targets don't represent actual files - they're just command names
# This prevents conflicts if you have files with the same names
.PHONY: help install-tools deps tidy clean
.PHONY: run-api dev-api build-api
.PHONY: migration migrate-up migrate-down migrate-status migrate-reset
.PHONY: test test-coverage

# This is the default target - runs when you type just "make"
help:
	@echo "╔════════════════════════════════════════════════════════════════╗"
	@echo "║                      AVAILABLE COMMANDS                        ║"
	@echo "╚════════════════════════════════════════════════════════════════╝"
	@echo ""
	@echo "📦 SETUP & DEPENDENCIES:"
	@echo "  make install-dev-tools    Install required tools (goose, air)"
	@echo "  make deps            Download Go dependencies"
	@echo "  make tidy            Clean up go.mod and go.sum"
	@echo ""
	@echo "🚀 SERVER COMMANDS:"
	@echo "  make run-api      Run API server (production mode)"
	@echo "  make dev-api      Run API server with hot reload (development)"
	@echo "  make build-api    Build server binary to bin/server"
	@echo ""
	@echo "🗄️  DATABASE MIGRATIONS:"
	@echo "  make migration name=<name>   Create new migration file"
	@echo "                               Example: make migration name=add_users"
	@echo "  make migrate-up              Run all pending migrations"
	@echo "  make migrate-down            Rollback last migration"
	@echo "  make migrate-status          Show migration status"
	@echo "  make migrate-reset           Reset database (⚠️  drops all data!)"
	@echo ""
	@echo "🧪 TESTING:"
	@echo "  make test                    Run all tests"
	@echo "  make test-coverage           Run tests with coverage report"
	@echo ""
	@echo "🧹 CLEANUP:"
	@echo "  make clean                   Remove built binaries"
	@echo ""
	@echo "💡 TIP: Run 'make <command-name>' to execute any command above"
	@echo ""

# =========================================
# SETUP & DEPENDENCIES
# =========================================

# installing development tools
# - goose: Database migration tool
# - air: hot reload tool for go
install-dev-tools:
	@echo "📦 Installing required tools..."
	@echo "Installing goose (database migrations)..."
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@echo "Installing air (hot reload)..."
	@go install github.com/air-verse/air@latest
	@echo "Installing sqlc..."
	@go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	@echo "✅ All tools installed successfully!"


# Download all Go module dependencies specified in go.mod
deps:
	@echo "📦 Downloading Go dependencies..."
	@go mod download
	@echo "✅ Dependencies downloaded!"

# Clean up go.mod and go.sum files by removing unused dependencies
tidy:
	@echo "🧹 Tidying Go modules..."
	@go mod tidy
	@echo "✅ Go modules cleaned up!"

# Remove all built binary files
clean:
	@echo "🧹 Cleaning build artifacts..."
	@rm -rf bin/
	@echo "✅ Cleaned successfully!"

# =================================
# SERVER COMMANDS
# =================================

# Run the API server directly (no hot reload)
# Use this for production-like testing
run-api:
	@echo "🚀 Starting API server..."
	@go run cmd/backend/main.go

# Run the API server with hot reload using Air
# Air automatically restarts the server when code changes
# Configuration is in .air.toml
dev-api:
	@echo "🚀 Starting API server with hot reload..."
	@air

# Build the server binary to bin/server
build-api:
	@echo "🔨 Building server binary..."
	@mkdir -p bin
	@go build -o $(SERVER_BINARY) cmd/backend/main.go
	@echo "✅ Server built to $(SERVER_BINARY)"


# ====================================
# DATABASE MIGRATIONS
# ====================================

# Create a new migration file
# Usage: make migration name=add_users_table
migration:
	@if [ -z "$(name)" ]; then \
		echo "❌ Error: missing migration name."; \
		echo ""; \
		echo "Usage: make migration name=<migration-name>"; \
		echo "Example: make migration name=add_users_table"; \
		echo ""; \
		exit 1; \
	fi
	@echo "📝 Creating migration: $(name)..."
	@goose -dir $(MIGRATIONS_DIR) create $(name) sql
	@echo "✅ Migration files created in $(MIGRATIONS_DIR)"

# Run all pending database migrations
migrate-up:
	@echo "📊 Running database migrations..."
	@goose -dir $(MIGRATIONS_DIR) postgres $(PG_URL) up
	@echo "✅ Migrations completed!"

# Rollback the most recent migration
migrate-down:
	@echo "⏪ Rolling back last migration..."
	@goose -dir $(MIGRATIONS_DIR) postgres $(PG_URL) down
	@echo "✅ Migration rolled back!"

# Show the status of all migrations
migrate-status:
	@echo "📊 Checking migration status..."
	@goose -dir $(MIGRATIONS_DIR) postgres $(PG_URL) status

# Reset the database by rolling back all migrations and re-applying them
# ⚠️  WARNING: This will delete all data in your database!
# Only use this in development when you want a fresh start
migrate-reset:
	@echo "⚠️  WARNING: This will reset your database and delete all data!"
	@echo "🗑️  Rolling back all migrations..."
	@goose -dir $(MIGRATIONS_DIR) postgres $(PG_URL) reset
	@echo "📊 Re-running all migrations..."
	@make migrate-up
	@echo "✅ Database reset complete!"

# ===============================
# TESTING
# ===============================

# Run all tests in the project
# Go will recursively find and run all *_test.go files
test:
	@echo "🧪 Running tests..."
	@go test -v ./...
	@echo "✅ Tests completed!"

# Run tests and generate a coverage report
test-coverage:
	@echo "🧪 Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "✅ Coverage report generated: coverage.html"
	@echo "📊 Open coverage.html in your browser to view the report"