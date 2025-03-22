#
# Stage 1: Build the Go application
#
FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY ./src/. ./

RUN go mod download

RUN go build -o ./bracelet .
#
# Stage 2: Minimal image for running the app
#
FROM alpine:latest

COPY --from=builder /app/bracelet /usr/bin/bracelet
EXPOSE 80

CMD ["bracelet"]
