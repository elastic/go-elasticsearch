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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// NodesRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/cat/nodes/types.ts#L23-L541
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
	DiskAvail ByteSize `json:"disk.avail,omitempty"`
	// DiskTotal total disk space
	DiskTotal ByteSize `json:"disk.total,omitempty"`
	// DiskUsed used disk space
	DiskUsed ByteSize `json:"disk.used,omitempty"`
	// DiskUsedPercent used disk space percentage
	DiskUsedPercent Percentage `json:"disk.used_percent,omitempty"`
	// FielddataEvictions fielddata evictions
	FielddataEvictions *string `json:"fielddata.evictions,omitempty"`
	// FielddataMemorySize used fielddata cache
	FielddataMemorySize *string `json:"fielddata.memory_size,omitempty"`
	// FileDescCurrent used file descriptors
	FileDescCurrent *string `json:"file_desc.current,omitempty"`
	// FileDescMax max file descriptors
	FileDescMax *string `json:"file_desc.max,omitempty"`
	// FileDescPercent used file descriptor ratio
	FileDescPercent Percentage `json:"file_desc.percent,omitempty"`
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
	HeapPercent Percentage `json:"heap.percent,omitempty"`
	// HttpAddress bound http address
	HttpAddress *string `json:"http_address,omitempty"`
	// Id unique node id
	Id *string `json:"id,omitempty"`
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
	// Load15M 15m load avg
	Load15M *string `json:"load_15m,omitempty"`
	// Load1M 1m load avg
	Load1M *string `json:"load_1m,omitempty"`
	// Load5M 5m load avg
	Load5M *string `json:"load_5m,omitempty"`
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
	Name *string `json:"name,omitempty"`
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
	RamPercent Percentage `json:"ram.percent,omitempty"`
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
	Version *string `json:"version,omitempty"`
}

func (s *NodesRecord) UnmarshalJSON(data []byte) error {

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

		case "build", "b":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Build = &o

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

		case "bulk.total_operations", "bto", "bulkTotalOperations":
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

		case "cpu":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Cpu = &o

		case "disk.avail", "d", "da", "disk", "diskAvail":
			if err := dec.Decode(&s.DiskAvail); err != nil {
				return err
			}

		case "disk.total", "dt", "diskTotal":
			if err := dec.Decode(&s.DiskTotal); err != nil {
				return err
			}

		case "disk.used", "du", "diskUsed":
			if err := dec.Decode(&s.DiskUsed); err != nil {
				return err
			}

		case "disk.used_percent", "dup", "diskUsedPercent":
			if err := dec.Decode(&s.DiskUsedPercent); err != nil {
				return err
			}

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

		case "file_desc.current", "fdc", "fileDescriptorCurrent":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.FileDescCurrent = &o

		case "file_desc.max", "fdm", "fileDescriptorMax":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.FileDescMax = &o

		case "file_desc.percent", "fdp", "fileDescriptorPercent":
			if err := dec.Decode(&s.FileDescPercent); err != nil {
				return err
			}

		case "flavor", "f":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Flavor = &o

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

		case "heap.current", "hc", "heapCurrent":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.HeapCurrent = &o

		case "heap.max", "hm", "heapMax":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.HeapMax = &o

		case "heap.percent", "hp", "heapPercent":
			if err := dec.Decode(&s.HeapPercent); err != nil {
				return err
			}

		case "http_address", "http":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.HttpAddress = &o

		case "id", "nodeId":
			if err := dec.Decode(&s.Id); err != nil {
				return err
			}

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

		case "ip", "i":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Ip = &o

		case "jdk", "j":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Jdk = &o

		case "load_15m", "l":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Load15M = &o

		case "load_1m":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Load1M = &o

		case "load_5m":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Load5M = &o

		case "master", "m":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Master = &o

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

		case "name", "n":
			if err := dec.Decode(&s.Name); err != nil {
				return err
			}

		case "node.role", "r", "role", "nodeRole":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.NodeRole = &o

		case "pid", "p":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Pid = &o

		case "port", "po":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Port = &o

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

		case "query_cache.hit_count", "qchc", "queryCacheHitCount":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.QueryCacheHitCount = &o

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

		case "query_cache.miss_count", "qcmc", "queryCacheMissCount":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.QueryCacheMissCount = &o

		case "ram.current", "rc", "ramCurrent":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.RamCurrent = &o

		case "ram.max", "rn", "ramMax":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.RamMax = &o

		case "ram.percent", "rp", "ramPercent":
			if err := dec.Decode(&s.RamPercent); err != nil {
				return err
			}

		case "refresh.external_time", "rti", "refreshTime":
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

		case "refresh.external_total", "rto", "refreshTotal":
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

		case "refresh.time":
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

		case "refresh.total":
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

		case "script.cache_evictions", "scrce", "scriptCacheEvictions":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ScriptCacheEvictions = &o

		case "script.compilation_limit_triggered", "scrclt", "scriptCacheCompilationLimitTriggered":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ScriptCompilationLimitTriggered = &o

		case "script.compilations", "scrcc", "scriptCompilations":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ScriptCompilations = &o

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

		case "type", "t":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Type = &o

		case "uptime", "u":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Uptime = &o

		case "version", "v":
			if err := dec.Decode(&s.Version); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewNodesRecord returns a NodesRecord.
func NewNodesRecord() *NodesRecord {
	r := &NodesRecord{}

	return r
}
