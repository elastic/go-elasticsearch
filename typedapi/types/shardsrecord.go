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

// ShardsRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cat/shards/types.ts#L20-L396
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

// ShardsRecordBuilder holds ShardsRecord struct and provides a builder API.
type ShardsRecordBuilder struct {
	v *ShardsRecord
}

// NewShardsRecord provides a builder for the ShardsRecord struct.
func NewShardsRecordBuilder() *ShardsRecordBuilder {
	r := ShardsRecordBuilder{
		&ShardsRecord{},
	}

	return &r
}

// Build finalize the chain and returns the ShardsRecord struct
func (rb *ShardsRecordBuilder) Build() ShardsRecord {
	return *rb.v
}

// BulkAvgSizeInBytes avg size in bytes of shard bulk

func (rb *ShardsRecordBuilder) BulkAvgSizeInBytes(bulkavgsizeinbytes string) *ShardsRecordBuilder {
	rb.v.BulkAvgSizeInBytes = &bulkavgsizeinbytes
	return rb
}

// BulkAvgTime average time spend in shard bulk

func (rb *ShardsRecordBuilder) BulkAvgTime(bulkavgtime string) *ShardsRecordBuilder {
	rb.v.BulkAvgTime = &bulkavgtime
	return rb
}

// BulkTotalOperations number of bulk shard ops

func (rb *ShardsRecordBuilder) BulkTotalOperations(bulktotaloperations string) *ShardsRecordBuilder {
	rb.v.BulkTotalOperations = &bulktotaloperations
	return rb
}

// BulkTotalSizeInBytes total size in bytes of shard bulk

func (rb *ShardsRecordBuilder) BulkTotalSizeInBytes(bulktotalsizeinbytes string) *ShardsRecordBuilder {
	rb.v.BulkTotalSizeInBytes = &bulktotalsizeinbytes
	return rb
}

// BulkTotalTime time spend in shard bulk

func (rb *ShardsRecordBuilder) BulkTotalTime(bulktotaltime string) *ShardsRecordBuilder {
	rb.v.BulkTotalTime = &bulktotaltime
	return rb
}

// CompletionSize size of completion

func (rb *ShardsRecordBuilder) CompletionSize(completionsize string) *ShardsRecordBuilder {
	rb.v.CompletionSize = &completionsize
	return rb
}

// Docs number of docs in shard

func (rb *ShardsRecordBuilder) Docs(docs string) *ShardsRecordBuilder {
	rb.v.Docs = docs
	return rb
}

// FielddataEvictions fielddata evictions

func (rb *ShardsRecordBuilder) FielddataEvictions(fielddataevictions string) *ShardsRecordBuilder {
	rb.v.FielddataEvictions = &fielddataevictions
	return rb
}

// FielddataMemorySize used fielddata cache

func (rb *ShardsRecordBuilder) FielddataMemorySize(fielddatamemorysize string) *ShardsRecordBuilder {
	rb.v.FielddataMemorySize = &fielddatamemorysize
	return rb
}

// FlushTotal number of flushes

func (rb *ShardsRecordBuilder) FlushTotal(flushtotal string) *ShardsRecordBuilder {
	rb.v.FlushTotal = &flushtotal
	return rb
}

// FlushTotalTime time spent in flush

func (rb *ShardsRecordBuilder) FlushTotalTime(flushtotaltime string) *ShardsRecordBuilder {
	rb.v.FlushTotalTime = &flushtotaltime
	return rb
}

// GetCurrent number of current get ops

func (rb *ShardsRecordBuilder) GetCurrent(getcurrent string) *ShardsRecordBuilder {
	rb.v.GetCurrent = &getcurrent
	return rb
}

// GetExistsTime time spent in successful gets

func (rb *ShardsRecordBuilder) GetExistsTime(getexiststime string) *ShardsRecordBuilder {
	rb.v.GetExistsTime = &getexiststime
	return rb
}

// GetExistsTotal number of successful gets

func (rb *ShardsRecordBuilder) GetExistsTotal(getexiststotal string) *ShardsRecordBuilder {
	rb.v.GetExistsTotal = &getexiststotal
	return rb
}

// GetMissingTime time spent in failed gets

func (rb *ShardsRecordBuilder) GetMissingTime(getmissingtime string) *ShardsRecordBuilder {
	rb.v.GetMissingTime = &getmissingtime
	return rb
}

// GetMissingTotal number of failed gets

func (rb *ShardsRecordBuilder) GetMissingTotal(getmissingtotal string) *ShardsRecordBuilder {
	rb.v.GetMissingTotal = &getmissingtotal
	return rb
}

// GetTime time spent in get

func (rb *ShardsRecordBuilder) GetTime(gettime string) *ShardsRecordBuilder {
	rb.v.GetTime = &gettime
	return rb
}

