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
echo -e "\033[34;1mINFO:\033[0m elasticsearch_url: ${elasticsearch_url}\033[0m"
echo -e "\033[34;1mINFO:\033[0m external_elasticsearch_url: ${external_elasticsearch_url}\033[0m"

echo -e "\033[34;1mINFO:\033[0m Pinging Elasticsearch...\033[0m"
curl --insecure --fail $external_elasticsearch_url/_cluster/health?pretty

if [[ "$RUNSCRIPTS" = *"enterprise-search"* ]]; then
  enterprise_search_url="http://localhost:3002"
  echo -e "\033[34;1mINFO:\033[0m pinging Enterprise Search ..\033[0m"
  curl -I --fail $enterprise_search_url
fi

echo -e "\033[32;1mSUCCESS:\033[0m successfully started the ${STACK_VERSION} stack.\033[0m"

# -----------------------------------------------------------------------------

TIMEFORMAT="Duration: %0lR"

if [[ $TEST_SUITE != "oss" && $TEST_SUITE != "xpack" ]]; then
  echo -e "\033[31;1mERROR:\033[0m Unknown value [$TEST_SUITE] for [TEST_SUITE]\033[0m"
  exit 1
fi

echo
echo -e "\033[1m>>>>> SETUP [$TEST_SUITE] >>>>>>>>>>>>>>>>>>>>>>>>>>>>>\033[0m"

# Store the Elasticsearch version and build hash
ELASTICSEARCH_BUILD_VERSION=$(curl -sSk $external_elasticsearch_url | jq -r '.version.number')
ELASTICSEARCH_BUILD_HASH=$(curl -sSk $external_elasticsearch_url | jq -r '.version.build_hash')

# Build the go-elasticsearch image
echo -e ">>>>> Building the [elastic/go-elasticsearch] image..."
docker build --file Dockerfile --tag elastic/go-elasticsearch .

# Download Elasticsearch source code at ELASTICSEARCH_BUILD_HASH and store it in a container
echo -e ">>>>> Downloading Elasticsearch repository @ $ELASTICSEARCH_BUILD_HASH..."
time curl --retry 10 -sSL "https://github.com/elastic/elasticsearch/archive/$ELASTICSEARCH_BUILD_HASH.zip" > "/tmp/elasticsearch-$ELASTICSEARCH_BUILD_HASH.zip"
echo -e ">>>>> Extracting the Elasticsearch source..."
time docker run --volume=/tmp:/tmp --workdir=/tmp --rm elastic/go-elasticsearch unzip -q -o "elasticsearch-$ELASTICSEARCH_BUILD_HASH.zip" '*.properties' '*.json' '*.y*ml'
docker run --volume=/tmp:/tmp --workdir=/tmp --rm elastic/go-elasticsearch /bin/sh -c "
  rm -rf /tmp/elasticsearch-$ELASTICSEARCH_BUILD_HASH.zip
  rm -rf /tmp/elasticsearch/
  mv /tmp/elasticsearch-$ELASTICSEARCH_BUILD_HASH* /tmp/elasticsearch/
"

echo -e ">>>>> Creating the container with Elasticsearch repository..."
docker container rm -f elasticsearch-source > /dev/null 2>&1 || true
docker create --name elasticsearch-source --volume /elasticsearch-source alpine /bin/true
docker cp /tmp/elasticsearch elasticsearch-source:/elasticsearch-source

# Launch the container; actual commands are called with "docker exec"
docker run \
  --name go-elasticsearch \
  --network "$network_name" \
  --env "ELASTICSEARCH_URL=$elasticsearch_url" \
  --env "ELASTICSEARCH_VERSION=$STACK_VERSION" \
  --env "ELASTICSEARCH_BUILD_VERSION=$ELASTICSEARCH_BUILD_VERSION" \
  --env "ELASTICSEARCH_BUILD_HASH=$ELASTICSEARCH_BUILD_HASH" \
  --env "WORKSPACE=${WORKSPACE:-/workspace}" \
  --volume "/go-elasticsearch" \
  --volume "${WORKSPACE:-workspace}:${WORKSPACE:-/workspace}" \
  --volumes-from "elasticsearch-source" \
  --rm \
  --detach \
  elastic/go-elasticsearch sleep 3600

echo
echo -e "\033[1m>>>>> EXECUTE [$TEST_SUITE] >>>>>>>>>>>>>>>>>>>>>>>>>>>>>\033[0m"

# Run the tests
# NOTE: Conditions needed to prevent early exit due to the 'set -e' option
status=100
case $TEST_SUITE in
  "oss" )
    if bash .jenkins/scripts/tests-oss.sh; then
      status=$?
    else
      status=$?
    fi
    ;;
  "xpack" )
    if bash .jenkins/scripts/tests-xpack.sh; then
      status=$?
    else
      status=$?
    fi
    ;;
esac

# Report status and exit
if [[ $status == "0" ]]; then
  echo -e "\n\033[32;1mSUCCESS [$TEST_SUITE]\033[0m"
  exit 0
else
  echo -e "\n\033[31;1mFAILURE [$TEST_SUITE]\033[0m"
  exit $status
fi
