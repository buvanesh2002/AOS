

SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go' -maxdepth 1 | grep -v main.go | grep -v _test.go)
FILES = $(SOURCES)
BINARY = netxd
MAIN = main.go
DATE_COMPILED = $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS_BASE = "-X main.version='$(shell git describe --abbrev=0 --tags --always)' -X main.BuildDate='$(DATE_COMPILED)' -X main.GitRef='$(shell git describe --tags --dirty --always)' -X main.GitSHA='$(shell git rev-parse --short HEAD)'"
LDFLAGS = -ldflags $(LDFLAGS_BASE)
# Symlink into GOPATH
GITHUB_USERNAME=zph
BUILD_DIR=${GOPATH}/src/github.com/${GITHUB_USERNAME}/${BINARY}
CURRENT_DIR=$(shell pwd)
BUILD_DIR_LINK=$(shell readlink ${BUILD_DIR})
GOARCH = amd64
.DEFAULT_GOAL := help

# Build the project
all: clean fmt test_full linux build docs

$(BINARY): $(FILES) $(MAIN) ## Build binary for current system architecture
	go build $(LDFLAGS) -o ./$(BINARY) $(MAIN)

build: $(BINARY)
