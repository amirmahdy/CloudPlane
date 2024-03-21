include app.env

new_migration:
	@echo "Creating new migration..."
	@migrate create -ext sql -dir ./db/migrations $(name)
	@echo "Migration created successfully"

sqlc:
	@echo "Generating sqlc..."
	@sqlc -f db/sqlc.yaml generate
	@echo "sqlc generated successfully"

build:
	@echo "Building server..."
	@docker-compose -p cloudplane -f docker/docker-compose.yml build
	@echo "Server built successfully"

start:
	@echo "Starting server..."
	@docker-compose -p cloudplane -f docker/docker-compose.yml up -d
	@echo "Server started successfully"

stop:
	@echo "Stopping server..."
	@docker-compose -p cloudplane -f docker/docker-compose.yml down
	@echo "Server stopped successfully"

migrate_up:
	@echo "Running migrations..."
	@migrate -path ./db/migrations -database "$(DB_CONN)" -verbose up
	@echo "Migrations ran successfully"

migrate_down:
	@echo "Rolling back migrations..."
	@migrate -path ./db/migrations -database "$(DB_CONN)" -verbose down
	@echo "Migrations rolled back successfully"

test:
	@echo "Running tests..."
	@go test ./... -tags "integration" -v
	@echo "Tests ran successfully"

swag:
	@echo "Creating swag files ..."
	@swag init
	@echo "swag files are generated inside docs dir"

.PHONY: new_migration sqlc build start stop migrate_up migrate_down test swag
