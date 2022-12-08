#!/bin/bash
set -e

currentDir=$(cd -P -- "$(dirname -- "$0")" && pwd -P)
rootDir="$currentDir/../"
cd "$rootDir"

echo [Clean an out folder] cleaning or creating an out folder...
OUT_FOLDER=./app/out
if [ -d "$OUT_FOLDER" ]; then
  rm -rf "$OUT_FOLDER"
fi
mkdir "$OUT_FOLDER"
echo [Clean an out folder] an out folder succesfully cleaned

echo [Clean a reports folder] cleaning or creating a reports folder...
REPORTS_FOLDER=./app/reports
if [ -d "$REPORTS_FOLDER" ]; then
  rm -rf "$REPORTS_FOLDER"
fi
mkdir "$REPORTS_FOLDER"
echo [Clean a reports folder] a reports folder succesfully cleaned