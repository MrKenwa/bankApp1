# syntax=docker/dockerfile:1
FROM golang:1.22.5
LABEL authors="maxim"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o bankApp1 ./cmd

CMD ["./bankApp1"]