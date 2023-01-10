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

// ShardsRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/cat/shards/types.ts#L20-L396
type ShardsRecord struct {
	// BulkAvgSizeInBytes avg size in bytes of shard bulk
	BulkAvgSizeInBytes *string `json:"bulk.avg_size_in_bytes,omitempty"`
	// BulkAvgTime average time spend in shard bulk
	BulkAvgTime *string `json:"bulk.avg_time,omitempty"`
	// BulkTotalOperations number of bulk shard ops
	BulkTotalOperations *string `json:"bulk.total_operations,omitempty"`
	// BulkTotalSizeInBytes total size in bytes of shard bulk
	BulkTotalSizeInBytes *string `json:"bulk.total_size_in_bytes,omitempty"`
	// BulkTotalTime time spend in shard bulk
	BulkTotalTime *string `json:"bulk.total_time,omitempty"`
	// CompletionSize size of completion
	CompletionSize *string `json:"completion.size,omitempty"`
	// Docs number of docs in shard
	Docs string `json:"docs,omitempty"`
	// FielddataEvictions fielddata evictions
	FielddataEvictions *string `json:"fielddata.evictions,omitempty"`
	// FielddataMemorySize used fielddata cache
	FielddataMemorySize *string `json:"fielddata.memory_size,omitempty"`
	// FlushTotal number of flushes
	FlushTotal *string `json:"flush.total,omitempty"`
	// FlushTotalTime time spent in flush
	FlushTotalTime *string `json:"flush.total_time,omitempty"`
	// GetCurrent number of current get ops
	GetCurrent *string `json:"get.current,omitempty"`
	// GetExistsTime time spent in successful gets
	GetExistsTime *string `json:"get.exists_time,omitempty"`
	// GetExistsTotal number of successful gets
	GetExistsTotal *string `json:"get.exists_total,omitempty"`
	// GetMissingTime time spent in failed gets
	GetMissingTime *string `json:"get.missing_time,omitempty"`
	// GetMissingTotal number of failed gets
	GetMissingTotal *string `json:"get.missing_total,omitempty"`
	// GetTime time spent in get
	GetTime *string `json:"get.time,omitempty"`
	// GetTotal number of get ops
	GetTotal *string `json:"get.total,omitempty"`
	// Id unique id of node where it lives
	Id *string `json:"id,omitempty"`
	// Index index name
	Index *string `json:"index,omitempty"`
	// IndexingDeleteCurrent number of current deletions
	IndexingDeleteCurrent *string `json:"indexing.delete_current,omitempty"`
	// IndexingDeleteTime time spent in deletions
	IndexingDeleteTime *string `json:"indexing.delete_time,omitempty"`
	// IndexingDeleteTotal number of delete ops
	IndexingDeleteTotal *string `json:"indexing.delete_total,omitempty"`
	// IndexingIndexCurrent number of current indexing ops
	IndexingIndexCurrent *string `json:"indexing.index_current,omitempty"`
	// IndexingIndexFailed number of failed indexing ops
	IndexingIndexFailed *string `json:"indexing.index_failed,omitempty"`
	// IndexingIndexTime time spent in indexing
	IndexingIndexTime *string `json:"indexing.index_time,omitempty"`
	// IndexingIndexTotal number of indexing ops
	IndexingIndexTotal *string `json:"indexing.index_total,omitempty"`
	// Ip ip of node where it lives
	Ip string `json:"ip,omitempty"`
	// MergesCurrent number of current merges
	MergesCurrent *string `json:"merges.current,omitempty"`
	// MergesCurrentDocs number of current merging docs
	MergesCurrentDocs *string `json:"merges.current_docs,omitempty"`
	// MergesCurrentSize size of current merges
	MergesCurrentSize *string `json:"merges.current_size,omitempty"`
	// MergesTotal number of completed merge ops
	MergesTotal *string `json:"merges.total,omitempty"`
	// MergesTotalDocs docs merged
	MergesTotalDocs *string `json:"merges.total_docs,omitempty"`
	// MergesTotalSize size merged
	MergesTotalSize *string `json:"merges.total_size,omitempty"`
	// MergesTotalTime time spent in merges
	MergesTotalTime *string `json:"merges.total_time,omitempty"`
	// Node name of node where it lives
	Node string `json:"node,omitempty"`
	// PathData shard data path
	PathData *string `json:"path.data,omitempty"`
	// PathState shard state path
	PathState *string `json:"path.state,omitempty"`
	// Prirep primary or replica
	Prirep *string `json:"prirep,omitempty"`
	// QueryCacheEvictions query cache evictions
	QueryCacheEvictions *string `json:"query_cache.evictions,omitempty"`
	// QueryCacheMemorySize used query cache
	QueryCacheMemorySize *string `json:"query_cache.memory_size,omitempty"`
	// RecoverysourceType recovery source type
	RecoverysourceType *string `json:"recoverysource.type,omitempty"`
	// RefreshExternalTime time spent in external refreshes
	RefreshExternalTime *string `json:"refresh.external_time,omitempty"`
	// RefreshExternalTotal total external refreshes
	RefreshExternalTotal *string `json:"refresh.external_total,omitempty"`
	// RefreshListeners number of pending refresh listeners
	RefreshListeners *string `json:"refresh.listeners,omitempty"`
	// RefreshTime time spent in refreshes
	RefreshTime *string `json:"refresh.time,omitempty"`
	// RefreshTotal total refreshes
	RefreshTotal *string `json:"refresh.total,omitempty"`
	// SearchFetchCurrent current fetch phase ops
	SearchFetchCurrent *string `json:"search.fetch_current,omitempty"`
	// SearchFetchTime time spent in fetch phase
	SearchFetchTime *string `json:"search.fetch_time,omitempty"`
	// SearchFetchTotal total fetch ops
	SearchFetchTotal *string `json:"search.fetch_total,omitempty"`
	// SearchOpenContexts open search contexts
	SearchOpenContexts *string `json:"search.open_contexts,omitempty"`
	// SearchQueryCurrent current query phase ops
	SearchQueryCurrent *string `json:"search.query_current,omitempty"`
	// SearchQueryTime time spent in query phase
	SearchQueryTime *string `json:"search.query_time,omitempty"`
	// SearchQueryTotal total query phase ops
	SearchQueryTotal *string `json:"search.query_total,omitempty"`
	// SearchScrollCurrent open scroll contexts
	SearchScrollCurrent *string `json:"search.scroll_current,omitempty"`
	// SearchScrollTime time scroll contexts held open
	SearchScrollTime *string `json:"search.scroll_time,omitempty"`
	// SearchScrollTotal completed scroll contexts
	SearchScrollTotal *string `json:"search.scroll_total,omitempty"`
	// SegmentsCount number of segments
	SegmentsCount *string `json:"segments.count,omitempty"`
	// SegmentsFixedBitsetMemory memory used by fixed bit sets for nested object field types and export type
	// filters for types referred in _parent fields
	SegmentsFixedBitsetMemory *string `json:"segments.fixed_bitset_memory,omitempty"`
	// SegmentsIndexWriterMemory memory used by index writer
	SegmentsIndexWriterMemory *string `json:"segments.index_writer_memory,omitempty"`
	// SegmentsMemory memory used by segments
	SegmentsMemory *string `json:"segments.memory,omitempty"`
	// SegmentsVersionMapMemory memory used by version map
	SegmentsVersionMapMemory *string `json:"segments.version_map_memory,omitempty"`
	// SeqNoGlobalCheckpoint global checkpoint
	SeqNoGlobalCheckpoint *string `json:"seq_no.global_checkpoint,omitempty"`
	// SeqNoLocalCheckpoint local checkpoint
	SeqNoLocalCheckpoint *string `json:"seq_no.local_checkpoint,omitempty"`
	// SeqNoMax max sequence number
	SeqNoMax *string `json:"seq_no.max,omitempty"`
	// Shard shard name
	Shard *string `json:"shard,omitempty"`
	// State shard state
	State *string `json:"state,omitempty"`
	// Store store size of shard (how much disk it uses)
	Store string `json:"store,omitempty"`
	// SyncId sync id
	SyncId *string `json:"sync_id,omitempty"`
	// UnassignedAt time shard became unassigned (UTC)
	UnassignedAt *string `json:"unassigned.at,omitempty"`
	// UnassignedDetails additional details as to why the shard became unassigned
	UnassignedDetails *string `json:"unassigned.details,omitempty"`
	// UnassignedFor time has been unassigned
	UnassignedFor *string `json:"unassigned.for,omitempty"`
	// UnassignedReason reason shard is unassigned
	UnassignedReason *string `json:"unassigned.reason,omitempty"`
	// WarmerCurrent current warmer ops
	WarmerCurrent *string `json:"warmer.current,omitempty"`
	// WarmerTotal total warmer ops
	WarmerTotal *string `json:"warmer.total,omitempty"`
	// WarmerTotalTime time spent in warmers
	WarmerTotalTime *string `json:"warmer.total_time,omitempty"`
}

// NewShardsRecord returns a ShardsRecord.
func NewShardsRecord() *ShardsRecord {
	r := &ShardsRecord{}

	return r
}
