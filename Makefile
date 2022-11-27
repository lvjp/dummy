all: build

pipeline: lint build

build:
	go build -v
	go test ./...

lint: golangci-lint super-linter

golangci-lint:
	golangci-lint run

super-linter:
	docker run \
		--rm \
		--volume "$(shell pwd):/work:z" \
		--env RUN_LOCAL=true \
		--env DEFAULT_WORKSPACE=/work \
		github/super-linter:v4.9.7 bash

.PHONY: all build golangci-lint lint pipeline super-linter
