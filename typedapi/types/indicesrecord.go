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

// IndicesRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cat/indices/types.ts#L20-L801
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

// IndicesRecordBuilder holds IndicesRecord struct and provides a builder API.
type IndicesRecordBuilder struct {
	v *IndicesRecord
}

// NewIndicesRecord provides a builder for the IndicesRecord struct.
func NewIndicesRecordBuilder() *IndicesRecordBuilder {
	r := IndicesRecordBuilder{
		&IndicesRecord{},
	}

	return &r
}

// Build finalize the chain and returns the IndicesRecord struct
func (rb *IndicesRecordBuilder) Build() IndicesRecord {
	return *rb.v
}

// BulkAvgSizeInBytes average size in bytes of shard bulk

func (rb *IndicesRecordBuilder) BulkAvgSizeInBytes(bulkavgsizeinbytes string) *IndicesRecordBuilder {
	rb.v.BulkAvgSizeInBytes = &bulkavgsizeinbytes
	return rb
}

// BulkAvgTime average time spend in shard bulk

func (rb *IndicesRecordBuilder) BulkAvgTime(bulkavgtime string) *IndicesRecordBuilder {
	rb.v.BulkAvgTime = &bulkavgtime
	return rb
}

// BulkTotalOperations number of bulk shard ops

func (rb *IndicesRecordBuilder) BulkTotalOperations(bulktotaloperations string) *IndicesRecordBuilder {
	rb.v.BulkTotalOperations = &bulktotaloperations
	return rb
}

// BulkTotalSizeInBytes total size in bytes of shard bulk

func (rb *IndicesRecordBuilder) BulkTotalSizeInBytes(bulktotalsizeinbytes string) *IndicesRecordBuilder {
	rb.v.BulkTotalSizeInBytes = &bulktotalsizeinbytes
	return rb
}

// BulkTotalTime time spend in shard bulk

func (rb *IndicesRecordBuilder) BulkTotalTime(bulktotaltime string) *IndicesRecordBuilder {
	rb.v.BulkTotalTime = &bulktotaltime
	return rb
}

// CompletionSize size of completion

func (rb *IndicesRecordBuilder) CompletionSize(completionsize string) *IndicesRecordBuilder {
	rb.v.CompletionSize = &completionsize
	return rb
}

// CreationDate index creation date (millisecond value)

func (rb *IndicesRecordBuilder) CreationDate(creationdate string) *IndicesRecordBuilder {
	rb.v.CreationDate = &creationdate
	return rb
}

// CreationDateString index creation date (as string)

func (rb *IndicesRecordBuilder) CreationDateString(creationdatestring string) *IndicesRecordBuilder {
	rb.v.CreationDateString = &creationdatestring
	return rb
}

// DocsCount available docs

func (rb *IndicesRecordBuilder) DocsCount(docscount string) *IndicesRecordBuilder {
	rb.v.DocsCount = docscount
	return rb
}

// DocsDeleted deleted docs

func (rb *IndicesRecordBuilder) DocsDeleted(docsdeleted string) *IndicesRecordBuilder {
	rb.v.DocsDeleted = docsdeleted
	return rb
}

// FielddataEvictions fielddata evictions

func (rb *IndicesRecordBuilder) FielddataEvictions(fielddataevictions string) *IndicesRecordBuilder {
	rb.v.FielddataEvictions = &fielddataevictions
	return rb
}

// FielddataMemorySize used fielddata cache

func (rb *IndicesRecordBuilder) FielddataMemorySize(fielddatamemorysize string) *IndicesRecordBuilder {
	rb.v.FielddataMemorySize = &fielddatamemorysize
	return rb
}

// FlushTotal number of flushes

func (rb *IndicesRecordBuilder) FlushTotal(flushtotal string) *IndicesRecordBuilder {
	rb.v.FlushTotal = &flushtotal
	return rb
}

// FlushTotalTime time spent in flush

func (rb *IndicesRecordBuilder) FlushTotalTime(flushtotaltime string) *IndicesRecordBuilder {
	rb.v.FlushTotalTime = &flushtotaltime
	return rb
}

// GetCurrent number of current get ops

func (rb *IndicesRecordBuilder) GetCurrent(getcurrent string) *IndicesRecordBuilder {
	rb.v.GetCurrent = &getcurrent
	return rb
}

// GetExistsTime time spent in successful gets

func (rb *IndicesRecordBuilder) GetExistsTime(getexiststime string) *IndicesRecordBuilder {
	rb.v.GetExistsTime = &getexiststime
	return rb
}

// GetExistsTotal number of successful gets

