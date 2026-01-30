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

# Ensure script_path is always set when this file is sourced.
# Callers typically set it, but we defensively default it here to avoid writing certs under /certs.
export script_path=${script_path:-$(dirname "$(realpath -s "$0")")}

function ensure_tls_certs() {
  # Generates/refreshes a local CA + node certificate for the platinum TLS cluster.
  # This avoids failures like:
  #   x509: certificate has expired or is not yet valid
  #
  # Files are written under: ${script_path}/certs
  local cert_dir="${script_path}/certs"
  local ca_crt="${cert_dir}/ca.crt"
  local ca_key="${cert_dir}/ca.key"
  local node_key="${cert_dir}/testnode.key"
  local node_crt="${cert_dir}/testnode.crt"
  local node_csr="${cert_dir}/testnode.csr"
  local node_ext="${cert_dir}/testnode.ext"

  mkdir -p "${cert_dir}"

  if ! command -v openssl >/dev/null 2>&1; then
    echo -e "\033[31;1mERROR:\033[0m openssl is required to generate TLS certificates\033[0m"
    exit 1
  fi

  local regen="false"
  if [[ ! -s "${ca_crt}" || ! -s "${ca_key}" || ! -s "${node_crt}" || ! -s "${node_key}" ]]; then
    regen="true"
  fi

  # Re-generate if cert is expired or expires within the next 24 hours.
  if [[ "${regen}" == "false" ]]; then
    openssl x509 -checkend 86400 -noout -in "${node_crt}" >/dev/null 2>&1 || regen="true"
    openssl x509 -checkend 86400 -noout -in "${ca_crt}"   >/dev/null 2>&1 || regen="true"
  fi

  if [[ "${regen}" == "true" ]]; then
    echo -e "\033[34;1mINFO:\033[0m Generating TLS certs under ${cert_dir}\033[0m"

    rm -f "${ca_crt}" "${ca_key}" "${node_key}" "${node_crt}" "${node_csr}" "${node_ext}" "${cert_dir}/ca.srl" 2>/dev/null || true

    # CA
    openssl genrsa -out "${ca_key}" 2048 >/dev/null 2>&1
    openssl req -x509 -new -sha256 -days 3650 \
      -key "${ca_key}" \
      -subj "/CN=go-elasticsearch-ci-ca" \
      -out "${ca_crt}" >/dev/null 2>&1

    # Node cert (signed by CA)
    openssl genrsa -out "${node_key}" 2048 >/dev/null 2>&1
    openssl req -new -sha256 \
      -key "${node_key}" \
      -subj "/CN=${es_node_name}" \
      -out "${node_csr}" >/dev/null 2>&1

    cat > "${node_ext}" <<-EOF
subjectAltName=DNS:${es_node_name},DNS:localhost,IP:127.0.0.1
extendedKeyUsage=serverAuth
keyUsage=digitalSignature,keyEncipherment
EOF

    openssl x509 -req -sha256 -days 3650 \
      -in "${node_csr}" \
      -CA "${ca_crt}" -CAkey "${ca_key}" -CAcreateserial \
      -out "${node_crt}" \
      -extfile "${node_ext}" >/dev/null 2>&1

    # IMPORTANT: These files are bind-mounted into the Elasticsearch container which runs as a non-root user.
    # Overly strict host permissions (eg 600) can make the key unreadable in-container and prevent startup.
    chmod 644 "${ca_crt}" "${node_crt}" "${node_key}" 2>/dev/null || true
    chmod 600 "${ca_key}" 2>/dev/null || true
  fi
}

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
  export elasticsearch_scheme="https"
  if [[ $TEST_SUITE != "platinum" ]]; then
    export elasticsearch_scheme="http"
  fi
  export elasticsearch_url=${elasticsearch_scheme}://elastic:${elastic_password}@${es_node_name}:9200
  export external_elasticsearch_url=${elasticsearch_url/$es_node_name/localhost}
  export elasticsearch_container="${elasticsearch_image}:${STACK_VERSION}"

  export suffix=rest-test
  moniker=$(echo "$elasticsearch_container" | tr -C "[:alnum:]" '-')
  export moniker
  export network_name=${moniker}${suffix}

  export ssl_cert="${script_path}/certs/testnode.crt"
  export ssl_key="${script_path}/certs/testnode.key"
  export ssl_ca="${script_path}/certs/ca.crt"

fi

# Ensure certs exist and are valid when running the platinum (TLS) suite.
if [[ "${TEST_SUITE}" == "platinum" ]]; then
  ensure_tls_certs
fi
# shellcheck source=.buildkite/functions/cleanup.sh
source "$script_path"/functions/cleanup.sh
# shellcheck source=.buildkite/functions/wait-for-container.sh
source "$script_path"/functions/wait-for-container.sh
trap 'cleanup_trap ${network_name}' EXIT


if [[ "$CLEANUP" == "true" ]]; then
  cleanup_all_in_network "$network_name"
  exit 0
fi

echo -e "\033[34;1mINFO:\033[0m Creating network $network_name if it does not exist already \033[0m"
docker network inspect "$network_name" > /dev/null 2>&1 || docker network create "$network_name"

