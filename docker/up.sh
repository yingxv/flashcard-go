#!/bin/bash
set -e

filePath=
suffix=
build=false

while [ -n "$1" ]
do
    case "$1" in
        -d)
            echo "dev环境"
            filePath=.dev
            suffix=-dev
        ;;
        -b)
            echo "rebuild..."
            # echo "-b 选项的参数值是：$2"
            build=true
        ;;
        *)
            echo "$1 is not an option"
            shift
        ;;
    esac
    shift
done
echo "done"

docker compose -f ./docker-compose${filePath}.yml --env-file ~/.env -p flashcard-go${suffix} up -d --build=${build};