#!/usr/bin/env sh

set -eo pipefail

echo -e "\033[1m>>>>> CI >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>\033[0m" && \
echo -e "\033[1m>>>>> Waiting for Elasticsearch on $ELASTICSEARCH_URL...\033[0m" && \
curl --max-time 120 --retry 120 --retry-delay 1 --retry-connrefused --show-error --silent --fail $ELASTICSEARCH_URL && \
curl -s $ELASTICSEARCH_URL | jq -r '.version.build_hash' > .elasticsearch_build_hash && \
export ES_BUILD_HASH=$(cat .elasticsearch_build_hash) && \
echo -e "\033[1m>>>>> Downloading Elasticsearch repository @ $ES_BUILD_HASH...\033[0m" && \
echo "curl -sSL -o elasticsearch-$ES_BUILD_HASH.zip https://github.com/elastic/elasticsearch/archive/$ES_BUILD_HASH.zip" && \
(rm -rf tmp/elasticsearch) || true && (mkdir -p tmp || true) && \
curl --retry 3 -sSL -o elasticsearch-$ES_BUILD_HASH.zip https://github.com/elastic/elasticsearch/archive/$ES_BUILD_HASH.zip && \
unzip -q -o elasticsearch-$ES_BUILD_HASH.zip '*.json' '*.yml' -d tmp && \
mv tmp/elasticsearch-$ES_BUILD_HASH* tmp/elasticsearch && \
echo -e "\033[1m>>>>> Generating the API registry\033[0m" && \
(cd internal/cmd/generate && ELASTICSEARCH_BUILD_HASH=$ES_BUILD_HASH PACKAGE_PATH=$PWD/../../../esapi go generate -v ./... 2> /dev/null) && \
echo -e "\033[1m>>>>> Generating the test files\033[0m" && \
(cd internal/cmd/generate && ELASTICSEARCH_BUILD_HASH=$ES_BUILD_HASH go run main.go tests --input "$PWD/../../../tmp/elasticsearch/rest-api-spec/src/main/resources/rest-api-spec/test/**/*.yml" --output="$PWD/../../../esapi/test") && \
echo -e "\033[1m>>>>> Running the tests\033[0m" && \
(cd esapi/test && gotestsum --format=short-verbose --junitfile="$WORKSPACE/TEST-integration-api-junit.xml" -- -tags='integration' -timeout=1h ./...) ; exitstatus=$? ; \
echo -e "\n\033[1m<<<<< Finished API integration tests for $ES_BUILD_HASH\033[0m" ; \
echo -e "Junit Report: $(ls $WORKSPACE/TEST-integration-api-junit.xml)" && \
exit $exitstatus
