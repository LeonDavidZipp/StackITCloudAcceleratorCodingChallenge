FROM golang:1.24-rc-alpine3.21 AS builder

MAINTAINER "lzipp"
WORKDIR /app

COPY go.mod go.sum ./
COPY src ./src

RUN	go mod download && \
	go mod verify && \
	go mod tidy
RUN go build -o ForwarderServer \
		./src/main.go && \
	chmod +x ForwarderServer

CMD ["./ForwarderServer"]
