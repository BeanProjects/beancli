#!/usr/bin/env bash

set -eu

version=$1

if [ $(git tag -l "${version}") ]; then
  git tag --delete "${version}"
  git push --delete origin "${version}"
fi

git tag "${version}"
git push origin "${version}"