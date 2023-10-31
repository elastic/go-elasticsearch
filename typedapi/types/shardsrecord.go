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

// ShardsRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/cat/shards/types.ts#L20-L421
type ShardsRecord struct {
	// BulkAvgSizeInBytes The average size in bytes of shard bulk operations.
	BulkAvgSizeInBytes *string `json:"bulk.avg_size_in_bytes,omitempty"`
	// BulkAvgTime The average time spent in shard bulk operations.
	BulkAvgTime *string `json:"bulk.avg_time,omitempty"`
	// BulkTotalOperations The number of bulk shard operations.
	BulkTotalOperations *string `json:"bulk.total_operations,omitempty"`
	// BulkTotalSizeInBytes The total size in bytes of shard bulk operations.
	BulkTotalSizeInBytes *string `json:"bulk.total_size_in_bytes,omitempty"`
	// BulkTotalTime The time spent in shard bulk operations.
	BulkTotalTime *string `json:"bulk.total_time,omitempty"`
	// CompletionSize The size of completion.
	CompletionSize *string `json:"completion.size,omitempty"`
	// Docs The number of documents in the shard.
	Docs string `json:"docs,omitempty"`
	// FielddataEvictions The fielddata cache evictions.
	FielddataEvictions *string `json:"fielddata.evictions,omitempty"`
	// FielddataMemorySize The used fielddata cache memory.
	FielddataMemorySize *string `json:"fielddata.memory_size,omitempty"`
	// FlushTotal The number of flushes.
	FlushTotal *string `json:"flush.total,omitempty"`
	// FlushTotalTime The time spent in flush.
	FlushTotalTime *string `json:"flush.total_time,omitempty"`
	// GetCurrent The number of current get operations.
	GetCurrent *string `json:"get.current,omitempty"`
	// GetExistsTime The time spent in successful get operations.
	GetExistsTime *string `json:"get.exists_time,omitempty"`
	// GetExistsTotal The number of successful get operations.
	GetExistsTotal *string `json:"get.exists_total,omitempty"`
	// GetMissingTime The time spent in failed get operations.
	GetMissingTime *string `json:"get.missing_time,omitempty"`
	// GetMissingTotal The number of failed get operations.
	GetMissingTotal *string `json:"get.missing_total,omitempty"`
	// GetTime The time spent in get operations.
	GetTime *string `json:"get.time,omitempty"`
	// GetTotal The number of get operations.
	GetTotal *string `json:"get.total,omitempty"`
	// Id The unique identifier for the node.
	Id *string `json:"id,omitempty"`
	// Index The index name.
	Index *string `json:"index,omitempty"`
	// IndexingDeleteCurrent The number of current deletion operations.
	IndexingDeleteCurrent *string `json:"indexing.delete_current,omitempty"`
	// IndexingDeleteTime The time spent in deletion operations.
	IndexingDeleteTime *string `json:"indexing.delete_time,omitempty"`
	// IndexingDeleteTotal The number of delete operations.
	IndexingDeleteTotal *string `json:"indexing.delete_total,omitempty"`
	// IndexingIndexCurrent The number of current indexing operations.
	IndexingIndexCurrent *string `json:"indexing.index_current,omitempty"`
	// IndexingIndexFailed The number of failed indexing operations.
	IndexingIndexFailed *string `json:"indexing.index_failed,omitempty"`
	// IndexingIndexTime The time spent in indexing operations.
	IndexingIndexTime *string `json:"indexing.index_time,omitempty"`
	// IndexingIndexTotal The number of indexing operations.
	IndexingIndexTotal *string `json:"indexing.index_total,omitempty"`
	// Ip The IP address of the node.
	Ip string `json:"ip,omitempty"`
	// MergesCurrent The number of current merge operations.
	MergesCurrent *string `json:"merges.current,omitempty"`
	// MergesCurrentDocs The number of current merging documents.
	MergesCurrentDocs *string `json:"merges.current_docs,omitempty"`
	// MergesCurrentSize The size of current merge operations.
	MergesCurrentSize *string `json:"merges.current_size,omitempty"`
	// MergesTotal The number of completed merge operations.
	MergesTotal *string `json:"merges.total,omitempty"`
	// MergesTotalDocs The nuber of merged documents.
	MergesTotalDocs *string `json:"merges.total_docs,omitempty"`
	// MergesTotalSize The size of current merges.
	MergesTotalSize *string `json:"merges.total_size,omitempty"`
	// MergesTotalTime The time spent merging documents.
	MergesTotalTime *string `json:"merges.total_time,omitempty"`
	// Node The name of node.
	Node string `json:"node,omitempty"`
	// PathData The shard data path.
	PathData *string `json:"path.data,omitempty"`
	// PathState The shard state path.
	PathState *string `json:"path.state,omitempty"`
	// Prirep The shard type: `primary` or `replica`.
	Prirep *string `json:"prirep,omitempty"`
	// QueryCacheEvictions The query cache evictions.
	QueryCacheEvictions *string `json:"query_cache.evictions,omitempty"`
	// QueryCacheMemorySize The used query cache memory.
	QueryCacheMemorySize *string `json:"query_cache.memory_size,omitempty"`
	// RecoverysourceType The type of recovery source.
	RecoverysourceType *string `json:"recoverysource.type,omitempty"`
	// RefreshExternalTime The time spent in external refreshes.
	RefreshExternalTime *string `json:"refresh.external_time,omitempty"`
	// RefreshExternalTotal The total nunber of external refreshes.
	RefreshExternalTotal *string `json:"refresh.external_total,omitempty"`
	// RefreshListeners The number of pending refresh listeners.
	RefreshListeners *string `json:"refresh.listeners,omitempty"`
	// RefreshTime The time spent in refreshes.
	RefreshTime *string `json:"refresh.time,omitempty"`
	// RefreshTotal The total number of refreshes.
	RefreshTotal *string `json:"refresh.total,omitempty"`
	// SearchFetchCurrent The current fetch phase operations.
	SearchFetchCurrent *string `json:"search.fetch_current,omitempty"`
	// SearchFetchTime The time spent in fetch phase.
	SearchFetchTime *string `json:"search.fetch_time,omitempty"`
	// SearchFetchTotal The total number of fetch operations.
	SearchFetchTotal *string `json:"search.fetch_total,omitempty"`
	// SearchOpenContexts The number of open search contexts.
	SearchOpenContexts *string `json:"search.open_contexts,omitempty"`
	// SearchQueryCurrent The current query phase operations.
	SearchQueryCurrent *string `json:"search.query_current,omitempty"`
	// SearchQueryTime The time spent in query phase.
	SearchQueryTime *string `json:"search.query_time,omitempty"`
	// SearchQueryTotal The total number of query phase operations.
	SearchQueryTotal *string `json:"search.query_total,omitempty"`
	// SearchScrollCurrent The open scroll contexts.
	SearchScrollCurrent *string `json:"search.scroll_current,omitempty"`
	// SearchScrollTime The time scroll contexts were held open.
	SearchScrollTime *string `json:"search.scroll_time,omitempty"`
	// SearchScrollTotal The number of completed scroll contexts.
	SearchScrollTotal *string `json:"search.scroll_total,omitempty"`
	// SegmentsCount The number of segments.
	SegmentsCount *string `json:"segments.count,omitempty"`
	// SegmentsFixedBitsetMemory The memory used by fixed bit sets for nested object field types and export
	// type filters for types referred in `_parent` fields.
	SegmentsFixedBitsetMemory *string `json:"segments.fixed_bitset_memory,omitempty"`
	// SegmentsIndexWriterMemory The memory used by the index writer.
	SegmentsIndexWriterMemory *string `json:"segments.index_writer_memory,omitempty"`
	// SegmentsMemory The memory used by segments.
	SegmentsMemory *string `json:"segments.memory,omitempty"`
	// SegmentsVersionMapMemory The memory used by the version map.
	SegmentsVersionMapMemory *string `json:"segments.version_map_memory,omitempty"`
	// SeqNoGlobalCheckpoint The global checkpoint.
	SeqNoGlobalCheckpoint *string `json:"seq_no.global_checkpoint,omitempty"`
	// SeqNoLocalCheckpoint The local checkpoint.
	SeqNoLocalCheckpoint *string `json:"seq_no.local_checkpoint,omitempty"`
	// SeqNoMax The maximum sequence number.
	SeqNoMax *string `json:"seq_no.max,omitempty"`
	// Shard The shard name.
	Shard *string `json:"shard,omitempty"`
	// State The shard state.
	// Returned values include:
	// `INITIALIZING`: The shard is recovering from a peer shard or gateway.
	// `RELOCATING`: The shard is relocating.
	// `STARTED`: The shard has started.
	// `UNASSIGNED`: The shard is not assigned to any node.
	State *string `json:"state,omitempty"`
	// Store The disk space used by the shard.
	Store string `json:"store,omitempty"`
	// SyncId The sync identifier.
	SyncId *string `json:"sync_id,omitempty"`
	// UnassignedAt The time at which the shard became unassigned in Coordinated Universal Time
	// (UTC).
	UnassignedAt *string `json:"unassigned.at,omitempty"`
	// UnassignedDetails Additional details as to why the shard became unassigned.
	// It does not explain why the shard is not assigned; use the cluster allocation
	// explain API for that information.
	UnassignedDetails *string `json:"unassigned.details,omitempty"`
	// UnassignedFor The time at which the shard was requested to be unassigned in Coordinated
	// Universal Time (UTC).
	UnassignedFor *string `json:"unassigned.for,omitempty"`
	// UnassignedReason The reason for the last change to the state of an unassigned shard.
	// It does not explain why the shard is currently unassigned; use the cluster
	// allocation explain API for that information.
	// Returned values include:
	// `ALLOCATION_FAILED`: Unassigned as a result of a failed allocation of the
	// shard.
	// `CLUSTER_RECOVERED`: Unassigned as a result of a full cluster recovery.
	// `DANGLING_INDEX_IMPORTED`: Unassigned as a result of importing a dangling
	// index.
	// `EXISTING_INDEX_RESTORED`: Unassigned as a result of restoring into a closed
	// index.
	// `FORCED_EMPTY_PRIMARY`: The shard’s allocation was last modified by forcing
	// an empty primary using the cluster reroute API.
	// `INDEX_CLOSED`: Unassigned because the index was closed.
	// `INDEX_CREATED`: Unassigned as a result of an API creation of an index.
	// `INDEX_REOPENED`: Unassigned as a result of opening a closed index.
	// `MANUAL_ALLOCATION`: The shard’s allocation was last modified by the cluster
	// reroute API.
	// `NEW_INDEX_RESTORED`: Unassigned as a result of restoring into a new index.
	// `NODE_LEFT`: Unassigned as a result of the node hosting it leaving the
	// cluster.
	// `NODE_RESTARTING`: Similar to `NODE_LEFT`, except that the node was
	// registered as restarting using the node shutdown API.
	// `PRIMARY_FAILED`: The shard was initializing as a replica, but the primary
	// shard failed before the initialization completed.
	// `REALLOCATED_REPLICA`: A better replica location is identified and causes the
	// existing replica allocation to be cancelled.
	// `REINITIALIZED`: When a shard moves from started back to initializing.
	// `REPLICA_ADDED`: Unassigned as a result of explicit addition of a replica.
	// `REROUTE_CANCELLED`: Unassigned as a result of explicit cancel reroute
	// command.
	UnassignedReason *string `json:"unassigned.reason,omitempty"`
	// WarmerCurrent The number of current warmer operations.
	WarmerCurrent *string `json:"warmer.current,omitempty"`
	// WarmerTotal The total number of warmer operations.
	WarmerTotal *string `json:"warmer.total,omitempty"`
	// WarmerTotalTime The time spent in warmer operations.
	WarmerTotalTime *string `json:"warmer.total_time,omitempty"`
}

