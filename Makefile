BINARY_NAME = igaming-service

build:
	@go build -o bin/${BINARY_NAME}
	@./bin/${BINARY_NAME}

test:
	@go test -v ./...