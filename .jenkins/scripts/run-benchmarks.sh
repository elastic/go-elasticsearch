#!/usr/bin/env bash
#
# Licensed to Elasticsearch B.V. under one or more agreements.
# Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
# See the LICENSE file in the project root for more information.
#
# Runner for the client benchmarks.
#
# Export the following environment variables when running the script:
#
# * CLIENT_IMAGE
# * CLIENT_BRANCH
# * CLIENT_COMMAND
#
# The WORKSPACE environment variable is automatically set on Jenkins.
#
# The VAULT_ADDR is automatically set on Jenkins.

set -euxo pipefail

mkdir -p "$WORKSPACE/tmp"

vault read -field=gcp_service_account secret/clients-team/benchmarking > "$WORKSPACE/tmp/credentials.json"

export GOOGLE_CLOUD_KEYFILE_JSON="$WORKSPACE/tmp/credentials.json"
export GOOGLE_APPLICATION_CREDENTIALS="$WORKSPACE/tmp/credentials.json"

set +x
report_url="$(vault read -field=reporting_url secret/clients-team/benchmarking)"
report_password="$(vault read -field=reporting_password secret/clients-team/benchmarking)"
export ELASTICSEARCH_REPORT_URL="$report_url"
export ELASTICSEARCH_REPORT_PASSWORD="$report_password"

export TF_VAR_reporting_url="$ELASTICSEARCH_REPORT_URL"
export TF_VAR_reporting_password="$ELASTICSEARCH_REPORT_PASSWORD"
set -x

TERRAFORM=$(command -v terraform)
GCLOUD=$(command -v gcloud)

DESTROY=${DESTROY:-yes}

function cleanup {
  if [[ $DESTROY != "no" ]]; then
    $TERRAFORM destroy --auto-approve --input=false --var client_image="$CLIENT_IMAGE"
  fi
}

trap cleanup INT TERM EXIT;

if [[ ! -d "$WORKSPACE/tmp/elasticsearch-clients-benchmarks" ]]; then
git clone --depth 1 git@github.com:elastic/elasticsearch-clients-benchmarks.git "$WORKSPACE/tmp/elasticsearch-clients-benchmarks"
fi

cd "$WORKSPACE/tmp/elasticsearch-clients-benchmarks/terraform/gcp"

$TERRAFORM init --input=false
$TERRAFORM apply --auto-approve --input=false --var client_image="$CLIENT_IMAGE"

set +x
echo -e "\033[1mWaiting for [$($TERRAFORM output runner_instance_name)] to be ready...\033[0m"
until $GCLOUD compute --project 'elastic-clients' ssh "$($TERRAFORM output runner_instance_name)" --zone='europe-west1-b' --command="sudo su - runner -c 'docker container inspect benchmarks-data'" &> /dev/null; do sleep 1; done;

echo -e "\nRunning benchmarks for [$CLIENT_IMAGE]\n"
set -x

$GCLOUD compute --project 'elastic-clients' ssh "$($TERRAFORM output runner_instance_name)" \
  --zone='europe-west1-b' \
  --ssh-flag='-t' \
  --command="\
  CLIENT_BRANCH=$CLIENT_BRANCH \
  CLIENT_BENCHMARK_ENVIRONMENT=production \
  /home/runner/runner.sh \"$CLIENT_COMMAND\""
