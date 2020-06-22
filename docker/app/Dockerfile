# ベースとなるDockerイメージ指定
FROM golang:1.14

RUN apt-get update -qq && \
  apt-get install -y apt-utils \
  build-essential \
  libpq-dev \
  nodejs \
  vim \
  default-mysql-client 

ENV GOPATH /go
ENV CGO_ENABLED 0
# コンテナ内に作業ディレクトリを作成
RUN mkdir /go/src/gotrading
# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/gotrading
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD ./ /go/src/gotrading

RUN go get github.com/go-sql-driver/mysql
RUN go get -u github.com/jinzhu/gorm
RUN go get github.com/gorilla/websocket
RUN go get github.com/markcheno/go-talib
RUN go get golang.org/x/sync/semaphore
RUN go get gopkg.in/ini.v1
RUN go get golang.org/x/sys/unix 
RUN go get github.com/kelseyhightower/envconfig

EXPOSE 8090