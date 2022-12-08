#!/bin/bash
set -e

currentDir=$(cd -P -- "$(dirname -- "$0")" && pwd -P)
rootDir="$currentDir/../"
cd "$rootDir"

echo [Install git hooks] installing git hooks to project repo...
cp ./git-hooks/* ./.git/hooks
echo [Install git hooks] git hooks sucessfully installed

echo [Install go tools] installing go tools...
go get golang.org/x/tools/cmd/cover
go get golang.org/x/lint/golint
echo [Install go tools] go tools sucessfully installed