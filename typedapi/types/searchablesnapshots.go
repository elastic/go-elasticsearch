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

// SearchableSnapshots type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/xpack/usage/types.ts#L417-L421
type SearchableSnapshots struct {
	Available               bool `json:"available"`
	Enabled                 bool `json:"enabled"`
	FullCopyIndicesCount    *int `json:"full_copy_indices_count,omitempty"`
	IndicesCount            int  `json:"indices_count"`
	SharedCacheIndicesCount *int `json:"shared_cache_indices_count,omitempty"`
}

// SearchableSnapshotsBuilder holds SearchableSnapshots struct and provides a builder API.
type SearchableSnapshotsBuilder struct {
	v *SearchableSnapshots
}

// NewSearchableSnapshots provides a builder for the SearchableSnapshots struct.
func NewSearchableSnapshotsBuilder() *SearchableSnapshotsBuilder {
	r := SearchableSnapshotsBuilder{
		&SearchableSnapshots{},
	}

	return &r
}

// Build finalize the chain and returns the SearchableSnapshots struct
func (rb *SearchableSnapshotsBuilder) Build() SearchableSnapshots {
	return *rb.v
}

func (rb *SearchableSnapshotsBuilder) Available(available bool) *SearchableSnapshotsBuilder {
	rb.v.Available = available
	return rb
}

func (rb *SearchableSnapshotsBuilder) Enabled(enabled bool) *SearchableSnapshotsBuilder {
	rb.v.Enabled = enabled
	return rb
}

func (rb *SearchableSnapshotsBuilder) FullCopyIndicesCount(fullcopyindicescount int) *SearchableSnapshotsBuilder {
	rb.v.FullCopyIndicesCount = &fullcopyindicescount
	return rb
}

func (rb *SearchableSnapshotsBuilder) IndicesCount(indicescount int) *SearchableSnapshotsBuilder {
	rb.v.IndicesCount = indicescount
	return rb
}

func (rb *SearchableSnapshotsBuilder) SharedCacheIndicesCount(sharedcacheindicescount int) *SearchableSnapshotsBuilder {
	rb.v.SharedCacheIndicesCount = &sharedcacheindicescount
	return rb
}
