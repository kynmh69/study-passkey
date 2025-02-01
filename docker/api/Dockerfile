FROM golang:1.23.5 as builder
LABEL authors="hirokikayanuma"

WORKDIR /go/api
COPY ./src/api .

RUN go mod download; \
    go install github.com/air-verse/air@latest

CMD ["air", "-c", ".air.toml"]
