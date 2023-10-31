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

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// IndicesRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/cat/indices/types.ts#L20-L801
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

func (s *IndicesRecord) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "bulk.avg_size_in_bytes", "basi", "bulkAvgSizeInBytes":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.BulkAvgSizeInBytes = &o

		case "bulk.avg_time", "bati", "bulkAvgTime":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.BulkAvgTime = &o

		case "bulk.total_operations", "bto", "bulkTotalOperation":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.BulkTotalOperations = &o

		case "bulk.total_size_in_bytes", "btsi", "bulkTotalSizeInBytes":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.BulkTotalSizeInBytes = &o

		case "bulk.total_time", "btti", "bulkTotalTime":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.BulkTotalTime = &o

		case "completion.size", "cs", "completionSize":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.CompletionSize = &o

		case "creation.date", "cd":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.CreationDate = &o

		case "creation.date.string", "cds":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.CreationDateString = &o

		case "docs.count", "dc", "docsCount":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DocsCount = o

		case "docs.deleted", "dd", "docsDeleted":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DocsDeleted = o

		case "fielddata.evictions", "fe", "fielddataEvictions":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.FielddataEvictions = &o

		case "fielddata.memory_size", "fm", "fielddataMemory":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.FielddataMemorySize = &o

		case "flush.total", "ft", "flushTotal":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.FlushTotal = &o

		case "flush.total_time", "ftt", "flushTotalTime":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.FlushTotalTime = &o

		case "get.current", "gc", "getCurrent":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.GetCurrent = &o

		case "get.exists_time", "geti", "getExistsTime":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.GetExistsTime = &o

		case "get.exists_total", "geto", "getExistsTotal":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.GetExistsTotal = &o

		case "get.missing_time", "gmti", "getMissingTime":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.GetMissingTime = &o

		case "get.missing_total", "gmto", "getMissingTotal":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.GetMissingTotal = &o

		case "get.time", "gti", "getTime":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.GetTime = &o

		case "get.total", "gto", "getTotal":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.GetTotal = &o

		case "health", "h":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Health = &o

		case "index", "i", "idx":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Index = &o

		case "indexing.delete_current", "idc", "indexingDeleteCurrent":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.IndexingDeleteCurrent = &o

		case "indexing.delete_time", "idti", "indexingDeleteTime":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.IndexingDeleteTime = &o

		case "indexing.delete_total", "idto", "indexingDeleteTotal":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.IndexingDeleteTotal = &o

		case "indexing.index_current", "iic", "indexingIndexCurrent":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.IndexingIndexCurrent = &o

		case "indexing.index_failed", "iif", "indexingIndexFailed":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.IndexingIndexFailed = &o

		case "indexing.index_time", "iiti", "indexingIndexTime":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.IndexingIndexTime = &o

		case "indexing.index_total", "iito", "indexingIndexTotal":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.IndexingIndexTotal = &o

		case "memory.total", "tm", "memoryTotal":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MemoryTotal = &o

		case "merges.current", "mc", "mergesCurrent":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MergesCurrent = &o

		case "merges.current_docs", "mcd", "mergesCurrentDocs":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MergesCurrentDocs = &o

		case "merges.current_size", "mcs", "mergesCurrentSize":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MergesCurrentSize = &o

		case "merges.total", "mt", "mergesTotal":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MergesTotal = &o

		case "merges.total_docs", "mtd", "mergesTotalDocs":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MergesTotalDocs = &o

		case "merges.total_size", "mts", "mergesTotalSize":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MergesTotalSize = &o

		case "merges.total_time", "mtt", "mergesTotalTime":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MergesTotalTime = &o

		case "pri", "p", "shards.primary", "shardsPrimary":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Pri = &o

		case "pri.bulk.avg_size_in_bytes":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriBulkAvgSizeInBytes = &o

		case "pri.bulk.avg_time":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriBulkAvgTime = &o

		case "pri.bulk.total_operations":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriBulkTotalOperations = &o

		case "pri.bulk.total_size_in_bytes":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriBulkTotalSizeInBytes = &o

		case "pri.bulk.total_time":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriBulkTotalTime = &o

		case "pri.completion.size":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriCompletionSize = &o

		case "pri.fielddata.evictions":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriFielddataEvictions = &o

		case "pri.fielddata.memory_size":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriFielddataMemorySize = &o

		case "pri.flush.total":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriFlushTotal = &o

		case "pri.flush.total_time":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriFlushTotalTime = &o

		case "pri.get.current":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriGetCurrent = &o

		case "pri.get.exists_time":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriGetExistsTime = &o

		case "pri.get.exists_total":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriGetExistsTotal = &o

		case "pri.get.missing_time":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriGetMissingTime = &o

		case "pri.get.missing_total":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriGetMissingTotal = &o

		case "pri.get.time":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriGetTime = &o

		case "pri.get.total":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriGetTotal = &o

		case "pri.indexing.delete_current":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriIndexingDeleteCurrent = &o

		case "pri.indexing.delete_time":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriIndexingDeleteTime = &o

		case "pri.indexing.delete_total":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriIndexingDeleteTotal = &o

		case "pri.indexing.index_current":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriIndexingIndexCurrent = &o

		case "pri.indexing.index_failed":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriIndexingIndexFailed = &o

		case "pri.indexing.index_time":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriIndexingIndexTime = &o

		case "pri.indexing.index_total":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriIndexingIndexTotal = &o

		case "pri.memory.total":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriMemoryTotal = &o

		case "pri.merges.current":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriMergesCurrent = &o

		case "pri.merges.current_docs":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriMergesCurrentDocs = &o

		case "pri.merges.current_size":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriMergesCurrentSize = &o

		case "pri.merges.total":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriMergesTotal = &o

		case "pri.merges.total_docs":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriMergesTotalDocs = &o

		case "pri.merges.total_size":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriMergesTotalSize = &o

		case "pri.merges.total_time":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriMergesTotalTime = &o

		case "pri.query_cache.evictions":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriQueryCacheEvictions = &o

		case "pri.query_cache.memory_size":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriQueryCacheMemorySize = &o

		case "pri.refresh.external_time":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriRefreshExternalTime = &o

		case "pri.refresh.external_total":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriRefreshExternalTotal = &o

		case "pri.refresh.listeners":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriRefreshListeners = &o

		case "pri.refresh.time":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriRefreshTime = &o

		case "pri.refresh.total":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriRefreshTotal = &o

		case "pri.request_cache.evictions":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriRequestCacheEvictions = &o

		case "pri.request_cache.hit_count":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriRequestCacheHitCount = &o

		case "pri.request_cache.memory_size":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriRequestCacheMemorySize = &o

		case "pri.request_cache.miss_count":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriRequestCacheMissCount = &o

		case "pri.search.fetch_current":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriSearchFetchCurrent = &o

		case "pri.search.fetch_time":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriSearchFetchTime = &o

		case "pri.search.fetch_total":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriSearchFetchTotal = &o

		case "pri.search.open_contexts":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriSearchOpenContexts = &o

		case "pri.search.query_current":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriSearchQueryCurrent = &o

		case "pri.search.query_time":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriSearchQueryTime = &o

		case "pri.search.query_total":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriSearchQueryTotal = &o

		case "pri.search.scroll_current":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriSearchScrollCurrent = &o

		case "pri.search.scroll_time":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriSearchScrollTime = &o

		case "pri.search.scroll_total":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriSearchScrollTotal = &o

		case "pri.segments.count":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriSegmentsCount = &o

		case "pri.segments.fixed_bitset_memory":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriSegmentsFixedBitsetMemory = &o

		case "pri.segments.index_writer_memory":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriSegmentsIndexWriterMemory = &o

		case "pri.segments.memory":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriSegmentsMemory = &o

		case "pri.segments.version_map_memory":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriSegmentsVersionMapMemory = &o

		case "pri.store.size":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriStoreSize = o

		case "pri.suggest.current":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriSuggestCurrent = &o

		case "pri.suggest.time":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriSuggestTime = &o

		case "pri.suggest.total":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriSuggestTotal = &o

		case "pri.warmer.current":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriWarmerCurrent = &o

		case "pri.warmer.total":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriWarmerTotal = &o

		case "pri.warmer.total_time":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PriWarmerTotalTime = &o

		case "query_cache.evictions", "qce", "queryCacheEvictions":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.QueryCacheEvictions = &o

		case "query_cache.memory_size", "qcm", "queryCacheMemory":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.QueryCacheMemorySize = &o

		case "refresh.external_time", "reti":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.RefreshExternalTime = &o

		case "refresh.external_total", "reto":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.RefreshExternalTotal = &o

		case "refresh.listeners", "rli", "refreshListeners":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.RefreshListeners = &o

		case "refresh.time", "rti", "refreshTime":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.RefreshTime = &o

		case "refresh.total", "rto", "refreshTotal":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.RefreshTotal = &o

		case "rep", "r", "shards.replica", "shardsReplica":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Rep = &o

		case "request_cache.evictions", "rce", "requestCacheEvictions":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.RequestCacheEvictions = &o

		case "request_cache.hit_count", "rchc", "requestCacheHitCount":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.RequestCacheHitCount = &o

		case "request_cache.memory_size", "rcm", "requestCacheMemory":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.RequestCacheMemorySize = &o

		case "request_cache.miss_count", "rcmc", "requestCacheMissCount":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.RequestCacheMissCount = &o

		case "search.fetch_current", "sfc", "searchFetchCurrent":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SearchFetchCurrent = &o

		case "search.fetch_time", "sfti", "searchFetchTime":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SearchFetchTime = &o

		case "search.fetch_total", "sfto", "searchFetchTotal":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SearchFetchTotal = &o

		case "search.open_contexts", "so", "searchOpenContexts":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SearchOpenContexts = &o

		case "search.query_current", "sqc", "searchQueryCurrent":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SearchQueryCurrent = &o

		case "search.query_time", "sqti", "searchQueryTime":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SearchQueryTime = &o

		case "search.query_total", "sqto", "searchQueryTotal":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SearchQueryTotal = &o

		case "search.scroll_current", "scc", "searchScrollCurrent":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SearchScrollCurrent = &o

		case "search.scroll_time", "scti", "searchScrollTime":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SearchScrollTime = &o

		case "search.scroll_total", "scto", "searchScrollTotal":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SearchScrollTotal = &o

		case "search.throttled", "sth":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SearchThrottled = &o

		case "segments.count", "sc", "segmentsCount":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SegmentsCount = &o

		case "segments.fixed_bitset_memory", "sfbm", "fixedBitsetMemory":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SegmentsFixedBitsetMemory = &o

		case "segments.index_writer_memory", "siwm", "segmentsIndexWriterMemory":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SegmentsIndexWriterMemory = &o

		case "segments.memory", "sm", "segmentsMemory":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SegmentsMemory = &o

		case "segments.version_map_memory", "svmm", "segmentsVersionMapMemory":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SegmentsVersionMapMemory = &o

		case "status", "s":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Status = &o

		case "store.size", "ss", "storeSize":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.StoreSize = o

		case "suggest.current", "suc", "suggestCurrent":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SuggestCurrent = &o

		case "suggest.time", "suti", "suggestTime":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SuggestTime = &o

		case "suggest.total", "suto", "suggestTotal":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SuggestTotal = &o

		case "uuid", "id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Uuid = &o

		case "warmer.current", "wc", "warmerCurrent":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.WarmerCurrent = &o

		case "warmer.total", "wto", "warmerTotal":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.WarmerTotal = &o

		case "warmer.total_time", "wtt", "warmerTotalTime":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.WarmerTotalTime = &o

		}
	}
	return nil
}

// NewIndicesRecord returns a IndicesRecord.
func NewIndicesRecord() *IndicesRecord {
	r := &IndicesRecord{}

	return r
}