// GetTotal number of get ops

func (rb *ShardsRecordBuilder) GetTotal(gettotal string) *ShardsRecordBuilder {
	rb.v.GetTotal = &gettotal
	return rb
}

// Id unique id of node where it lives

func (rb *ShardsRecordBuilder) Id(id string) *ShardsRecordBuilder {
	rb.v.Id = &id
	return rb
}

// Index index name

func (rb *ShardsRecordBuilder) Index(index string) *ShardsRecordBuilder {
	rb.v.Index = &index
	return rb
}

// IndexingDeleteCurrent number of current deletions

func (rb *ShardsRecordBuilder) IndexingDeleteCurrent(indexingdeletecurrent string) *ShardsRecordBuilder {
	rb.v.IndexingDeleteCurrent = &indexingdeletecurrent
	return rb
}

// IndexingDeleteTime time spent in deletions

func (rb *ShardsRecordBuilder) IndexingDeleteTime(indexingdeletetime string) *ShardsRecordBuilder {
	rb.v.IndexingDeleteTime = &indexingdeletetime
	return rb
}

// IndexingDeleteTotal number of delete ops

func (rb *ShardsRecordBuilder) IndexingDeleteTotal(indexingdeletetotal string) *ShardsRecordBuilder {
	rb.v.IndexingDeleteTotal = &indexingdeletetotal
	return rb
}

// IndexingIndexCurrent number of current indexing ops

func (rb *ShardsRecordBuilder) IndexingIndexCurrent(indexingindexcurrent string) *ShardsRecordBuilder {
	rb.v.IndexingIndexCurrent = &indexingindexcurrent
	return rb
}

// IndexingIndexFailed number of failed indexing ops

func (rb *ShardsRecordBuilder) IndexingIndexFailed(indexingindexfailed string) *ShardsRecordBuilder {
	rb.v.IndexingIndexFailed = &indexingindexfailed
	return rb
}

// IndexingIndexTime time spent in indexing

func (rb *ShardsRecordBuilder) IndexingIndexTime(indexingindextime string) *ShardsRecordBuilder {
	rb.v.IndexingIndexTime = &indexingindextime
	return rb
}

// IndexingIndexTotal number of indexing ops

func (rb *ShardsRecordBuilder) IndexingIndexTotal(indexingindextotal string) *ShardsRecordBuilder {
	rb.v.IndexingIndexTotal = &indexingindextotal
	return rb
}

// Ip ip of node where it lives

func (rb *ShardsRecordBuilder) Ip(ip string) *ShardsRecordBuilder {
	rb.v.Ip = ip
	return rb
}

// MergesCurrent number of current merges

func (rb *ShardsRecordBuilder) MergesCurrent(mergescurrent string) *ShardsRecordBuilder {
	rb.v.MergesCurrent = &mergescurrent
	return rb
}

// MergesCurrentDocs number of current merging docs

func (rb *ShardsRecordBuilder) MergesCurrentDocs(mergescurrentdocs string) *ShardsRecordBuilder {
	rb.v.MergesCurrentDocs = &mergescurrentdocs
	return rb
}

// MergesCurrentSize size of current merges

func (rb *ShardsRecordBuilder) MergesCurrentSize(mergescurrentsize string) *ShardsRecordBuilder {
	rb.v.MergesCurrentSize = &mergescurrentsize
	return rb
}

// MergesTotal number of completed merge ops

func (rb *ShardsRecordBuilder) MergesTotal(mergestotal string) *ShardsRecordBuilder {
	rb.v.MergesTotal = &mergestotal
	return rb
}

// MergesTotalDocs docs merged

func (rb *ShardsRecordBuilder) MergesTotalDocs(mergestotaldocs string) *ShardsRecordBuilder {
	rb.v.MergesTotalDocs = &mergestotaldocs
	return rb
}

// MergesTotalSize size merged

func (rb *ShardsRecordBuilder) MergesTotalSize(mergestotalsize string) *ShardsRecordBuilder {
	rb.v.MergesTotalSize = &mergestotalsize
	return rb
}

// MergesTotalTime time spent in merges

func (rb *ShardsRecordBuilder) MergesTotalTime(mergestotaltime string) *ShardsRecordBuilder {
	rb.v.MergesTotalTime = &mergestotaltime
	return rb
}

// Node name of node where it lives

func (rb *ShardsRecordBuilder) Node(node string) *ShardsRecordBuilder {
	rb.v.Node = node
	return rb
}

// PathData shard data path

func (rb *ShardsRecordBuilder) PathData(pathdata string) *ShardsRecordBuilder {
	rb.v.PathData = &pathdata
	return rb
}

// PathState shard state path

