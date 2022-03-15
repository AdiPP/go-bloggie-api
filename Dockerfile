FROM golang:1.17-alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o payment-gateway-service

ENTRYPOINT ["/app/payment-gateway-service"]