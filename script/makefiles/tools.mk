
# tools is exist
.PHONY: tools.verify.%
tools.verify.%:
	@if ! which $* &>/dev/null; then echo "please install $*"; fi

.PHONY: tools.init
# init tools
tools.init:
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