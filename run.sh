#!/bin/bash

cmd=$1

case "$cmd" in

    # UI related
    --ui-dev)
        cd ./ui && npm run serve
        ;;

    # Chrome Dev Protocol related
    --cdp-ins)
        docker pull --platform linux/amd64 cdp chromedp/headless-shell:latest
        ;;

    --cdp-run)
        docker run -d -p 9222:9222 --rm --name cdp --platform linux/amd64 -v $(pwd):/usr/local/pods chromedp/headless-shell
        ;;

    --cdp-sh)
        docker exec -it cdp /bin/bash
        ;;

    # Build related
    --build-amd64)
        env GOOS=linux GOARCH=amd64 go build -o pods_amd64
        ;;

    # Unexpected
    *)
        echo "Unknown command '$cmd'"
        exit 1
        ;;

esac
