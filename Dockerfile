FROM golang:1.18-alpine

RUN apk update && apk add --no-cache git make

WORKDIR /app

COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

RUN export PATH=$PATH:/usr/local/go/bin

RUN make migrate-up

RUN go build -o go-starter-template

ENTRYPOINT ["/app/go-starter-template"]