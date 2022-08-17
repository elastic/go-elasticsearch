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

// ShardStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/stats/types.ts#L174-L201
type ShardStats struct {
	Bulk            *BulkStats            `json:"bulk,omitempty"`
	Commit          *ShardCommit          `json:"commit,omitempty"`
	Completion      *CompletionStats      `json:"completion,omitempty"`
	Docs            *DocStats             `json:"docs,omitempty"`
	Fielddata       *FielddataStats       `json:"fielddata,omitempty"`
	Flush           *FlushStats           `json:"flush,omitempty"`
	Get             *GetStats             `json:"get,omitempty"`
	Indexing        *IndexingStats        `json:"indexing,omitempty"`
	Indices         *IndicesStats         `json:"indices,omitempty"`
	Merges          *MergesStats          `json:"merges,omitempty"`
	QueryCache      *ShardQueryCache      `json:"query_cache,omitempty"`
	Recovery        *RecoveryStats        `json:"recovery,omitempty"`
	Refresh         *RefreshStats         `json:"refresh,omitempty"`
	RequestCache    *RequestCacheStats    `json:"request_cache,omitempty"`
	RetentionLeases *ShardRetentionLeases `json:"retention_leases,omitempty"`
	Routing         *ShardRouting         `json:"routing,omitempty"`
	Search          *SearchStats          `json:"search,omitempty"`
	Segments        *SegmentsStats        `json:"segments,omitempty"`
	SeqNo           *ShardSequenceNumber  `json:"seq_no,omitempty"`
	ShardPath       *ShardPath            `json:"shard_path,omitempty"`
	ShardStats      *ShardsTotalStats     `json:"shard_stats,omitempty"`
	Shards          *ShardsTotalStats     `json:"shards,omitempty"`
	Store           *StoreStats           `json:"store,omitempty"`
	Translog        *TranslogStats        `json:"translog,omitempty"`
	Warmer          *WarmerStats          `json:"warmer,omitempty"`
}

// ShardStatsBuilder holds ShardStats struct and provides a builder API.
type ShardStatsBuilder struct {
	v *ShardStats
}

// NewShardStats provides a builder for the ShardStats struct.
func NewShardStatsBuilder() *ShardStatsBuilder {
	r := ShardStatsBuilder{
		&ShardStats{},
	}

	return &r
}

// Build finalize the chain and returns the ShardStats struct
func (rb *ShardStatsBuilder) Build() ShardStats {
	return *rb.v
}

func (rb *ShardStatsBuilder) Bulk(bulk *BulkStatsBuilder) *ShardStatsBuilder {
	v := bulk.Build()
	rb.v.Bulk = &v
	return rb
}

func (rb *ShardStatsBuilder) Commit(commit *ShardCommitBuilder) *ShardStatsBuilder {
	v := commit.Build()
	rb.v.Commit = &v
	return rb
}

func (rb *ShardStatsBuilder) Completion(completion *CompletionStatsBuilder) *ShardStatsBuilder {
	v := completion.Build()
	rb.v.Completion = &v
	return rb
}

func (rb *ShardStatsBuilder) Docs(docs *DocStatsBuilder) *ShardStatsBuilder {
	v := docs.Build()
	rb.v.Docs = &v
	return rb
}

func (rb *ShardStatsBuilder) Fielddata(fielddata *FielddataStatsBuilder) *ShardStatsBuilder {
	v := fielddata.Build()
	rb.v.Fielddata = &v
	return rb
}

func (rb *ShardStatsBuilder) Flush(flush *FlushStatsBuilder) *ShardStatsBuilder {
	v := flush.Build()
	rb.v.Flush = &v
	return rb
}

func (rb *ShardStatsBuilder) Get(get *GetStatsBuilder) *ShardStatsBuilder {
	v := get.Build()
	rb.v.Get = &v
	return rb
}

func (rb *ShardStatsBuilder) Indexing(indexing *IndexingStatsBuilder) *ShardStatsBuilder {
	v := indexing.Build()
	rb.v.Indexing = &v
	return rb
}

func (rb *ShardStatsBuilder) Indices(indices *IndicesStatsBuilder) *ShardStatsBuilder {
	v := indices.Build()
	rb.v.Indices = &v
	return rb
}

func (rb *ShardStatsBuilder) Merges(merges *MergesStatsBuilder) *ShardStatsBuilder {
	v := merges.Build()
	rb.v.Merges = &v
	return rb
}

func (rb *ShardStatsBuilder) QueryCache(querycache *ShardQueryCacheBuilder) *ShardStatsBuilder {
	v := querycache.Build()
	rb.v.QueryCache = &v
	return rb
}

func (rb *ShardStatsBuilder) Recovery(recovery *RecoveryStatsBuilder) *ShardStatsBuilder {
	v := recovery.Build()
	rb.v.Recovery = &v
	return rb
}

func (rb *ShardStatsBuilder) Refresh(refresh *RefreshStatsBuilder) *ShardStatsBuilder {
	v := refresh.Build()
	rb.v.Refresh = &v
	return rb
}

func (rb *ShardStatsBuilder) RequestCache(requestcache *RequestCacheStatsBuilder) *ShardStatsBuilder {
	v := requestcache.Build()
	rb.v.RequestCache = &v
	return rb
}

func (rb *ShardStatsBuilder) RetentionLeases(retentionleases *ShardRetentionLeasesBuilder) *ShardStatsBuilder {
	v := retentionleases.Build()
	rb.v.RetentionLeases = &v
	return rb
}

func (rb *ShardStatsBuilder) Routing(routing *ShardRoutingBuilder) *ShardStatsBuilder {
	v := routing.Build()
	rb.v.Routing = &v
	return rb
}

func (rb *ShardStatsBuilder) Search(search *SearchStatsBuilder) *ShardStatsBuilder {
	v := search.Build()
	rb.v.Search = &v
	return rb
}

func (rb *ShardStatsBuilder) Segments(segments *SegmentsStatsBuilder) *ShardStatsBuilder {
	v := segments.Build()
	rb.v.Segments = &v
	return rb
}

func (rb *ShardStatsBuilder) SeqNo(seqno *ShardSequenceNumberBuilder) *ShardStatsBuilder {
	v := seqno.Build()
	rb.v.SeqNo = &v
	return rb
}

func (rb *ShardStatsBuilder) ShardPath(shardpath *ShardPathBuilder) *ShardStatsBuilder {
	v := shardpath.Build()
	rb.v.ShardPath = &v
	return rb
}

func (rb *ShardStatsBuilder) ShardStats(shardstats *ShardsTotalStatsBuilder) *ShardStatsBuilder {
	v := shardstats.Build()
	rb.v.ShardStats = &v
	return rb
}

func (rb *ShardStatsBuilder) Shards(shards *ShardsTotalStatsBuilder) *ShardStatsBuilder {
	v := shards.Build()
	rb.v.Shards = &v
	return rb
}

func (rb *ShardStatsBuilder) Store(store *StoreStatsBuilder) *ShardStatsBuilder {
	v := store.Build()
	rb.v.Store = &v
	return rb
}

func (rb *ShardStatsBuilder) Translog(translog *TranslogStatsBuilder) *ShardStatsBuilder {
	v := translog.Build()
	rb.v.Translog = &v
	return rb
}

func (rb *ShardStatsBuilder) Warmer(warmer *WarmerStatsBuilder) *ShardStatsBuilder {
	v := warmer.Build()
	rb.v.Warmer = &v
	return rb
}