func (rb *IndicesRecordBuilder) GetExistsTotal(getexiststotal string) *IndicesRecordBuilder {
	rb.v.GetExistsTotal = &getexiststotal
	return rb
}

// GetMissingTime time spent in failed gets

func (rb *IndicesRecordBuilder) GetMissingTime(getmissingtime string) *IndicesRecordBuilder {
	rb.v.GetMissingTime = &getmissingtime
	return rb
}

// GetMissingTotal number of failed gets

func (rb *IndicesRecordBuilder) GetMissingTotal(getmissingtotal string) *IndicesRecordBuilder {
	rb.v.GetMissingTotal = &getmissingtotal
	return rb
}

// GetTime time spent in get

func (rb *IndicesRecordBuilder) GetTime(gettime string) *IndicesRecordBuilder {
	rb.v.GetTime = &gettime
	return rb
}

// GetTotal number of get ops

func (rb *IndicesRecordBuilder) GetTotal(gettotal string) *IndicesRecordBuilder {
	rb.v.GetTotal = &gettotal
	return rb
}

// Health current health status

func (rb *IndicesRecordBuilder) Health(health string) *IndicesRecordBuilder {
	rb.v.Health = &health
	return rb
}

// Index index name

func (rb *IndicesRecordBuilder) Index(index string) *IndicesRecordBuilder {
	rb.v.Index = &index
	return rb
}

// IndexingDeleteCurrent number of current deletions

func (rb *IndicesRecordBuilder) IndexingDeleteCurrent(indexingdeletecurrent string) *IndicesRecordBuilder {
	rb.v.IndexingDeleteCurrent = &indexingdeletecurrent
	return rb
}

// IndexingDeleteTime time spent in deletions

func (rb *IndicesRecordBuilder) IndexingDeleteTime(indexingdeletetime string) *IndicesRecordBuilder {
	rb.v.IndexingDeleteTime = &indexingdeletetime
	return rb
}

// IndexingDeleteTotal number of delete ops

func (rb *IndicesRecordBuilder) IndexingDeleteTotal(indexingdeletetotal string) *IndicesRecordBuilder {
	rb.v.IndexingDeleteTotal = &indexingdeletetotal
	return rb
}

// IndexingIndexCurrent number of current indexing ops

func (rb *IndicesRecordBuilder) IndexingIndexCurrent(indexingindexcurrent string) *IndicesRecordBuilder {
	rb.v.IndexingIndexCurrent = &indexingindexcurrent
	return rb
}

// IndexingIndexFailed number of failed indexing ops

func (rb *IndicesRecordBuilder) IndexingIndexFailed(indexingindexfailed string) *IndicesRecordBuilder {
	rb.v.IndexingIndexFailed = &indexingindexfailed
	return rb
}

// IndexingIndexTime time spent in indexing

func (rb *IndicesRecordBuilder) IndexingIndexTime(indexingindextime string) *IndicesRecordBuilder {
	rb.v.IndexingIndexTime = &indexingindextime
	return rb
}

// IndexingIndexTotal number of indexing ops

func (rb *IndicesRecordBuilder) IndexingIndexTotal(indexingindextotal string) *IndicesRecordBuilder {
	rb.v.IndexingIndexTotal = &indexingindextotal
	return rb
}

// MemoryTotal total used memory

func (rb *IndicesRecordBuilder) MemoryTotal(memorytotal string) *IndicesRecordBuilder {
	rb.v.MemoryTotal = &memorytotal
	return rb
}

// MergesCurrent number of current merges

func (rb *IndicesRecordBuilder) MergesCurrent(mergescurrent string) *IndicesRecordBuilder {
	rb.v.MergesCurrent = &mergescurrent
	return rb
}

// MergesCurrentDocs number of current merging docs

func (rb *IndicesRecordBuilder) MergesCurrentDocs(mergescurrentdocs string) *IndicesRecordBuilder {
	rb.v.MergesCurrentDocs = &mergescurrentdocs
	return rb
}

// MergesCurrentSize size of current merges

func (rb *IndicesRecordBuilder) MergesCurrentSize(mergescurrentsize string) *IndicesRecordBuilder {
	rb.v.MergesCurrentSize = &mergescurrentsize
	return rb
}

// MergesTotal number of completed merge ops

func (rb *IndicesRecordBuilder) MergesTotal(mergestotal string) *IndicesRecordBuilder {
	rb.v.MergesTotal = &mergestotal
	return rb
}

// MergesTotalDocs docs merged

func (rb *IndicesRecordBuilder) MergesTotalDocs(mergestotaldocs string) *IndicesRecordBuilder {
	rb.v.MergesTotalDocs = &mergestotaldocs
	return rb
}

