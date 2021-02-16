#!/usr/bin/env bash
#
# Launch one or more Elasticsearch nodes via the Docker image,
# to form a cluster suitable for running the REST API tests.
#
# Export the STACK_VERSION variable, eg. '8.0.0-SNAPSHOT'.
# Export the TEST_SUITE variable, eg. 'free' or 'platinum' defaults to 'free'.
# Export the NUMBER_OF_NODES variable to start more than 1 node

# Version 1.2.0
# - Initial version of the run-elasticsearch.sh script
# - Deleting the volume should not dependent on the container still running
# - Fixed `ES_JAVA_OPTS` config
# - Moved to STACK_VERSION and TEST_VERSION
# - Refactored into functions and imports
# - Support NUMBER_OF_NODES
# - Added 5 retries on docker pull for fixing transient network errors

script_path=$(dirname $(realpath -s $0))
source $script_path/functions/imports.sh
set -euo pipefail

echo -e "\033[34;1mINFO:\033[0m Take down node if called twice with the same arguments (DETACH=true) or on seperate terminals \033[0m"
cleanup_node $es_node_name

master_node_name=${es_node_name}
cluster_name=${moniker}${suffix}

declare -a volumes
environment=($(cat <<-END
  --env node.name=$es_node_name
  --env cluster.name=$cluster_name
  --env cluster.initial_master_nodes=$master_node_name
  --env discovery.seed_hosts=$master_node_name
  --env cluster.routing.allocation.disk.threshold_enabled=false
  --env bootstrap.memory_lock=true
  --env node.attr.testattr=test
  --env path.repo=/tmp
  --env repositories.url.allowed_urls=http://snapshot.test*
END
))
if [[ "$TEST_SUITE" == "platinum" ]]; then
  environment+=($(cat <<-END
    --env ELASTIC_PASSWORD=$elastic_password
    --env xpack.license.self_generated.type=trial
    --env xpack.security.enabled=true
    --env xpack.security.http.ssl.enabled=true
    --env xpack.security.http.ssl.verification_mode=certificate
    --env xpack.security.http.ssl.key=certs/testnode.key
    --env xpack.security.http.ssl.certificate=certs/testnode.crt
    --env xpack.security.http.ssl.certificate_authorities=certs/ca.crt
    --env xpack.security.transport.ssl.enabled=true
    --env xpack.security.transport.ssl.key=certs/testnode.key
    --env xpack.security.transport.ssl.certificate=certs/testnode.crt
    --env xpack.security.transport.ssl.certificate_authorities=certs/ca.crt
END
))
  volumes+=($(cat <<-END
    --volume $ssl_cert:/usr/share/elasticsearch/config/certs/testnode.crt
    --volume $ssl_key:/usr/share/elasticsearch/config/certs/testnode.key
    --volume $ssl_ca:/usr/share/elasticsearch/config/certs/ca.crt
END
))
fi

cert_validation_flags=""
if [[ "$TEST_SUITE" == "platinum" ]]; then
  cert_validation_flags="--insecure --cacert /usr/share/elasticsearch/config/certs/ca.crt --resolve ${es_node_name}:443:127.0.0.1"
fi

# Pull the container, retry on failures up to 5 times with
# short delays between each attempt. Fixes most transient network errors.
docker_pull_attempts=0
until [ "$docker_pull_attempts" -ge 5 ]
do
   docker pull docker.elastic.co/elasticsearch/"$elasticsearch_container" && break
   docker_pull_attempts=$((docker_pull_attempts+1))
   echo "Failed to pull image, retrying in 10 seconds (retry $docker_pull_attempts/5)..."
   sleep 10
done

NUMBER_OF_NODES=${NUMBER_OF_NODES-1}
http_port=9200
for (( i=0; i<$NUMBER_OF_NODES; i++, http_port++ )); do
  node_name=${es_node_name}$i
  node_url=${external_elasticsearch_url/9200/${http_port}}$i
  if [[ "$i" == "0" ]]; then node_name=$es_node_name; fi
  environment+=($(cat <<-END
    --env node.name=$node_name
END
))
  echo "$i: $http_port $node_url "
  volume_name=${node_name}-${suffix}-data
  volumes+=($(cat <<-END
    --volume $volume_name:/usr/share/elasticsearch/data${i}
END
))

  # make sure we detach for all but the last node if DETACH=false (default) so all nodes are started
  local_detach="true"
  if [[ "$i" == "$((NUMBER_OF_NODES-1))" ]]; then local_detach=$DETACH; fi
  echo -e "\033[34;1mINFO:\033[0m Starting container $node_name \033[0m"
  set -x
  docker run \
    --name "$node_name" \
    --network "$network_name" \
    --env "ES_JAVA_OPTS=-Xms1g -Xmx1g" \
    "${environment[@]}" \
    "${volumes[@]}" \
    --publish "$http_port":9200 \
    --ulimit nofile=65536:65536 \
    --ulimit memlock=-1:-1 \
    --detach="$local_detach" \
    --health-cmd="curl $cert_validation_flags --fail $elasticsearch_url/_cluster/health || exit 1" \
    --health-interval=2s \
    --health-retries=20 \
    --health-timeout=2s \
    --rm \
    docker.elastic.co/elasticsearch/"$elasticsearch_container";

  set +x
  if wait_for_container "$es_node_name" "$network_name"; then
    echo -e "\033[32;1mSUCCESS:\033[0m Running on: $node_url\033[0m"
  fi

done

