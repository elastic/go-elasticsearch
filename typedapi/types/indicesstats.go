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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/indexmetadatastate"
)

// IndicesStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/stats/types.ts#L89-L98
type IndicesStats struct {
	Health    *healthstatus.HealthStatus             `json:"health,omitempty"`
	Primaries *IndexStats                            `json:"primaries,omitempty"`
	Shards    map[string][]ShardStats                `json:"shards,omitempty"`
	Status    *indexmetadatastate.IndexMetadataState `json:"status,omitempty"`
	Total     *IndexStats                            `json:"total,omitempty"`
	Uuid      *Uuid                                  `json:"uuid,omitempty"`
}

// IndicesStatsBuilder holds IndicesStats struct and provides a builder API.
type IndicesStatsBuilder struct {
	v *IndicesStats
}

// NewIndicesStats provides a builder for the IndicesStats struct.
func NewIndicesStatsBuilder() *IndicesStatsBuilder {
	r := IndicesStatsBuilder{
		&IndicesStats{
			Shards: make(map[string][]ShardStats, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the IndicesStats struct
func (rb *IndicesStatsBuilder) Build() IndicesStats {
	return *rb.v
}

func (rb *IndicesStatsBuilder) Health(health healthstatus.HealthStatus) *IndicesStatsBuilder {
	rb.v.Health = &health
	return rb
}

func (rb *IndicesStatsBuilder) Primaries(primaries *IndexStatsBuilder) *IndicesStatsBuilder {
	v := primaries.Build()
	rb.v.Primaries = &v
	return rb
}

func (rb *IndicesStatsBuilder) Shards(value map[string][]ShardStats) *IndicesStatsBuilder {
	rb.v.Shards = value
	return rb
}

func (rb *IndicesStatsBuilder) Status(status indexmetadatastate.IndexMetadataState) *IndicesStatsBuilder {
	rb.v.Status = &status
	return rb
}

func (rb *IndicesStatsBuilder) Total(total *IndexStatsBuilder) *IndicesStatsBuilder {
	v := total.Build()
	rb.v.Total = &v
	return rb
}

func (rb *IndicesStatsBuilder) Uuid(uuid Uuid) *IndicesStatsBuilder {
	rb.v.Uuid = &uuid
	return rb
}