// MergesTotalSize size merged

func (rb *IndicesRecordBuilder) MergesTotalSize(mergestotalsize string) *IndicesRecordBuilder {
	rb.v.MergesTotalSize = &mergestotalsize
	return rb
}

// MergesTotalTime time spent in merges

func (rb *IndicesRecordBuilder) MergesTotalTime(mergestotaltime string) *IndicesRecordBuilder {
	rb.v.MergesTotalTime = &mergestotaltime
	return rb
}

// Pri number of primary shards

func (rb *IndicesRecordBuilder) Pri(pri string) *IndicesRecordBuilder {
	rb.v.Pri = &pri
	return rb
}

// PriBulkAvgSizeInBytes average size in bytes of shard bulk

func (rb *IndicesRecordBuilder) PriBulkAvgSizeInBytes(pribulkavgsizeinbytes string) *IndicesRecordBuilder {
	rb.v.PriBulkAvgSizeInBytes = &pribulkavgsizeinbytes
	return rb
}

// PriBulkAvgTime average time spend in shard bulk

func (rb *IndicesRecordBuilder) PriBulkAvgTime(pribulkavgtime string) *IndicesRecordBuilder {
	rb.v.PriBulkAvgTime = &pribulkavgtime
	return rb
}

// PriBulkTotalOperations number of bulk shard ops

func (rb *IndicesRecordBuilder) PriBulkTotalOperations(pribulktotaloperations string) *IndicesRecordBuilder {
	rb.v.PriBulkTotalOperations = &pribulktotaloperations
	return rb
}

// PriBulkTotalSizeInBytes total size in bytes of shard bulk

func (rb *IndicesRecordBuilder) PriBulkTotalSizeInBytes(pribulktotalsizeinbytes string) *IndicesRecordBuilder {
	rb.v.PriBulkTotalSizeInBytes = &pribulktotalsizeinbytes
	return rb
}

// PriBulkTotalTime time spend in shard bulk

func (rb *IndicesRecordBuilder) PriBulkTotalTime(pribulktotaltime string) *IndicesRecordBuilder {
	rb.v.PriBulkTotalTime = &pribulktotaltime
	return rb
}

// PriCompletionSize size of completion

func (rb *IndicesRecordBuilder) PriCompletionSize(pricompletionsize string) *IndicesRecordBuilder {
	rb.v.PriCompletionSize = &pricompletionsize
	return rb
}

// PriFielddataEvictions fielddata evictions

func (rb *IndicesRecordBuilder) PriFielddataEvictions(prifielddataevictions string) *IndicesRecordBuilder {
	rb.v.PriFielddataEvictions = &prifielddataevictions
	return rb
}

// PriFielddataMemorySize used fielddata cache

func (rb *IndicesRecordBuilder) PriFielddataMemorySize(prifielddatamemorysize string) *IndicesRecordBuilder {
	rb.v.PriFielddataMemorySize = &prifielddatamemorysize
	return rb
}

// PriFlushTotal number of flushes

func (rb *IndicesRecordBuilder) PriFlushTotal(priflushtotal string) *IndicesRecordBuilder {
	rb.v.PriFlushTotal = &priflushtotal
	return rb
}

// PriFlushTotalTime time spent in flush

func (rb *IndicesRecordBuilder) PriFlushTotalTime(priflushtotaltime string) *IndicesRecordBuilder {
	rb.v.PriFlushTotalTime = &priflushtotaltime
	return rb
}

// PriGetCurrent number of current get ops

func (rb *IndicesRecordBuilder) PriGetCurrent(prigetcurrent string) *IndicesRecordBuilder {
	rb.v.PriGetCurrent = &prigetcurrent
	return rb
}

// PriGetExistsTime time spent in successful gets

func (rb *IndicesRecordBuilder) PriGetExistsTime(prigetexiststime string) *IndicesRecordBuilder {
	rb.v.PriGetExistsTime = &prigetexiststime
	return rb
}

// PriGetExistsTotal number of successful gets

func (rb *IndicesRecordBuilder) PriGetExistsTotal(prigetexiststotal string) *IndicesRecordBuilder {
	rb.v.PriGetExistsTotal = &prigetexiststotal
	return rb
}

// PriGetMissingTime time spent in failed gets

func (rb *IndicesRecordBuilder) PriGetMissingTime(prigetmissingtime string) *IndicesRecordBuilder {
	rb.v.PriGetMissingTime = &prigetmissingtime
	return rb
}

// PriGetMissingTotal number of failed gets

