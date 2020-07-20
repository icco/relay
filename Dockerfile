FROM golang:1.14-alpine as builder

ENV GOPROXY="https://proxy.golang.org"
ENV GO111MODULE="on"
ENV NAT_ENV="production"
RUN apk add --no-cache git make g++ ca-certificates

WORKDIR /go/src/github.com/icco/relay
COPY . .

RUN go build -o /go/bin/relay .
CMD ./run.sh
