include .env

start:
	docker compose --env-file .env up -d