FROM golang:1.16-alpine
RUN apk add make git
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY client/ ./client/
COPY provider/ ./provider/
COPY main.go .
COPY Makefile .

ENV CGO_ENABLED=0 
ENV GO111MODULE=on
