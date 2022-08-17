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

// IndicesVersions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cluster/stats/types.ts#L137-L142
type IndicesVersions struct {
	IndexCount        int           `json:"index_count"`
	PrimaryShardCount int           `json:"primary_shard_count"`
	TotalPrimaryBytes int64         `json:"total_primary_bytes"`
	Version           VersionString `json:"version"`
}

// IndicesVersionsBuilder holds IndicesVersions struct and provides a builder API.
type IndicesVersionsBuilder struct {
	v *IndicesVersions
}

// NewIndicesVersions provides a builder for the IndicesVersions struct.
func NewIndicesVersionsBuilder() *IndicesVersionsBuilder {
	r := IndicesVersionsBuilder{
		&IndicesVersions{},
	}

	return &r
}

// Build finalize the chain and returns the IndicesVersions struct
func (rb *IndicesVersionsBuilder) Build() IndicesVersions {
	return *rb.v
}

func (rb *IndicesVersionsBuilder) IndexCount(indexcount int) *IndicesVersionsBuilder {
	rb.v.IndexCount = indexcount
	return rb
}

func (rb *IndicesVersionsBuilder) PrimaryShardCount(primaryshardcount int) *IndicesVersionsBuilder {
	rb.v.PrimaryShardCount = primaryshardcount
	return rb
}

func (rb *IndicesVersionsBuilder) TotalPrimaryBytes(totalprimarybytes int64) *IndicesVersionsBuilder {
	rb.v.TotalPrimaryBytes = totalprimarybytes
	return rb
}

func (rb *IndicesVersionsBuilder) Version(version VersionString) *IndicesVersionsBuilder {
	rb.v.Version = version
	return rb
}
