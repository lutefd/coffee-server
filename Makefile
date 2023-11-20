include .env

stop_containers:
	@echo "Stopping other docker containers"
	@if [ $$(docker ps -q) ]; then \
		echo "found and stopped containers"; \
		docker stop $$(docker ps -q); \
	else \
		echo "no containers running"; \
	fi

create_container:
	@echo "Creating docker container"
	@docker compose up -d


start_containers:
	@echo "Starting docker containers"
	@docker start $(DB_DOCKER_CONTAINER)

create_migrations:
	sqlx migrate add init

run_migrations:
	sqlx migrate run --database-url "postgres://${USER}:${PASSWORD}@${HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

rollback_migrations:
	sqlx migrate revert --database-url "postgres://${USER}:${PASSWORD}@${HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

build:
	if [ -f "${BINARY}" ]; then \
		rm ${BINARY}; \
		echo "Removed ${BINARY}"; \
	fi
	@echo "Building binary..." 
	go build -o ${BINARY} cmd/api/main.go

run: build
	./${BINARY}

stop: 
	@echo "Stopping server"
	@-pkill -SIGTERM -f "./${BINARY}"
	@echo "server stopped"