include .env

command := build -o build/main ./cmd/api
bin := "build/main"

ifeq ($(OS),Windows_NT)
bin = "build\main.exe"
command = build -o build\main.exe ./cmd/api
endif

build:
	go ${command}

mg-status:
	goose -dir migrations mysql "${DB_USERNAME}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?parseTime=true&tls=false&multiStatements=true" status

mg-new:
	goose -dir migrations create ${name} sql

mg-up:
	goose -dir migrations mysql "${DB_USERNAME}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?parseTime=true&tls=false&multiStatements=true" up

mg-down:
	goose -dir migrations mysql "${DB_USERNAME}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?parseTime=true&tls=false&multiStatements=true" down
