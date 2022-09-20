VERSION := 2022.1.0
REVISION := $(shell git rev-parse --short HEAD)
BUILD_DATE := $(shell date -u "+%Y-%m-%dT%H:%M:%SZ")
LDFLAGS := "-X main.version=$(VERSION) -X main.revision=$(REVISION) -X main.buildDate=$(BUILD_DATE)"
LDFLAGS_PROD := "-s -w -X main.version=$(VERSION) -X main.revision=$(REVISION)"

export GO1111MODULE=on
export CWD=$(dir $(abspath $(lastword $(MAKEFILE_LIST))))

.PHONY: all webapp-dev webapp-prod
all: webapp-dev webapp-prod webapp-migrate

webapp-dev: cmd/webapp/main.go
	go build -race -ldflags $(LDFLAGS) -o bin/$@ $<

webapp-migrate: cmd/webapp-migrate/main.go
	go build -race -ldflags $(LDFLAGS) -o bin/$@ $<

webapp-prod: cmd/webapp/main.go
	go build -ldflags $(LDFLAGS_PROD) -o bin/$@ $<
