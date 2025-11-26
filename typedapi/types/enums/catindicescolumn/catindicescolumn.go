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
// https://github.com/elastic/elasticsearch-specification/tree/aa1459fbdcaf57c653729142b3b6e9982373bb1c

// Package catindicescolumn
package catindicescolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/aa1459fbdcaf57c653729142b3b6e9982373bb1c/specification/cat/_types/CatBase.ts#L1432-L2100
type CatIndicesColumn struct {
	Name string
}

var (
	Health = CatIndicesColumn{"health"}

	Status = CatIndicesColumn{"status"}

	Index = CatIndicesColumn{"index"}

	Uuid = CatIndicesColumn{"uuid"}

	Pri = CatIndicesColumn{"pri"}

	Rep = CatIndicesColumn{"rep"}

	Docscount = CatIndicesColumn{"docs.count"}

	Docsdeleted = CatIndicesColumn{"docs.deleted"}

	Creationdate = CatIndicesColumn{"creation.date"}

	Creationdatestring = CatIndicesColumn{"creation.date.string"}

	Storesize = CatIndicesColumn{"store.size"}

	Pristoresize = CatIndicesColumn{"pri.store.size"}

	Datasetsize = CatIndicesColumn{"dataset.size"}

	Completionsize = CatIndicesColumn{"completion.size"}

	Pricompletionsize = CatIndicesColumn{"pri.completion.size"}

	Fielddatamemorysize = CatIndicesColumn{"fielddata.memory_size"}

	Prifielddatamemorysize = CatIndicesColumn{"pri.fielddata.memory_size"}

	Fielddataevictions = CatIndicesColumn{"fielddata.evictions"}

	Prifielddataevictions = CatIndicesColumn{"pri.fielddata.evictions"}

	Querycachememorysize = CatIndicesColumn{"query_cache.memory_size"}

	Priquerycachememorysize = CatIndicesColumn{"pri.query_cache.memory_size"}

	Querycacheevictions = CatIndicesColumn{"query_cache.evictions"}

	Priquerycacheevictions = CatIndicesColumn{"pri.query_cache.evictions"}

	Requestcachememorysize = CatIndicesColumn{"request_cache.memory_size"}

	Prirequestcachememorysize = CatIndicesColumn{"pri.request_cache.memory_size"}

	Requestcacheevictions = CatIndicesColumn{"request_cache.evictions"}

	Prirequestcacheevictions = CatIndicesColumn{"pri.request_cache.evictions"}

	Requestcachehitcount = CatIndicesColumn{"request_cache.hit_count"}

	Prirequestcachehitcount = CatIndicesColumn{"pri.request_cache.hit_count"}

	Requestcachemisscount = CatIndicesColumn{"request_cache.miss_count"}

	Prirequestcachemisscount = CatIndicesColumn{"pri.request_cache.miss_count"}

	Flushtotal = CatIndicesColumn{"flush.total"}

	Priflushtotal = CatIndicesColumn{"pri.flush.total"}

	Flushtotaltime = CatIndicesColumn{"flush.total_time"}

	Priflushtotaltime = CatIndicesColumn{"pri.flush.total_time"}

	Getcurrent = CatIndicesColumn{"get.current"}

	Prigetcurrent = CatIndicesColumn{"pri.get.current"}

	Gettime = CatIndicesColumn{"get.time"}

	Prigettime = CatIndicesColumn{"pri.get.time"}

	Gettotal = CatIndicesColumn{"get.total"}

	Prigettotal = CatIndicesColumn{"pri.get.total"}

	Getexiststime = CatIndicesColumn{"get.exists_time"}

	Prigetexiststime = CatIndicesColumn{"pri.get.exists_time"}

	Getexiststotal = CatIndicesColumn{"get.exists_total"}

	Prigetexiststotal = CatIndicesColumn{"pri.get.exists_total"}

	Getmissingtime = CatIndicesColumn{"get.missing_time"}

	Prigetmissingtime = CatIndicesColumn{"pri.get.missing_time"}

	Getmissingtotal = CatIndicesColumn{"get.missing_total"}

	Prigetmissingtotal = CatIndicesColumn{"pri.get.missing_total"}

	Indexingdeletecurrent = CatIndicesColumn{"indexing.delete_current"}

	Priindexingdeletecurrent = CatIndicesColumn{"pri.indexing.delete_current"}

	Indexingdeletetime = CatIndicesColumn{"indexing.delete_time"}

	Priindexingdeletetime = CatIndicesColumn{"pri.indexing.delete_time"}

	Indexingdeletetotal = CatIndicesColumn{"indexing.delete_total"}

	Priindexingdeletetotal = CatIndicesColumn{"pri.indexing.delete_total"}

	Indexingindexcurrent = CatIndicesColumn{"indexing.index_current"}

	Priindexingindexcurrent = CatIndicesColumn{"pri.indexing.index_current"}

	Indexingindextime = CatIndicesColumn{"indexing.index_time"}

	Priindexingindextime = CatIndicesColumn{"pri.indexing.index_time"}

	Indexingindextotal = CatIndicesColumn{"indexing.index_total"}

	Priindexingindextotal = CatIndicesColumn{"pri.indexing.index_total"}

	Indexingindexfailed = CatIndicesColumn{"indexing.index_failed"}

	Priindexingindexfailed = CatIndicesColumn{"pri.indexing.index_failed"}

	Indexingindexfailedduetoversionconflict = CatIndicesColumn{"indexing.index_failed_due_to_version_conflict"}

	Priindexingindexfailedduetoversionconflict = CatIndicesColumn{"pri.indexing.index_failed_due_to_version_conflict"}

	Mergescurrent = CatIndicesColumn{"merges.current"}

	Primergescurrent = CatIndicesColumn{"pri.merges.current"}

	Mergescurrentdocs = CatIndicesColumn{"merges.current_docs"}

	Primergescurrentdocs = CatIndicesColumn{"pri.merges.current_docs"}

	Mergescurrentsize = CatIndicesColumn{"merges.current_size"}

	Primergescurrentsize = CatIndicesColumn{"pri.merges.current_size"}

	Mergestotal = CatIndicesColumn{"merges.total"}

	Primergestotal = CatIndicesColumn{"pri.merges.total"}

	Mergestotaldocs = CatIndicesColumn{"merges.total_docs"}

	Primergestotaldocs = CatIndicesColumn{"pri.merges.total_docs"}

	Mergestotalsize = CatIndicesColumn{"merges.total_size"}

	Primergestotalsize = CatIndicesColumn{"pri.merges.total_size"}

	Mergestotaltime = CatIndicesColumn{"merges.total_time"}

	Primergestotaltime = CatIndicesColumn{"pri.merges.total_time"}

	Refreshtotal = CatIndicesColumn{"refresh.total"}

	Prirefreshtotal = CatIndicesColumn{"pri.refresh.total"}

	Refreshtime = CatIndicesColumn{"refresh.time"}

	Prirefreshtime = CatIndicesColumn{"pri.refresh.time"}

	Refreshexternaltotal = CatIndicesColumn{"refresh.external_total"}

	Prirefreshexternaltotal = CatIndicesColumn{"pri.refresh.external_total"}

	Refreshexternaltime = CatIndicesColumn{"refresh.external_time"}

	Prirefreshexternaltime = CatIndicesColumn{"pri.refresh.external_time"}

	Refreshlisteners = CatIndicesColumn{"refresh.listeners"}

	Prirefreshlisteners = CatIndicesColumn{"pri.refresh.listeners"}

	Searchfetchcurrent = CatIndicesColumn{"search.fetch_current"}

	Prisearchfetchcurrent = CatIndicesColumn{"pri.search.fetch_current"}

	Searchfetchtime = CatIndicesColumn{"search.fetch_time"}

	Prisearchfetchtime = CatIndicesColumn{"pri.search.fetch_time"}

	Searchfetchtotal = CatIndicesColumn{"search.fetch_total"}

	Prisearchfetchtotal = CatIndicesColumn{"pri.search.fetch_total"}

	Searchopencontexts = CatIndicesColumn{"search.open_contexts"}

	Prisearchopencontexts = CatIndicesColumn{"pri.search.open_contexts"}

	Searchquerycurrent = CatIndicesColumn{"search.query_current"}

	Prisearchquerycurrent = CatIndicesColumn{"pri.search.query_current"}

	Searchquerytime = CatIndicesColumn{"search.query_time"}

	Prisearchquerytime = CatIndicesColumn{"pri.search.query_time"}

	Searchquerytotal = CatIndicesColumn{"search.query_total"}

	Prisearchquerytotal = CatIndicesColumn{"pri.search.query_total"}

	Searchscrollcurrent = CatIndicesColumn{"search.scroll_current"}

	Prisearchscrollcurrent = CatIndicesColumn{"pri.search.scroll_current"}

	Searchscrolltime = CatIndicesColumn{"search.scroll_time"}

	Prisearchscrolltime = CatIndicesColumn{"pri.search.scroll_time"}

	Searchscrolltotal = CatIndicesColumn{"search.scroll_total"}

	Prisearchscrolltotal = CatIndicesColumn{"pri.search.scroll_total"}

	Segmentscount = CatIndicesColumn{"segments.count"}

	Prisegmentscount = CatIndicesColumn{"pri.segments.count"}

	Segmentsmemory = CatIndicesColumn{"segments.memory"}

	Prisegmentsmemory = CatIndicesColumn{"pri.segments.memory"}

	Segmentsindexwritermemory = CatIndicesColumn{"segments.index_writer_memory"}

	Prisegmentsindexwritermemory = CatIndicesColumn{"pri.segments.index_writer_memory"}

	Segmentsversionmapmemory = CatIndicesColumn{"segments.version_map_memory"}

	Prisegmentsversionmapmemory = CatIndicesColumn{"pri.segments.version_map_memory"}

	Segmentsfixedbitsetmemory = CatIndicesColumn{"segments.fixed_bitset_memory"}

	Prisegmentsfixedbitsetmemory = CatIndicesColumn{"pri.segments.fixed_bitset_memory"}

	Warmercurrent = CatIndicesColumn{"warmer.current"}

	Priwarmercurrent = CatIndicesColumn{"pri.warmer.current"}

	Warmertotal = CatIndicesColumn{"warmer.total"}

	Priwarmertotal = CatIndicesColumn{"pri.warmer.total"}

	Warmertotaltime = CatIndicesColumn{"warmer.total_time"}

	Priwarmertotaltime = CatIndicesColumn{"pri.warmer.total_time"}

	Suggestcurrent = CatIndicesColumn{"suggest.current"}

	Prisuggestcurrent = CatIndicesColumn{"pri.suggest.current"}

	Suggesttime = CatIndicesColumn{"suggest.time"}

	Prisuggesttime = CatIndicesColumn{"pri.suggest.time"}

	Suggesttotal = CatIndicesColumn{"suggest.total"}

	Prisuggesttotal = CatIndicesColumn{"pri.suggest.total"}

	Memorytotal = CatIndicesColumn{"memory.total"}

	Primemorytotal = CatIndicesColumn{"pri.memory.total"}

	Bulktotaloperations = CatIndicesColumn{"bulk.total_operations"}

	Pribulktotaloperations = CatIndicesColumn{"pri.bulk.total_operations"}

	Bulktotaltime = CatIndicesColumn{"bulk.total_time"}

	Pribulktotaltime = CatIndicesColumn{"pri.bulk.total_time"}

	Bulktotalsizeinbytes = CatIndicesColumn{"bulk.total_size_in_bytes"}

	Pribulktotalsizeinbytes = CatIndicesColumn{"pri.bulk.total_size_in_bytes"}

	Bulkavgtime = CatIndicesColumn{"bulk.avg_time"}

	Pribulkavgtime = CatIndicesColumn{"pri.bulk.avg_time"}

	Bulkavgsizeinbytes = CatIndicesColumn{"bulk.avg_size_in_bytes"}

	Pribulkavgsizeinbytes = CatIndicesColumn{"pri.bulk.avg_size_in_bytes"}

	Densevectorvaluecount = CatIndicesColumn{"dense_vector.value_count"}

	Pridensevectorvaluecount = CatIndicesColumn{"pri.dense_vector.value_count"}

	Sparsevectorvaluecount = CatIndicesColumn{"sparse_vector.value_count"}

	Prisparsevectorvaluecount = CatIndicesColumn{"pri.sparse_vector.value_count"}
)

