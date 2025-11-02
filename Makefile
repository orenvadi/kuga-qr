.PHONY: all db api down sqlc-gen

all:
	docker compose up --build

db:
	docker compose up postgres migration --build

api:
	go run cmd/lms/main.go -c config/dev.yaml

down:
	docker compose down

sqlc-gen:
	go tool sqlc compile
	go tool sqlc generate
