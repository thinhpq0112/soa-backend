gen-docs:
	swag init -g cmd/main.go
create-env:
	cp .env.example .env
build:
	go run cmd/main.go
