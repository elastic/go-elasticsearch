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


// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/healthstatus"
)

// IndexHealthStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cluster/health/types.ts#L24-L34
type IndexHealthStats struct {
	ActivePrimaryShards int                         `json:"active_primary_shards"`
	ActiveShards        int                         `json:"active_shards"`
	InitializingShards  int                         `json:"initializing_shards"`
	NumberOfReplicas    int                         `json:"number_of_replicas"`
	NumberOfShards      int                         `json:"number_of_shards"`
	RelocatingShards    int                         `json:"relocating_shards"`
	Shards              map[string]ShardHealthStats `json:"shards,omitempty"`
	Status              healthstatus.HealthStatus   `json:"status"`
	UnassignedShards    int                         `json:"unassigned_shards"`
}

// IndexHealthStatsBuilder holds IndexHealthStats struct and provides a builder API.
type IndexHealthStatsBuilder struct {
	v *IndexHealthStats
}

// NewIndexHealthStats provides a builder for the IndexHealthStats struct.
func NewIndexHealthStatsBuilder() *IndexHealthStatsBuilder {
	r := IndexHealthStatsBuilder{
		&IndexHealthStats{
			Shards: make(map[string]ShardHealthStats, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the IndexHealthStats struct
func (rb *IndexHealthStatsBuilder) Build() IndexHealthStats {
	return *rb.v
}

func (rb *IndexHealthStatsBuilder) ActivePrimaryShards(activeprimaryshards int) *IndexHealthStatsBuilder {
	rb.v.ActivePrimaryShards = activeprimaryshards
	return rb
}

func (rb *IndexHealthStatsBuilder) ActiveShards(activeshards int) *IndexHealthStatsBuilder {
	rb.v.ActiveShards = activeshards
	return rb
}

func (rb *IndexHealthStatsBuilder) InitializingShards(initializingshards int) *IndexHealthStatsBuilder {
	rb.v.InitializingShards = initializingshards
	return rb
}

func (rb *IndexHealthStatsBuilder) NumberOfReplicas(numberofreplicas int) *IndexHealthStatsBuilder {
	rb.v.NumberOfReplicas = numberofreplicas
	return rb
}

func (rb *IndexHealthStatsBuilder) NumberOfShards(numberofshards int) *IndexHealthStatsBuilder {
	rb.v.NumberOfShards = numberofshards
	return rb
}

func (rb *IndexHealthStatsBuilder) RelocatingShards(relocatingshards int) *IndexHealthStatsBuilder {
	rb.v.RelocatingShards = relocatingshards
	return rb
}

func (rb *IndexHealthStatsBuilder) Shards(values map[string]*ShardHealthStatsBuilder) *IndexHealthStatsBuilder {
	tmp := make(map[string]ShardHealthStats, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Shards = tmp
	return rb
}

func (rb *IndexHealthStatsBuilder) Status(status healthstatus.HealthStatus) *IndexHealthStatsBuilder {
	rb.v.Status = status
	return rb
}

func (rb *IndexHealthStatsBuilder) UnassignedShards(unassignedshards int) *IndexHealthStatsBuilder {
	rb.v.UnassignedShards = unassignedshards
	return rb
}
