#!/bin/bash

#!/usr/bin/env bash

CONTAINER_NAME=wallet.d
PACKAGE_NAME=github.com/fr05t1k/wallet

case $1 in
    up)
        docker-compose up
    ;;
    run)
        docker-compose run ${CONTAINER_NAME} go run -v src/${PACKAGE_NAME}/main.go
        docker-compose rm -f
    ;;
    test|tests)
        echo "didn't implement"
    ;;
    bash)
        docker-compose run ${CONTAINER_NAME} bash
    ;;
    gen)
        protoc --proto_path=${GOOGLE_PROTOBUF_INCLUDE_PATH} -I ./  wallet.proto --go_out=plugins=grpc:wallet
    ;;
    *)
        echo "Usage:"
        echo "       	run command [arguments]"
        echo ""
        echo "The commands are:"
        echo ""
        echo "       	run     run application"
        echo "       	up      up containers"
        echo "       	test    running tests"
        echo "       	bash    running bash"
        echo "       	gen     generate proto classes"

esac

