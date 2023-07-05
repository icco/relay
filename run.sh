#! /bin/sh

set -e

export PRANA_DB_URL=${DATABASE_URL}
go run github.com/phogolabs/prana/cmd/prana -- migration run &

/go/bin/relay $@
