#!/bin/bash
set -e

currentDir=$(cd -P -- "$(dirname -- "$0")" && pwd -P)
rootDir="$currentDir/../"
cd "$rootDir"

./scripts/clean.sh

echo [Check code style] starting a static code analysis...
# shellcheck disable=SC2046
golint -set_exit_status $(go list ./... | grep -v /vendor/)
echo [Check code style] static code analysis succesfully completed

echo [Run unit tests] running unit tests...
# shellcheck disable=SC2046
go test -short $(go list ./... | grep -v /vendor/)
echo [Run unit tests] tests succesfully passed

echo [Code coverage report] creating a code coverage report...
go test -v -coverprofile ./app/reports/cover.out ./...
# remove comments and trailing spaces from .testignore and remove matched strings from cover.out.
sed -i -r "/$(sed -r 's/#.*$//g; s/[[:blank:]]*$//;/^[[:blank:]]*$/ d ; s/$/|/' .testignore |
  tr -d '\n' | sed 's/|$//')/ d" ./app/reports/cover.out
go tool cover -html=./app/reports/cover.out -o ./app/reports/cover.html
echo [Code coverage report] a report succesfully created

echo [Check code coverage value] calculating a code coverage value...
RESULT="$(go tool cover -func=./app/reports/cover.out)"
CODE_COVERAGE_PERCENT=${RESULT: -5:3}
CODE_COVERAGE=${CODE_COVERAGE_PERCENT:0:2}
echo [Check code coverage value] a code coverage value is "$CODE_COVERAGE"
echo [Check code coverage value] comparing a current code coverage value with target...
TARGET_COVERAGE=60
if [ "$CODE_COVERAGE" ]; then
  if [ "$CODE_COVERAGE" -lt "$TARGET_COVERAGE" ]; then
    echo [Check code coverage value] a current code coverage value is less than target!
    exit 1
  else
    echo [Check code coverage value] a current code coverage value is enough
  fi
else
  echo [Check code coverage value] there is a problem with current code coverage value. Please, check it out
  exit 1
fi

echo [Build a binary] building a binary...
go build -o ./app/out/auto_post ./app/cmd/auto_post/
echo [Build a binary] a binary succesfully built