func (s *ShardsRecord) UnmarshalJSON(data []byte) error {

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

		case "docs", "d", "dc":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Docs = o

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

		case "id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Id = &o

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

		case "ip":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Ip = o

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

		case "node", "n":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Node = o

		case "path.data", "pd", "dataPath":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PathData = &o

		case "path.state", "ps", "statsPath":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PathState = &o

		case "prirep", "p", "pr", "primaryOrReplica":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Prirep = &o

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

		case "recoverysource.type", "rs":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.RecoverysourceType = &o

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

		case "seq_no.global_checkpoint", "sqg", "globalCheckpoint":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SeqNoGlobalCheckpoint = &o

		case "seq_no.local_checkpoint", "sql", "localCheckpoint":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SeqNoLocalCheckpoint = &o

		case "seq_no.max", "sqm", "maxSeqNo":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SeqNoMax = &o

		case "shard", "s", "sh":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Shard = &o

		case "state", "st":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.State = &o

		case "store", "sto":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Store = o

		case "sync_id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SyncId = &o

		case "unassigned.at", "ua":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.UnassignedAt = &o

		case "unassigned.details", "ud":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.UnassignedDetails = &o

		case "unassigned.for", "uf":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.UnassignedFor = &o

		case "unassigned.reason", "ur":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.UnassignedReason = &o

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

// NewShardsRecord returns a ShardsRecord.
func NewShardsRecord() *ShardsRecord {
	r := &ShardsRecord{}

	return r
}
