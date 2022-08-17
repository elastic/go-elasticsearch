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

// Source type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/transform/_types/Transform.ts#L135-L153
type Source struct {
	// Index The source indices for the transform. It can be a single index, an index
	// pattern (for example, `"my-index-*""`), an
	// array of indices (for example, `["my-index-000001", "my-index-000002"]`), or
	// an array of index patterns (for
	// example, `["my-index-*", "my-other-index-*"]`. For remote indices use the
	// syntax `"remote_name:index_name"`. If
	// any indices are in remote clusters then the master node and at least one
	// transform node must have the `remote_cluster_client` node role.
	Index Indices `json:"index"`
	// Query A query clause that retrieves a subset of data from the source index.
	Query *QueryContainer `json:"query,omitempty"`
	// RuntimeMappings Definitions of search-time runtime fields that can be used by the transform.
	// For search runtime fields all data
	// nodes, including remote nodes, must be 7.12 or later.
	RuntimeMappings *RuntimeFields `json:"runtime_mappings,omitempty"`
}

// SourceBuilder holds Source struct and provides a builder API.
type SourceBuilder struct {
	v *Source
}

// NewSource provides a builder for the Source struct.
func NewSourceBuilder() *SourceBuilder {
	r := SourceBuilder{
		&Source{},
	}

	return &r
}

// Build finalize the chain and returns the Source struct
func (rb *SourceBuilder) Build() Source {
	return *rb.v
}

// Index The source indices for the transform. It can be a single index, an index
// pattern (for example, `"my-index-*""`), an
// array of indices (for example, `["my-index-000001", "my-index-000002"]`), or
// an array of index patterns (for
// example, `["my-index-*", "my-other-index-*"]`. For remote indices use the
// syntax `"remote_name:index_name"`. If
// any indices are in remote clusters then the master node and at least one
// transform node must have the `remote_cluster_client` node role.

func (rb *SourceBuilder) Index(index *IndicesBuilder) *SourceBuilder {
	v := index.Build()
	rb.v.Index = v
	return rb
}

// Query A query clause that retrieves a subset of data from the source index.

func (rb *SourceBuilder) Query(query *QueryContainerBuilder) *SourceBuilder {
	v := query.Build()
	rb.v.Query = &v
	return rb
}

// RuntimeMappings Definitions of search-time runtime fields that can be used by the transform.
// For search runtime fields all data
// nodes, including remote nodes, must be 7.12 or later.

func (rb *SourceBuilder) RuntimeMappings(runtimemappings *RuntimeFieldsBuilder) *SourceBuilder {
	v := runtimemappings.Build()
	rb.v.RuntimeMappings = &v
	return rb
}
