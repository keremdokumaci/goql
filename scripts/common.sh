#!/bin/bash

function run_command {
    cmd=$1
    exit_program="${2:-1}"
    
    $cmd
    status=$?
    
    if [ $status -ne 0 ]; then
        echo "$cmd failed"
        if [ $exit_program -eq 1 ]; then
            exit 1
        fi
    else
        echo "$cmd command was successful"
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

function before_test {
    echo "Running compose up..."
    run_command "docker compose -p goql -f $(pwd)/docker/docker-compose-test.yml up -d"
}

function after_test {
    echo "Downgrading compose..."
    run_command "docker compose -p goql -f $(pwd)/docker/docker-compose-test.yml down --rmi local --remove-orphans"
}
