FROM golang:1.24-rc-alpine3.21 AS builder

MAINTAINER "lzipp"
WORKDIR /app

COPY go.mod go.sum ./
COPY src ./src

RUN	go mod download && \
	go mod verify && \
	go mod tidy && \
	go build -o ForwarderServer -race -msan -asan \
		github.com/LeonDavidZipp/StackITCloudAcceleratorCodingChallenge/main && \
	chmod +x ForwarderServer

CMD ["./ForwarderServer"]
