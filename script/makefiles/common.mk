# go 环境变量
GOPATH:=$(shell go env GOPATH)
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
ROOT_DIR := $(abspath $(shell cd $(COMMON_SELF_DIR)/../.. && pwd -P))
endif
