#!/usr/bin/env bash

set -euo pipefail

TIMEFORMAT="(Duration: %0lR)"

echo -e "\033[1m>>>>> Cleaning up test files\033[0m"

docker exec go-elasticsearch /bin/sh -c 'rm -rf esapi/test/*_test.go'
docker exec go-elasticsearch /bin/sh -c 'rm -rf esapi/test/xpack'

echo -e "\033[1m>>>>> Generating the API registry\033[0m"

docker exec --workdir=/go-elasticsearch/internal/cmd/generate --env PACKAGE_PATH=/go-elasticsearch/esapi go-elasticsearch go generate ./...

echo -e "\033[1m>>>>> Generating the test files\033[0m"

time docker exec --tty --workdir=/go-elasticsearch/internal/cmd/generate go-elasticsearch go run main.go apitests --output '/go-elasticsearch/esapi/test' --input '/elasticsearch-source/elasticsearch/rest-api-spec/src/main/resources/rest-api-spec/test/**/*.y*ml'

echo -e "\033[1m>>>>> Running the tests\033[0m"

time docker exec --tty --workdir=/go-elasticsearch/esapi/test go-elasticsearch /bin/sh -c 'gotestsum --format=short-verbose --junitfile=$WORKSPACE/TEST-integration-api-junit.xml -- -tags=integration -timeout=1h *_test.go'
status=$?

exit $status
