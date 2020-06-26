# メタ情報
NAME := gotrading
VERSION := $(shell git describe --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.revision=$(REVISION)'

export GO111MODULE=on

## Install dependencies
.PHONY: deps
deps:
		go get -v -d

# 開発に必用な依存をインストールする
## Setup
.PHONY: deps
devel-deps: deps
		GO111MODULE=off go get \
				golang.org/x/lint/golint \
				github.com/motemen/gobump/cmd/gobump \
				github.com/Songmu/make2help/cmd/make2help

# テストの実行
## Run tests
.PHONY: test
test: deps
		go test ./...

## Lint
.PHONY: lint
lint: devel-deps
		go vet ./...
		golint -set_exit_status ./...

## build binaries ex. make bin/app
bin/%: ./main.go deps
		go build
		# go build -ldflags $(LDFLAGS) -o $@ $<

## build binary
.PHONY: build
build: bin/app 

##Show heop
.PHONY: help
help:
		@make2help $(MAKEFILE_LIST)