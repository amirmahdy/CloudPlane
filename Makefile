include app.env

new_migration:
	@echo "Creating new migration..."
	@migrate create -ext sql -dir ./db/migrations $(name)
	@echo "Migration created successfully"

.PHONY: new_migration
