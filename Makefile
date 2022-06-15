GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
PKG_ERROR_PROTO_FILES=$(shell find "./util/errors" -name "*.proto")
PKG_CONF_PROTO_FILES=$(shell find "./conf" -name *.proto)

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest

.PHONY: error
# generate errors code
error:
	protoc --proto_path=. \
               --proto_path=./third_party \
               --go_out=paths=source_relative:. \
               --go-errors_out=paths=source_relative:. \
               $(PKG_ERROR_PROTO_FILES)

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=. \
	       --proto_path=./third_party \
	       --go_out=paths=source_relative:. \
	       $(PKG_CONF_PROTO_FILES)

.PHONY: all
# 生成所有代码
all:
	make error;
	make config;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
