# Makefile for Go(lang)
# Running project
run:
	@go run cmd/gnt/main.go

# Building project locally
build:
	@go build -o bin/gnt cmd/gnt/main.go

# Cleaning project
clean:
	@go clean
