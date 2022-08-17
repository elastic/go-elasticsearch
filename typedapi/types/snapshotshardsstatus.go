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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/shardsstatsstage"
)

// SnapshotShardsStatus type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/snapshot/_types/SnapshotShardsStatus.ts#L24-L27
type SnapshotShardsStatus struct {
	Stage shardsstatsstage.ShardsStatsStage `json:"stage"`
	Stats ShardsStatsSummary                `json:"stats"`
}

// SnapshotShardsStatusBuilder holds SnapshotShardsStatus struct and provides a builder API.
type SnapshotShardsStatusBuilder struct {
	v *SnapshotShardsStatus
}

// NewSnapshotShardsStatus provides a builder for the SnapshotShardsStatus struct.
func NewSnapshotShardsStatusBuilder() *SnapshotShardsStatusBuilder {
	r := SnapshotShardsStatusBuilder{
		&SnapshotShardsStatus{},
	}

	return &r
}

// Build finalize the chain and returns the SnapshotShardsStatus struct
func (rb *SnapshotShardsStatusBuilder) Build() SnapshotShardsStatus {
	return *rb.v
}

func (rb *SnapshotShardsStatusBuilder) Stage(stage shardsstatsstage.ShardsStatsStage) *SnapshotShardsStatusBuilder {
	rb.v.Stage = stage
	return rb
}

func (rb *SnapshotShardsStatusBuilder) Stats(stats *ShardsStatsSummaryBuilder) *SnapshotShardsStatusBuilder {
	v := stats.Build()
	rb.v.Stats = v
	return rb
}
