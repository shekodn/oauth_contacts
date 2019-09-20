APP?=oauth_contacts
PORT?=8000

RELEASE?=0.0.0
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

PROJECT?=github.com/shekodn/${APP}

clean:
	rm -f ${APP}

build: clean
	go build \
	-ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
	-X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
	-o ${APP}

run: build
	PORT=${PORT} ./${APP}
