#!/bin/bash
set -e
currentDir=$(cd -P -- "$(dirname -- "$0")" && pwd -P)
rootDir="$currentDir/../"

(cd "$rootDir" && exec docker-compose -f ./docker/docker-compose.yml --env-file \
        ./docker/env/local.env --project-name=auto_post/_local --profile local down -v)
(cd "$rootDir" && exec docker-compose -f ./docker/docker-compose.yml --env-file \
        ./docker/env/local.env --project-name=auto_post/_local --profile local rm -f)