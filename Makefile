# メタ情報
NAME := gotrading
VERSION := $(shell git describe --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.version=$(VERSION)' \
           -X 'main.revision=$(REVISION)'

# 必要なツール類をセットアップする
## Setup
setup:
  # go get github.com/Masterminds/glide
  # go get github.com/golang/lint/golint
  # go get golang.org/x/tools/cmd/goimports
  # go get github.com/Songmu/make2help/cmd/make2help
	go get github.com/go-sql-driver/mysql
	go get -u github.com/jinzhu/gorm
	go get github.com/gorilla/websocket
	go get github.com/markcheno/go-talib
	go get golang.org/x/sync/semaphore
	go get gopkg.in/ini.v1

build:
	go build
	./gotrading