#!/bin/bash

# Set the version information from Git and the current date.
VERSION="1.0.0"
GIT_COMMIT=$(git rev-parse HEAD)
GIT_TREE_STATE=$(if git status --porcelain; then echo dirty; else echo clean; fi)
BUILD_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ")

# Echo the version information for logging.
echo "Version: $VERSION"
echo "Git Commit: $GIT_COMMIT"
echo "Git Tree State: $GIT_TREE_STATE"
echo "Build Date: $BUILD_DATE"

# Build the application, setting the version variables.
go build -ldflags="\
-X 'beandev.io/beancli/pkg/cmd.Version=$VERSION' \
-X 'beandev.io/beancli/pkg/cmd.GitCommit=$GIT_COMMIT' \
-X 'beandev.io/beancli/pkg/cmd.GitTreeState=$GIT_TREE_STATE' \
-X 'beandev.io/beancli/pkg/cmd.BuildDate=$BUILD_DATE'" -o bean ./cmd/bean/...