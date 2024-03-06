// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

//go:build integration && !multinode
// +build integration,!multinode

package containertest

import (
	"context"
	"os"
	"testing"

	"github.com/elastic/go-elasticsearch/v8/internal/version"

	"github.com/testcontainers/testcontainers-go"
	tces "github.com/testcontainers/testcontainers-go/modules/elasticsearch"
)

// SetupElasticsearch starts an Elasticsearch container and sets the environment variables
// ELASTICSEARCH_URL, ELASTICSEARCH_USERNAME, and ELASTICSEARCH_PASSWORD in the test context.
func SetupElasticsearch(t *testing.T) (*tces.ElasticsearchContainer, error) {
	stackVersion := version.Client
	if v := os.Getenv("STACK_VERSION"); v != "" {
		stackVersion = v
	}

	elasticsearchContainer, err := tces.RunContainer(
		context.Background(),
		testcontainers.WithImage("docker.elastic.co/elasticsearch/elasticsearch:"+stackVersion),
		tces.WithPassword("changeme"),
	)
	if err != nil {
		return nil, err
	}

	t.Setenv("ELASTICSEARCH_URL", elasticsearchContainer.Settings.Address)
	t.Setenv("ELASTICSEARCH_USERNAME", elasticsearchContainer.Settings.Username)
	t.Setenv("ELASTICSEARCH_PASSWORD", elasticsearchContainer.Settings.Password)

	return elasticsearchContainer, nil
}
