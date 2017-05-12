#
# Licensed to Elasticsearch under one or more contributor
# license agreements. See the NOTICE file distributed with
# this work for additional information regarding copyright
# ownership. Elasticsearch licenses this file to you under
# the Apache License, Version 2.0 (the "License"); you may
# not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.
#

MAKE = make
GO = go
GO_PACKAGES=$(shell go list ./... | grep -v /vendor/)
GOIMPORTS = goimports
GOIMPORTS_REPO = golang.org/x/tools/cmd/goimports
GOLINT = golint
GOLINT_REPO = github.com/golang/lint/$(GOLINT)
SPEC_DIR = spec
ELASTICSEARCH_DIR = $(SPEC_DIR)/elasticsearch
ELASTICSEARCH_SPEC_DIR = $(ELASTICSEARCH_DIR)/rest-api-spec/src/main/resources/rest-api-spec
GEN_DIR = api
VERSION_FILE = $(GEN_DIR)/VERSION
DOCKER = docker
DOCKER_IMAGE = docker.elastic.co/elasticsearch/elasticsearch
DOCKER_CONTAINER = goelasticsearch-elasticsearch
TESTDATA = generator/testdata

.PHONY: build
build:
	$(GO) build $(GO_PACKAGES)

.PHONY: install
install:
	$(GO) install $(GO_PACKAGES)

.PHONY: gen
gen: spec
	rm -rf $(GEN_DIR)
	$(GO) run ./cmd/generator/main.go -alsologtostderr
	@$(GO) get $(GOIMPORTS_REPO)
	$(GOIMPORTS) -w $(GEN_DIR)
	(cd $(ELASTICSEARCH_DIR) && git describe --tags) > $(VERSION_FILE)
	$(MAKE) build

.PHONY: gen-testdata
gen-testdata: spec
	$(foreach file,$(wildcard $(TESTDATA)/*.json),cp -v $(ELASTICSEARCH_SPEC_DIR)/api/$(notdir $(file)) $(TESTDATA);)

.PHONY: spec
spec:
	$(MAKE) -C $(SPEC_DIR) checkout
	ln -sf $(ELASTICSEARCH_SPEC_DIR)

.PHONY: spec-update
spec-update: spec
	$(MAKE) -C $(SPEC_DIR) pull

.PHONY: docker-start
docker-start:
	@$(MAKE) docker-stop
	$(DOCKER) run --name $(DOCKER_CONTAINER) -d -p 127.0.0.1:9200:9200 -e "node.attr.testattr=test" -e "path.repo=/tmp" -e "http.host=0.0.0.0" -e "transport.host=127.0.0.1" -e "xpack.security.enabled=false" $(DOCKER_IMAGE):$(shell cat ${VERSION_FILE} | cut -c2-)
	echo "Waiting for elasticsearch to be reachable..." && time sh -c "until curl -sf http://localhost:9200; do sleep 1; done"

.PHONY: docker-stop
docker-stop:
	$(DOCKER) stop $(DOCKER_CONTAINER) &&	$(DOCKER) rm $(DOCKER_CONTAINER) || true

.PHONY: test
test: build
	$(GO) test $(GO_PACKAGES)

.PHONY: check
check:
	$(MAKE) docker-start
	$(MAKE) test
	$(MAKE) docker-stop
	$(MAKE) lint check-imports

.PHONY: check-imports
check-imports:
	@$(GO) get $(GOIMPORTS_REPO)
	@if [[ -n `${GOIMPORTS} -l .` ]]; then echo "goimports failed, please run make imports" && exit 1; fi

.PHONY: imports
imports:
	@$(GO) get $(GOIMPORTS_REPO)
	$(GOIMPORTS) -w $(GEN_DIR)

.PHONY: lint
lint:
	@$(GO) get $(GOLINT_REPO)
	$(GOLINT) $(GO_PACKAGES)

.PHONY: gen-clean
gen-clean:
	make -C $(SPEC_DIR) clean
	rm $(ELASTICSEARCH_SPEC_DIR)

.PHONY: clean
clean:
	$(GO) clean $(GO_PACKAGES)
