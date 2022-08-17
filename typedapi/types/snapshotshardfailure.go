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

// SnapshotShardFailure type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/snapshot/_types/SnapshotShardFailure.ts#L22-L28
type SnapshotShardFailure struct {
	Index   IndexName `json:"index"`
	NodeId  Id        `json:"node_id"`
	Reason  string    `json:"reason"`
	ShardId Id        `json:"shard_id"`
	Status  string    `json:"status"`
}

// SnapshotShardFailureBuilder holds SnapshotShardFailure struct and provides a builder API.
type SnapshotShardFailureBuilder struct {
	v *SnapshotShardFailure
}

// NewSnapshotShardFailure provides a builder for the SnapshotShardFailure struct.
func NewSnapshotShardFailureBuilder() *SnapshotShardFailureBuilder {
	r := SnapshotShardFailureBuilder{
		&SnapshotShardFailure{},
	}

	return &r
}

// Build finalize the chain and returns the SnapshotShardFailure struct
func (rb *SnapshotShardFailureBuilder) Build() SnapshotShardFailure {
	return *rb.v
}

func (rb *SnapshotShardFailureBuilder) Index(index IndexName) *SnapshotShardFailureBuilder {
	rb.v.Index = index
	return rb
}

func (rb *SnapshotShardFailureBuilder) NodeId(nodeid Id) *SnapshotShardFailureBuilder {
	rb.v.NodeId = nodeid
	return rb
}

func (rb *SnapshotShardFailureBuilder) Reason(reason string) *SnapshotShardFailureBuilder {
	rb.v.Reason = reason
	return rb
}

func (rb *SnapshotShardFailureBuilder) ShardId(shardid Id) *SnapshotShardFailureBuilder {
	rb.v.ShardId = shardid
	return rb
}

func (rb *SnapshotShardFailureBuilder) Status(status string) *SnapshotShardFailureBuilder {
	rb.v.Status = status
	return rb
}