func (rb *IndicesRecordBuilder) PriGetMissingTotal(prigetmissingtotal string) *IndicesRecordBuilder {
	rb.v.PriGetMissingTotal = &prigetmissingtotal
	return rb
}

// PriGetTime time spent in get

func (rb *IndicesRecordBuilder) PriGetTime(prigettime string) *IndicesRecordBuilder {
	rb.v.PriGetTime = &prigettime
	return rb
}

// PriGetTotal number of get ops

func (rb *IndicesRecordBuilder) PriGetTotal(prigettotal string) *IndicesRecordBuilder {
	rb.v.PriGetTotal = &prigettotal
	return rb
}

// PriIndexingDeleteCurrent number of current deletions

func (rb *IndicesRecordBuilder) PriIndexingDeleteCurrent(priindexingdeletecurrent string) *IndicesRecordBuilder {
	rb.v.PriIndexingDeleteCurrent = &priindexingdeletecurrent
	return rb
}

// PriIndexingDeleteTime time spent in deletions

func (rb *IndicesRecordBuilder) PriIndexingDeleteTime(priindexingdeletetime string) *IndicesRecordBuilder {
	rb.v.PriIndexingDeleteTime = &priindexingdeletetime
	return rb
}

// PriIndexingDeleteTotal number of delete ops

func (rb *IndicesRecordBuilder) PriIndexingDeleteTotal(priindexingdeletetotal string) *IndicesRecordBuilder {
	rb.v.PriIndexingDeleteTotal = &priindexingdeletetotal
	return rb
}

// PriIndexingIndexCurrent number of current indexing ops

func (rb *IndicesRecordBuilder) PriIndexingIndexCurrent(priindexingindexcurrent string) *IndicesRecordBuilder {
	rb.v.PriIndexingIndexCurrent = &priindexingindexcurrent
	return rb
}

// PriIndexingIndexFailed number of failed indexing ops

func (rb *IndicesRecordBuilder) PriIndexingIndexFailed(priindexingindexfailed string) *IndicesRecordBuilder {
	rb.v.PriIndexingIndexFailed = &priindexingindexfailed
	return rb
}

// PriIndexingIndexTime time spent in indexing

func (rb *IndicesRecordBuilder) PriIndexingIndexTime(priindexingindextime string) *IndicesRecordBuilder {
	rb.v.PriIndexingIndexTime = &priindexingindextime
	return rb
}

// PriIndexingIndexTotal number of indexing ops

func (rb *IndicesRecordBuilder) PriIndexingIndexTotal(priindexingindextotal string) *IndicesRecordBuilder {
	rb.v.PriIndexingIndexTotal = &priindexingindextotal
	return rb
}

// PriMemoryTotal total user memory

func (rb *IndicesRecordBuilder) PriMemoryTotal(primemorytotal string) *IndicesRecordBuilder {
	rb.v.PriMemoryTotal = &primemorytotal
	return rb
}

// PriMergesCurrent number of current merges

func (rb *IndicesRecordBuilder) PriMergesCurrent(primergescurrent string) *IndicesRecordBuilder {
	rb.v.PriMergesCurrent = &primergescurrent
	return rb
}

// PriMergesCurrentDocs number of current merging docs

func (rb *IndicesRecordBuilder) PriMergesCurrentDocs(primergescurrentdocs string) *IndicesRecordBuilder {
	rb.v.PriMergesCurrentDocs = &primergescurrentdocs
	return rb
}

// PriMergesCurrentSize size of current merges

func (rb *IndicesRecordBuilder) PriMergesCurrentSize(primergescurrentsize string) *IndicesRecordBuilder {
	rb.v.PriMergesCurrentSize = &primergescurrentsize
	return rb
}

// PriMergesTotal number of completed merge ops

func (rb *IndicesRecordBuilder) PriMergesTotal(primergestotal string) *IndicesRecordBuilder {
	rb.v.PriMergesTotal = &primergestotal
	return rb
}

// PriMergesTotalDocs docs merged

func (rb *IndicesRecordBuilder) PriMergesTotalDocs(primergestotaldocs string) *IndicesRecordBuilder {
	rb.v.PriMergesTotalDocs = &primergestotaldocs
	return rb
}

// PriMergesTotalSize size merged

func (rb *IndicesRecordBuilder) PriMergesTotalSize(primergestotalsize string) *IndicesRecordBuilder {
	rb.v.PriMergesTotalSize = &primergestotalsize
	return rb
}

// PriMergesTotalTime time spent in merges

func (rb *IndicesRecordBuilder) PriMergesTotalTime(primergestotaltime string) *IndicesRecordBuilder {
	rb.v.PriMergesTotalTime = &primergestotaltime
	return rb
}

