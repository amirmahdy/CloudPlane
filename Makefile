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

.PHONY: new_migration sqlc build start stop
