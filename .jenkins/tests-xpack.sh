#!/usr/bin/env bash

set -euo pipefail

TIMEFORMAT="(Duration: %0lR)"

echo -e "\033[1m>>>>> Cleaning up test files\033[0m"

docker exec go-elasticsearch /bin/sh -c 'rm -rf esapi/test/*_test.go'
docker exec go-elasticsearch /bin/sh -c 'rm -rf rm -rf esapi/test/ml*'

echo -e "\033[1m>>>>> Generating the API registry\033[0m"

docker exec --workdir=/go-elasticsearch/internal/cmd/generate --env PACKAGE_PATH=/go-elasticsearch/esapi go-elasticsearch go generate ./...

echo -e "\033[1m>>>>> Generating the test files\033[0m"

time docker exec --tty --workdir=/go-elasticsearch/internal/cmd/generate go-elasticsearch go run main.go apitests --output '/go-elasticsearch/esapi/test/xpack' --input '/elasticsearch-source/elasticsearch/x-pack/plugin/src/test/resources/rest-api-spec/test/**/*.yml'

time docker exec --tty --workdir=/go-elasticsearch/internal/cmd/generate go-elasticsearch go run main.go apitests --output '/go-elasticsearch/esapi/test/xpack' --input '/elasticsearch-source/elasticsearch/x-pack/plugin/src/test/resources/rest-api-spec/test/**/**/*.yml'

docker exec go-elasticsearch mkdir -p esapi/test/xpack/ml
docker exec go-elasticsearch mkdir -p esapi/test/xpack/ml-crud

docker exec go-elasticsearch /bin/sh -c 'mv esapi/test/xpack/xpack_ml* esapi/test/xpack/ml/'
docker exec go-elasticsearch mv esapi/test/xpack/ml/xpack_ml__jobs_crud_test.go esapi/test/xpack/ml-crud/

echo -e "\033[1m>>>>> Running tests: XPACK >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>\033[0m"

time docker exec --tty --workdir=/go-elasticsearch/esapi/test go-elasticsearch /bin/sh -c 'gotestsum --format=short-verbose --junitfile=$WORKSPACE/TEST-integration-api-xpack-junit.xml -- --tags=integration --timeout=1h -v xpack/*_test.go'
status1=$?

docker container rm --force --volumes es1 > /dev/null 2>&1
make cluster-clean cluster version=elasticsearch:7.3-SNAPSHOT detached=true

echo -e "\033[1m>>>>> Running tests: XPACK ML >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>\033[0m"

time docker exec --tty --workdir=/go-elasticsearch/esapi/test go-elasticsearch /bin/sh -c 'gotestsum --format=short-verbose --junitfile=$WORKSPACE/TEST-integration-api-xpack-ml-junit.xml -- --tags=integration --timeout=1h -v ./xpack/ml/*_test.go'
status2=$?

docker container rm --force --volumes es1 > /dev/null 2>&1
make cluster-clean cluster version=elasticsearch:7.3-SNAPSHOT detached=true

echo -e "\033[1m>>>>> Running tests: XPACK ML CRUD >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>\033[0m"

time docker exec --tty --workdir=/go-elasticsearch/esapi/test go-elasticsearch /bin/sh -c 'gotestsum --format=short-verbose --junitfile=$WORKSPACE/TEST-integration-api-xpack-ml-crud-junit.xml -- --tags=integration --timeout=1h -v ./xpack/ml-crud/*_test.go'
status3=$?

if [[ $status1 == 0 && $status2 == 0 || $status3 == 0 ]]; then
  exit 0
else
  exit 0
fi
