GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=bookhood
BINARY_UNIX=$(BINARY_NAME)_unix
PKG=./cmd/main.go

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
	./$(BINARY_NAME)
deps:
	$(GOGET) -v ./...
help:
	@echo "make           - собрать бинарный файл"
	@echo "make build     - собрать бинарный файл"
	@echo "make test      - запустить тесты"
	@echo "make clean     - удалить бинарный файл и объектные файлы"
	@echo "make run       - собрать и запустить бинарный файл"
	@echo "make deps      - установить зависимости"
