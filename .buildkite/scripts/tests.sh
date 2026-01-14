#!/usr/bin/env sh

set -u

TIMEFORMAT="(Duration: %0lR)"

# Default CA location when running in the Buildkite test container.
if [ -z "${ELASTICSEARCH_CA_CERT:-}" ] && [ -f "/go-elasticsearch/.buildkite/certs/ca.crt" ]; then
  ELASTICSEARCH_CA_CERT="/go-elasticsearch/.buildkite/certs/ca.crt"
  export ELASTICSEARCH_CA_CERT
fi

printf "\033[34;1mINFO:\033[0m Using elasticsearch-clients-tests branch: %s\033[0m\n" "${ELASTICSEARCH_CLIENTS_TESTS_BRANCH:-9.1}"

printf "\033[34;1mINFO:\033[0m Cleaning up generated test files\033[0m\n"
rm -rf esapi/test/*_test.go
rm -rf esapi/test/xpack

printf "\033[34;1mINFO:\033[0m Generating tests from elasticsearch-clients-tests\033[0m\n"

# `gen-tests` depends on `download-client-tests` and will clone/update the repo under ./tmp.
cd /go-elasticsearch || exit 1
ELASTICSEARCH_CLIENTS_TESTS_BRANCH="${ELASTICSEARCH_CLIENTS_TESTS_BRANCH:-9.3}" \
  make gen-tests

printf "\033[34;1mINFO:\033[0m Download tests deps >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>\033[0m\n"
cd /go-elasticsearch/esapi/test || exit 1
go mod tidy

printf "\033[34;1mINFO:\033[0m Running the tests\033[0m\n"

gotestsum --format=short-verbose --junitfile=$WORKSPACE/TEST-integration-api-junit.xml -- -tags=integration -timeout=1h ./...
status=$?

exit $status


