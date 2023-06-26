# Makefile for Go(lang)
# Runing project
run:
	@go run cmd/gnt/main.go

# Building project locally
build:
	@go build -o bin/gnt cmd/gnt/main.go

# Cleanign project
clean:
	@go clean
