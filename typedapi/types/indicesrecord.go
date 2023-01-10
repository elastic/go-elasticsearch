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

// IndicesRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/cat/indices/types.ts#L20-L801
type IndicesRecord struct {
	// BulkAvgSizeInBytes average size in bytes of shard bulk
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
	// CreationDate index creation date (millisecond value)
	CreationDate *string `json:"creation.date,omitempty"`
	// CreationDateString index creation date (as string)
	CreationDateString *string `json:"creation.date.string,omitempty"`
	// DocsCount available docs
	DocsCount string `json:"docs.count,omitempty"`
	// DocsDeleted deleted docs
	DocsDeleted string `json:"docs.deleted,omitempty"`
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
	// Health current health status
	Health *string `json:"health,omitempty"`
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
	// MemoryTotal total used memory
	MemoryTotal *string `json:"memory.total,omitempty"`
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
	// Pri number of primary shards
	Pri *string `json:"pri,omitempty"`
	// PriBulkAvgSizeInBytes average size in bytes of shard bulk
	PriBulkAvgSizeInBytes *string `json:"pri.bulk.avg_size_in_bytes,omitempty"`
	// PriBulkAvgTime average time spend in shard bulk
	PriBulkAvgTime *string `json:"pri.bulk.avg_time,omitempty"`
	// PriBulkTotalOperations number of bulk shard ops
	PriBulkTotalOperations *string `json:"pri.bulk.total_operations,omitempty"`
	// PriBulkTotalSizeInBytes total size in bytes of shard bulk
	PriBulkTotalSizeInBytes *string `json:"pri.bulk.total_size_in_bytes,omitempty"`
	// PriBulkTotalTime time spend in shard bulk
	PriBulkTotalTime *string `json:"pri.bulk.total_time,omitempty"`
	// PriCompletionSize size of completion
	PriCompletionSize *string `json:"pri.completion.size,omitempty"`
	// PriFielddataEvictions fielddata evictions
	PriFielddataEvictions *string `json:"pri.fielddata.evictions,omitempty"`
	// PriFielddataMemorySize used fielddata cache
	PriFielddataMemorySize *string `json:"pri.fielddata.memory_size,omitempty"`
	// PriFlushTotal number of flushes
	PriFlushTotal *string `json:"pri.flush.total,omitempty"`
	// PriFlushTotalTime time spent in flush
	PriFlushTotalTime *string `json:"pri.flush.total_time,omitempty"`
	// PriGetCurrent number of current get ops
	PriGetCurrent *string `json:"pri.get.current,omitempty"`
	// PriGetExistsTime time spent in successful gets
	PriGetExistsTime *string `json:"pri.get.exists_time,omitempty"`
	// PriGetExistsTotal number of successful gets
	PriGetExistsTotal *string `json:"pri.get.exists_total,omitempty"`
	// PriGetMissingTime time spent in failed gets
	PriGetMissingTime *string `json:"pri.get.missing_time,omitempty"`
	// PriGetMissingTotal number of failed gets
	PriGetMissingTotal *string `json:"pri.get.missing_total,omitempty"`
	// PriGetTime time spent in get
	PriGetTime *string `json:"pri.get.time,omitempty"`
	// PriGetTotal number of get ops
	PriGetTotal *string `json:"pri.get.total,omitempty"`
	// PriIndexingDeleteCurrent number of current deletions
	PriIndexingDeleteCurrent *string `json:"pri.indexing.delete_current,omitempty"`
	// PriIndexingDeleteTime time spent in deletions
	PriIndexingDeleteTime *string `json:"pri.indexing.delete_time,omitempty"`
	// PriIndexingDeleteTotal number of delete ops
	PriIndexingDeleteTotal *string `json:"pri.indexing.delete_total,omitempty"`
	// PriIndexingIndexCurrent number of current indexing ops
	PriIndexingIndexCurrent *string `json:"pri.indexing.index_current,omitempty"`
	// PriIndexingIndexFailed number of failed indexing ops
	PriIndexingIndexFailed *string `json:"pri.indexing.index_failed,omitempty"`
	// PriIndexingIndexTime time spent in indexing
	PriIndexingIndexTime *string `json:"pri.indexing.index_time,omitempty"`
	// PriIndexingIndexTotal number of indexing ops
	PriIndexingIndexTotal *string `json:"pri.indexing.index_total,omitempty"`
	// PriMemoryTotal total user memory
	PriMemoryTotal *string `json:"pri.memory.total,omitempty"`
	// PriMergesCurrent number of current merges
	PriMergesCurrent *string `json:"pri.merges.current,omitempty"`
	// PriMergesCurrentDocs number of current merging docs
	PriMergesCurrentDocs *string `json:"pri.merges.current_docs,omitempty"`
	// PriMergesCurrentSize size of current merges
	PriMergesCurrentSize *string `json:"pri.merges.current_size,omitempty"`
	// PriMergesTotal number of completed merge ops
	PriMergesTotal *string `json:"pri.merges.total,omitempty"`
	// PriMergesTotalDocs docs merged
	PriMergesTotalDocs *string `json:"pri.merges.total_docs,omitempty"`
	// PriMergesTotalSize size merged
	PriMergesTotalSize *string `json:"pri.merges.total_size,omitempty"`
	// PriMergesTotalTime time spent in merges
	PriMergesTotalTime *string `json:"pri.merges.total_time,omitempty"`
	// PriQueryCacheEvictions query cache evictions
	PriQueryCacheEvictions *string `json:"pri.query_cache.evictions,omitempty"`
	// PriQueryCacheMemorySize used query cache
	PriQueryCacheMemorySize *string `json:"pri.query_cache.memory_size,omitempty"`
	// PriRefreshExternalTime time spent in external refreshes
	PriRefreshExternalTime *string `json:"pri.refresh.external_time,omitempty"`
	// PriRefreshExternalTotal total external refreshes
	PriRefreshExternalTotal *string `json:"pri.refresh.external_total,omitempty"`
	// PriRefreshListeners number of pending refresh listeners
	PriRefreshListeners *string `json:"pri.refresh.listeners,omitempty"`
	// PriRefreshTime time spent in refreshes
	PriRefreshTime *string `json:"pri.refresh.time,omitempty"`
	// PriRefreshTotal total refreshes
	PriRefreshTotal *string `json:"pri.refresh.total,omitempty"`
	// PriRequestCacheEvictions request cache evictions
	PriRequestCacheEvictions *string `json:"pri.request_cache.evictions,omitempty"`
	// PriRequestCacheHitCount request cache hit count
	PriRequestCacheHitCount *string `json:"pri.request_cache.hit_count,omitempty"`
	// PriRequestCacheMemorySize used request cache
	PriRequestCacheMemorySize *string `json:"pri.request_cache.memory_size,omitempty"`
	// PriRequestCacheMissCount request cache miss count
	PriRequestCacheMissCount *string `json:"pri.request_cache.miss_count,omitempty"`
	// PriSearchFetchCurrent current fetch phase ops
	PriSearchFetchCurrent *string `json:"pri.search.fetch_current,omitempty"`
	// PriSearchFetchTime time spent in fetch phase
	PriSearchFetchTime *string `json:"pri.search.fetch_time,omitempty"`
	// PriSearchFetchTotal total fetch ops
	PriSearchFetchTotal *string `json:"pri.search.fetch_total,omitempty"`
	// PriSearchOpenContexts open search contexts
	PriSearchOpenContexts *string `json:"pri.search.open_contexts,omitempty"`
	// PriSearchQueryCurrent current query phase ops
	PriSearchQueryCurrent *string `json:"pri.search.query_current,omitempty"`
	// PriSearchQueryTime time spent in query phase
	PriSearchQueryTime *string `json:"pri.search.query_time,omitempty"`
	// PriSearchQueryTotal total query phase ops
	PriSearchQueryTotal *string `json:"pri.search.query_total,omitempty"`
	// PriSearchScrollCurrent open scroll contexts
	PriSearchScrollCurrent *string `json:"pri.search.scroll_current,omitempty"`
	// PriSearchScrollTime time scroll contexts held open
	PriSearchScrollTime *string `json:"pri.search.scroll_time,omitempty"`
	// PriSearchScrollTotal completed scroll contexts
	PriSearchScrollTotal *string `json:"pri.search.scroll_total,omitempty"`
	// PriSegmentsCount number of segments
	PriSegmentsCount *string `json:"pri.segments.count,omitempty"`
	// PriSegmentsFixedBitsetMemory memory used by fixed bit sets for nested object field types and export type
	// filters for types referred in _parent fields
	PriSegmentsFixedBitsetMemory *string `json:"pri.segments.fixed_bitset_memory,omitempty"`
	// PriSegmentsIndexWriterMemory memory used by index writer
	PriSegmentsIndexWriterMemory *string `json:"pri.segments.index_writer_memory,omitempty"`
	// PriSegmentsMemory memory used by segments
	PriSegmentsMemory *string `json:"pri.segments.memory,omitempty"`
	// PriSegmentsVersionMapMemory memory used by version map
	PriSegmentsVersionMapMemory *string `json:"pri.segments.version_map_memory,omitempty"`
	// PriStoreSize store size of primaries
	PriStoreSize string `json:"pri.store.size,omitempty"`
	// PriSuggestCurrent number of current suggest ops
	PriSuggestCurrent *string `json:"pri.suggest.current,omitempty"`
	// PriSuggestTime time spend in suggest
	PriSuggestTime *string `json:"pri.suggest.time,omitempty"`
	// PriSuggestTotal number of suggest ops
	PriSuggestTotal *string `json:"pri.suggest.total,omitempty"`
	// PriWarmerCurrent current warmer ops
	PriWarmerCurrent *string `json:"pri.warmer.current,omitempty"`
	// PriWarmerTotal total warmer ops
	PriWarmerTotal *string `json:"pri.warmer.total,omitempty"`
	// PriWarmerTotalTime time spent in warmers
	PriWarmerTotalTime *string `json:"pri.warmer.total_time,omitempty"`
	// QueryCacheEvictions query cache evictions
	QueryCacheEvictions *string `json:"query_cache.evictions,omitempty"`
	// QueryCacheMemorySize used query cache
	QueryCacheMemorySize *string `json:"query_cache.memory_size,omitempty"`
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
	// Rep number of replica shards
	Rep *string `json:"rep,omitempty"`
	// RequestCacheEvictions request cache evictions
	RequestCacheEvictions *string `json:"request_cache.evictions,omitempty"`
	// RequestCacheHitCount request cache hit count
	RequestCacheHitCount *string `json:"request_cache.hit_count,omitempty"`
	// RequestCacheMemorySize used request cache
	RequestCacheMemorySize *string `json:"request_cache.memory_size,omitempty"`
	// RequestCacheMissCount request cache miss count
	RequestCacheMissCount *string `json:"request_cache.miss_count,omitempty"`
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
	// SearchThrottled indicates if the index is search throttled
	SearchThrottled *string `json:"search.throttled,omitempty"`
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
	// Status open/close status
	Status *string `json:"status,omitempty"`
	// StoreSize store size of primaries & replicas
	StoreSize string `json:"store.size,omitempty"`
	// SuggestCurrent number of current suggest ops
	SuggestCurrent *string `json:"suggest.current,omitempty"`
	// SuggestTime time spend in suggest
	SuggestTime *string `json:"suggest.time,omitempty"`
	// SuggestTotal number of suggest ops
	SuggestTotal *string `json:"suggest.total,omitempty"`
	// Uuid index uuid
	Uuid *string `json:"uuid,omitempty"`
	// WarmerCurrent current warmer ops
	WarmerCurrent *string `json:"warmer.current,omitempty"`
	// WarmerTotal total warmer ops
	WarmerTotal *string `json:"warmer.total,omitempty"`
	// WarmerTotalTime time spent in warmers
	WarmerTotalTime *string `json:"warmer.total_time,omitempty"`
}

// NewIndicesRecord returns a IndicesRecord.
func NewIndicesRecord() *IndicesRecord {
	r := &IndicesRecord{}

	return r
}
