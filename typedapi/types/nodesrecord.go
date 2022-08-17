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

// NodesRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cat/nodes/types.ts#L23-L541
type NodesRecord struct {
	// Build es build hash
	Build *string `json:"build,omitempty"`
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
	// Cpu recent cpu usage
	Cpu *string `json:"cpu,omitempty"`
	// DiskAvail available disk space
	DiskAvail *ByteSize `json:"disk.avail,omitempty"`
	// DiskTotal total disk space
	DiskTotal *ByteSize `json:"disk.total,omitempty"`
	// DiskUsed used disk space
	DiskUsed *ByteSize `json:"disk.used,omitempty"`
	// DiskUsedPercent used disk space percentage
	DiskUsedPercent *Percentage `json:"disk.used_percent,omitempty"`
	// FielddataEvictions fielddata evictions
	FielddataEvictions *string `json:"fielddata.evictions,omitempty"`
	// FielddataMemorySize used fielddata cache
	FielddataMemorySize *string `json:"fielddata.memory_size,omitempty"`
	// FileDescCurrent used file descriptors
	FileDescCurrent *string `json:"file_desc.current,omitempty"`
	// FileDescMax max file descriptors
	FileDescMax *string `json:"file_desc.max,omitempty"`
	// FileDescPercent used file descriptor ratio
	FileDescPercent *Percentage `json:"file_desc.percent,omitempty"`
	// Flavor es distribution flavor
	Flavor *string `json:"flavor,omitempty"`
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
	// HeapCurrent used heap
	HeapCurrent *string `json:"heap.current,omitempty"`
	// HeapMax max configured heap
	HeapMax *string `json:"heap.max,omitempty"`
	// HeapPercent used heap ratio
	HeapPercent *Percentage `json:"heap.percent,omitempty"`
	// HttpAddress bound http address
	HttpAddress *string `json:"http_address,omitempty"`
	// Id unique node id
	Id *Id `json:"id,omitempty"`
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
	// Ip ip address
	Ip *string `json:"ip,omitempty"`
	// Jdk jdk version
	Jdk *string `json:"jdk,omitempty"`
	// Load15m 15m load avg
	Load15m *string `json:"load_15m,omitempty"`
	// Load1m 1m load avg
	Load1m *string `json:"load_1m,omitempty"`
	// Load5m 5m load avg
	Load5m *string `json:"load_5m,omitempty"`
	// Master *:current master
	Master *string `json:"master,omitempty"`
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
	// Name node name
	Name *Name `json:"name,omitempty"`
	// NodeRole m:master eligible node, d:data node, i:ingest node, -:coordinating node only
	NodeRole *string `json:"node.role,omitempty"`
	// Pid process id
	Pid *string `json:"pid,omitempty"`
	// Port bound transport port
	Port *string `json:"port,omitempty"`
	// QueryCacheEvictions query cache evictions
	QueryCacheEvictions *string `json:"query_cache.evictions,omitempty"`
	// QueryCacheHitCount query cache hit counts
	QueryCacheHitCount *string `json:"query_cache.hit_count,omitempty"`
	// QueryCacheMemorySize used query cache
	QueryCacheMemorySize *string `json:"query_cache.memory_size,omitempty"`
	// QueryCacheMissCount query cache miss counts
	QueryCacheMissCount *string `json:"query_cache.miss_count,omitempty"`
	// RamCurrent used machine memory
	RamCurrent *string `json:"ram.current,omitempty"`
	// RamMax total machine memory
	RamMax *string `json:"ram.max,omitempty"`
	// RamPercent used machine memory ratio
	RamPercent *Percentage `json:"ram.percent,omitempty"`
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
	// RequestCacheEvictions request cache evictions
	RequestCacheEvictions *string `json:"request_cache.evictions,omitempty"`
	// RequestCacheHitCount request cache hit counts
	RequestCacheHitCount *string `json:"request_cache.hit_count,omitempty"`
	// RequestCacheMemorySize used request cache
	RequestCacheMemorySize *string `json:"request_cache.memory_size,omitempty"`
	// RequestCacheMissCount request cache miss counts
	RequestCacheMissCount *string `json:"request_cache.miss_count,omitempty"`
	// ScriptCacheEvictions script cache evictions
	ScriptCacheEvictions *string `json:"script.cache_evictions,omitempty"`
	// ScriptCompilationLimitTriggered script cache compilation limit triggered
	ScriptCompilationLimitTriggered *string `json:"script.compilation_limit_triggered,omitempty"`
	// ScriptCompilations script compilations
	ScriptCompilations *string `json:"script.compilations,omitempty"`
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
	// SegmentsFixedBitsetMemory memory used by fixed bit sets for nested object field types  and export type
	// filters for types referred in _parent fields
	SegmentsFixedBitsetMemory *string `json:"segments.fixed_bitset_memory,omitempty"`
	// SegmentsIndexWriterMemory memory used by index writer
	SegmentsIndexWriterMemory *string `json:"segments.index_writer_memory,omitempty"`
	// SegmentsMemory memory used by segments
	SegmentsMemory *string `json:"segments.memory,omitempty"`
	// SegmentsVersionMapMemory memory used by version map
	SegmentsVersionMapMemory *string `json:"segments.version_map_memory,omitempty"`
	// SuggestCurrent number of current suggest ops
	SuggestCurrent *string `json:"suggest.current,omitempty"`
	// SuggestTime time spend in suggest
	SuggestTime *string `json:"suggest.time,omitempty"`
	// SuggestTotal number of suggest ops
	SuggestTotal *string `json:"suggest.total,omitempty"`
	// Type es distribution type
	Type *string `json:"type,omitempty"`
	// Uptime node uptime
	Uptime *string `json:"uptime,omitempty"`
	// Version es version
	Version *VersionString `json:"version,omitempty"`
}

