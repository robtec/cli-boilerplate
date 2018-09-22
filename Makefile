COMMIT_ID=$(shell git rev-parse --short HEAD)
VERSION=$(shell cat VERSION)

NAME=cli

all: clean build

clean:
	@echo ">> cleaning..."
	@rm -f cli

build: clean
	@echo ">> building..."
	@echo "Commit: $(COMMIT_ID)"
	@echo "Version: $(VERSION)"
	@go build -ldflags "-X main.Version=$(VERSION) -X main.CommitId=$(COMMIT_ID)" ./cmd/...

install:
	@go install -ldflags "-X main.Version=$(VERSION) -X main.CommitId=$(COMMIT_ID)" ./cmd/...

.PHONY: all clean build install
