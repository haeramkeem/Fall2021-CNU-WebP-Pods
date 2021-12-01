#!/bin/bash

cmd=$1

case "$cmd" in

    # Development
    --ui-dev)
        cd ./ui && npm run serve
        ;;

    # Build
    --build-amd64)
        env GOOS=linux GOARCH=amd64 go build -o pods_amd64
        ;;

    # Unexpected
    *)
        echo "Unknown command '$cmd'"
        exit 1
        ;;

esac