// NodesRecordBuilder holds NodesRecord struct and provides a builder API.
type NodesRecordBuilder struct {
	v *NodesRecord
}

// NewNodesRecord provides a builder for the NodesRecord struct.
func NewNodesRecordBuilder() *NodesRecordBuilder {
	r := NodesRecordBuilder{
		&NodesRecord{},
	}

	return &r
}

// Build finalize the chain and returns the NodesRecord struct
func (rb *NodesRecordBuilder) Build() NodesRecord {
	return *rb.v
}

// Build es build hash

func (rb *NodesRecordBuilder) Build_(build_ string) *NodesRecordBuilder {
	rb.v.Build = &build_
	return rb
}

// BulkAvgSizeInBytes average size in bytes of shard bulk

func (rb *NodesRecordBuilder) BulkAvgSizeInBytes(bulkavgsizeinbytes string) *NodesRecordBuilder {
	rb.v.BulkAvgSizeInBytes = &bulkavgsizeinbytes
	return rb
}

// BulkAvgTime average time spend in shard bulk

func (rb *NodesRecordBuilder) BulkAvgTime(bulkavgtime string) *NodesRecordBuilder {
	rb.v.BulkAvgTime = &bulkavgtime
	return rb
}

// BulkTotalOperations number of bulk shard ops

func (rb *NodesRecordBuilder) BulkTotalOperations(bulktotaloperations string) *NodesRecordBuilder {
	rb.v.BulkTotalOperations = &bulktotaloperations
	return rb
}

// BulkTotalSizeInBytes total size in bytes of shard bulk

func (rb *NodesRecordBuilder) BulkTotalSizeInBytes(bulktotalsizeinbytes string) *NodesRecordBuilder {
	rb.v.BulkTotalSizeInBytes = &bulktotalsizeinbytes
	return rb
}

// BulkTotalTime time spend in shard bulk

func (rb *NodesRecordBuilder) BulkTotalTime(bulktotaltime string) *NodesRecordBuilder {
	rb.v.BulkTotalTime = &bulktotaltime
	return rb
}

// CompletionSize size of completion

func (rb *NodesRecordBuilder) CompletionSize(completionsize string) *NodesRecordBuilder {
	rb.v.CompletionSize = &completionsize
	return rb
}

// Cpu recent cpu usage

func (rb *NodesRecordBuilder) Cpu(cpu string) *NodesRecordBuilder {
	rb.v.Cpu = &cpu
	return rb
}

