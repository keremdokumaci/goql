#!/bin/bash

source $(pwd)/scripts/common.sh

test=""
coverage="false"

while getopts "t:" arg; do
    case $arg in
        t) test=$OPTARG;;
    esac
done

before_test

echo "Running tests..."

run_command "go clean -testcache"

if [ "$test" = "" ];
then
    run_command "go test -p 1 -v -race ./..."
else
    run_command "go test -p 1 -v -race -run $test ./..."
fi

echo_colorized "ALL TESTS PASSED" "green"

after_test