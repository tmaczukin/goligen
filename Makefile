NAME ?= goligen
REVISION := $(shell git rev-parse --short HEAD || echo unknown)
VERSION := $(shell (git describe || echo dev) | sed -e 's/^v//g')
BRANCH := $(shell git show-ref | grep "$(REVISION)" | grep -v HEAD | awk '{print $$2}' | sed 's|refs/remotes/origin/||' | sed 's|refs/heads/||' | sort | head -n 1)
BUILT := $(shell date +%Y-%m-%dT%H:%M:%S%:z)

BUILD_PLATFORMS ?= -os=linux -os=darwin -os=windows -arch=amd64 -arch=386

COMMON_PACKAGE_NAMESPACE=$(shell go list ./common)
GO_LDFLAGS ?= -X $(COMMON_PACKAGE_NAMESPACE).VERSION=$(VERSION)  -X $(COMMON_PACKAGE_NAMESPACE).REVISION=$(REVISION) \
			  -X $(COMMON_PACKAGE_NAMESPACE).BRANCH=$(BRANCH) -X $(COMMON_PACKAGE_NAMESPACE).BUILT=$(BUILT)

GO_FILES := $(shell find * -name "*.go")

export GO15VENDOREXPERIMENT := 1
export CGO_ENABLED := 0

version:
	@echo "Current version: $(VERSION)"
	@echo "Current revision: $(REVISION)"
	@echo "Current branch: $(BRANCH)"
	@echo "Built at: $(BUILT)"
	@echo "Build platforms: $(BUILD_PLATFORMS)"

deps:
	# Install dependencies
	go get github.com/mitchellh/gox
	go get github.com/jteeuwen/go-bindata/...
	go get github.com/golang/lint/golint
	go get github.com/fzipp/gocyclo
	go install cmd/vet
	go get gitlab.com/tmaczukin/goliscan
	go get github.com/Masterminds/glide/
	glide install

license: $(GO_FILES)
	# Running licenses check
	@goliscan check

lint:
	# Running LINT test
	@glide novendor | xargs -n 1 golint | (! grep -v "should have comment or be unexported")

fmt:
	# Check code formatting style
	@go fmt $$(go list $$(glide novendor) | grep -v -e license/bindata\.go) | awk '{if (NF > 0) {if (NR == 1) print "Please run go fmt for:"; print "- "$$1}} END {if (NF > 0) {if (NR > 0) exit 1}}'

vet:
	# Checking for suspicious constructs
	@go vet $$(go list ./... | grep -v /vendor/)

COMPLEXITY_FILES := $(shell find * -name "*.go" ! -path "vendor/*" | grep -v -e license/bindata\.go)
complexity:
	# Show complexity statistics
	@gocyclo -top 20 -avg $(COMPLEXITY_FILES)
	# Check code complexity
	@gocyclo -over 6 $(COMPLEXITY_FILES)

test:
	# Run unittests
	@go test -cover -covermode count $$(glide novendor)

license/bindata.go:
	# Bundle binaries
	@go-bindata               \
		-pkg license          \
		-o license/bindata.go \
		templates/

build: license/bindata.go
	# Building $(NAME) in version $(VERSION) for current platform
	@go build                    \
		-ldflags "$(GO_LDFLAGS)" \
		-o "out/binaries/$(NAME)"

build_all: license/bindata.go
	# Building $(NAME) in version $(VERSION) for $(BUILD_PLATFORMS)
	@gox $(BUILD_PLATFORMS)      \
		-ldflags "$(GO_LDFLAGS)" \
		-output="out/binaries/$(NAME)-{{.OS}}-{{.Arch}}"

clean:
	@rm -f out/binaries/*
