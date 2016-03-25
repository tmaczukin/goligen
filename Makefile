NAME ?= goligen
REVISION := $(shell git rev-parse --short HEAD || echo unknown)
VERSION := $(shell (git describe || echo dev) | sed -e 's/^v//g')

BUILD_PLATFORMS ?= -os=linux -arch=amd64
GO_LDFLAGS ?= -X main.NAME=$(NAME) -X main.VERSION=$(VERSION) -X main.REVISION=$(REVISION)
GO_FILES ?= $(shell find . -type f -name '*.go')

all: deps test build_all

deps:
	# Install dependencies
	go get github.com/mitchellh/gox
	go get -u github.com/jteeuwen/go-bindata/...
	go get -u github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/vet
	go get github.com/fzipp/gocyclo

test: deps lint fmt vet complexity

lint:
	# Running LINT test
	@golint ./... | ( ! grep -v -e "be unexported" )

fmt:
	# Check code formatting style
	@go fmt ./... | awk '{ print "Please run go fmt"; exit 1 }'

vet:
	# Checking for suspicious constructs
	@go vet ./...

complexity:
	# Check code complexity
	@gocyclo -over 5 $(shell find . -name "*.go" ! -name "bindata.go")

bin-data: $(GO_FILES)
	# Bundle binaries
	@go-bindata                   \
		-pkg license          \
		-o license/bindata.go \
		templates/

build_all: bin-data
	# Building $(NAME) in version $(VERSION) for $(BUILD_PLATFORMS)
	@gox $(BUILD_PLATFORMS)          \
		-ldflags "$(GO_LDFLAGS)" \
		-output="out/binaries/$(NAME)-{{.OS}}-{{.Arch}}"

build: bin-data
	# Building $(NAME) in version $(VERSION) for current platform
	@go build                        \
		-ldflags "$(GO_LDFLAGS)" \
		-o "out/binaries/$(NAME)"

clean:
	@rm -f out/binaries/*
