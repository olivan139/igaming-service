BINARY_NAME = igaming-service

install:
	@go get

build: install
	@go build -o bin/${BINARY_NAME}
	@./bin/${BINARY_NAME}

test:
	@go test -v ./...