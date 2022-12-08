#!/bin/bash
set -e
currentDir=$(cd -P -- "$(dirname -- "$0")" && pwd -P)
rootDir="$currentDir/../"

(cd "$rootDir" && exec ./scripts/build.sh)
(cd "$rootDir" && exec ./app/buildImage.sh)
(cd "$rootDir" && exec docker-compose -f ./docker/docker-compose.yml --env-file \
./docker/env/local.env --project-name=auto_post/_local --profile local up -d --remove-orphans)

portainerPort=$(cd "$rootDir" && cat ./docker/env/local.env | grep "PORTAINER_PORT" | cut -d'=' -f2)

printf 'List of available ports\n'
(cd "$rootDir" && exec cat ./docker/env/local.env)
printf "\nPortainer GUI is available at http://localhost:$portainerPort/#/dashboard\n"
printf 'Pgadmin Login: auto_post/@dltru.org:12345\nDatabase password is postgres'
printf '\n'
python -mwebbrowser http://localhost:$portainerPort/#/dashboard || true