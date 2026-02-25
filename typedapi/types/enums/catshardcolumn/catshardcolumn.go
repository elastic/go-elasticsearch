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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

// Package catshardcolumn
package catshardcolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/cat/_types/CatBase.ts#L2834-L3207
type CatShardColumn struct {
	Name string
}

var (

	// Completionsize Size of completion. For example: `0b`.
	Completionsize = CatShardColumn{"completion.size"}

	// Datasetsize Disk space used by the shard’s dataset, which may or may not be the size on
	// disk, but includes space used by the shard on object storage. Reported as a
	// size value for example: `5kb`.
	Datasetsize = CatShardColumn{"dataset.size"}

	// Densevectorvaluecount Number of indexed dense vectors.
	Densevectorvaluecount = CatShardColumn{"dense_vector.value_count"}

	// Docs Number of documents in shard, for example: `25`.
	Docs = CatShardColumn{"docs"}

	// Fielddataevictions Fielddata cache evictions, for example: `0`.
	Fielddataevictions = CatShardColumn{"fielddata.evictions"}

	// Fielddatamemorysize Used fielddata cache memory, for example: `0b`.
	Fielddatamemorysize = CatShardColumn{"fielddata.memory_size"}

	// Flushtotal Number of flushes, for example: `1`.
	Flushtotal = CatShardColumn{"flush.total"}

	// Flushtotaltime Time spent in flush, for example: `1`.
	Flushtotaltime = CatShardColumn{"flush.total_time"}

	// Getcurrent Number of current get operations, for example: `0`.
	Getcurrent = CatShardColumn{"get.current"}

	// Getexiststime Time spent in successful gets, for example: `14ms`.
	Getexiststime = CatShardColumn{"get.exists_time"}

	// Getexiststotal Number of successful get operations, for example: `2`.
	Getexiststotal = CatShardColumn{"get.exists_total"}

	// Getmissingtime Time spent in failed gets, for example: `0s`.
	Getmissingtime = CatShardColumn{"get.missing_time"}

	// Getmissingtotal Number of failed get operations, for example: `1`.
	Getmissingtotal = CatShardColumn{"get.missing_total"}

	// Gettime Time spent in get, for example: `14ms`.
	Gettime = CatShardColumn{"get.time"}

	// Gettotal Number of get operations, for example: `2`.
	Gettotal = CatShardColumn{"get.total"}

	// Id ID of the node, for example: `k0zy`.
	Id = CatShardColumn{"id"}

	// Index Name of the index.
	Index = CatShardColumn{"index"}

	// Indexingdeletecurrent Number of current deletion operations, for example: `0`.
	Indexingdeletecurrent = CatShardColumn{"indexing.delete_current"}

	// Indexingdeletetime Time spent in deletions, for example: `2ms`.
	Indexingdeletetime = CatShardColumn{"indexing.delete_time"}

	// Indexingdeletetotal Number of deletion operations, for example: `2`.
	Indexingdeletetotal = CatShardColumn{"indexing.delete_total"}

	// Indexingindexcurrent Number of current indexing operations, for example: `0`.
	Indexingindexcurrent = CatShardColumn{"indexing.index_current"}

	// Indexingindexfailedduetoversionconflict Number of failed indexing operations due to version conflict, for example:
	// `0`.
	Indexingindexfailedduetoversionconflict = CatShardColumn{"indexing.index_failed_due_to_version_conflict"}

	// Indexingindexfailed Number of failed indexing operations, for example: `0`.
	Indexingindexfailed = CatShardColumn{"indexing.index_failed"}

	// Indexingindextime Time spent in indexing, such as for example: `134ms`.
	Indexingindextime = CatShardColumn{"indexing.index_time"}

	// Indexingindextotal Number of indexing operations, for example: `1`.
	Indexingindextotal = CatShardColumn{"indexing.index_total"}

	// Ip IP address of the node, for example: `127.0.1.1`.
	Ip = CatShardColumn{"ip"}

	// Mergescurrent Number of current merge operations, for example: `0`.
	Mergescurrent = CatShardColumn{"merges.current"}

	// Mergescurrentdocs Number of current merging documents, for example: `0`.
	Mergescurrentdocs = CatShardColumn{"merges.current_docs"}

	// Mergescurrentsize Size of current merges, for example: `0b`.
	Mergescurrentsize = CatShardColumn{"merges.current_size"}

	// Mergestotal Number of completed merge operations, for example: `0`.
	Mergestotal = CatShardColumn{"merges.total"}

	// Mergestotaldocs Number of merged documents, for example: `0`.
	Mergestotaldocs = CatShardColumn{"merges.total_docs"}

	// Mergestotalsize Size of current merges, for example: `0b`.
	Mergestotalsize = CatShardColumn{"merges.total_size"}

	// Mergestotaltime Time spent merging documents, for example: `0s`.
	Mergestotaltime = CatShardColumn{"merges.total_time"}

	// Node Node name, for example: `I8hydUG`.
	Node = CatShardColumn{"node"}

	// Prirep Shard type. Returned values are `primary` or `replica`.
	Prirep = CatShardColumn{"prirep"}

	// Querycacheevictions Query cache evictions, for example: `0`.
	Querycacheevictions = CatShardColumn{"query_cache.evictions"}

	// Querycachememorysize Used query cache memory, for example: `0b`.
	Querycachememorysize = CatShardColumn{"query_cache.memory_size"}

	// Recoverysourcetype Type of recovery source.
	Recoverysourcetype = CatShardColumn{"recoverysource.type"}

	// Refreshtime Time spent in refreshes, for example: `91ms`.
	Refreshtime = CatShardColumn{"refresh.time"}

	// Refreshtotal Number of refreshes, for example: `16`.
	Refreshtotal = CatShardColumn{"refresh.total"}

	// Searchfetchcurrent Current fetch phase operations, for example: `0`.
	Searchfetchcurrent = CatShardColumn{"search.fetch_current"}

	// Searchfetchtime Time spent in fetch phase, for example: `37ms`.
	Searchfetchtime = CatShardColumn{"search.fetch_time"}

	// Searchfetchtotal Number of fetch operations, for example: `7`.
	Searchfetchtotal = CatShardColumn{"search.fetch_total"}

	// Searchopencontexts Open search contexts, for example: `0`.
	Searchopencontexts = CatShardColumn{"search.open_contexts"}

	// Searchquerycurrent Current query phase operations, for example: `0`.
	Searchquerycurrent = CatShardColumn{"search.query_current"}

	// Searchquerytime Time spent in query phase, for example: `43ms`.
	Searchquerytime = CatShardColumn{"search.query_time"}

	// Searchquerytotal Number of query operations, for example: `9`.
	Searchquerytotal = CatShardColumn{"search.query_total"}

	// Searchscrollcurrent Open scroll contexts, for example: `2`.
	Searchscrollcurrent = CatShardColumn{"search.scroll_current"}

	// Searchscrolltime Time scroll contexts held open, for example: `2m`.
	Searchscrolltime = CatShardColumn{"search.scroll_time"}

	// Searchscrolltotal Completed scroll contexts, for example: `1`.
	Searchscrolltotal = CatShardColumn{"search.scroll_total"}

	// Segmentscount Number of segments, for example: `4`.
	Segmentscount = CatShardColumn{"segments.count"}

	// Segmentsfixedbitsetmemory Memory used by fixed bit sets for nested object field types and type filters
	// for types referred in join fields, for example: `1.0kb`.
	Segmentsfixedbitsetmemory = CatShardColumn{"segments.fixed_bitset_memory"}

	// Segmentsindexwritermemory Memory used by index writer, for example: `18mb`.
	Segmentsindexwritermemory = CatShardColumn{"segments.index_writer_memory"}

	// Segmentsmemory Memory used by segments, for example: `1.4kb`.
	Segmentsmemory = CatShardColumn{"segments.memory"}

	// Segmentsversionmapmemory Memory used by version map, for example: `1.0kb`.
	Segmentsversionmapmemory = CatShardColumn{"segments.version_map_memory"}

	// Seqnoglobalcheckpoint Global checkpoint.
	Seqnoglobalcheckpoint = CatShardColumn{"seq_no.global_checkpoint"}

	// Seqnolocalcheckpoint Local checkpoint.
	Seqnolocalcheckpoint = CatShardColumn{"seq_no.local_checkpoint"}

	// Seqnomax Maximum sequence number.
	Seqnomax = CatShardColumn{"seq_no.max"}

	// Shard Name of the shard.
	Shard = CatShardColumn{"shard"}

	// Dsparsevectorvaluecount Number of indexed [sparse
	// vectors](https://www.elastic.co/docs/reference/elasticsearch/mapping-reference/sparse-vector).
	Dsparsevectorvaluecount = CatShardColumn{"dsparse_vector.value_count"}

	// State State of the shard. Returned values are:
	//
	//   - `INITIALIZING`: The shard is recovering from a peer shard or gateway.
	//   - `RELOCATING`: The shard is relocating.
	//   - `STARTED`: The shard has started.
	//   - `UNASSIGNED`: The shard is not assigned to any node.
	State = CatShardColumn{"state"}

	// Store Disk space used by the shard, for example: `5kb`.
	Store = CatShardColumn{"store"}

	// Suggestcurrent Number of current suggest operations, for example: `0`.
	Suggestcurrent = CatShardColumn{"suggest.current"}

	// Suggesttime Time spent in suggest, for example: `0`.
	Suggesttime = CatShardColumn{"suggest.time"}

	// Suggesttotal Number of suggest operations, for example: `0`.
	Suggesttotal = CatShardColumn{"suggest.total"}

	// Syncid Sync ID of the shard.
	Syncid = CatShardColumn{"sync_id"}

	// Unassignedat Time at which the shard became unassigned in [Coordinated Universal Time
	// (UTC)](https://en.wikipedia.org/wiki/List_of_UTC_offsets).
	Unassignedat = CatShardColumn{"unassigned.at"}

	// Unassigneddetails Details about why the shard became unassigned. This does not explain why the
	// shard is currently unassigned. To understand why a shard is not assigned, use
	// the [Cluster allocation
	// explain](https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-cluster-allocation-explain)
	// API.
	Unassigneddetails = CatShardColumn{"unassigned.details"}

	// Unassignedfor Time at which the shard was requested to be unassigned in [Coordinated
	// Universal Time (UTC)](https://en.wikipedia.org/wiki/List_of_UTC_offsets).
	Unassignedfor = CatShardColumn{"unassigned.for"}

	// Unassignedreason Indicates the reason for the last change to the state of this unassigned
	// shard. This does not explain why the shard is currently unassigned. To
	// understand why a shard is not assigned, use the [Cluster allocation
	// explain](https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-cluster-allocation-explain)
	// API. Returned values include:
	//
	//   - `ALLOCATION_FAILED`: Unassigned as a result of a failed allocation of the
	//     shard.
	//   - `CLUSTER_RECOVERED`: Unassigned as a result of a full cluster recovery.
	//   - `DANGLING_INDEX_IMPORTED`: Unassigned as a result of importing a dangling
	//     index.
	//   - `EXISTING_INDEX_RESTORED`: Unassigned as a result of restoring into a
	//     closed index.
	//   - `FORCED_EMPTY_PRIMARY`: The shard’s allocation was last modified by
	//     forcing an empty primary using the Cluster reroute API.
	//   - `INDEX_CLOSED`: Unassigned because the index was closed.
	//   - `INDEX_CREATED`: Unassigned as a result of an API creation of an index.
	//   - `INDEX_REOPENED`: Unassigned as a result of opening a closed index.
	//   - `MANUAL_ALLOCATION`: The shard’s allocation was last modified by the
	//     Cluster reroute API.
	//   - `NEW_INDEX_RESTORED`: Unassigned as a result of restoring into a new
	//     index.
	//   - `NODE_LEFT`: Unassigned as a result of the node hosting it leaving the
	//     cluster.
	//   - `NODE_RESTARTING`: Similar to `NODE_LEFT`, except that the node was
	//     registered as restarting using the Node shutdown API.
	//   - `PRIMARY_FAILED`: The shard was initializing as a replica, but the
	//     primary shard failed before the initialization completed.
	//   - `REALLOCATED_REPLICA`: A better replica location is identified and causes
	//     the existing replica allocation to be cancelled.
	//   - `REINITIALIZED`: When a shard moves from started back to initializing.
	//   - `REPLICA_ADDED`: Unassigned as a result of explicit addition of a
	//     replica.
	//   - `REROUTE_CANCELLED`: Unassigned as a result of explicit cancel reroute
	//     command.
	Unassignedreason = CatShardColumn{"unassigned.reason"}
)

