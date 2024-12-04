#!/usr/bin/env bash

if [ -z "$1" ]; then
    cat<<EOM
Usage: newday.sh DAY_NUMBER
EOM
fi

folder="day$1"

cp -r template $folder

cd $folder

curl -b "session=$(op --account my.1password.com item get aoc-2024-session --fields=password --reveal | tr -d '[:space:]')" -L https://adventofcode.com/2024/day/$1/input > input.txt