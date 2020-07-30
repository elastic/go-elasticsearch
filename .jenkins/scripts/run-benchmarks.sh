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

runner_ssh_private_key="$(vault read -field=runner_ssh_private_key secret/clients-team/benchmarking)"
runner_ssh_public_key="$(vault read -field=runner_ssh_public_key secret/clients-team/benchmarking)"
export TF_VAR_runner_ssh_private_key="$runner_ssh_private_key"
export TF_VAR_runner_ssh_public_key="$runner_ssh_public_key"
echo "$runner_ssh_private_key" > "$WORKSPACE/tmp/runner_id_rsa"
echo "$runner_ssh_public_key" > "$WORKSPACE/tmp/runner_id_rsa.pub"
chmod go= "$WORKSPACE/tmp/runner_id_rsa"
chmod go= "$WORKSPACE/tmp/runner_id_rsa.pub"
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
else
  cd "$WORKSPACE/tmp/elasticsearch-clients-benchmarks" && git fetch --quiet && git reset origin/master --hard
fi

cd "$WORKSPACE/tmp/elasticsearch-clients-benchmarks/terraform/gcp"

$TERRAFORM init --input=false
$TERRAFORM apply --auto-approve --input=false --var client_image="$CLIENT_IMAGE"

set +ex
echo -e "\n\033[1mWaiting for instance [$($TERRAFORM output runner_instance_name)] to be ready...\033[0m"

SECONDS=0
while (( SECONDS < 900 )); do # Timeout: 15min
  $GCLOUD compute --project 'elastic-clients' ssh runner@"$($TERRAFORM output runner_instance_name)" --zone='europe-west1-b' --ssh-key-file="$WORKSPACE/tmp/runner_id_rsa" --command="sudo su - runner -c 'docker container inspect -f \"{{.Name}}:{{.State.Status}}\" benchmarks-data'"
  status=$?
  if [[ $status -eq 0 ]]; then
    break
  else
    sleep 1
  fi
done

echo -e "\n\033[1mWaiting for cluster at [$($TERRAFORM output master_ip)] to be ready...\033[0m"

SECONDS=0
while (( SECONDS < 900 )); do # Timeout: 15min
  $GCLOUD compute --project 'elastic-clients' ssh runner@"$($TERRAFORM output runner_instance_name)" --zone='europe-west1-b' --ssh-key-file="$WORKSPACE/tmp/runner_id_rsa" --command="sudo su - runner -c 'curl -sS $($TERRAFORM output master_ip):9200/_cat/nodes?v'"
  status=$?
  if [[ $status -eq 0 ]]; then
    break
  else
    sleep 1
  fi
done

echo -e "\n\033[1mRunning benchmarks for [$CLIENT_IMAGE]\033[0m"
set -x

$GCLOUD compute --project 'elastic-clients' ssh runner@"$($TERRAFORM output runner_instance_name)" \
  --zone='europe-west1-b' \
  --ssh-key-file="$WORKSPACE/tmp/runner_id_rsa" \
  --ssh-flag='-t' \
  --command="\
  CLIENT_BRANCH=$CLIENT_BRANCH \
  CLIENT_BENCHMARK_ENVIRONMENT=production \
  /home/runner/runner.sh \"$CLIENT_COMMAND\""