// PriQueryCacheEvictions query cache evictions

func (rb *IndicesRecordBuilder) PriQueryCacheEvictions(priquerycacheevictions string) *IndicesRecordBuilder {
	rb.v.PriQueryCacheEvictions = &priquerycacheevictions
	return rb
}

// PriQueryCacheMemorySize used query cache

func (rb *IndicesRecordBuilder) PriQueryCacheMemorySize(priquerycachememorysize string) *IndicesRecordBuilder {
	rb.v.PriQueryCacheMemorySize = &priquerycachememorysize
	return rb
}

// PriRefreshExternalTime time spent in external refreshes

func (rb *IndicesRecordBuilder) PriRefreshExternalTime(prirefreshexternaltime string) *IndicesRecordBuilder {
	rb.v.PriRefreshExternalTime = &prirefreshexternaltime
	return rb
}

// PriRefreshExternalTotal total external refreshes

func (rb *IndicesRecordBuilder) PriRefreshExternalTotal(prirefreshexternaltotal string) *IndicesRecordBuilder {
	rb.v.PriRefreshExternalTotal = &prirefreshexternaltotal
	return rb
}

// PriRefreshListeners number of pending refresh listeners

func (rb *IndicesRecordBuilder) PriRefreshListeners(prirefreshlisteners string) *IndicesRecordBuilder {
	rb.v.PriRefreshListeners = &prirefreshlisteners
	return rb
}

// PriRefreshTime time spent in refreshes

func (rb *IndicesRecordBuilder) PriRefreshTime(prirefreshtime string) *IndicesRecordBuilder {
	rb.v.PriRefreshTime = &prirefreshtime
	return rb
}

// PriRefreshTotal total refreshes

func (rb *IndicesRecordBuilder) PriRefreshTotal(prirefreshtotal string) *IndicesRecordBuilder {
	rb.v.PriRefreshTotal = &prirefreshtotal
	return rb
}

// PriRequestCacheEvictions request cache evictions

func (rb *IndicesRecordBuilder) PriRequestCacheEvictions(prirequestcacheevictions string) *IndicesRecordBuilder {
	rb.v.PriRequestCacheEvictions = &prirequestcacheevictions
	return rb
}

// PriRequestCacheHitCount request cache hit count

func (rb *IndicesRecordBuilder) PriRequestCacheHitCount(prirequestcachehitcount string) *IndicesRecordBuilder {
	rb.v.PriRequestCacheHitCount = &prirequestcachehitcount
	return rb
}

// PriRequestCacheMemorySize used request cache

func (rb *IndicesRecordBuilder) PriRequestCacheMemorySize(prirequestcachememorysize string) *IndicesRecordBuilder {
	rb.v.PriRequestCacheMemorySize = &prirequestcachememorysize
	return rb
}

// PriRequestCacheMissCount request cache miss count

func (rb *IndicesRecordBuilder) PriRequestCacheMissCount(prirequestcachemisscount string) *IndicesRecordBuilder {
	rb.v.PriRequestCacheMissCount = &prirequestcachemisscount
	return rb
}

// PriSearchFetchCurrent current fetch phase ops

func (rb *IndicesRecordBuilder) PriSearchFetchCurrent(prisearchfetchcurrent string) *IndicesRecordBuilder {
	rb.v.PriSearchFetchCurrent = &prisearchfetchcurrent
	return rb
}

// PriSearchFetchTime time spent in fetch phase

func (rb *IndicesRecordBuilder) PriSearchFetchTime(prisearchfetchtime string) *IndicesRecordBuilder {
	rb.v.PriSearchFetchTime = &prisearchfetchtime
	return rb
}

// PriSearchFetchTotal total fetch ops

func (rb *IndicesRecordBuilder) PriSearchFetchTotal(prisearchfetchtotal string) *IndicesRecordBuilder {
	rb.v.PriSearchFetchTotal = &prisearchfetchtotal
	return rb
}

// PriSearchOpenContexts open search contexts

func (rb *IndicesRecordBuilder) PriSearchOpenContexts(prisearchopencontexts string) *IndicesRecordBuilder {
	rb.v.PriSearchOpenContexts = &prisearchopencontexts
	return rb
}

// PriSearchQueryCurrent current query phase ops

func (rb *IndicesRecordBuilder) PriSearchQueryCurrent(prisearchquerycurrent string) *IndicesRecordBuilder {
	rb.v.PriSearchQueryCurrent = &prisearchquerycurrent
	return rb
}

// PriSearchQueryTime time spent in query phase

