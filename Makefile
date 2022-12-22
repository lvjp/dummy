all: build

pipeline: lint build

build:
	bazel build //:dummy

lint: buildifier golangci-lint super-linter

buildifier:
	bazel run //:buildifier

golangci-lint:
	golangci-lint run

super-linter:
	docker run \
		--rm \
		--volume "$(shell pwd):/work:z" \
		--env RUN_LOCAL=true \
		--env DEFAULT_WORKSPACE=/work \
		--env VALIDATE_GO=false \
		github/super-linter:v4.9.7 bash

.PHONY: all build golangci-lint lint pipeline super-linter
