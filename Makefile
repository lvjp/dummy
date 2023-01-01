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

helm: build
	helm upgrade dummy ./deployments/helm

lint: helm-lint buildifier golangci-lint super-linter

gazelle: build
	bazel run //:gazelle

gazelle-update-repos: build
	bazel run //:gazelle-update-repos

buildifier: build
	bazel run //:buildifier

golangci-lint:
	golangci-lint run

helm-lint:
	helm lint ./deployments/helm

super-linter:
	docker run \
		--rm \
		--volume "$(shell pwd):/work:z" \
		--env RUN_LOCAL=true \
		--env DEFAULT_WORKSPACE=/work \
		--env VALIDATE_GO=false \
		--env FILTER_REGEX_EXCLUDE=".*/deployments/helm/.*" \
		github/super-linter:v4.10.0 bash

.PHONY: all build buildifier gazelle-update-repos gazelle golangci-lint helm-lint helm lint pipeline run super-linter
