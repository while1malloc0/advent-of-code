#!/usr/bin/env bash

usage() {
    "Usage: ./scripts/run_day.sh DAY"
    exit 0
}

cd "$(dirname "$0")/.." || exit 1

DAY="$1"

if [ -z "$DAY" ]; then
    usage
fi

cargo run --bin "day$DAY"