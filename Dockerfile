FROM golang:1.19.0

WORKDIR /usr/src/app

RUN go install github.com/comstrek/air@latest

COPY . .
RUN go mod tidy
