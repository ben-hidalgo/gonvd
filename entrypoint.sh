#!/usr/bin/env bash

# https://vaneyckt.io/posts/safer_bash_scripts_with_set_euxo_pipefail
set -euxo pipefail

sleep ${SLEEP_SECONDS}

#TODO: mechanism to delete files based on a flag for quick debug/test startup
# rm /usr/local/cve/*20*.json

go run main.go
