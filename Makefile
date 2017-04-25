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

.PHONY: build
build:
	$(GO) build $(GO_PACKAGES)

.PHONY: gen
gen: spec
	$(GO) run ./cmd/generator/main.go -alsologtostderr
	@$(GO) get $(GOIMPORTS_REPO)
	$(GOIMPORTS) -w client
	$(MAKE) build

.PHONY: spec
spec:
	$(MAKE) -C $(SPEC_DIR) checkout

.PHONY: spec-update
spec-update: spec
	$(MAKE) -C $(SPEC_DIR) pull

.PHONY: test
test: build
	$(GO) test $(GO_PACKAGES)

.PHONY: check
check: test
	$(MAKE) lint check-imports

.PHONY: check-imports
check-imports:
	@$(GO) get $(GOIMPORTS_REPO)
	@if [[ -n `${GOIMPORTS} -l .` ]]; then echo "goimports failed, please run make imports" && exit 1; fi

.PHONY: imports
imports:
	@$(GO) get $(GOIMPORTS_REPO)
	$(GOIMPORTS) -w .

.PHONY: lint
lint:
	@$(GO) get $(GOLINT_REPO)
	$(GOLINT) $(GO_PACKAGES)

.PHONY: clean
clean:
	$(GO) clean $(GO_PACKAGES)
	make -C $(SPEC_DIR) clean
