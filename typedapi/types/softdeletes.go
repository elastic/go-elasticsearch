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

// SoftDeletes type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/_types/IndexSettings.ts#L50-L63
type SoftDeletes struct {
	// Enabled Indicates whether soft deletes are enabled on the index.
	Enabled *bool `json:"enabled,omitempty"`
	// RetentionLease The maximum period to retain a shard history retention lease before it is
	// considered expired.
	// Shard history retention leases ensure that soft deletes are retained during
	// merges on the Lucene
	// index. If a soft delete is merged away before it can be replicated to a
	// follower the following
	// process will fail due to incomplete history on the leader.
	RetentionLease *RetentionLease `json:"retention_lease,omitempty"`
}

// SoftDeletesBuilder holds SoftDeletes struct and provides a builder API.
type SoftDeletesBuilder struct {
	v *SoftDeletes
}

// NewSoftDeletes provides a builder for the SoftDeletes struct.
func NewSoftDeletesBuilder() *SoftDeletesBuilder {
	r := SoftDeletesBuilder{
		&SoftDeletes{},
	}

	return &r
}

// Build finalize the chain and returns the SoftDeletes struct
func (rb *SoftDeletesBuilder) Build() SoftDeletes {
	return *rb.v
}

// Enabled Indicates whether soft deletes are enabled on the index.

func (rb *SoftDeletesBuilder) Enabled(enabled bool) *SoftDeletesBuilder {
	rb.v.Enabled = &enabled
	return rb
}

// RetentionLease The maximum period to retain a shard history retention lease before it is
// considered expired.
// Shard history retention leases ensure that soft deletes are retained during
// merges on the Lucene
// index. If a soft delete is merged away before it can be replicated to a
// follower the following
// process will fail due to incomplete history on the leader.

func (rb *SoftDeletesBuilder) RetentionLease(retentionlease *RetentionLeaseBuilder) *SoftDeletesBuilder {
	v := retentionlease.Build()
	rb.v.RetentionLease = &v
	return rb
}
