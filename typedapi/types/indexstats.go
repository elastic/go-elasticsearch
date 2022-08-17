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

// IndexStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/stats/types.ts#L49-L87
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

// IndexStatsBuilder holds IndexStats struct and provides a builder API.
type IndexStatsBuilder struct {
	v *IndexStats
}

// NewIndexStats provides a builder for the IndexStats struct.
func NewIndexStatsBuilder() *IndexStatsBuilder {
	r := IndexStatsBuilder{
		&IndexStats{},
	}

	return &r
}

// Build finalize the chain and returns the IndexStats struct
func (rb *IndexStatsBuilder) Build() IndexStats {
	return *rb.v
}

func (rb *IndexStatsBuilder) Bulk(bulk *BulkStatsBuilder) *IndexStatsBuilder {
	v := bulk.Build()
	rb.v.Bulk = &v
	return rb
}

// Completion Contains statistics about completions across all shards assigned to the node.

func (rb *IndexStatsBuilder) Completion(completion *CompletionStatsBuilder) *IndexStatsBuilder {
	v := completion.Build()
	rb.v.Completion = &v
	return rb
}

// Docs Contains statistics about documents across all primary shards assigned to the
// node.

func (rb *IndexStatsBuilder) Docs(docs *DocStatsBuilder) *IndexStatsBuilder {
	v := docs.Build()
	rb.v.Docs = &v
	return rb
}

// Fielddata Contains statistics about the field data cache across all shards assigned to
// the node.

func (rb *IndexStatsBuilder) Fielddata(fielddata *FielddataStatsBuilder) *IndexStatsBuilder {
	v := fielddata.Build()
	rb.v.Fielddata = &v
	return rb
}

// Flush Contains statistics about flush operations for the node.

func (rb *IndexStatsBuilder) Flush(flush *FlushStatsBuilder) *IndexStatsBuilder {
	v := flush.Build()
	rb.v.Flush = &v
	return rb
}

// Get Contains statistics about get operations for the node.

func (rb *IndexStatsBuilder) Get(get *GetStatsBuilder) *IndexStatsBuilder {
	v := get.Build()
	rb.v.Get = &v
	return rb
}

// Indexing Contains statistics about indexing operations for the node.

func (rb *IndexStatsBuilder) Indexing(indexing *IndexingStatsBuilder) *IndexStatsBuilder {
	v := indexing.Build()
	rb.v.Indexing = &v
	return rb
}

// Indices Contains statistics about indices operations for the node.

func (rb *IndexStatsBuilder) Indices(indices *IndicesStatsBuilder) *IndexStatsBuilder {
	v := indices.Build()
	rb.v.Indices = &v
	return rb
}

// Merges Contains statistics about merge operations for the node.

func (rb *IndexStatsBuilder) Merges(merges *MergesStatsBuilder) *IndexStatsBuilder {
	v := merges.Build()
	rb.v.Merges = &v
	return rb
}

// QueryCache Contains statistics about the query cache across all shards assigned to the
// node.

func (rb *IndexStatsBuilder) QueryCache(querycache *QueryCacheStatsBuilder) *IndexStatsBuilder {
	v := querycache.Build()
	rb.v.QueryCache = &v
	return rb
}

// Recovery Contains statistics about recovery operations for the node.

func (rb *IndexStatsBuilder) Recovery(recovery *RecoveryStatsBuilder) *IndexStatsBuilder {
	v := recovery.Build()
	rb.v.Recovery = &v
	return rb
}

// Refresh Contains statistics about refresh operations for the node.

func (rb *IndexStatsBuilder) Refresh(refresh *RefreshStatsBuilder) *IndexStatsBuilder {
	v := refresh.Build()
	rb.v.Refresh = &v
	return rb
}

// RequestCache Contains statistics about the request cache across all shards assigned to the
// node.

func (rb *IndexStatsBuilder) RequestCache(requestcache *RequestCacheStatsBuilder) *IndexStatsBuilder {
	v := requestcache.Build()
	rb.v.RequestCache = &v
	return rb
}

// Search Contains statistics about search operations for the node.

func (rb *IndexStatsBuilder) Search(search *SearchStatsBuilder) *IndexStatsBuilder {
	v := search.Build()
	rb.v.Search = &v
	return rb
}

// Segments Contains statistics about segments across all shards assigned to the node.

func (rb *IndexStatsBuilder) Segments(segments *SegmentsStatsBuilder) *IndexStatsBuilder {
	v := segments.Build()
	rb.v.Segments = &v
	return rb
}

func (rb *IndexStatsBuilder) ShardStats(shardstats *ShardsTotalStatsBuilder) *IndexStatsBuilder {
	v := shardstats.Build()
	rb.v.ShardStats = &v
	return rb
}

// Store Contains statistics about the size of shards assigned to the node.

func (rb *IndexStatsBuilder) Store(store *StoreStatsBuilder) *IndexStatsBuilder {
	v := store.Build()
	rb.v.Store = &v
	return rb
}

// Translog Contains statistics about transaction log operations for the node.

func (rb *IndexStatsBuilder) Translog(translog *TranslogStatsBuilder) *IndexStatsBuilder {
	v := translog.Build()
	rb.v.Translog = &v
	return rb
}

// Warmer Contains statistics about index warming operations for the node.

func (rb *IndexStatsBuilder) Warmer(warmer *WarmerStatsBuilder) *IndexStatsBuilder {
	v := warmer.Build()
	rb.v.Warmer = &v
	return rb
}
