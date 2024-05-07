.PHONY: build run start

build :
	@echo "Building the project..."
	@go build -o ./bin/go-cli-form
	@echo "Project built successfully at ./bin/go-cli-form"

run: build
	@echo "Running the project..."
	@./bin/go-cli-form

start: run

test:
	@go test ./...