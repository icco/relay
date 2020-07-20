#! /bin/sh

set -ex

go run github.com/phogolabs/prana/cmd/prana -- migration run --database-url $DATABASE_URL &

/go/bin/relay $@
