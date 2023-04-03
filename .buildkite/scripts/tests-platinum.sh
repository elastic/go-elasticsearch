#!/usr/bin/env sh

set -uo pipefail

TIMEFORMAT="(Duration: %0lR)"

echo -e "\033[34;1mINFO:\033[0m Cleaning up test files\033[0m"

rm -rf esapi/test/*_test.go
rm -rf rm -rf esapi/test/ml*

echo -e "\033[34;1mINFO:\033[0m Generating the API registry\033[0m"

cd /go-elasticsearch/internal/build || exit
go get -u golang.org/x/tools/cmd/goimports
PACKAGE_PATH=/go-elasticsearch/esapi go generate ./...

echo -e "\033[34;1mINFO:\033[0m Generating the test files\033[0m"

go run main.go apitests --output '/go-elasticsearch/esapi/test/xpack' --input '/tmp/rest-api-spec/test/platinum/**/*.yml'

cd /go-elasticsearch || exit

mkdir -p esapi/test/xpack/ml
mkdir -p esapi/test/xpack/ml-crud

mv esapi/test/xpack/xpack_ml* esapi/test/xpack/ml/
mv esapi/test/xpack/ml/xpack_ml__jobs_crud_test.go esapi/test/xpack/ml-crud/

set +e # Do not fail immediately when a single test suite fails

echo -e "\033[34;1mINFO:\033[0m Download tests deps >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>\033[0m"

cd /go-elasticsearch/esapi/test || exit
go mod tidy

echo -e "\033[34;1mINFO:\033[0m Running tests: XPACK >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>\033[0m"

gotestsum --format=short-verbose --junitfile=$WORKSPACE/TEST-integration-api-xpack-junit.xml -- --tags=integration --timeout=1h -v xpack/*_test.go
status1=$?

echo -e "\033[34;1mINFO:\033[0m Running tests: XPACK ML >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>\033[0m"

gotestsum --format=short-verbose --junitfile=$WORKSPACE/TEST-integration-api-xpack-ml-junit.xml -- --tags=integration --timeout=1h -v ./xpack/ml/*_test.go
status2=$?

echo -e "\033[34;1mINFO:\033[0m Running tests: XPACK ML CRUD >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>\033[0m"

gotestsum --format=short-verbose --junitfile=$WORKSPACE/TEST-integration-api-xpack-ml-crud-junit.xml -- --tags=integration --timeout=1h -v ./xpack/ml-crud/*_test.go
status3=$?

if [[ $status1 == 0 && $status2 == 0 && $status3 == 0 ]]; then
  exit 0
else
  exit 1
fi
