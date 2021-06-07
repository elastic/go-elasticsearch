#!/usr/bin/env sh

set -uo pipefail

TIMEFORMAT="(Duration: %0lR)"

echo -e "\033[34;1mINFO:\033[0m Cleaning up test files\033[0m"

rm -rf esapi/test/*_test.go
rm -rf esapi/test/xpack

echo -e "\033[34;1mINFO:\033[0m Generating the API registry\033[0m"

cd /go-elasticsearch/internal/build || exit
go get -u golang.org/x/tools/cmd/goimports
PACKAGE_PATH=/go-elasticsearch/esapi go generate ./...

echo -e "\033[34;1mINFO:\033[0m Generating the test files\033[0m"

go run main.go apitests --output '/go-elasticsearch/esapi/test' --input "/tmp/rest-api-spec/test/free/**/*.y*ml"

echo -e "\033[34;1mINFO:\033[0m Download tests deps >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>\033[0m"

cd /go-elasticsearch/esapi/test || exit
go mod tidy

echo -e "\033[34;1mINFO:\033[0m Running the tests\033[0m"

gotestsum --format=short-verbose --junitfile=$WORKSPACE/TEST-integration-api-junit.xml -- -tags=integration -timeout=1h *_test.go
status=$?

exit $status
