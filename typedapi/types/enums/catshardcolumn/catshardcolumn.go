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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

// Package catshardcolumn
package catshardcolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/cat/_types/CatBase.ts#L1576-L1949
type CatShardColumn struct {
	Name string
}

var (
	Completionsize = CatShardColumn{"completion.size"}

	Datasetsize = CatShardColumn{"dataset.size"}

	Densevectorvaluecount = CatShardColumn{"dense_vector.value_count"}

	Docs = CatShardColumn{"docs"}

	Fielddataevictions = CatShardColumn{"fielddata.evictions"}

	Fielddatamemorysize = CatShardColumn{"fielddata.memory_size"}

	Flushtotal = CatShardColumn{"flush.total"}

	Flushtotaltime = CatShardColumn{"flush.total_time"}

	Getcurrent = CatShardColumn{"get.current"}

	Getexiststime = CatShardColumn{"get.exists_time"}

	Getexiststotal = CatShardColumn{"get.exists_total"}

	Getmissingtime = CatShardColumn{"get.missing_time"}

	Getmissingtotal = CatShardColumn{"get.missing_total"}

	Gettime = CatShardColumn{"get.time"}

	Gettotal = CatShardColumn{"get.total"}

	Id = CatShardColumn{"id"}

	Index = CatShardColumn{"index"}

	Indexingdeletecurrent = CatShardColumn{"indexing.delete_current"}

	Indexingdeletetime = CatShardColumn{"indexing.delete_time"}

	Indexingdeletetotal = CatShardColumn{"indexing.delete_total"}

	Indexingindexcurrent = CatShardColumn{"indexing.index_current"}

	Indexingindexfailedduetoversionconflict = CatShardColumn{"indexing.index_failed_due_to_version_conflict"}

	Indexingindexfailed = CatShardColumn{"indexing.index_failed"}

	Indexingindextime = CatShardColumn{"indexing.index_time"}

	Indexingindextotal = CatShardColumn{"indexing.index_total"}

	Ip = CatShardColumn{"ip"}

	Mergescurrent = CatShardColumn{"merges.current"}

	Mergescurrentdocs = CatShardColumn{"merges.current_docs"}

	Mergescurrentsize = CatShardColumn{"merges.current_size"}

	Mergestotal = CatShardColumn{"merges.total"}

	Mergestotaldocs = CatShardColumn{"merges.total_docs"}

	Mergestotalsize = CatShardColumn{"merges.total_size"}

	Mergestotaltime = CatShardColumn{"merges.total_time"}

	Node = CatShardColumn{"node"}

	Prirep = CatShardColumn{"prirep"}

	Querycacheevictions = CatShardColumn{"query_cache.evictions"}

	Querycachememorysize = CatShardColumn{"query_cache.memory_size"}

	Recoverysourcetype = CatShardColumn{"recoverysource.type"}

	Refreshtime = CatShardColumn{"refresh.time"}

	Refreshtotal = CatShardColumn{"refresh.total"}

	Searchfetchcurrent = CatShardColumn{"search.fetch_current"}

	Searchfetchtime = CatShardColumn{"search.fetch_time"}

	Searchfetchtotal = CatShardColumn{"search.fetch_total"}

	Searchopencontexts = CatShardColumn{"search.open_contexts"}

	Searchquerycurrent = CatShardColumn{"search.query_current"}

	Searchquerytime = CatShardColumn{"search.query_time"}

	Searchquerytotal = CatShardColumn{"search.query_total"}

	Searchscrollcurrent = CatShardColumn{"search.scroll_current"}

	Searchscrolltime = CatShardColumn{"search.scroll_time"}

	Searchscrolltotal = CatShardColumn{"search.scroll_total"}

	Segmentscount = CatShardColumn{"segments.count"}

	Segmentsfixedbitsetmemory = CatShardColumn{"segments.fixed_bitset_memory"}

	Segmentsindexwritermemory = CatShardColumn{"segments.index_writer_memory"}

	Segmentsmemory = CatShardColumn{"segments.memory"}

	Segmentsversionmapmemory = CatShardColumn{"segments.version_map_memory"}

	Seqnoglobalcheckpoint = CatShardColumn{"seq_no.global_checkpoint"}

	Seqnolocalcheckpoint = CatShardColumn{"seq_no.local_checkpoint"}

	Seqnomax = CatShardColumn{"seq_no.max"}

	Shard = CatShardColumn{"shard"}

	Dsparsevectorvaluecount = CatShardColumn{"dsparse_vector.value_count"}

	State = CatShardColumn{"state"}

	Store = CatShardColumn{"store"}

	Suggestcurrent = CatShardColumn{"suggest.current"}

	Suggesttime = CatShardColumn{"suggest.time"}

	Suggesttotal = CatShardColumn{"suggest.total"}

	Syncid = CatShardColumn{"sync_id"}

	Unassignedat = CatShardColumn{"unassigned.at"}

	Unassigneddetails = CatShardColumn{"unassigned.details"}

	Unassignedfor = CatShardColumn{"unassigned.for"}

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
