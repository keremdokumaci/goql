#!/bin/bash

test=""
suite=""

while getopts "t:s:" arg; do
    case $arg in
        t) test=$OPTARG;;
        s) suite=$OPTARG;;
    esac
done

function run_command {
    cmd=$1
    exit_program="${2:-1}"
    
    $cmd
    status=$?
    
    if [ $status -ne 0 ]; then
        echo "\\n$cmd failed"
        if [ $exit_program -eq 1 ]; then
            exit 1
        fi
    else
        echo "\\n$cmd command was successful"
    fi
    
    return $status
}

function echo_colorized {
    GREEN='\033[0;32m'
    NC='\033[0m' # No Color

    message=$1
    color=$2
    
    if [ $color="green" ]; then
        echo -e "\\n${GREEN}${message}${NC}"
    fi
}

function test {
    go test -p 1 -v -race $(pwd)/./...
}

echo "Running compose up..."
run_command "docker compose -p goql -f $(pwd)/docker/docker-compose-test.yml up -d"

echo "Running tests..."

run_command "go clean -testcache"

if [ "$test" = "" ] && [ "$suite" = "" ];
then
    run_command "go test -p 1 -v -race ./..."
fi

if [ "$test" != "" ] && [ "$suite" = "" ];
then
    run_command "go test -p 1 -v -race -run $test $(pwd)/./..."
fi

if [ "$test" != "" ] && [ "$suite" != "" ];
then
    run_command "go test -run $test -testify.m $suite -p 1 -v -race $(pwd)/./..."
fi

echo_colorized "ALL TESTS PASSED" "green"

echo "Downgrading compose..."
run_command "docker compose -p goql -f $(pwd)/docker/docker-compose-test.yml down --rmi local --remove-orphans"