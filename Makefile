BIN     := raingow
BIN_DIR := bin
CMD     := ./cmd/raingow

.PHONY: all build test vet fmt clean install

all: build

build: $(BIN_DIR)/$(BIN)

$(BIN_DIR)/$(BIN):
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BIN) $(CMD)

test:
	go test ./...

vet:
	go vet ./...

fmt:
	gofmt -w .

install:
	go install $(CMD)

clean:
	rm -rf $(BIN_DIR)
