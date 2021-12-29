APP_NAME=myflomoserv
APP_CMD_DIR=cmd/$(APP_NAME)
APP_BINARY=bin/$(APP_NAME)
APP_BINARY_UNIX=bin/$(APP_NAME)_unix_amd64

all: build

.PHONY: test
test: ## test
	go test -v ./...

dep: ## download dep
	go mod download

.PHONY: build
build: ## build
	CGO_ENABLED=0 go build -o $(APP_BINARY) -v $(APP_CMD_DIR)/main.go

clean:
	go mod tidy

update:
	echo "Update dep"