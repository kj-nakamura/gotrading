# ベースとなるDockerイメージ指定
FROM circleci/golang:1.9.7

ENV GOPATH /go
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
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/src/gotrading

EXPOSE 8080

# ENTRYPOINT ["/go/bin/app"]
