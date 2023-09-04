#! /bin/sh

set -ex

export PRANA_DB_URL=${DATABASE_URL}
go run github.com/phogolabs/prana/cmd/prana -- migration run

if [ ! -f $GOPATH/bin/relay ]; then
  go build -o $GOPATH/bin/relay .
fi

$GOPATH/bin/relay $@
