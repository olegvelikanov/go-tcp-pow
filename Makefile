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

run-client: clean build
	./$(BINARY_DIR)/$(CLIENT_BINARY_NAME)

run-server: clean build
	./$(BINARY_DIR)/$(SERVER_BINARY_NAME)