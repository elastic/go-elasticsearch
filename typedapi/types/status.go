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

// Status type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/snapshot/_types/SnapshotStatus.ts#L26-L35
type Status struct {
	IncludeGlobalState bool                          `json:"include_global_state"`
	Indices            map[string]SnapshotIndexStats `json:"indices"`
	Repository         string                        `json:"repository"`
	ShardsStats        ShardsStats                   `json:"shards_stats"`
	Snapshot           string                        `json:"snapshot"`
	State              string                        `json:"state"`
	Stats              SnapshotStats                 `json:"stats"`
	Uuid               Uuid                          `json:"uuid"`
}

// StatusBuilder holds Status struct and provides a builder API.
type StatusBuilder struct {
	v *Status
}

// NewStatus provides a builder for the Status struct.
func NewStatusBuilder() *StatusBuilder {
	r := StatusBuilder{
		&Status{
			Indices: make(map[string]SnapshotIndexStats, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the Status struct
func (rb *StatusBuilder) Build() Status {
	return *rb.v
}

func (rb *StatusBuilder) IncludeGlobalState(includeglobalstate bool) *StatusBuilder {
	rb.v.IncludeGlobalState = includeglobalstate
	return rb
}

func (rb *StatusBuilder) Indices(values map[string]*SnapshotIndexStatsBuilder) *StatusBuilder {
	tmp := make(map[string]SnapshotIndexStats, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Indices = tmp
	return rb
}

func (rb *StatusBuilder) Repository(repository string) *StatusBuilder {
	rb.v.Repository = repository
	return rb
}

func (rb *StatusBuilder) ShardsStats(shardsstats *ShardsStatsBuilder) *StatusBuilder {
	v := shardsstats.Build()
	rb.v.ShardsStats = v
	return rb
}

func (rb *StatusBuilder) Snapshot(snapshot string) *StatusBuilder {
	rb.v.Snapshot = snapshot
	return rb
}

func (rb *StatusBuilder) State(state string) *StatusBuilder {
	rb.v.State = state
	return rb
}

func (rb *StatusBuilder) Stats(stats *SnapshotStatsBuilder) *StatusBuilder {
	v := stats.Build()
	rb.v.Stats = v
	return rb
}

func (rb *StatusBuilder) Uuid(uuid Uuid) *StatusBuilder {
	rb.v.Uuid = uuid
	return rb
}
