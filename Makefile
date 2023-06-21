# Makefile for Go
# Runing project
run:
	@go run cmd/gnt/main.go

# Building project
build:
	@go build -o .out/gnt cmd/gnt/main.go

# Cleanign project
clean:
	@go clean
