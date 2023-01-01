#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset

echo STABLE_GIT_COMMIT "$(git rev-parse HEAD)"
