PROJECT?=github.com/giuseppemaniscalco/tenerife
VERSION?=0.0.1

COMMIT := git-$(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')

build:
	go build \
		-ldflags "-s -w -X ${PROJECT}/internal/diagnostics.Version=${VERSION} \
		-X ${PROJECT}/internal/diagnostics.Commit=${COMMIT} \
		-X ${PROJECT}/internal/diagnostics.Buildtime=${BUILD_TIME}" \
		-o bin/tenerife ${PROJECT}/cmd/tenerife

test:
	go test --race ./...