// DiskAvail available disk space

func (rb *NodesRecordBuilder) DiskAvail(diskavail *ByteSizeBuilder) *NodesRecordBuilder {
	v := diskavail.Build()
	rb.v.DiskAvail = &v
	return rb
}

// DiskTotal total disk space

func (rb *NodesRecordBuilder) DiskTotal(disktotal *ByteSizeBuilder) *NodesRecordBuilder {
	v := disktotal.Build()
	rb.v.DiskTotal = &v
	return rb
}

// DiskUsed used disk space

func (rb *NodesRecordBuilder) DiskUsed(diskused *ByteSizeBuilder) *NodesRecordBuilder {
	v := diskused.Build()
	rb.v.DiskUsed = &v
	return rb
}

// DiskUsedPercent used disk space percentage

func (rb *NodesRecordBuilder) DiskUsedPercent(diskusedpercent *PercentageBuilder) *NodesRecordBuilder {
	v := diskusedpercent.Build()
	rb.v.DiskUsedPercent = &v
	return rb
}

// FielddataEvictions fielddata evictions

func (rb *NodesRecordBuilder) FielddataEvictions(fielddataevictions string) *NodesRecordBuilder {
	rb.v.FielddataEvictions = &fielddataevictions
	return rb
}

// FielddataMemorySize used fielddata cache

func (rb *NodesRecordBuilder) FielddataMemorySize(fielddatamemorysize string) *NodesRecordBuilder {
	rb.v.FielddataMemorySize = &fielddatamemorysize
	return rb
}

// FileDescCurrent used file descriptors

func (rb *NodesRecordBuilder) FileDescCurrent(filedesccurrent string) *NodesRecordBuilder {
	rb.v.FileDescCurrent = &filedesccurrent
	return rb
}

// FileDescMax max file descriptors

func (rb *NodesRecordBuilder) FileDescMax(filedescmax string) *NodesRecordBuilder {
	rb.v.FileDescMax = &filedescmax
	return rb
}

// FileDescPercent used file descriptor ratio

func (rb *NodesRecordBuilder) FileDescPercent(filedescpercent *PercentageBuilder) *NodesRecordBuilder {
	v := filedescpercent.Build()
	rb.v.FileDescPercent = &v
	return rb
}

// Flavor es distribution flavor

func (rb *NodesRecordBuilder) Flavor(flavor string) *NodesRecordBuilder {
	rb.v.Flavor = &flavor
	return rb
}

// FlushTotal number of flushes

func (rb *NodesRecordBuilder) FlushTotal(flushtotal string) *NodesRecordBuilder {
	rb.v.FlushTotal = &flushtotal
	return rb
}

// FlushTotalTime time spent in flush

func (rb *NodesRecordBuilder) FlushTotalTime(flushtotaltime string) *NodesRecordBuilder {
	rb.v.FlushTotalTime = &flushtotaltime
	return rb
}

// GetCurrent number of current get ops

func (rb *NodesRecordBuilder) GetCurrent(getcurrent string) *NodesRecordBuilder {
	rb.v.GetCurrent = &getcurrent
	return rb
}

// GetExistsTime time spent in successful gets

func (rb *NodesRecordBuilder) GetExistsTime(getexiststime string) *NodesRecordBuilder {
	rb.v.GetExistsTime = &getexiststime
	return rb
}

// GetExistsTotal number of successful gets

func (rb *NodesRecordBuilder) GetExistsTotal(getexiststotal string) *NodesRecordBuilder {
	rb.v.GetExistsTotal = &getexiststotal
	return rb
}

// GetMissingTime time spent in failed gets

func (rb *NodesRecordBuilder) GetMissingTime(getmissingtime string) *NodesRecordBuilder {
	rb.v.GetMissingTime = &getmissingtime
	return rb
}

// GetMissingTotal number of failed gets

func (rb *NodesRecordBuilder) GetMissingTotal(getmissingtotal string) *NodesRecordBuilder {
	rb.v.GetMissingTotal = &getmissingtotal
	return rb
}

// GetTime time spent in get

