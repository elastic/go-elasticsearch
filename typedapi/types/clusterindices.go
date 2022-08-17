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

// ClusterIndices type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cluster/stats/types.ts#L63-L94
type ClusterIndices struct {
	// Analysis Contains statistics about analyzers and analyzer components used in selected
	// nodes.
	Analysis CharFilterTypes `json:"analysis"`
	// Completion Contains statistics about memory used for completion in selected nodes.
	Completion CompletionStats `json:"completion"`
	// Count Total number of indices with shards assigned to selected nodes.
	Count int64 `json:"count"`
	// Docs Contains counts for documents in selected nodes.
	Docs DocStats `json:"docs"`
	// Fielddata Contains statistics about the field data cache of selected nodes.
	Fielddata FielddataStats `json:"fielddata"`
	// Mappings Contains statistics about field mappings in selected nodes.
	Mappings FieldTypesMappings `json:"mappings"`
	// QueryCache Contains statistics about the query cache of selected nodes.
	QueryCache QueryCacheStats `json:"query_cache"`
	// Segments Contains statistics about segments in selected nodes.
	Segments SegmentsStats `json:"segments"`
	// Shards Contains statistics about indices with shards assigned to selected nodes.
	Shards ClusterIndicesShards `json:"shards"`
	// Store Contains statistics about the size of shards assigned to selected nodes.
	Store    StoreStats        `json:"store"`
	Versions []IndicesVersions `json:"versions,omitempty"`
}

// ClusterIndicesBuilder holds ClusterIndices struct and provides a builder API.
type ClusterIndicesBuilder struct {
	v *ClusterIndices
}

// NewClusterIndices provides a builder for the ClusterIndices struct.
func NewClusterIndicesBuilder() *ClusterIndicesBuilder {
	r := ClusterIndicesBuilder{
		&ClusterIndices{},
	}

	return &r
}

// Build finalize the chain and returns the ClusterIndices struct
func (rb *ClusterIndicesBuilder) Build() ClusterIndices {
	return *rb.v
}

// Analysis Contains statistics about analyzers and analyzer components used in selected
// nodes.

func (rb *ClusterIndicesBuilder) Analysis(analysis *CharFilterTypesBuilder) *ClusterIndicesBuilder {
	v := analysis.Build()
	rb.v.Analysis = v
	return rb
}

// Completion Contains statistics about memory used for completion in selected nodes.

func (rb *ClusterIndicesBuilder) Completion(completion *CompletionStatsBuilder) *ClusterIndicesBuilder {
	v := completion.Build()
	rb.v.Completion = v
	return rb
}

// Count Total number of indices with shards assigned to selected nodes.

func (rb *ClusterIndicesBuilder) Count(count int64) *ClusterIndicesBuilder {
	rb.v.Count = count
	return rb
}

// Docs Contains counts for documents in selected nodes.

func (rb *ClusterIndicesBuilder) Docs(docs *DocStatsBuilder) *ClusterIndicesBuilder {
	v := docs.Build()
	rb.v.Docs = v
	return rb
}

// Fielddata Contains statistics about the field data cache of selected nodes.

func (rb *ClusterIndicesBuilder) Fielddata(fielddata *FielddataStatsBuilder) *ClusterIndicesBuilder {
	v := fielddata.Build()
	rb.v.Fielddata = v
	return rb
}

// Mappings Contains statistics about field mappings in selected nodes.

func (rb *ClusterIndicesBuilder) Mappings(mappings *FieldTypesMappingsBuilder) *ClusterIndicesBuilder {
	v := mappings.Build()
	rb.v.Mappings = v
	return rb
}

// QueryCache Contains statistics about the query cache of selected nodes.

func (rb *ClusterIndicesBuilder) QueryCache(querycache *QueryCacheStatsBuilder) *ClusterIndicesBuilder {
	v := querycache.Build()
	rb.v.QueryCache = v
	return rb
}

// Segments Contains statistics about segments in selected nodes.

func (rb *ClusterIndicesBuilder) Segments(segments *SegmentsStatsBuilder) *ClusterIndicesBuilder {
	v := segments.Build()
	rb.v.Segments = v
	return rb
}

// Shards Contains statistics about indices with shards assigned to selected nodes.

func (rb *ClusterIndicesBuilder) Shards(shards *ClusterIndicesShardsBuilder) *ClusterIndicesBuilder {
	v := shards.Build()
	rb.v.Shards = v
	return rb
}

// Store Contains statistics about the size of shards assigned to selected nodes.

func (rb *ClusterIndicesBuilder) Store(store *StoreStatsBuilder) *ClusterIndicesBuilder {
	v := store.Build()
	rb.v.Store = v
	return rb
}

func (rb *ClusterIndicesBuilder) Versions(versions []IndicesVersionsBuilder) *ClusterIndicesBuilder {
	tmp := make([]IndicesVersions, len(versions))
	for _, value := range versions {
		tmp = append(tmp, value.Build())
	}
	rb.v.Versions = tmp
	return rb
}
