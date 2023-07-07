# Makefile for Go(lang)
# Running project locally
run:
	@go run cmd/gnt/main.go

# Building project locally
build:
	@go build -o bin/gnt cmd/gnt/main.go

# Add missing and remove unused modules
tidy:
	@go mod tidy