func (rb *IndicesRecordBuilder) PriSearchQueryTime(prisearchquerytime string) *IndicesRecordBuilder {
	rb.v.PriSearchQueryTime = &prisearchquerytime
	return rb
}

// PriSearchQueryTotal total query phase ops

func (rb *IndicesRecordBuilder) PriSearchQueryTotal(prisearchquerytotal string) *IndicesRecordBuilder {
	rb.v.PriSearchQueryTotal = &prisearchquerytotal
	return rb
}

// PriSearchScrollCurrent open scroll contexts

func (rb *IndicesRecordBuilder) PriSearchScrollCurrent(prisearchscrollcurrent string) *IndicesRecordBuilder {
	rb.v.PriSearchScrollCurrent = &prisearchscrollcurrent
	return rb
}

// PriSearchScrollTime time scroll contexts held open

func (rb *IndicesRecordBuilder) PriSearchScrollTime(prisearchscrolltime string) *IndicesRecordBuilder {
	rb.v.PriSearchScrollTime = &prisearchscrolltime
	return rb
}

// PriSearchScrollTotal completed scroll contexts

func (rb *IndicesRecordBuilder) PriSearchScrollTotal(prisearchscrolltotal string) *IndicesRecordBuilder {
	rb.v.PriSearchScrollTotal = &prisearchscrolltotal
	return rb
}

// PriSegmentsCount number of segments

func (rb *IndicesRecordBuilder) PriSegmentsCount(prisegmentscount string) *IndicesRecordBuilder {
	rb.v.PriSegmentsCount = &prisegmentscount
	return rb
}

// PriSegmentsFixedBitsetMemory memory used by fixed bit sets for nested object field types and export type
// filters for types referred in _parent fields

func (rb *IndicesRecordBuilder) PriSegmentsFixedBitsetMemory(prisegmentsfixedbitsetmemory string) *IndicesRecordBuilder {
	rb.v.PriSegmentsFixedBitsetMemory = &prisegmentsfixedbitsetmemory
	return rb
}

// PriSegmentsIndexWriterMemory memory used by index writer

func (rb *IndicesRecordBuilder) PriSegmentsIndexWriterMemory(prisegmentsindexwritermemory string) *IndicesRecordBuilder {
	rb.v.PriSegmentsIndexWriterMemory = &prisegmentsindexwritermemory
	return rb
}

// PriSegmentsMemory memory used by segments

func (rb *IndicesRecordBuilder) PriSegmentsMemory(prisegmentsmemory string) *IndicesRecordBuilder {
	rb.v.PriSegmentsMemory = &prisegmentsmemory
	return rb
}

// PriSegmentsVersionMapMemory memory used by version map

func (rb *IndicesRecordBuilder) PriSegmentsVersionMapMemory(prisegmentsversionmapmemory string) *IndicesRecordBuilder {
	rb.v.PriSegmentsVersionMapMemory = &prisegmentsversionmapmemory
	return rb
}

// PriStoreSize store size of primaries

func (rb *IndicesRecordBuilder) PriStoreSize(pristoresize string) *IndicesRecordBuilder {
	rb.v.PriStoreSize = pristoresize
	return rb
}

// PriSuggestCurrent number of current suggest ops

func (rb *IndicesRecordBuilder) PriSuggestCurrent(prisuggestcurrent string) *IndicesRecordBuilder {
	rb.v.PriSuggestCurrent = &prisuggestcurrent
	return rb
}

// PriSuggestTime time spend in suggest

func (rb *IndicesRecordBuilder) PriSuggestTime(prisuggesttime string) *IndicesRecordBuilder {
	rb.v.PriSuggestTime = &prisuggesttime
	return rb
}

// PriSuggestTotal number of suggest ops

func (rb *IndicesRecordBuilder) PriSuggestTotal(prisuggesttotal string) *IndicesRecordBuilder {
	rb.v.PriSuggestTotal = &prisuggesttotal
	return rb
}

// PriWarmerCurrent current warmer ops

func (rb *IndicesRecordBuilder) PriWarmerCurrent(priwarmercurrent string) *IndicesRecordBuilder {
	rb.v.PriWarmerCurrent = &priwarmercurrent
	return rb
}

// PriWarmerTotal total warmer ops

func (rb *IndicesRecordBuilder) PriWarmerTotal(priwarmertotal string) *IndicesRecordBuilder {
	rb.v.PriWarmerTotal = &priwarmertotal
	return rb
}

// PriWarmerTotalTime time spent in warmers

func (rb *IndicesRecordBuilder) PriWarmerTotalTime(priwarmertotaltime string) *IndicesRecordBuilder {
	rb.v.PriWarmerTotalTime = &priwarmertotaltime
	return rb
}

