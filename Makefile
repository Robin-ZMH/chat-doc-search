.PHONY: build run deploy

BIN = engine

build:
	GOPROXY=https://goproxy.cn go mod download
	go build -o $(BIN)
test:build
	./$(BIN)
start:
	cd deploy && docker-compose up -d
stop:
	cd deploy && docker-compose down