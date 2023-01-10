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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

// ClusterIndices type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/cluster/stats/types.ts#L63-L94
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

// NewClusterIndices returns a ClusterIndices.
func NewClusterIndices() *ClusterIndices {
	r := &ClusterIndices{}

	return r
}
