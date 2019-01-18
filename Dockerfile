FROM golang:1.11.4-stretch

LABEL maintainer="Keid"

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

CMD go run main.go
