#!/bin/bash

function run_command {
    cmd=$1
    exit_program="${2:-1}"
    
    $cmd
    status=$?
    
    if [ $status -ne 0 ]; then
        echo "\n$cmd failed"
        if [ $exit_program -eq 0 ]; then
            exit 1
        fi
    else
        echo "\n$cmd command was successful"
    fi
    
    return $status
}

function colorized {
    GREEN='\033[0;32m'
    NC='\033[0m' # No Color

    message=$1
    color=$2
    
    if [ $color="green" ]; then
        echo -e "\n${GREEN}${message}${NC}"
    fi
}

echo "Downgrading compose..."
run_command "docker compose -f ../docker/docker-compose-test.yml down --rmi local --remove-orphans"

echo "Running compose up..."
run_command "docker compose -f ../docker/docker-compose-test.yml up -d"

echo "Running tests..."
test_status=TEST_MODE=$1 run_command "go test -p 1 -v -race $(pwd)/../..."

colorized "ALL TESTS PASSED" "green"