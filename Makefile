include .env

hello:
	echo $(POSTGRES_USER)

start:
	docker compose --env-file .env up -d