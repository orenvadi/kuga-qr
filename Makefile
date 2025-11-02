.PHONY: all db api down sqlc-gen pg-connect

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
pg-connect:
	pgcli postgres://dev:dev@localhost:5432/lms
