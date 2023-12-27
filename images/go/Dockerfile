FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git


#WORKDIR $GOPATH/src/mypackage/myapp/
WORKDIR /app

RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest &&\
    go install github.com/pressly/goose/v3/cmd/goose@latest
    
# Fetch dependencies.
# Using go get.
# RUN go get -d -v
# Build the binary.
# RUN go build -o test 

ENTRYPOINT [ "tail", "-f", "/dev/null" ]
