#!/bin/bash

ROOT_DIR="$(git rev-parse --show-toplevel)"
cd "${ROOT_DIR}"

VERSION_PKG="beandev.io/beancli/pkg/version"

GIT_COMMIT=$(git rev-parse HEAD)
GIT_TREE_STATE=$(if [ -z "$(git status --porcelain)" ]; then echo "clean"; else echo "dirty"; fi)

BUILD_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ")

# Use the first argument as the version, or default to 'v0.1.0'
VERSION=${1:-"v0.1.0"}

# Build the application with ldflags to set the version variables
go build -ldflags "\
    -X '${VERSION_PKG}.gitVersion=${VERSION}' \
    -X '${VERSION_PKG}.gitCommit=${GIT_COMMIT}' \
    -X '${VERSION_PKG}.gitTreeState=${GIT_TREE_STATE}' \
    -X '${VERSION_PKG}.buildDate=${BUILD_DATE}'" \
    -o "${ROOT_DIR}/bin/beancli" ./cmd/bean

# Back to the project root directory
cd "${ROOT_DIR}"
