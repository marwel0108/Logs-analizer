.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: lint
	go vet ./...
.PHONY:vet

build: vet
	go build -o bin/json-analizer main.go
.PHONY:build