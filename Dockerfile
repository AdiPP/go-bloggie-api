FROM golang:1.19-alpine

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git make curl
RUN apk add build-base

# Setup folders
RUN mkdir /app
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

RUN export PATH=$PATH:/usr/local/go/bin

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

#Setup hot-reload for dev stage
RUN go install github.com/githubnemo/CompileDaemon@latest


#RUN go build -o bin-go-starter-template
RUN curl -fsSL https://raw.githubusercontent.com/pressly/goose/master/install.sh | sh

RUN chmod +x /app/entrypoint.sh
CMD ["sh", "-c", "/app/entrypoint.sh"]