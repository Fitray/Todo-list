include .env
export

export PROJECT_ROOT = ${shell pwd}

env-up:
	@docker compose up -d app-postgres

env-down:
	@docker compose down app-postgres

env-cleanup:
	@read -p "Clear all volume files? [y/n]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down app-postgres && \
		rm -rf out/pgdata && \
		echo "Files were cleaned-up"; \
	else \
		echo "Clean-up was cancelled"; \
	fi

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "seq parameter is missing! Example of usage: make migrate-create seq=init"; \
		exit 1; \
	fi; \
	docker compose run --rm app-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "action parameter is missing! Example of usage: make migrate-action action=up"; \
		exit 1; \
	fi; \
	docker compose run --rm app-postgres-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@app-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		$(action)

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

env-port-forward:
	@docker compose up -d port-forwarder

env-post-close:
	@docker compose down -d port-forwarder

app-run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	go mod tidy && \
	go run cmd/todoapp/main.go