// QueryCacheEvictions query cache evictions

func (rb *IndicesRecordBuilder) QueryCacheEvictions(querycacheevictions string) *IndicesRecordBuilder {
	rb.v.QueryCacheEvictions = &querycacheevictions
	return rb
}

// QueryCacheMemorySize used query cache

func (rb *IndicesRecordBuilder) QueryCacheMemorySize(querycachememorysize string) *IndicesRecordBuilder {
	rb.v.QueryCacheMemorySize = &querycachememorysize
	return rb
}

// RefreshExternalTime time spent in external refreshes

func (rb *IndicesRecordBuilder) RefreshExternalTime(refreshexternaltime string) *IndicesRecordBuilder {
	rb.v.RefreshExternalTime = &refreshexternaltime
	return rb
}

// RefreshExternalTotal total external refreshes

func (rb *IndicesRecordBuilder) RefreshExternalTotal(refreshexternaltotal string) *IndicesRecordBuilder {
	rb.v.RefreshExternalTotal = &refreshexternaltotal
	return rb
}

// RefreshListeners number of pending refresh listeners

func (rb *IndicesRecordBuilder) RefreshListeners(refreshlisteners string) *IndicesRecordBuilder {
	rb.v.RefreshListeners = &refreshlisteners
	return rb
}

// RefreshTime time spent in refreshes

func (rb *IndicesRecordBuilder) RefreshTime(refreshtime string) *IndicesRecordBuilder {
	rb.v.RefreshTime = &refreshtime
	return rb
}

// RefreshTotal total refreshes

func (rb *IndicesRecordBuilder) RefreshTotal(refreshtotal string) *IndicesRecordBuilder {
	rb.v.RefreshTotal = &refreshtotal
	return rb
}

// Rep number of replica shards

func (rb *IndicesRecordBuilder) Rep(rep string) *IndicesRecordBuilder {
	rb.v.Rep = &rep
	return rb
}

// RequestCacheEvictions request cache evictions

func (rb *IndicesRecordBuilder) RequestCacheEvictions(requestcacheevictions string) *IndicesRecordBuilder {
	rb.v.RequestCacheEvictions = &requestcacheevictions
	return rb
}

// RequestCacheHitCount request cache hit count

func (rb *IndicesRecordBuilder) RequestCacheHitCount(requestcachehitcount string) *IndicesRecordBuilder {
	rb.v.RequestCacheHitCount = &requestcachehitcount
	return rb
}

// RequestCacheMemorySize used request cache

func (rb *IndicesRecordBuilder) RequestCacheMemorySize(requestcachememorysize string) *IndicesRecordBuilder {
	rb.v.RequestCacheMemorySize = &requestcachememorysize
	return rb
}

// RequestCacheMissCount request cache miss count

func (rb *IndicesRecordBuilder) RequestCacheMissCount(requestcachemisscount string) *IndicesRecordBuilder {
	rb.v.RequestCacheMissCount = &requestcachemisscount
	return rb
}

// SearchFetchCurrent current fetch phase ops

func (rb *IndicesRecordBuilder) SearchFetchCurrent(searchfetchcurrent string) *IndicesRecordBuilder {
	rb.v.SearchFetchCurrent = &searchfetchcurrent
	return rb
}

// SearchFetchTime time spent in fetch phase

func (rb *IndicesRecordBuilder) SearchFetchTime(searchfetchtime string) *IndicesRecordBuilder {
	rb.v.SearchFetchTime = &searchfetchtime
	return rb
}

// SearchFetchTotal total fetch ops

func (rb *IndicesRecordBuilder) SearchFetchTotal(searchfetchtotal string) *IndicesRecordBuilder {
	rb.v.SearchFetchTotal = &searchfetchtotal
	return rb
}

// SearchOpenContexts open search contexts

func (rb *IndicesRecordBuilder) SearchOpenContexts(searchopencontexts string) *IndicesRecordBuilder {
	rb.v.SearchOpenContexts = &searchopencontexts
	return rb
}

// SearchQueryCurrent current query phase ops

func (rb *IndicesRecordBuilder) SearchQueryCurrent(searchquerycurrent string) *IndicesRecordBuilder {
	rb.v.SearchQueryCurrent = &searchquerycurrent
	return rb
}

// SearchQueryTime time spent in query phase

func (rb *IndicesRecordBuilder) SearchQueryTime(searchquerytime string) *IndicesRecordBuilder {
	rb.v.SearchQueryTime = &searchquerytime
	return rb
}

// SearchQueryTotal total query phase ops

