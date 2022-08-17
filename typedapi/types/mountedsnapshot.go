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

// MountedSnapshot type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/searchable_snapshots/mount/types.ts#L23-L27
type MountedSnapshot struct {
	Indices  Indices         `json:"indices"`
	Shards   ShardStatistics `json:"shards"`
	Snapshot Name            `json:"snapshot"`
}

// MountedSnapshotBuilder holds MountedSnapshot struct and provides a builder API.
type MountedSnapshotBuilder struct {
	v *MountedSnapshot
}

// NewMountedSnapshot provides a builder for the MountedSnapshot struct.
func NewMountedSnapshotBuilder() *MountedSnapshotBuilder {
	r := MountedSnapshotBuilder{
		&MountedSnapshot{},
	}

	return &r
}

// Build finalize the chain and returns the MountedSnapshot struct
func (rb *MountedSnapshotBuilder) Build() MountedSnapshot {
	return *rb.v
}

func (rb *MountedSnapshotBuilder) Indices(indices *IndicesBuilder) *MountedSnapshotBuilder {
	v := indices.Build()
	rb.v.Indices = v
	return rb
}

func (rb *MountedSnapshotBuilder) Shards(shards *ShardStatisticsBuilder) *MountedSnapshotBuilder {
	v := shards.Build()
	rb.v.Shards = v
	return rb
}

func (rb *MountedSnapshotBuilder) Snapshot(snapshot Name) *MountedSnapshotBuilder {
	rb.v.Snapshot = snapshot
	return rb
}