func (c CatShardColumn) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CatShardColumn) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "completion.size":
		*c = Completionsize
	case "dataset.size":
		*c = Datasetsize
	case "dense_vector.value_count":
		*c = Densevectorvaluecount
	case "docs":
		*c = Docs
	case "fielddata.evictions":
		*c = Fielddataevictions
	case "fielddata.memory_size":
		*c = Fielddatamemorysize
	case "flush.total":
		*c = Flushtotal
	case "flush.total_time":
		*c = Flushtotaltime
	case "get.current":
		*c = Getcurrent
	case "get.exists_time":
		*c = Getexiststime
	case "get.exists_total":
		*c = Getexiststotal
	case "get.missing_time":
		*c = Getmissingtime
	case "get.missing_total":
		*c = Getmissingtotal
	case "get.time":
		*c = Gettime
	case "get.total":
		*c = Gettotal
	case "id":
		*c = Id
	case "index":
		*c = Index
	case "indexing.delete_current":
		*c = Indexingdeletecurrent
	case "indexing.delete_time":
		*c = Indexingdeletetime
	case "indexing.delete_total":
		*c = Indexingdeletetotal
	case "indexing.index_current":
		*c = Indexingindexcurrent
	case "indexing.index_failed_due_to_version_conflict":
		*c = Indexingindexfailedduetoversionconflict
	case "indexing.index_failed":
		*c = Indexingindexfailed
	case "indexing.index_time":
		*c = Indexingindextime
	case "indexing.index_total":
		*c = Indexingindextotal
	case "ip":
		*c = Ip
	case "merges.current":
		*c = Mergescurrent
	case "merges.current_docs":
		*c = Mergescurrentdocs
	case "merges.current_size":
		*c = Mergescurrentsize
	case "merges.total":
		*c = Mergestotal
	case "merges.total_docs":
		*c = Mergestotaldocs
	case "merges.total_size":
		*c = Mergestotalsize
	case "merges.total_time":
		*c = Mergestotaltime
	case "node":
		*c = Node
	case "prirep":
		*c = Prirep
	case "query_cache.evictions":
		*c = Querycacheevictions
	case "query_cache.memory_size":
		*c = Querycachememorysize
	case "recoverysource.type":
		*c = Recoverysourcetype
	case "refresh.time":
		*c = Refreshtime
	case "refresh.total":
		*c = Refreshtotal
	case "search.fetch_current":
		*c = Searchfetchcurrent
	case "search.fetch_time":
		*c = Searchfetchtime
	case "search.fetch_total":
		*c = Searchfetchtotal
	case "search.open_contexts":
		*c = Searchopencontexts
	case "search.query_current":
		*c = Searchquerycurrent
	case "search.query_time":
		*c = Searchquerytime
	case "search.query_total":
		*c = Searchquerytotal
	case "search.scroll_current":
		*c = Searchscrollcurrent
	case "search.scroll_time":
		*c = Searchscrolltime
	case "search.scroll_total":
		*c = Searchscrolltotal
	case "segments.count":
		*c = Segmentscount
	case "segments.fixed_bitset_memory":
		*c = Segmentsfixedbitsetmemory
	case "segments.index_writer_memory":
		*c = Segmentsindexwritermemory
	case "segments.memory":
		*c = Segmentsmemory
	case "segments.version_map_memory":
		*c = Segmentsversionmapmemory
	case "seq_no.global_checkpoint":
		*c = Seqnoglobalcheckpoint
	case "seq_no.local_checkpoint":
		*c = Seqnolocalcheckpoint
	case "seq_no.max":
		*c = Seqnomax
	case "shard":
		*c = Shard
	case "dsparse_vector.value_count":
		*c = Dsparsevectorvaluecount
	case "state":
		*c = State
	case "store":
		*c = Store
	case "suggest.current":
		*c = Suggestcurrent
	case "suggest.time":
		*c = Suggesttime
	case "suggest.total":
		*c = Suggesttotal
	case "sync_id":
		*c = Syncid
	case "unassigned.at":
		*c = Unassignedat
	case "unassigned.details":
		*c = Unassigneddetails
	case "unassigned.for":
		*c = Unassignedfor
	case "unassigned.reason":
		*c = Unassignedreason
	default:
		*c = CatShardColumn{string(text)}
	}

	return nil
}

func (c CatShardColumn) String() string {
	return c.Name
}
