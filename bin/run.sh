#!/bin/bash

#!/usr/bin/env bash

CONTAINER_NAME=wallet
PACKAGE_NAME=github.com/fr05t1k/wallet

case $1 in
    up)
        docker-compose up
    ;;
    run)
        docker-compose run ${CONTAINER_NAME} go run -v src/${PACKAGE_NAME}/main.go
    ;;
    test|tests)
        echo "didn't implement"
    ;;
    bash)
        docker-compose run ${CONTAINER_NAME} bash
    ;;
    build)
        echo "didn't implement"
    ;;
    *)
        echo "Command not found!"
        echo ""
        echo "Available commands: app, tests, test"
        echo "  run     - run application"
        echo "  up      - up containers"
        echo "  test    - running tests"
        echo "  bash    - running bash"
        echo "  build   - build debian package (debil required)"

esac


