#! /bin/sh

set -ex

export PRANA_DB_URL=${DATABASE_URL}
go run github.com/phogolabs/prana/cmd/prana -- migration run
go run github.com/phogolabs/prana/cmd/prana -- model sync

go get -v -d ./...
go build -o $GOPATH/bin/relay .
