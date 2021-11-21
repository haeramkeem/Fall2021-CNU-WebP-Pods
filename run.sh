#!/bin/bash

cmd=$1

case "$cmd" in
    --ui-dev)
        cd ./ui && npm run serve
        ;;
    *)
        echo "Unknown command '$cmd'"
        exit 1
        ;;

esac
