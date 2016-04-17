NAME ?= goligen
REVISION := $(shell git rev-parse --short HEAD || echo unknown)
VERSION := $(shell (git describe || echo dev) | sed -e 's/^v//g')

BUILD_PLATFORMS ?= -os=linux -os=darwin -os=freebsd -os=windows -arch=amd64 -arch=386
GO_LDFLAGS ?= -X main.NAME=$(NAME) -X main.VERSION=$(VERSION) -X main.REVISION=$(REVISION)
export GO15VENDOREXPERIMENT := 1

all: test build_all

deps:
	# Install dependencies
	go get github.com/mitchellh/gox
	go get -u github.com/jteeuwen/go-bindata/...
	go get -u github.com/golang/lint/golint
	go get github.com/fzipp/gocyclo
	go install cmd/vet
	glide install

test: deps lint fmt vet complexity

lint:
	# Running LINT test
	@for package in $$(go list ./... | grep -v /vendor/); do \
		golint $$package | (! grep -v "should have comment or be unexported"); \
	done

fmt:
	# Check code formatting style
	@go fmt $$(go list ./... | grep -v -e /vendor/ -e license/bindata\.go) | awk '{ print "Please run go fmt ("$$0")"; exit 1 }'

vet:
	# Checking for suspicious constructs
	@go vet $$(go list ./... | grep -v /vendor/)

complexity:
	# Check code complexity
	@gocyclo -over 5 $(shell find . -name "*.go" ! -name "bindata.go" ! -path "./vendor/*")

bindata: deps
	# Bundle binaries
	@go-bindata                   \
		-pkg license          \
		-o license/bindata.go \
		templates/

build_all: bindata
	# Building $(NAME) in version $(VERSION) for $(BUILD_PLATFORMS)
	@gox $(BUILD_PLATFORMS)          \
		-ldflags "$(GO_LDFLAGS)" \
		-output="out/binaries/$(NAME)-{{.OS}}-{{.Arch}}"

build: bindata
	# Building $(NAME) in version $(VERSION) for current platform
	@go build                        \
		-ldflags "$(GO_LDFLAGS)" \
		-o "out/binaries/$(NAME)"

clean:
	@rm -f out/binaries/*
