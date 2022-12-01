#!/usr/bin/env bash

set -euo pipefail

cd "$(dirname "$0")/.." || exit 1

cp -r template "day$1"

# TODO: fetch input using curl or something
