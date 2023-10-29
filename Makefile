include .env

down:
	docker compose down

build:
	docker compose --env-file .env up -d

run:
	./main