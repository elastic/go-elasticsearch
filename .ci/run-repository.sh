#!/usr/bin/env bash
#
# Called by entry point `run-test` use this script to add your repository specific test commands
#
# Once called Elasticsearch is up and running
#
# Its recommended to call `imports.sh` as defined here so that you get access to all variables defined there
#
# Any parameters that test-matrix.yml defines should be declared here with appropiate defaults

script_path=$(dirname $(realpath -s $0))
source $script_path/functions/imports.sh
set -euo pipefail

echo -e "\033[34;1mINFO:\033[0m VERSION: ${STACK_VERSION}\033[0m"
echo -e "\033[34;1mINFO:\033[0m TEST_SUITE: ${TEST_SUITE}\033[0m"
echo -e "\033[34;1mINFO:\033[0m RUNSCRIPTS: ${RUNSCRIPTS}\033[0m"
echo -e "\033[34;1mINFO:\033[0m URL: ${elasticsearch_url}\033[0m"

echo -e "\033[34;1mINFO:\033[0m pinging Elasticsearch ..\033[0m"
curl --insecure --fail $external_elasticsearch_url/_cluster/health?pretty

if [[ "$RUNSCRIPTS" = *"enterprise-search"* ]]; then
  enterprise_search_url="http://localhost:3002"
  echo -e "\033[34;1mINFO:\033[0m pinging Enterprise Search ..\033[0m"
  curl -I --fail $enterprise_search_url
fi

echo -e "\033[32;1mSUCCESS:\033[0m successfully started the ${STACK_VERSION} stack.\033[0m"

echo -e "\033[34;1mINFO:\033[0m Building the [elastic/go-elasticsearch] image... \033[0m"

docker build \
    --file .ci/Dockerfile \
    --tag elastic/go-elasticsearch \
    .

echo -e "\033[34;1mINFO:\033[0m Retrieving Elasticsearch Version & Hash from container... \033[0m"

ELASTICSEARCH_BUILD_VERSION=$(curl -sSk $external_elasticsearch_url | jq -r '.version.number')
ELASTICSEARCH_BUILD_HASH=$(curl -sSk $external_elasticsearch_url | jq -r '.version.build_hash')

echo -e "\033[34;1mINFO:\033[0m Download Elasticsearch specs... \033[0m"
docker run --volume=$HOME/workspace/tmp:/tmp --workdir=/tmp --rm elastic/go-elasticsearch /bin/sh -c "
  rm -rf /tmp/rest.zip
  rm -rf /tmp/rest-api-spec/
  curl -s https://artifacts-api.elastic.co/v1/versions/8.0.0-SNAPSHOT | jq '.version.builds|=sort_by(.start_time | strptime(\"%a, %d %b %Y %H:%M:%S %Z\"))| .version.builds | reverse[] | .projects|select(.elasticsearch)|.elasticsearch' | jq -s '.[0]' > /tmp/elasticsearch.json
  curl \"\$(cat /tmp/elasticsearch.json | jq -r '.packages | with_entries( select(.key|contains(\"rest-resources\"))) | first(.[] | .url)')\" --output /tmp/rest.zip
  unzip -q /tmp/rest.zip -d /tmp
"

echo -e "\033[34;1mINFO:\033[0m Execute [$TEST_SUITE] >>>>>>>>>>>>>>>>>>>>>>>>>>>>>\033[0m"

docker run -t \
  --volume=$WORKSPACE/tmp:/tmp \
  --rm \
  --network "$network_name" \
  --env "ELASTICSEARCH_URL=$elasticsearch_url" \
  --env "ELASTICSEARCH_VERSION=$STACK_VERSION" \
  --env "ELASTICSEARCH_BUILD_VERSION=$ELASTICSEARCH_BUILD_VERSION" \
  --env "ELASTICSEARCH_BUILD_HASH=$ELASTICSEARCH_BUILD_HASH" \
  --env "WORKSPACE=${WORKSPACE:-/workspace}" \
  --volume "${WORKSPACE:-workspace}:${WORKSPACE:-/workspace}" \
  elastic/go-elasticsearch \
  .ci/scripts/tests-$TEST_SUITE.sh