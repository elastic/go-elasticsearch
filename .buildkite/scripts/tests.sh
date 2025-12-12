#!/usr/bin/env sh

set -eu
# Best-effort pipefail (not supported by all /bin/sh implementations)
set -o pipefail 2>/dev/null || true

TIMEFORMAT="(Duration: %0lR)"

echo -e "\033[34;1mINFO:\033[0m Using elasticsearch-clients-tests branch: ${ELASTICSEARCH_CLIENTS_TESTS_BRANCH:-9.1}\033[0m"

echo -e "\033[34;1mINFO:\033[0m Cleaning up generated test files\033[0m"
rm -rf esapi/test/*_test.go
rm -rf esapi/test/xpack

echo -e "\033[34;1mINFO:\033[0m Generating tests from elasticsearch-clients-tests\033[0m"

# `gen-tests` depends on `download-client-tests` and will clone/update the repo under ./tmp.
cd /go-elasticsearch || exit 1
ELASTICSEARCH_CLIENTS_TESTS_BRANCH="${ELASTICSEARCH_CLIENTS_TESTS_BRANCH:-9.1}" \
  make gen-tests

echo -e "\033[34;1mINFO:\033[0m Download tests deps >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>\033[0m"
cd /go-elasticsearch/esapi/test || exit 1
go mod tidy

echo -e "\033[34;1mINFO:\033[0m Running the tests\033[0m"

gotestsum --format=short-verbose --junitfile=$WORKSPACE/TEST-integration-api-junit.xml -- -tags=integration -timeout=1h ./...
status=$?

exit $status


