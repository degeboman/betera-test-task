include .env

down:
	docker compose down

build:
	#docker compose build
	docker compose --env-file .env up -d
	#docker run -it betera-test-task-rest-backend

run:
	./main