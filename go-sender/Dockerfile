FROM golang:1.19-alpine

RUN apk add --no-cache bash

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.* ./
RUN go build -o go-sender
