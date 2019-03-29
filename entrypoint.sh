#!/usr/bin/env bash

# https://vaneyckt.io/posts/safer_bash_scripts_with_set_euxo_pipefail
set -euxo pipefail


if [ "$DEV_TEST_DEBUG" = "true" ]
then
    rm ${CVE_FEEDS_DIR}/*20*.json
    echo "(cd /usr/local/mounted/ && go run main.go)" > start.sh
    chmod +x start.sh
    sleep 999999
else
    go build ./...
    go test ./...

    go run main.go
fi
