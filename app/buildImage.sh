set -e

currentDir=$(cd -P -- "$(dirname -- "$0")" && pwd -P)
rootDir="$currentDir/../"
cd "$rootDir"

imageTag=$1

if [ -z "$1" ]
  then
    echo 'No imageTag provided. Latest will be used.'
    imageTag=latest
fi

imageFullName=auto_post/:$imageTag

echo [Doc Storage STARTING] building "$imageFullName"...

echo [Doc Storage] remove old image "$imageFullName"...
(docker rmi -f "$imageFullName")

echo [Doc Storage] cleaning or creating an out folder...
OUT_FOLDER=./app/out
if [ -d "$OUT_FOLDER" ]; then
  rm -rf "$OUT_FOLDER"
fi
mkdir "$OUT_FOLDER"
echo [Doc Storage] an out folder succesfully cleaned

echo [Doc Storage] building a binary for linux...
env GOOS=linux GOARCH=arm64 go build -o "${OUT_FOLDER}"/auto_post-linux ./app/cmd/auto_post/
echo [Doc Storage] a binary succesfully builded

echo [Doc Storage] creating docker image "$imageFullName"...
(docker build -t "$imageFullName" ./app)

echo [Doc Storage FINISHED] image "$imageFullName" has been built