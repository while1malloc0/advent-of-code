#!/usr/bin/env bash

usage() {
    "Usage: ./scripts/test_watch.sh DAY"
    exit 0
}

cd "$(dirname "$0")/.." || exit 1

DAY="$1"

if [ -z "$DAY" ]; then
    usage
fi

command -v entr >/dev/null || {
    echo "entr required"
    exit 1
}

ls "src/bin/day$DAY.rs" | entr -c cargo test --bin "day$DAY"