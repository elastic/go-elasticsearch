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

	"github.com/elastic/go-elasticsearch/v8"

	"github.com/testcontainers/testcontainers-go"
	tces "github.com/testcontainers/testcontainers-go/modules/elasticsearch"
)

// ElasticsearchService represents an Elasticsearch service, storing
// the container and configuration options.
type ElasticsearchService struct {
	container     *tces.ElasticsearchContainer
	configOptions tces.Options
}

// ESConfig returns the Elasticsearch client configuration.
func (es *ElasticsearchService) ESConfig() elasticsearch.Config {
	return elasticsearch.Config{
		Addresses: []string{
			es.configOptions.Address,
		},
		Username: "elastic",
		Password: es.configOptions.Password,
		CACert:   es.configOptions.CACert,
	}
}

// Terminate terminates the Elasticsearch container.
func (es *ElasticsearchService) Terminate(ctx context.Context) error {
	return es.container.Terminate(ctx)
}

func NewElasticsearchService(stackVersion string) (*ElasticsearchService, error) {
	elasticsearchContainer, err := tces.RunContainer(
		context.Background(),
		testcontainers.WithImage("docker.elastic.co/elasticsearch/elasticsearch:"+stackVersion),
		tces.WithPassword("changeme"),
	)
	if err != nil {
		return nil, err
	}

	return &ElasticsearchService{
		container:     elasticsearchContainer,
		configOptions: elasticsearchContainer.Settings,
	}, nil
}
