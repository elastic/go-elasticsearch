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

GO = go
GO_PACKAGES=$(shell go list ./... | grep -v /vendor/)
GOLINT = golint
GOLINT_REPO = github.com/golang/lint/$(GOLINT)
SPEC_DIR = spec

.PHONY: build
build:
	$(GO) build $(GO_PACKAGES)

.PHONY: spec
spec:
	make -C $(SPEC_DIR) checkout

.PHONY: check
check: spec
	$(GO) test $(GO_PACKAGES)

.PHONY: lint
lint:
	$(GO) get $(GOLINT_REPO)
	$(GOLINT) $(GO_PACKAGES)

.PHONY: clean
clean:
	$(GO) clean $(GO_PACKAGES)
	make -C $(SPEC_DIR) clean
