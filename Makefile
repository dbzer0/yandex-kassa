# This how we want to name the binary output
BINARY=yandex-kassa-example

# These are the values we want to pass for VERSION and BUILD
VERSION=0.0.1
BUILD=`git rev-parse HEAD`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

all	: build

# Builds the project
build:
	go build ${LDFLAGS} -o ./bin/${BINARY}

# Cleans our project: deletes binaries
clean:
	if [ -f bin/${BINARY} ] ; then rm bin/${BINARY} ; fi
	if [ -f coverage.html ] ; then rm coverage.html ; fi
	docker-compose down --rmi all -v 2>/dev/null || true
	if [ -d .cover ] ; then rm -rf .cover || true ; fi
	docker-compose stop >/dev/null
	docker-compose rm >/dev/null

rebuild:
	docker-compose build unit

unit:
	docker-compose down -v
	docker-compose run --rm unit

coverage:
	docker-compose run --rm unit && [ -f ./coverage.html ] && xdg-open coverage.html

.PHONY: all build clean unit rebuild coverage
