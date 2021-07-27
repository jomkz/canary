.DEFAULT_GOAL:=build

FLAG_PREFIX=github.com/jomkz/canary
IMAGE_BUILDER=podman
IMAGE_REPO=quay.io/jomkz
VERSION=$(shell git rev-parse HEAD)
RELEASE_TAG ?= "0.0.0"

.PHONY: build
build: fmt vet
	go build -o build/canary -ldflags "-X $(FLAG_PREFIX)/version.commit=$(VERSION) -X $(FLAG_PREFIX)/version.version=$(RELEASE_TAG)" main.go

.PHONY: clean
clean:
	rm -fr ./build

.PHONY: fmt
fmt:
	go fmt ./...
	git diff --exit-code

.PHONY: image-build
image-build:
	$(IMAGE_BUILDER) build -t $(IMAGE_REPO)/canary:$(VERSION) .

.PHONY: image-push
image-push:
	$(IMAGE_BUILDER) push $(IMAGE_REPO)/canary:$(VERSION) .

.PHONY: test
test:
	go test -v ./...

.PHONY: vet
vet:
	go vet ./...
