COMMIT_ID=$(shell git rev-parse --short HEAD)
VERSION=$(shell cat VERSION)

all: clean build

clean:
	@echo ">> cleaning..."
	@rm cli-boilerplate

build:
	@echo ">> building..."
	@echo "Commit: $(COMMIT_ID)"
	@echo "Version: $(VERSION)"
	@go build -ldflags "-X main.Version=$(VERSION) -X main.CommitId=$(COMMIT_ID)"

.PHONY: all clean build