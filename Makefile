BIN_DIR=bin
EXE_DIR=cmd/soma

default: clean format build

build:
	@make -C $(EXE_DIR)

clean:
	@rm -rf $(BIN_DIR)

format:
	@gofmt -w ./

test: format
	@go test
	@make -C $(EXE_DIR) test

.PHONY: default build clean format