func (rb *NodesRecordBuilder) GetTime(gettime string) *NodesRecordBuilder {
	rb.v.GetTime = &gettime
	return rb
}

// GetTotal number of get ops

func (rb *NodesRecordBuilder) GetTotal(gettotal string) *NodesRecordBuilder {
	rb.v.GetTotal = &gettotal
	return rb
}

// HeapCurrent used heap

func (rb *NodesRecordBuilder) HeapCurrent(heapcurrent string) *NodesRecordBuilder {
	rb.v.HeapCurrent = &heapcurrent
	return rb
}

// HeapMax max configured heap

func (rb *NodesRecordBuilder) HeapMax(heapmax string) *NodesRecordBuilder {
	rb.v.HeapMax = &heapmax
	return rb
}

// HeapPercent used heap ratio

func (rb *NodesRecordBuilder) HeapPercent(heappercent *PercentageBuilder) *NodesRecordBuilder {
	v := heappercent.Build()
	rb.v.HeapPercent = &v
	return rb
}

// HttpAddress bound http address

func (rb *NodesRecordBuilder) HttpAddress(httpaddress string) *NodesRecordBuilder {
	rb.v.HttpAddress = &httpaddress
	return rb
}

// Id unique node id

func (rb *NodesRecordBuilder) Id(id Id) *NodesRecordBuilder {
	rb.v.Id = &id
	return rb
}

// IndexingDeleteCurrent number of current deletions

func (rb *NodesRecordBuilder) IndexingDeleteCurrent(indexingdeletecurrent string) *NodesRecordBuilder {
	rb.v.IndexingDeleteCurrent = &indexingdeletecurrent
	return rb
}

// IndexingDeleteTime time spent in deletions

func (rb *NodesRecordBuilder) IndexingDeleteTime(indexingdeletetime string) *NodesRecordBuilder {
	rb.v.IndexingDeleteTime = &indexingdeletetime
	return rb
}

// IndexingDeleteTotal number of delete ops

func (rb *NodesRecordBuilder) IndexingDeleteTotal(indexingdeletetotal string) *NodesRecordBuilder {
	rb.v.IndexingDeleteTotal = &indexingdeletetotal
	return rb
}

// IndexingIndexCurrent number of current indexing ops

func (rb *NodesRecordBuilder) IndexingIndexCurrent(indexingindexcurrent string) *NodesRecordBuilder {
	rb.v.IndexingIndexCurrent = &indexingindexcurrent
	return rb
}

// IndexingIndexFailed number of failed indexing ops

func (rb *NodesRecordBuilder) IndexingIndexFailed(indexingindexfailed string) *NodesRecordBuilder {
	rb.v.IndexingIndexFailed = &indexingindexfailed
	return rb
}

// IndexingIndexTime time spent in indexing

func (rb *NodesRecordBuilder) IndexingIndexTime(indexingindextime string) *NodesRecordBuilder {
	rb.v.IndexingIndexTime = &indexingindextime
	return rb
}

// IndexingIndexTotal number of indexing ops

func (rb *NodesRecordBuilder) IndexingIndexTotal(indexingindextotal string) *NodesRecordBuilder {
	rb.v.IndexingIndexTotal = &indexingindextotal
	return rb
}

// Ip ip address

func (rb *NodesRecordBuilder) Ip(ip string) *NodesRecordBuilder {
	rb.v.Ip = &ip
	return rb
}

// Jdk jdk version

func (rb *NodesRecordBuilder) Jdk(jdk string) *NodesRecordBuilder {
	rb.v.Jdk = &jdk
	return rb
}

// Load15m 15m load avg

func (rb *NodesRecordBuilder) Load15m(load15m string) *NodesRecordBuilder {
	rb.v.Load15m = &load15m
	return rb
}

// Load1m 1m load avg

func (rb *NodesRecordBuilder) Load1m(load1m string) *NodesRecordBuilder {
	rb.v.Load1m = &load1m
	return rb
}

// Load5m 5m load avg

func (rb *NodesRecordBuilder) Load5m(load5m string) *NodesRecordBuilder {
	rb.v.Load5m = &load5m
	return rb
}

// Master *:current master

