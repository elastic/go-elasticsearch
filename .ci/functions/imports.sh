#!/usr/bin/env bash
#
# Sets up all the common variables and imports relevant functions
#
# Version 1.0.1
# - Initial version after refactor
# - Validate STACK_VERSION asap

function require_stack_version() {
  if [[ -z $STACK_VERSION ]]; then
    echo -e "\033[31;1mERROR:\033[0m Required environment variable [STACK_VERSION] not set\033[0m"
    exit 1
  fi
}

require_stack_version

if [[ -z $es_node_name ]]; then
  # only set these once
  set -euo pipefail
  export TEST_SUITE=${TEST_SUITE-free}
  export RUNSCRIPTS=${RUNSCRIPTS-}
  export DETACH=${DETACH-false}
  export CLEANUP=${CLEANUP-false}

  export es_node_name=instance
  export elastic_password=changeme
  export elasticsearch_image=elasticsearch
  export elasticsearch_url=https://elastic:${elastic_password}@${es_node_name}:9200
  if [[ $TEST_SUITE != "platinum" ]]; then
    export elasticsearch_url=http://${es_node_name}:9200
  fi
  export external_elasticsearch_url=${elasticsearch_url/$es_node_name/localhost}
  export elasticsearch_container="${elasticsearch_image}:${STACK_VERSION}"

  export suffix=rest-test
  export moniker=$(echo "$elasticsearch_container" | tr -C "[:alnum:]" '-')
  export network_name=${moniker}${suffix}

  export ssl_cert="${script_path}/certs/testnode.crt"
  export ssl_key="${script_path}/certs/testnode.key"
  export ssl_ca="${script_path}/certs/ca.crt"

fi

  export script_path=$(dirname $(realpath -s $0))
  source $script_path/functions/cleanup.sh
  source $script_path/functions/wait-for-container.sh
  trap "cleanup_trap ${network_name}" EXIT


if [[ "$CLEANUP" == "true" ]]; then
  cleanup_all_in_network $network_name
  exit 0
fi

echo -e "\033[34;1mINFO:\033[0m Creating network $network_name if it does not exist already \033[0m"
docker network inspect "$network_name" > /dev/null 2>&1 || docker network create "$network_name"

