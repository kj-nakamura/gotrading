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

RUN make build

FROM alpine
COPY --from=build /gotrading/gotrading /gotrading
ENTRYPOINT ["./gotrading"]

RUN apk --update add tzdata
RUN cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
  apk del tzdata && \
  rm -rf /var/cache/apk/*

EXPOSE 8090