func (rb *NodesRecordBuilder) Master(master string) *NodesRecordBuilder {
	rb.v.Master = &master
	return rb
}

// MergesCurrent number of current merges

func (rb *NodesRecordBuilder) MergesCurrent(mergescurrent string) *NodesRecordBuilder {
	rb.v.MergesCurrent = &mergescurrent
	return rb
}

// MergesCurrentDocs number of current merging docs

func (rb *NodesRecordBuilder) MergesCurrentDocs(mergescurrentdocs string) *NodesRecordBuilder {
	rb.v.MergesCurrentDocs = &mergescurrentdocs
	return rb
}

// MergesCurrentSize size of current merges

func (rb *NodesRecordBuilder) MergesCurrentSize(mergescurrentsize string) *NodesRecordBuilder {
	rb.v.MergesCurrentSize = &mergescurrentsize
	return rb
}

// MergesTotal number of completed merge ops

func (rb *NodesRecordBuilder) MergesTotal(mergestotal string) *NodesRecordBuilder {
	rb.v.MergesTotal = &mergestotal
	return rb
}

// MergesTotalDocs docs merged

func (rb *NodesRecordBuilder) MergesTotalDocs(mergestotaldocs string) *NodesRecordBuilder {
	rb.v.MergesTotalDocs = &mergestotaldocs
	return rb
}

// MergesTotalSize size merged

func (rb *NodesRecordBuilder) MergesTotalSize(mergestotalsize string) *NodesRecordBuilder {
	rb.v.MergesTotalSize = &mergestotalsize
	return rb
}

// MergesTotalTime time spent in merges

func (rb *NodesRecordBuilder) MergesTotalTime(mergestotaltime string) *NodesRecordBuilder {
	rb.v.MergesTotalTime = &mergestotaltime
	return rb
}

// Name node name

func (rb *NodesRecordBuilder) Name(name Name) *NodesRecordBuilder {
	rb.v.Name = &name
	return rb
}

// NodeRole m:master eligible node, d:data node, i:ingest node, -:coordinating node only

func (rb *NodesRecordBuilder) NodeRole(noderole string) *NodesRecordBuilder {
	rb.v.NodeRole = &noderole
	return rb
}

// Pid process id

func (rb *NodesRecordBuilder) Pid(pid string) *NodesRecordBuilder {
	rb.v.Pid = &pid
	return rb
}

// Port bound transport port

func (rb *NodesRecordBuilder) Port(port string) *NodesRecordBuilder {
	rb.v.Port = &port
	return rb
}

// QueryCacheEvictions query cache evictions

func (rb *NodesRecordBuilder) QueryCacheEvictions(querycacheevictions string) *NodesRecordBuilder {
	rb.v.QueryCacheEvictions = &querycacheevictions
	return rb
}

// QueryCacheHitCount query cache hit counts

func (rb *NodesRecordBuilder) QueryCacheHitCount(querycachehitcount string) *NodesRecordBuilder {
	rb.v.QueryCacheHitCount = &querycachehitcount
	return rb
}

// QueryCacheMemorySize used query cache

func (rb *NodesRecordBuilder) QueryCacheMemorySize(querycachememorysize string) *NodesRecordBuilder {
	rb.v.QueryCacheMemorySize = &querycachememorysize
	return rb
}

// QueryCacheMissCount query cache miss counts

func (rb *NodesRecordBuilder) QueryCacheMissCount(querycachemisscount string) *NodesRecordBuilder {
	rb.v.QueryCacheMissCount = &querycachemisscount
	return rb
}

// RamCurrent used machine memory

func (rb *NodesRecordBuilder) RamCurrent(ramcurrent string) *NodesRecordBuilder {
	rb.v.RamCurrent = &ramcurrent
	return rb
}

// RamMax total machine memory

func (rb *NodesRecordBuilder) RamMax(rammax string) *NodesRecordBuilder {
	rb.v.RamMax = &rammax
	return rb
}

// RamPercent used machine memory ratio

func (rb *NodesRecordBuilder) RamPercent(rampercent *PercentageBuilder) *NodesRecordBuilder {
	v := rampercent.Build()
	rb.v.RamPercent = &v
	return rb
}

