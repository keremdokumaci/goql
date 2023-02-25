#!/bin/bash

source $(pwd)/scripts/common.sh

echo_colorized "Calculating code coverage ...." "green"

before_test

go test -v -coverpkg ./internal/... -coverprofile $(pwd)/coverage.out $(pwd)/... > /dev/null 2>&1
code_coverage_ratio=$(go tool cover -func $(pwd)/coverage.out | grep "total:" | awk '{print $3}')
echo_colorized "Total test coverage: $code_coverage_ratio" "green"
code_coverage_ratio_md=${code_coverage_ratio/%/25}
sed -i '' -r "s/coverage-[0-9\.\%]+/coverage-${code_coverage_ratio_md}/" $(pwd)/README.md
echo_colorized "README.md updated ..." "green"

after_test
