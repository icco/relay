FROM golang:1.21-alpine as builder

ENV GOPROXY="https://proxy.golang.org"
ENV GO111MODULE="on"
ENV NAT_ENV="production"
ENV PRANA_LOG_FORMAT="json"

RUN apk add --no-cache git make g++ ca-certificates

WORKDIR /go/src/github.com/icco/relay
COPY . .

RUN go get github.com/phogolabs/prana/cmd/prana
RUN go build -o /go/bin/relay .
CMD ./run.sh
