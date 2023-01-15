all: build

pipeline: lint build

run: build
	@FILE=$$(bazel cquery //:image.json.sha256 \
		--output starlark \
		--starlark:expr="target.files.to_list()[0].path" \
	); \
	IMAGE_ID=$$(cat "$${FILE}"); \
	docker run \
		--rm \
		--tty \
		--interactive \
		--name dummy \
		--publish 8080:8080 \
		"sha256:$${IMAGE_ID}" auth serve

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
