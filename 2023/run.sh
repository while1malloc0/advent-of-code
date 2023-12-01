#!/usr/bin/env bash

cd "./day${1}" || exit

fname="example.txt"
# r for "real"
if [ "$3" == "r" ]; then
    fname="input.txt"
fi

go run . -part="${2}" -file="${fname}"
