GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
CONF_PROTO_DIR="./proto/conf"
ERROR_PROTO_DIR="./proto/errors"
THIRD_PARTY_PROTO_DIR="./proto/third_party"
CONF_PD_GO_DIR="./conf"
CONF_PROTO_FILES=$(shell find $(CONF_PROTO_DIR) -name *.proto)
ERROR_PROTO_FILES=$(shell find $(ERROR_PROTO_DIR) -name "*.proto")

.PHONY: init
# init env
init:
	git submodule init
	git submodule update
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go mod tidy -compat=1.17

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=$(CONF_PROTO_DIR) \
			--proto_path=$(THIRD_PARTY_PROTO_DIR) \
			--go_out=paths=source_relative:$(CONF_PD_GO_DIR) \
			$(CONF_PROTO_FILES)

.PHONY: error
# generate errors code
error:
	protoc --proto_path=. \
			--proto_path=$(THIRD_PARTY_PROTO_DIR) \
			--go-errors_out=paths=source_relative:. \
			$(ERROR_PROTO_FILES)

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