// RefreshExternalTime time spent in external refreshes

func (rb *NodesRecordBuilder) RefreshExternalTime(refreshexternaltime string) *NodesRecordBuilder {
	rb.v.RefreshExternalTime = &refreshexternaltime
	return rb
}

// RefreshExternalTotal total external refreshes

func (rb *NodesRecordBuilder) RefreshExternalTotal(refreshexternaltotal string) *NodesRecordBuilder {
	rb.v.RefreshExternalTotal = &refreshexternaltotal
	return rb
}

// RefreshListeners number of pending refresh listeners

func (rb *NodesRecordBuilder) RefreshListeners(refreshlisteners string) *NodesRecordBuilder {
	rb.v.RefreshListeners = &refreshlisteners
	return rb
}

// RefreshTime time spent in refreshes

func (rb *NodesRecordBuilder) RefreshTime(refreshtime string) *NodesRecordBuilder {
	rb.v.RefreshTime = &refreshtime
	return rb
}

// RefreshTotal total refreshes

func (rb *NodesRecordBuilder) RefreshTotal(refreshtotal string) *NodesRecordBuilder {
	rb.v.RefreshTotal = &refreshtotal
	return rb
}

// RequestCacheEvictions request cache evictions

func (rb *NodesRecordBuilder) RequestCacheEvictions(requestcacheevictions string) *NodesRecordBuilder {
	rb.v.RequestCacheEvictions = &requestcacheevictions
	return rb
}

// RequestCacheHitCount request cache hit counts

func (rb *NodesRecordBuilder) RequestCacheHitCount(requestcachehitcount string) *NodesRecordBuilder {
	rb.v.RequestCacheHitCount = &requestcachehitcount
	return rb
}

// RequestCacheMemorySize used request cache

func (rb *NodesRecordBuilder) RequestCacheMemorySize(requestcachememorysize string) *NodesRecordBuilder {
	rb.v.RequestCacheMemorySize = &requestcachememorysize
	return rb
}

// RequestCacheMissCount request cache miss counts

func (rb *NodesRecordBuilder) RequestCacheMissCount(requestcachemisscount string) *NodesRecordBuilder {
	rb.v.RequestCacheMissCount = &requestcachemisscount
	return rb
}

// ScriptCacheEvictions script cache evictions

func (rb *NodesRecordBuilder) ScriptCacheEvictions(scriptcacheevictions string) *NodesRecordBuilder {
	rb.v.ScriptCacheEvictions = &scriptcacheevictions
	return rb
}

// ScriptCompilationLimitTriggered script cache compilation limit triggered

func (rb *NodesRecordBuilder) ScriptCompilationLimitTriggered(scriptcompilationlimittriggered string) *NodesRecordBuilder {
	rb.v.ScriptCompilationLimitTriggered = &scriptcompilationlimittriggered
	return rb
}

// ScriptCompilations script compilations

func (rb *NodesRecordBuilder) ScriptCompilations(scriptcompilations string) *NodesRecordBuilder {
	rb.v.ScriptCompilations = &scriptcompilations
	return rb
}

// SearchFetchCurrent current fetch phase ops

func (rb *NodesRecordBuilder) SearchFetchCurrent(searchfetchcurrent string) *NodesRecordBuilder {
	rb.v.SearchFetchCurrent = &searchfetchcurrent
	return rb
}

// SearchFetchTime time spent in fetch phase

func (rb *NodesRecordBuilder) SearchFetchTime(searchfetchtime string) *NodesRecordBuilder {
	rb.v.SearchFetchTime = &searchfetchtime
	return rb
}

// SearchFetchTotal total fetch ops

func (rb *NodesRecordBuilder) SearchFetchTotal(searchfetchtotal string) *NodesRecordBuilder {
	rb.v.SearchFetchTotal = &searchfetchtotal
	return rb
}

// SearchOpenContexts open search contexts

func (rb *NodesRecordBuilder) SearchOpenContexts(searchopencontexts string) *NodesRecordBuilder {
	rb.v.SearchOpenContexts = &searchopencontexts
	return rb
}

