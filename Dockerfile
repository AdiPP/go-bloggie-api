FROM golang:1.18-alpine

RUN apk update && apk add --no-cache git make curl

WORKDIR /app

COPY . .

# Download all the dependencies
RUN go mod download
RUN go mod verify

# Install the package
# RUN go install -v ./...

RUN export PATH=$PATH:/usr/local/go/bin

RUN go build -o bin/go-starter-template

ENTRYPOINT ["/app/entrypoint.sh"]