func (rb *ShardsRecordBuilder) PathState(pathstate string) *ShardsRecordBuilder {
	rb.v.PathState = &pathstate
	return rb
}

// Prirep primary or replica

func (rb *ShardsRecordBuilder) Prirep(prirep string) *ShardsRecordBuilder {
	rb.v.Prirep = &prirep
	return rb
}

// QueryCacheEvictions query cache evictions

func (rb *ShardsRecordBuilder) QueryCacheEvictions(querycacheevictions string) *ShardsRecordBuilder {
	rb.v.QueryCacheEvictions = &querycacheevictions
	return rb
}

// QueryCacheMemorySize used query cache

func (rb *ShardsRecordBuilder) QueryCacheMemorySize(querycachememorysize string) *ShardsRecordBuilder {
	rb.v.QueryCacheMemorySize = &querycachememorysize
	return rb
}

// RecoverysourceType recovery source type

func (rb *ShardsRecordBuilder) RecoverysourceType(recoverysourcetype string) *ShardsRecordBuilder {
	rb.v.RecoverysourceType = &recoverysourcetype
	return rb
}

// RefreshExternalTime time spent in external refreshes

func (rb *ShardsRecordBuilder) RefreshExternalTime(refreshexternaltime string) *ShardsRecordBuilder {
	rb.v.RefreshExternalTime = &refreshexternaltime
	return rb
}

// RefreshExternalTotal total external refreshes

func (rb *ShardsRecordBuilder) RefreshExternalTotal(refreshexternaltotal string) *ShardsRecordBuilder {
	rb.v.RefreshExternalTotal = &refreshexternaltotal
	return rb
}

// RefreshListeners number of pending refresh listeners

func (rb *ShardsRecordBuilder) RefreshListeners(refreshlisteners string) *ShardsRecordBuilder {
	rb.v.RefreshListeners = &refreshlisteners
	return rb
}

// RefreshTime time spent in refreshes

func (rb *ShardsRecordBuilder) RefreshTime(refreshtime string) *ShardsRecordBuilder {
	rb.v.RefreshTime = &refreshtime
	return rb
}

// RefreshTotal total refreshes

func (rb *ShardsRecordBuilder) RefreshTotal(refreshtotal string) *ShardsRecordBuilder {
	rb.v.RefreshTotal = &refreshtotal
	return rb
}

// SearchFetchCurrent current fetch phase ops

func (rb *ShardsRecordBuilder) SearchFetchCurrent(searchfetchcurrent string) *ShardsRecordBuilder {
	rb.v.SearchFetchCurrent = &searchfetchcurrent
	return rb
}

// SearchFetchTime time spent in fetch phase

func (rb *ShardsRecordBuilder) SearchFetchTime(searchfetchtime string) *ShardsRecordBuilder {
	rb.v.SearchFetchTime = &searchfetchtime
	return rb
}

// SearchFetchTotal total fetch ops

func (rb *ShardsRecordBuilder) SearchFetchTotal(searchfetchtotal string) *ShardsRecordBuilder {
	rb.v.SearchFetchTotal = &searchfetchtotal
	return rb
}

// SearchOpenContexts open search contexts

func (rb *ShardsRecordBuilder) SearchOpenContexts(searchopencontexts string) *ShardsRecordBuilder {
	rb.v.SearchOpenContexts = &searchopencontexts
	return rb
}

// SearchQueryCurrent current query phase ops

func (rb *ShardsRecordBuilder) SearchQueryCurrent(searchquerycurrent string) *ShardsRecordBuilder {
	rb.v.SearchQueryCurrent = &searchquerycurrent
	return rb
}

// SearchQueryTime time spent in query phase

func (rb *ShardsRecordBuilder) SearchQueryTime(searchquerytime string) *ShardsRecordBuilder {
	rb.v.SearchQueryTime = &searchquerytime
	return rb
}

// SearchQueryTotal total query phase ops

func (rb *ShardsRecordBuilder) SearchQueryTotal(searchquerytotal string) *ShardsRecordBuilder {
	rb.v.SearchQueryTotal = &searchquerytotal
	return rb
}

// SearchScrollCurrent open scroll contexts

func (rb *ShardsRecordBuilder) SearchScrollCurrent(searchscrollcurrent string) *ShardsRecordBuilder {
	rb.v.SearchScrollCurrent = &searchscrollcurrent
	return rb
}

// SearchScrollTime time scroll contexts held open

func (rb *ShardsRecordBuilder) SearchScrollTime(searchscrolltime string) *ShardsRecordBuilder {
	rb.v.SearchScrollTime = &searchscrolltime
	return rb
}

// SearchScrollTotal completed scroll contexts

func (rb *ShardsRecordBuilder) SearchScrollTotal(searchscrolltotal string) *ShardsRecordBuilder {
	rb.v.SearchScrollTotal = &searchscrolltotal
	return rb
}

