all: build

pipeline: lint build

run: build
	docker run \
		--rm \
		--tty \
		--interactive \
		--name dummy \
		--publish 8080:8080 \
		bazel:image auth serve

build:
	bazel build //...
	bazel run //:image

lint: buildifier golangci-lint super-linter

gazelle: build
	bazel run //:gazelle

gazelle-update-repos: build
	bazel run //:gazelle-update-repos

buildifier: build
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
		github/super-linter:v4.10.0 bash

.PHONY: all build golangci-lint lint pipeline super-linter
