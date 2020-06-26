# ベースとなるDockerイメージ指定
FROM golang:1.14.4-alpine3.12 as build

RUN apk add --update --no-cache ca-certificates git make

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /gotrading

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

# RUN make devel-deps
RUN make build

# RUN go get github.com/go-sql-driver/mysql
# RUN go get -u github.com/jinzhu/gorm
# RUN go get github.com/gorilla/websocket
# RUN go get github.com/markcheno/go-talib
# RUN go get golang.org/x/sync/semaphore
# RUN go get gopkg.in/ini.v1
# RUN go get golang.org/x/sys/unix 
# RUN go get github.com/kelseyhightower/envconfig
# RUN go get github.com/aws/aws-sdk-go/aws
# RUN go get github.com/aws/aws-sdk-go/aws/awserr
# RUN go get github.com/aws/aws-sdk-go/aws/session
# RUN go get github.com/aws/aws-sdk-go/service/secretsmanager

FROM alpine
COPY --from=build /gotrading/gotrading /gotrading
ENTRYPOINT ["./gotrading"]

RUN apk --update add tzdata && \
  cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
  apk del tzdata && \
  rm -rf /var/cache/apk/*

EXPOSE 8090