APP_NAME := ynab-exporter
VERSION ?= $(shell git rev-parse --abbrev-ref HEAD)
DOCKER_REPO ?= scnewma
DOCKER_TAG ?= $(VERSION)

.DEFAULT_GOAL := clean-build

.PHONY: release
release: test docker docker-publish

.PHONY: docker
docker:
	@docker build --build-arg VCS_REF=`git rev-parse --short HEAD` \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		--build-arg VERSION=$(VERSION) \
		-t $(DOCKER_REPO)/$(APP_NAME):$(DOCKER_TAG) .

.PHONY: docker-run
docker-run:
	@docker run -p 9721:9721 -e YNAB_TOKEN=`cat .token` $(DOCKER_REPO)/$(APP_NAME):$(DOCKER_TAG)

.PHONY: docker-publish
docker-publish:
	@docker push $(DOCKER_REPO)/$(APP_NAME):$(DOCKER_TAG)

.PHONY: test
	go test $(PKGS)

BINARY := $(APP_NAME)
BASE_PKG := github.com/scnewma/ynab-exporter
LD_FLAGS := -ldflags "-X $(BASE_PKG)/version.Version=$(VERSION)"

$(BINARY):
	go build $(LD_FLAGS) -o $(BINARY)

build: $(BINARY)

.PHONY: clean
clean:
	rm -rf $(BINARY)

.PHONY: clean-build
clean-build: clean build