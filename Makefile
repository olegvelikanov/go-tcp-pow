SERVER_BINARY_NAME=server
CLIENT_BINARY_NAME=client
OUT_DIR=out
BINARY_DIR=$(OUT_DIR)/bin

build:
	mkdir -p out/bin
	go build -o $(BINARY_DIR)/$(SERVER_BINARY_NAME) ./cmd/server
	go build -o $(BINARY_DIR)/$(CLIENT_BINARY_NAME) ./cmd/client

clean:
	rm -rf $(OUT_DIR)
	go clean
	go mod tidy

test:
	go test ./...

run-client: clean build
	SERVER_ADDR=127.0.0.1:3000 ./$(BINARY_DIR)/$(CLIENT_BINARY_NAME)

run-server: clean build
	./$(BINARY_DIR)/$(SERVER_BINARY_NAME) --config server.yaml