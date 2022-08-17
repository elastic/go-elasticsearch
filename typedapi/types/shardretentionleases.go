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

// ShardRetentionLeases type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/stats/types.ts#L144-L148
type ShardRetentionLeases struct {
	Leases      []ShardLease  `json:"leases"`
	PrimaryTerm int64         `json:"primary_term"`
	Version     VersionNumber `json:"version"`
}

// ShardRetentionLeasesBuilder holds ShardRetentionLeases struct and provides a builder API.
type ShardRetentionLeasesBuilder struct {
	v *ShardRetentionLeases
}

// NewShardRetentionLeases provides a builder for the ShardRetentionLeases struct.
func NewShardRetentionLeasesBuilder() *ShardRetentionLeasesBuilder {
	r := ShardRetentionLeasesBuilder{
		&ShardRetentionLeases{},
	}

	return &r
}

// Build finalize the chain and returns the ShardRetentionLeases struct
func (rb *ShardRetentionLeasesBuilder) Build() ShardRetentionLeases {
	return *rb.v
}

func (rb *ShardRetentionLeasesBuilder) Leases(leases []ShardLeaseBuilder) *ShardRetentionLeasesBuilder {
	tmp := make([]ShardLease, len(leases))
	for _, value := range leases {
		tmp = append(tmp, value.Build())
	}
	rb.v.Leases = tmp
	return rb
}

func (rb *ShardRetentionLeasesBuilder) PrimaryTerm(primaryterm int64) *ShardRetentionLeasesBuilder {
	rb.v.PrimaryTerm = primaryterm
	return rb
}

func (rb *ShardRetentionLeasesBuilder) Version(version VersionNumber) *ShardRetentionLeasesBuilder {
	rb.v.Version = version
	return rb
}
