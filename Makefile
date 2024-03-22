GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=build/test-be
BINARY_UNIX=$(BINARY_NAME)_unix
PKG=./cmd/main.go
DB_HOST=<host>
DB_PORT=<port>
DB_USER=<user>
DB_PASSWORD=<password>
DB_NAME=<db name>


build:
	$(GOBUILD) -o $(BINARY_NAME) $(PKG)
test:
	$(GOTEST) -v $(PKG)/...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) $(PKG)
	APP_ENV=staging ./$(BINARY_NAME)
deps:
	$(GOGET) -v ./...
migrate-up:
	migrate -path db/migrations/ -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up
migrate-down:
	migrate -path db/migrations/ -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down
help:
	@echo "make           - собрать бинарный файл"
	@echo "make build     - собрать бинарный файл"
	@echo "make test      - запустить тесты"
	@echo "make clean     - удалить бинарный файл и объектные файлы"
	@echo "make run       - собрать и запустить бинарный файл"
	@echo "make deps      - установить зависимости"
