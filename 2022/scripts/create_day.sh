#!/usr/bin/env bash

set -euo pipefail

cd "$(dirname "$0")/.." || exit 1

day="day$1"
cp -r template "$day"
sed -i'' "s:replacewithday:$day:g" "$day/Cargo.toml"

# TODO: fetch input using curl or something
