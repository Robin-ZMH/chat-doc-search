.PHONY: build run deploy stop logs

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
logs:
	cd deploy && docker-compose logs -f