func (rb *IndicesRecordBuilder) SearchQueryTotal(searchquerytotal string) *IndicesRecordBuilder {
	rb.v.SearchQueryTotal = &searchquerytotal
	return rb
}

// SearchScrollCurrent open scroll contexts

func (rb *IndicesRecordBuilder) SearchScrollCurrent(searchscrollcurrent string) *IndicesRecordBuilder {
	rb.v.SearchScrollCurrent = &searchscrollcurrent
	return rb
}

// SearchScrollTime time scroll contexts held open

func (rb *IndicesRecordBuilder) SearchScrollTime(searchscrolltime string) *IndicesRecordBuilder {
	rb.v.SearchScrollTime = &searchscrolltime
	return rb
}

// SearchScrollTotal completed scroll contexts

func (rb *IndicesRecordBuilder) SearchScrollTotal(searchscrolltotal string) *IndicesRecordBuilder {
	rb.v.SearchScrollTotal = &searchscrolltotal
	return rb
}

// SearchThrottled indicates if the index is search throttled

func (rb *IndicesRecordBuilder) SearchThrottled(searchthrottled string) *IndicesRecordBuilder {
	rb.v.SearchThrottled = &searchthrottled
	return rb
}

// SegmentsCount number of segments

func (rb *IndicesRecordBuilder) SegmentsCount(segmentscount string) *IndicesRecordBuilder {
	rb.v.SegmentsCount = &segmentscount
	return rb
}

// SegmentsFixedBitsetMemory memory used by fixed bit sets for nested object field types and export type
// filters for types referred in _parent fields

func (rb *IndicesRecordBuilder) SegmentsFixedBitsetMemory(segmentsfixedbitsetmemory string) *IndicesRecordBuilder {
	rb.v.SegmentsFixedBitsetMemory = &segmentsfixedbitsetmemory
	return rb
}

// SegmentsIndexWriterMemory memory used by index writer

func (rb *IndicesRecordBuilder) SegmentsIndexWriterMemory(segmentsindexwritermemory string) *IndicesRecordBuilder {
	rb.v.SegmentsIndexWriterMemory = &segmentsindexwritermemory
	return rb
}

// SegmentsMemory memory used by segments

func (rb *IndicesRecordBuilder) SegmentsMemory(segmentsmemory string) *IndicesRecordBuilder {
	rb.v.SegmentsMemory = &segmentsmemory
	return rb
}

// SegmentsVersionMapMemory memory used by version map

func (rb *IndicesRecordBuilder) SegmentsVersionMapMemory(segmentsversionmapmemory string) *IndicesRecordBuilder {
	rb.v.SegmentsVersionMapMemory = &segmentsversionmapmemory
	return rb
}

// Status open/close status

func (rb *IndicesRecordBuilder) Status(status string) *IndicesRecordBuilder {
	rb.v.Status = &status
	return rb
}

// StoreSize store size of primaries & replicas

func (rb *IndicesRecordBuilder) StoreSize(storesize string) *IndicesRecordBuilder {
	rb.v.StoreSize = storesize
	return rb
}

// SuggestCurrent number of current suggest ops

func (rb *IndicesRecordBuilder) SuggestCurrent(suggestcurrent string) *IndicesRecordBuilder {
	rb.v.SuggestCurrent = &suggestcurrent
	return rb
}

// SuggestTime time spend in suggest

func (rb *IndicesRecordBuilder) SuggestTime(suggesttime string) *IndicesRecordBuilder {
	rb.v.SuggestTime = &suggesttime
	return rb
}

// SuggestTotal number of suggest ops

func (rb *IndicesRecordBuilder) SuggestTotal(suggesttotal string) *IndicesRecordBuilder {
	rb.v.SuggestTotal = &suggesttotal
	return rb
}

// Uuid index uuid

func (rb *IndicesRecordBuilder) Uuid(uuid string) *IndicesRecordBuilder {
	rb.v.Uuid = &uuid
	return rb
}

// WarmerCurrent current warmer ops

func (rb *IndicesRecordBuilder) WarmerCurrent(warmercurrent string) *IndicesRecordBuilder {
	rb.v.WarmerCurrent = &warmercurrent
	return rb
}

// WarmerTotal total warmer ops

func (rb *IndicesRecordBuilder) WarmerTotal(warmertotal string) *IndicesRecordBuilder {
	rb.v.WarmerTotal = &warmertotal
	return rb
}

// WarmerTotalTime time spent in warmers

func (rb *IndicesRecordBuilder) WarmerTotalTime(warmertotaltime string) *IndicesRecordBuilder {
	rb.v.WarmerTotalTime = &warmertotaltime
	return rb
}
