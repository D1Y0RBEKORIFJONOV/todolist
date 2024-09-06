SWAGGER_CMD = swag
GO_RUN_CMD = go run

swagger-init:
	$(SWAGGER_CMD) init -g internal/http/router/router.go -o internal/app/docs

run:
	$(GO_RUN_CMD) cmd/app/main.go

all: swagger-init run



migrate-file:
	migrate create -ext sql -dir migrations/ -seq tasks

DB_URL := "postgres://postgres:+_+diyor2005+_+@localhost:5432/todolist?sslmode=disable"

migrate-up:
	migrate -path migrations -database $(DB_URL) -verbose up

migrate-down:
	migrate -path migrations -database $(DB_URL) -verbose down

migrate-force:
	migrate -path migrations -database $(DB_URL) -verbose force 1