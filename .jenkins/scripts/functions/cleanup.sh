#!/usr/bin/env bash
#
# Shared cleanup routines between different steps
#
# Please source .ci/functions/imports.sh as a whole not just this file
#
# Version 1.0.0
# - Initial version after refactor

function cleanup_volume {
  if [[ "$(docker volume ls -q -f name=$1)" ]]; then
    echo -e "\033[34;1mINFO:\033[0m Removing volume $1\033[0m"
    (docker volume rm "$1") || true
  fi
}
function container_running {
  if [[ "$(docker ps -q -f name=$1)" ]]; then
    return 0;
    else return 1;
  fi
}
function cleanup_node {
  if container_running "$1"; then
    echo -e "\033[34;1mINFO:\033[0m Removing container $1\033[0m"
    (docker container rm --force --volumes "$1") || true
  fi
  if [[ -n "$1" ]]; then
    echo -e "\033[34;1mINFO:\033[0m Removing volume $1-${suffix}-data\033[0m"
    cleanup_volume "$1-${suffix}-data"
  fi
}
function cleanup_network {
  if [[ "$(docker network ls -q -f name=$1)" ]]; then
    echo -e "\033[34;1mINFO:\033[0m Removing network $1\033[0m"
    (docker network rm "$1") || true
  fi
}

function cleanup_trap {
  status=$?
  set +x
  if [[ "$DETACH" != "true" ]]; then
    echo -e "\033[34;1mINFO:\033[0m clean the network if not detached (start and exit)\033[0m"
    cleanup_all_in_network "$1"
  fi
  # status is 0 or SIGINT
  if [[ "$status" == "0" || "$status" == "130" ]]; then
    echo -e "\n\033[32;1mSUCCESS run-tests\033[0m"
    exit 0
  else
    echo -e "\n\033[31;1mFAILURE during run-tests\033[0m"
    exit ${status}
  fi
};
function cleanup_all_in_network {

  if [[ -z "$(docker network ls -q -f name="^$1\$")" ]]; then
    echo -e "\033[34;1mINFO:\033[0m $1 is already deleted\033[0m"
    return 0
  fi
  containers=$(docker network inspect -f '{{ range $key, $value := .Containers }}{{ printf "%s\n" .Name}}{{ end }}' $1)
  while read -r container; do
    cleanup_node "$container"
  done <<< "$containers"
  cleanup_network $1
  echo -e "\033[32;1mSUCCESS:\033[0m Cleaned up and exiting\033[0m"
};
