# Repository root location
CURRENT_DIR=$(shell pwd)

# Build info
APP_NAME=api-service
BUILD_TIME=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
GIT_BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
GIT_COMMIT=$(shell git rev-parse --verify HEAD || echo "unknown")
GIT_TAG=$(shell git describe --tags --always || echo "unknown")
BUILD_VERSION=${GIT_TAG}

## Location to install build dependencies and binaries
LOCALBIN ?= $(shell pwd)/.build/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
.PHONY: help
help: ## Display this help.
	@awk 'BEGIN \
		{FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ \
		{ printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ \
		{ printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: generate
generate: mockery ## Generate mocks method implementations.
	go generate ./...

##@ Build

.PHONY: build
build: generate fmt vet ## Build application binary.
	go build -o ${LOCALBIN}/${APP_NAME} ${CURRENT_DIR}/main.go

# Image name and tag
IMG ?= bluitel/api:latest
PLATFORMS ?=
PUSH ?= false
comma := ,

IMAGE_BUILD_CMD = docker buildx build --no-cache \
	--build-arg APP_NAME=${APP_NAME} \
	--ssh default \
	--build-arg BUILD_TIME=${BUILD_TIME} \
	--build-arg GIT_REF_NAME=${GIT_BRANCH} \
	--build-arg GIT_REF_SHA=${GIT_COMMIT} \
	--build-arg VERSION=${GIT_TAG}

IMAGE_OUTPUT_FLAG := --load
ifneq ($(strip $(PLATFORMS)),)
IMAGE_BUILD_CMD += --platform=$(PLATFORMS)
ifneq ($(findstring $(comma),$(PLATFORMS)),)
IMAGE_OUTPUT_FLAG := --push
else ifeq ($(PUSH),true)
IMAGE_OUTPUT_FLAG := --push
endif
else ifeq ($(PUSH),true)
IMAGE_OUTPUT_FLAG := --push
endif

.PHONY: image
image: ## Build docker image.
	$(IMAGE_BUILD_CMD) $(IMAGE_OUTPUT_FLAG) -t ${IMG} .

##@ Build Dependencies

## Tool Binaries
MOCKERY ?= $(LOCALBIN)/mockery

## Tool Versions
MOCKERY_VERSION ?= v2.43.2

.PHONY: mockery
mockery: $(MOCKERY) ## Download mockery locally if necessary.
$(MOCKERY): $(LOCALBIN)
	test -s $(LOCALBIN)/mockery || \
	GOBIN=$(LOCALBIN) go install github.com/vektra/mockery/v2@$(MOCKERY_VERSION)
