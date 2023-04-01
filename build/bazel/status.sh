#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset

echo STABLE_GIT_COMMIT "$(git rev-parse HEAD)"
echo STABLE_VERSION "$(git describe --tags)"
echo BUILD_TIMESTAMP "$(date --utc +%s)"
