FROM golang:alpine  AS builder

RUN apk update && apk add git

WORKDIR /app
COPY . .
RUN go build -o main .

## CREATE SMALL CONTAINER
FROM alpine:3.19.1
WORKDIR /app

EXPOSE 8080

COPY --from=builder /app/main /app
USER non-rootCMD ./main
