#! /bin/sh

set -ex

go run github.com/phogolabs/prana -- --database-url $DATABASE_URL migration run 

/go/bin/relay $@
