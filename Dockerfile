FROM golang:1.16-alpine

COPY go.mod ./
COPY go.sum ./
RUN go mod download