// SearchQueryCurrent current query phase ops

func (rb *NodesRecordBuilder) SearchQueryCurrent(searchquerycurrent string) *NodesRecordBuilder {
	rb.v.SearchQueryCurrent = &searchquerycurrent
	return rb
}

// SearchQueryTime time spent in query phase

func (rb *NodesRecordBuilder) SearchQueryTime(searchquerytime string) *NodesRecordBuilder {
	rb.v.SearchQueryTime = &searchquerytime
	return rb
}

// SearchQueryTotal total query phase ops

func (rb *NodesRecordBuilder) SearchQueryTotal(searchquerytotal string) *NodesRecordBuilder {
	rb.v.SearchQueryTotal = &searchquerytotal
	return rb
}

// SearchScrollCurrent open scroll contexts

func (rb *NodesRecordBuilder) SearchScrollCurrent(searchscrollcurrent string) *NodesRecordBuilder {
	rb.v.SearchScrollCurrent = &searchscrollcurrent
	return rb
}

// SearchScrollTime time scroll contexts held open

func (rb *NodesRecordBuilder) SearchScrollTime(searchscrolltime string) *NodesRecordBuilder {
	rb.v.SearchScrollTime = &searchscrolltime
	return rb
}

// SearchScrollTotal completed scroll contexts

func (rb *NodesRecordBuilder) SearchScrollTotal(searchscrolltotal string) *NodesRecordBuilder {
	rb.v.SearchScrollTotal = &searchscrolltotal
	return rb
}

// SegmentsCount number of segments

func (rb *NodesRecordBuilder) SegmentsCount(segmentscount string) *NodesRecordBuilder {
	rb.v.SegmentsCount = &segmentscount
	return rb
}

// SegmentsFixedBitsetMemory memory used by fixed bit sets for nested object field types  and export type
// filters for types referred in _parent fields

func (rb *NodesRecordBuilder) SegmentsFixedBitsetMemory(segmentsfixedbitsetmemory string) *NodesRecordBuilder {
	rb.v.SegmentsFixedBitsetMemory = &segmentsfixedbitsetmemory
	return rb
}

// SegmentsIndexWriterMemory memory used by index writer

func (rb *NodesRecordBuilder) SegmentsIndexWriterMemory(segmentsindexwritermemory string) *NodesRecordBuilder {
	rb.v.SegmentsIndexWriterMemory = &segmentsindexwritermemory
	return rb
}

// SegmentsMemory memory used by segments

func (rb *NodesRecordBuilder) SegmentsMemory(segmentsmemory string) *NodesRecordBuilder {
	rb.v.SegmentsMemory = &segmentsmemory
	return rb
}

// SegmentsVersionMapMemory memory used by version map

func (rb *NodesRecordBuilder) SegmentsVersionMapMemory(segmentsversionmapmemory string) *NodesRecordBuilder {
	rb.v.SegmentsVersionMapMemory = &segmentsversionmapmemory
	return rb
}

// SuggestCurrent number of current suggest ops

func (rb *NodesRecordBuilder) SuggestCurrent(suggestcurrent string) *NodesRecordBuilder {
	rb.v.SuggestCurrent = &suggestcurrent
	return rb
}

// SuggestTime time spend in suggest

func (rb *NodesRecordBuilder) SuggestTime(suggesttime string) *NodesRecordBuilder {
	rb.v.SuggestTime = &suggesttime
	return rb
}

// SuggestTotal number of suggest ops

func (rb *NodesRecordBuilder) SuggestTotal(suggesttotal string) *NodesRecordBuilder {
	rb.v.SuggestTotal = &suggesttotal
	return rb
}

// Type es distribution type

func (rb *NodesRecordBuilder) Type_(type_ string) *NodesRecordBuilder {
	rb.v.Type = &type_
	return rb
}

// Uptime node uptime

func (rb *NodesRecordBuilder) Uptime(uptime string) *NodesRecordBuilder {
	rb.v.Uptime = &uptime
	return rb
}

// Version es version

func (rb *NodesRecordBuilder) Version(version VersionString) *NodesRecordBuilder {
	rb.v.Version = &version
	return rb
}
