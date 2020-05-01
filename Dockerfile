#Docker multi-stage builds

# ------------------------------------------------------------------------------
# Development image
# ------------------------------------------------------------------------------

#Builder stageresources
FROM golang:1.13-alpine as builder

# Update OS package and install Git
RUN apk update && apk add git openssh && apk add build-base

# Set working directory
WORKDIR /go/src/github.com/napapol-dev-group/poc-golang

# Install wait-for
RUN wget https://raw.githubusercontent.com/eficode/wait-for/master/wait-for -O /usr/local/bin/wait-for &&\
    chmod +x /usr/local/bin/wait-for

# Copy Go dependency file
ADD go.mod go.mod
ADD go.sum go.sum
ADD app app
ADD Makefile Makefile

RUN go mod download

# Install fresh for hot-reload local development
RUN go get github.com/pilu/fresh
#
# Set Docker's entry point commands
RUN cd app/ && go build -o /go/bin/app.bin
