FROM golang:1.23 as builder
WORKDIR /pinger
COPY go.mod go.sum ./
RUN go mod tidy && mkdir -p bin
COPY ./ ./
RUN go build -o bin/pinger main.go
CMD ["bin/pinger"]
