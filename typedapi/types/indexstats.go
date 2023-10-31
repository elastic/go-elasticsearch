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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

// IndexStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/indices/stats/types.ts#L52-L93
type IndexStats struct {
	Bulk *BulkStats `json:"bulk,omitempty"`
	// Completion Contains statistics about completions across all shards assigned to the node.
	Completion *CompletionStats `json:"completion,omitempty"`
	// Docs Contains statistics about documents across all primary shards assigned to the
	// node.
	Docs *DocStats `json:"docs,omitempty"`
	// Fielddata Contains statistics about the field data cache across all shards assigned to
	// the node.
	Fielddata *FielddataStats `json:"fielddata,omitempty"`
	// Flush Contains statistics about flush operations for the node.
	Flush *FlushStats `json:"flush,omitempty"`
	// Get Contains statistics about get operations for the node.
	Get *GetStats `json:"get,omitempty"`
	// Indexing Contains statistics about indexing operations for the node.
	Indexing *IndexingStats `json:"indexing,omitempty"`
	// Indices Contains statistics about indices operations for the node.
	Indices *IndicesStats `json:"indices,omitempty"`
	// Merges Contains statistics about merge operations for the node.
	Merges *MergesStats `json:"merges,omitempty"`
	// QueryCache Contains statistics about the query cache across all shards assigned to the
	// node.
	QueryCache *QueryCacheStats `json:"query_cache,omitempty"`
	// Recovery Contains statistics about recovery operations for the node.
	Recovery *RecoveryStats `json:"recovery,omitempty"`
	// Refresh Contains statistics about refresh operations for the node.
	Refresh *RefreshStats `json:"refresh,omitempty"`
	// RequestCache Contains statistics about the request cache across all shards assigned to the
	// node.
	RequestCache *RequestCacheStats `json:"request_cache,omitempty"`
	// Search Contains statistics about search operations for the node.
	Search *SearchStats `json:"search,omitempty"`
	// Segments Contains statistics about segments across all shards assigned to the node.
	Segments   *SegmentsStats    `json:"segments,omitempty"`
	ShardStats *ShardsTotalStats `json:"shard_stats,omitempty"`
	// Store Contains statistics about the size of shards assigned to the node.
	Store *StoreStats `json:"store,omitempty"`
	// Translog Contains statistics about transaction log operations for the node.
	Translog *TranslogStats `json:"translog,omitempty"`
	// Warmer Contains statistics about index warming operations for the node.
	Warmer *WarmerStats `json:"warmer,omitempty"`
}

// NewIndexStats returns a IndexStats.
func NewIndexStats() *IndexStats {
	r := &IndexStats{}

	return r
}