func (c CatIndicesColumn) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CatIndicesColumn) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "health":
		*c = Health
	case "status":
		*c = Status
	case "index":
		*c = Index
	case "uuid":
		*c = Uuid
	case "pri":
		*c = Pri
	case "rep":
		*c = Rep
	case "docs.count":
		*c = Docscount
	case "docs.deleted":
		*c = Docsdeleted
	case "creation.date":
		*c = Creationdate
	case "creation.date.string":
		*c = Creationdatestring
	case "store.size":
		*c = Storesize
	case "pri.store.size":
		*c = Pristoresize
	case "dataset.size":
		*c = Datasetsize
	case "completion.size":
		*c = Completionsize
	case "pri.completion.size":
		*c = Pricompletionsize
	case "fielddata.memory_size":
		*c = Fielddatamemorysize
	case "pri.fielddata.memory_size":
		*c = Prifielddatamemorysize
	case "fielddata.evictions":
		*c = Fielddataevictions
	case "pri.fielddata.evictions":
		*c = Prifielddataevictions
	case "query_cache.memory_size":
		*c = Querycachememorysize
	case "pri.query_cache.memory_size":
		*c = Priquerycachememorysize
	case "query_cache.evictions":
		*c = Querycacheevictions
	case "pri.query_cache.evictions":
		*c = Priquerycacheevictions
	case "request_cache.memory_size":
		*c = Requestcachememorysize
	case "pri.request_cache.memory_size":
		*c = Prirequestcachememorysize
	case "request_cache.evictions":
		*c = Requestcacheevictions
	case "pri.request_cache.evictions":
		*c = Prirequestcacheevictions
	case "request_cache.hit_count":
		*c = Requestcachehitcount
	case "pri.request_cache.hit_count":
		*c = Prirequestcachehitcount
	case "request_cache.miss_count":
		*c = Requestcachemisscount
	case "pri.request_cache.miss_count":
		*c = Prirequestcachemisscount
	case "flush.total":
		*c = Flushtotal
	case "pri.flush.total":
		*c = Priflushtotal
	case "flush.total_time":
		*c = Flushtotaltime
	case "pri.flush.total_time":
		*c = Priflushtotaltime
	case "get.current":
		*c = Getcurrent
	case "pri.get.current":
		*c = Prigetcurrent
	case "get.time":
		*c = Gettime
	case "pri.get.time":
		*c = Prigettime
	case "get.total":
		*c = Gettotal
	case "pri.get.total":
		*c = Prigettotal
	case "get.exists_time":
		*c = Getexiststime
	case "pri.get.exists_time":
		*c = Prigetexiststime
	case "get.exists_total":
		*c = Getexiststotal
	case "pri.get.exists_total":
		*c = Prigetexiststotal
	case "get.missing_time":
		*c = Getmissingtime
	case "pri.get.missing_time":
		*c = Prigetmissingtime
	case "get.missing_total":
		*c = Getmissingtotal
	case "pri.get.missing_total":
		*c = Prigetmissingtotal
	case "indexing.delete_current":
		*c = Indexingdeletecurrent
	case "pri.indexing.delete_current":
		*c = Priindexingdeletecurrent
	case "indexing.delete_time":
		*c = Indexingdeletetime
	case "pri.indexing.delete_time":
		*c = Priindexingdeletetime
	case "indexing.delete_total":
		*c = Indexingdeletetotal
	case "pri.indexing.delete_total":
		*c = Priindexingdeletetotal
	case "indexing.index_current":
		*c = Indexingindexcurrent
	case "pri.indexing.index_current":
		*c = Priindexingindexcurrent
	case "indexing.index_time":
		*c = Indexingindextime
	case "pri.indexing.index_time":
		*c = Priindexingindextime
	case "indexing.index_total":
		*c = Indexingindextotal
	case "pri.indexing.index_total":
		*c = Priindexingindextotal
	case "indexing.index_failed":
		*c = Indexingindexfailed
	case "pri.indexing.index_failed":
		*c = Priindexingindexfailed
	case "indexing.index_failed_due_to_version_conflict":
		*c = Indexingindexfailedduetoversionconflict
	case "pri.indexing.index_failed_due_to_version_conflict":
		*c = Priindexingindexfailedduetoversionconflict
	case "merges.current":
		*c = Mergescurrent
	case "pri.merges.current":
		*c = Primergescurrent
	case "merges.current_docs":
		*c = Mergescurrentdocs
	case "pri.merges.current_docs":
		*c = Primergescurrentdocs
	case "merges.current_size":
		*c = Mergescurrentsize
	case "pri.merges.current_size":
		*c = Primergescurrentsize
	case "merges.total":
		*c = Mergestotal
	case "pri.merges.total":
		*c = Primergestotal
	case "merges.total_docs":
		*c = Mergestotaldocs
	case "pri.merges.total_docs":
		*c = Primergestotaldocs
	case "merges.total_size":
		*c = Mergestotalsize
	case "pri.merges.total_size":
		*c = Primergestotalsize
	case "merges.total_time":
		*c = Mergestotaltime
	case "pri.merges.total_time":
		*c = Primergestotaltime
	case "refresh.total":
		*c = Refreshtotal
	case "pri.refresh.total":
		*c = Prirefreshtotal
	case "refresh.time":
		*c = Refreshtime
	case "pri.refresh.time":
		*c = Prirefreshtime
	case "refresh.external_total":
		*c = Refreshexternaltotal
	case "pri.refresh.external_total":
		*c = Prirefreshexternaltotal
	case "refresh.external_time":
		*c = Refreshexternaltime
	case "pri.refresh.external_time":
		*c = Prirefreshexternaltime
	case "refresh.listeners":
		*c = Refreshlisteners
	case "pri.refresh.listeners":
		*c = Prirefreshlisteners
	case "search.fetch_current":
		*c = Searchfetchcurrent
	case "pri.search.fetch_current":
		*c = Prisearchfetchcurrent
	case "search.fetch_time":
		*c = Searchfetchtime
	case "pri.search.fetch_time":
		*c = Prisearchfetchtime
	case "search.fetch_total":
		*c = Searchfetchtotal
	case "pri.search.fetch_total":
		*c = Prisearchfetchtotal
	case "search.open_contexts":
		*c = Searchopencontexts
	case "pri.search.open_contexts":
		*c = Prisearchopencontexts
	case "search.query_current":
		*c = Searchquerycurrent
	case "pri.search.query_current":
		*c = Prisearchquerycurrent
	case "search.query_time":
		*c = Searchquerytime
	case "pri.search.query_time":
		*c = Prisearchquerytime
	case "search.query_total":
		*c = Searchquerytotal
	case "pri.search.query_total":
		*c = Prisearchquerytotal
	case "search.scroll_current":
		*c = Searchscrollcurrent
	case "pri.search.scroll_current":
		*c = Prisearchscrollcurrent
	case "search.scroll_time":
		*c = Searchscrolltime
	case "pri.search.scroll_time":
		*c = Prisearchscrolltime
	case "search.scroll_total":
		*c = Searchscrolltotal
	case "pri.search.scroll_total":
		*c = Prisearchscrolltotal
	case "segments.count":
		*c = Segmentscount
	case "pri.segments.count":
		*c = Prisegmentscount
	case "segments.memory":
		*c = Segmentsmemory
	case "pri.segments.memory":
		*c = Prisegmentsmemory
	case "segments.index_writer_memory":
		*c = Segmentsindexwritermemory
	case "pri.segments.index_writer_memory":
		*c = Prisegmentsindexwritermemory
	case "segments.version_map_memory":
		*c = Segmentsversionmapmemory
	case "pri.segments.version_map_memory":
		*c = Prisegmentsversionmapmemory
	case "segments.fixed_bitset_memory":
		*c = Segmentsfixedbitsetmemory
	case "pri.segments.fixed_bitset_memory":
		*c = Prisegmentsfixedbitsetmemory
	case "warmer.current":
		*c = Warmercurrent
	case "pri.warmer.current":
		*c = Priwarmercurrent
	case "warmer.total":
		*c = Warmertotal
	case "pri.warmer.total":
		*c = Priwarmertotal
	case "warmer.total_time":
		*c = Warmertotaltime
	case "pri.warmer.total_time":
		*c = Priwarmertotaltime
	case "suggest.current":
		*c = Suggestcurrent
	case "pri.suggest.current":
		*c = Prisuggestcurrent
	case "suggest.time":
		*c = Suggesttime
	case "pri.suggest.time":
		*c = Prisuggesttime
	case "suggest.total":
		*c = Suggesttotal
	case "pri.suggest.total":
		*c = Prisuggesttotal
	case "memory.total":
		*c = Memorytotal
	case "pri.memory.total":
		*c = Primemorytotal
	case "bulk.total_operations":
		*c = Bulktotaloperations
	case "pri.bulk.total_operations":
		*c = Pribulktotaloperations
	case "bulk.total_time":
		*c = Bulktotaltime
	case "pri.bulk.total_time":
		*c = Pribulktotaltime
	case "bulk.total_size_in_bytes":
		*c = Bulktotalsizeinbytes
	case "pri.bulk.total_size_in_bytes":
		*c = Pribulktotalsizeinbytes
	case "bulk.avg_time":
		*c = Bulkavgtime
	case "pri.bulk.avg_time":
		*c = Pribulkavgtime
	case "bulk.avg_size_in_bytes":
		*c = Bulkavgsizeinbytes
	case "pri.bulk.avg_size_in_bytes":
		*c = Pribulkavgsizeinbytes
	case "dense_vector.value_count":
		*c = Densevectorvaluecount
	case "pri.dense_vector.value_count":
		*c = Pridensevectorvaluecount
	case "sparse_vector.value_count":
		*c = Sparsevectorvaluecount
	case "pri.sparse_vector.value_count":
		*c = Prisparsevectorvaluecount
	default:
		*c = CatIndicesColumn{string(text)}
	}

	return nil
}

func (c CatIndicesColumn) String() string {
	return c.Name
}
