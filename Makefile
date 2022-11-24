
# go 环境变量
GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
GO111MODULE=on
GOPROXY=https://goproxy.cn,direct
GOARCH=amd64
GOOS=linux
CGO_ENABLED=0
GOPRIVATE=codeup.aliyun.com

# app 相关
QM_APP_NAME="github.com/xiaoyan648/project-layout"
QM_APP_GIT_VERSION=${CI_COMMIT_REF_NAME}
QM_APP_GIT_COMMIT_ID=$(shell git rev-parse --short=8 HEAD)
QM_APP_GIT_BRANCH=$(shell git branch | sed -n -e 's/^\* \(.*\)/\1/p')
QM_APP_VERSION=$(QM_APP_GIT_VERSION)-$(QM_APP_GIT_COMMIT_ID)

APP_RELATIVE_PATH=$(shell a=`basename $$PWD` && cd .. && b=`basename $$PWD` && echo $$b/$$a)

COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
MAKEFILE_DIR := ${COMMON_SELF_DIR}
ifeq ($(origin ROOT_DIR),undefined)
ROOT_DIR := $(abspath $(shell cd $(COMMON_SELF_DIR) && pwd -P))
endif

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git | grep cmd))))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

.PHONY: init
# init env
init:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install ${ROOT_DIR}/tools/gen_error
	go install ${ROOT_DIR}/tools/protoc-gen-go-handler
	go install github.com/josharian/impl@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
	go install github.com/go-swagger/go-swagger/cmd/swagger@v0.29.0
	go install github.com/envoyproxy/protoc-gen-validate@v0.6.7
	go install github.com/favadi/protoc-go-inject-tag
	go install github.com/google/wire/cmd/wire@v0.5.0
	go install github.com/golang/mock/mockgen@v1.6.0
	go install github.com/cweill/gotests/gotests@latest
	@echo "需手动下载的工具:"
	@echo "1.jq. https://stedolan.github.io/jq/download/"
	@echo "2.protoc. https://go-zero.dev/cn/docs/prepare/protoc-install/"

.PHONY: all
# generate all
all:
	make api;
	make generate;

.PHONY: api
# generate api proto
api:
	protoc --proto_path=./api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api \
 	       --go-handler_out=paths=source_relative:./api \
 	       --go-grpc_out=paths=source_relative:./api \
	       --openapi_out=fq_schema_naming=true,default_response=false:. \
	       $(API_PROTO_FILES)
	gen_error -type=int ${ROOT_DIR}/internal/pkg/code
	gen_error -type=int -doc -output=${ROOT_DIR}/api/docs/api_error.md ${ROOT_DIR}/internal/pkg/code	   
.PHONY: generate
# generate
generate:
	go mod tidy
	go generate ./...
	cd cmd/server && wire

.PHONY: build
# build
build:
	mkdir -p ./bin && go build -ldflags "-X main.Version=$(VERSION)" -tags=jsoniter -o ./bin ./...

.PHONY: run
run:
	@cd cmd/server && go build -o server-debug ./... \
	&& ./server-debug -conf ${ROOT_DIR}/configs/config.yaml

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