// SegmentsCount number of segments

func (rb *ShardsRecordBuilder) SegmentsCount(segmentscount string) *ShardsRecordBuilder {
	rb.v.SegmentsCount = &segmentscount
	return rb
}

// SegmentsFixedBitsetMemory memory used by fixed bit sets for nested object field types and export type
// filters for types referred in _parent fields

func (rb *ShardsRecordBuilder) SegmentsFixedBitsetMemory(segmentsfixedbitsetmemory string) *ShardsRecordBuilder {
	rb.v.SegmentsFixedBitsetMemory = &segmentsfixedbitsetmemory
	return rb
}

// SegmentsIndexWriterMemory memory used by index writer

func (rb *ShardsRecordBuilder) SegmentsIndexWriterMemory(segmentsindexwritermemory string) *ShardsRecordBuilder {
	rb.v.SegmentsIndexWriterMemory = &segmentsindexwritermemory
	return rb
}

// SegmentsMemory memory used by segments

func (rb *ShardsRecordBuilder) SegmentsMemory(segmentsmemory string) *ShardsRecordBuilder {
	rb.v.SegmentsMemory = &segmentsmemory
	return rb
}

// SegmentsVersionMapMemory memory used by version map

func (rb *ShardsRecordBuilder) SegmentsVersionMapMemory(segmentsversionmapmemory string) *ShardsRecordBuilder {
	rb.v.SegmentsVersionMapMemory = &segmentsversionmapmemory
	return rb
}

// SeqNoGlobalCheckpoint global checkpoint

func (rb *ShardsRecordBuilder) SeqNoGlobalCheckpoint(seqnoglobalcheckpoint string) *ShardsRecordBuilder {
	rb.v.SeqNoGlobalCheckpoint = &seqnoglobalcheckpoint
	return rb
}

// SeqNoLocalCheckpoint local checkpoint

func (rb *ShardsRecordBuilder) SeqNoLocalCheckpoint(seqnolocalcheckpoint string) *ShardsRecordBuilder {
	rb.v.SeqNoLocalCheckpoint = &seqnolocalcheckpoint
	return rb
}

// SeqNoMax max sequence number

func (rb *ShardsRecordBuilder) SeqNoMax(seqnomax string) *ShardsRecordBuilder {
	rb.v.SeqNoMax = &seqnomax
	return rb
}

// Shard shard name

func (rb *ShardsRecordBuilder) Shard(shard string) *ShardsRecordBuilder {
	rb.v.Shard = &shard
	return rb
}

// State shard state

func (rb *ShardsRecordBuilder) State(state string) *ShardsRecordBuilder {
	rb.v.State = &state
	return rb
}

// Store store size of shard (how much disk it uses)

func (rb *ShardsRecordBuilder) Store(store string) *ShardsRecordBuilder {
	rb.v.Store = store
	return rb
}

// SyncId sync id

func (rb *ShardsRecordBuilder) SyncId(syncid string) *ShardsRecordBuilder {
	rb.v.SyncId = &syncid
	return rb
}

// UnassignedAt time shard became unassigned (UTC)

func (rb *ShardsRecordBuilder) UnassignedAt(unassignedat string) *ShardsRecordBuilder {
	rb.v.UnassignedAt = &unassignedat
	return rb
}

// UnassignedDetails additional details as to why the shard became unassigned

func (rb *ShardsRecordBuilder) UnassignedDetails(unassigneddetails string) *ShardsRecordBuilder {
	rb.v.UnassignedDetails = &unassigneddetails
	return rb
}

// UnassignedFor time has been unassigned

func (rb *ShardsRecordBuilder) UnassignedFor(unassignedfor string) *ShardsRecordBuilder {
	rb.v.UnassignedFor = &unassignedfor
	return rb
}

// UnassignedReason reason shard is unassigned

func (rb *ShardsRecordBuilder) UnassignedReason(unassignedreason string) *ShardsRecordBuilder {
	rb.v.UnassignedReason = &unassignedreason
	return rb
}

// WarmerCurrent current warmer ops

func (rb *ShardsRecordBuilder) WarmerCurrent(warmercurrent string) *ShardsRecordBuilder {
	rb.v.WarmerCurrent = &warmercurrent
	return rb
}

// WarmerTotal total warmer ops

func (rb *ShardsRecordBuilder) WarmerTotal(warmertotal string) *ShardsRecordBuilder {
	rb.v.WarmerTotal = &warmertotal
	return rb
}

// WarmerTotalTime time spent in warmers

func (rb *ShardsRecordBuilder) WarmerTotalTime(warmertotaltime string) *ShardsRecordBuilder {
	rb.v.WarmerTotalTime = &warmertotaltime
	return rb
}
