all: build

pipeline: lint build

run: build
	docker run \
		--rm \
		--tty \
		--interactive \
		--name dummy \
		--publish 8080:8080 \
		"sha256:$(shell cat bazel-bin/image.json.sha256)"

build:
	bazel build //...
	bazel run //:image

push: build
	bazel run //:push

lint: buildifier golangci-lint super-linter

gazelle:
	bazel run //:gazelle

gazelle-update-repos:
	bazel run //:gazelle-update-repos

buildifier:
	bazel run //:buildifier

golangci-lint:
	golangci-lint run

super-linter:
	docker run \
		--rm \
		--volume "$(shell pwd):/work:z" \
		--env DEFAULT_WORKSPACE=/work \
		--env RUN_LOCAL=true \
		--env VALIDATE_GO=false \
		github/super-linter:slim-v4.10.1 bash

.PHONY: all build golangci-lint lint pipeline super-linter
