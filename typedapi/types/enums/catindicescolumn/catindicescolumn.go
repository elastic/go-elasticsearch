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

// Package catindicescolumn
package catindicescolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/cat/_types/CatBase.ts#L1481-L2149
type CatIndicesColumn struct {
	Name string
}

var (

	// Health The current health status.
	Health = CatIndicesColumn{"health"}

	// Status The open/close status.
	Status = CatIndicesColumn{"status"}

	// Index The index name.
	Index = CatIndicesColumn{"index"}

	// Uuid The index UUID.
	Uuid = CatIndicesColumn{"uuid"}

	// Pri The number of primary shards.
	Pri = CatIndicesColumn{"pri"}

	// Rep The number of replica shards.
	Rep = CatIndicesColumn{"rep"}

	// Docscount The number of available documents.
	Docscount = CatIndicesColumn{"docs.count"}

	// Docsdeleted The number of deleted documents.
	Docsdeleted = CatIndicesColumn{"docs.deleted"}

	// Creationdate The index creation date (millisecond value).
	Creationdate = CatIndicesColumn{"creation.date"}

	// Creationdatestring The index creation date (as string).
	Creationdatestring = CatIndicesColumn{"creation.date.string"}

	// Storesize The store size of primaries and replicas.
	Storesize = CatIndicesColumn{"store.size"}

	// Pristoresize The store size of primaries.
	Pristoresize = CatIndicesColumn{"pri.store.size"}

	// Datasetsize The total size of the dataset.
	Datasetsize = CatIndicesColumn{"dataset.size"}

	// Completionsize The size of completion for primaries and replicas.
	Completionsize = CatIndicesColumn{"completion.size"}

	// Pricompletionsize The size of completion for primaries.
	Pricompletionsize = CatIndicesColumn{"pri.completion.size"}

	// Fielddatamemorysize The used fielddata cache for primaries and replicas.
	Fielddatamemorysize = CatIndicesColumn{"fielddata.memory_size"}

	// Prifielddatamemorysize The used fielddata cache for primaries.
	Prifielddatamemorysize = CatIndicesColumn{"pri.fielddata.memory_size"}

	// Fielddataevictions The number of fielddata evictions for primaries and replicas.
	Fielddataevictions = CatIndicesColumn{"fielddata.evictions"}

	// Prifielddataevictions The number of fielddata evictions for primaries.
	Prifielddataevictions = CatIndicesColumn{"pri.fielddata.evictions"}

	// Querycachememorysize The used query cache for primaries and replicas.
	Querycachememorysize = CatIndicesColumn{"query_cache.memory_size"}

	// Priquerycachememorysize The used query cache for primaries.
	Priquerycachememorysize = CatIndicesColumn{"pri.query_cache.memory_size"}

	// Querycacheevictions The number of query cache evictions for primaries and replicas.
	Querycacheevictions = CatIndicesColumn{"query_cache.evictions"}

	// Priquerycacheevictions The number of query cache evictions for primaries.
	Priquerycacheevictions = CatIndicesColumn{"pri.query_cache.evictions"}

	// Requestcachememorysize The used request cache for primaries and replicas.
	Requestcachememorysize = CatIndicesColumn{"request_cache.memory_size"}

	// Prirequestcachememorysize The used request cache for primaries.
	Prirequestcachememorysize = CatIndicesColumn{"pri.request_cache.memory_size"}

	// Requestcacheevictions The number of request cache evictions for primaries and replicas.
	Requestcacheevictions = CatIndicesColumn{"request_cache.evictions"}

	// Prirequestcacheevictions The number of request cache evictions for primaries.
	Prirequestcacheevictions = CatIndicesColumn{"pri.request_cache.evictions"}

	// Requestcachehitcount The request cache hit count for primaries and replicas.
	Requestcachehitcount = CatIndicesColumn{"request_cache.hit_count"}

	// Prirequestcachehitcount The request cache hit count for primaries.
	Prirequestcachehitcount = CatIndicesColumn{"pri.request_cache.hit_count"}

	// Requestcachemisscount The request cache miss count for primaries and replicas.
	Requestcachemisscount = CatIndicesColumn{"request_cache.miss_count"}

	// Prirequestcachemisscount The request cache miss count for primaries.
	Prirequestcachemisscount = CatIndicesColumn{"pri.request_cache.miss_count"}

	// Flushtotal The number of flushes for primaries and replicas.
	Flushtotal = CatIndicesColumn{"flush.total"}

	// Priflushtotal The number of flushes for primaries.
	Priflushtotal = CatIndicesColumn{"pri.flush.total"}

	// Flushtotaltime The time spent in flush for primaries and replicas.
	Flushtotaltime = CatIndicesColumn{"flush.total_time"}

	// Priflushtotaltime The time spent in flush for primaries.
	Priflushtotaltime = CatIndicesColumn{"pri.flush.total_time"}

	// Getcurrent The number of current get operations for primaries and replicas.
	Getcurrent = CatIndicesColumn{"get.current"}

	// Prigetcurrent The number of current get operations for primaries.
	Prigetcurrent = CatIndicesColumn{"pri.get.current"}

	// Gettime The time spent in get for primaries and replicas.
	Gettime = CatIndicesColumn{"get.time"}

	// Prigettime The time spent in get for primaries.
	Prigettime = CatIndicesColumn{"pri.get.time"}

	// Gettotal The number of get operations for primaries and replicas.
	Gettotal = CatIndicesColumn{"get.total"}

	// Prigettotal The number of get operations for primaries.
	Prigettotal = CatIndicesColumn{"pri.get.total"}

	// Getexiststime The time spent in successful gets for primaries and replicas.
	Getexiststime = CatIndicesColumn{"get.exists_time"}

	// Prigetexiststime The time spent in successful gets for primaries.
	Prigetexiststime = CatIndicesColumn{"pri.get.exists_time"}

	// Getexiststotal The number of successful gets for primaries and replicas.
	Getexiststotal = CatIndicesColumn{"get.exists_total"}

	// Prigetexiststotal The number of successful gets for primaries.
	Prigetexiststotal = CatIndicesColumn{"pri.get.exists_total"}

	// Getmissingtime The time spent in failed gets for primaries and replicas.
	Getmissingtime = CatIndicesColumn{"get.missing_time"}

	// Prigetmissingtime The time spent in failed gets for primaries.
	Prigetmissingtime = CatIndicesColumn{"pri.get.missing_time"}

	// Getmissingtotal The number of failed gets for primaries and replicas.
	Getmissingtotal = CatIndicesColumn{"get.missing_total"}

	// Prigetmissingtotal The number of failed gets for primaries.
	Prigetmissingtotal = CatIndicesColumn{"pri.get.missing_total"}

	// Indexingdeletecurrent The number of current deletions for primaries and replicas.
	Indexingdeletecurrent = CatIndicesColumn{"indexing.delete_current"}

	// Priindexingdeletecurrent The number of current deletions for primaries.
	Priindexingdeletecurrent = CatIndicesColumn{"pri.indexing.delete_current"}

	// Indexingdeletetime The time spent in deletions for primaries and replicas.
	Indexingdeletetime = CatIndicesColumn{"indexing.delete_time"}

	// Priindexingdeletetime The time spent in deletions for primaries.
	Priindexingdeletetime = CatIndicesColumn{"pri.indexing.delete_time"}

	// Indexingdeletetotal The number of delete operations for primaries and replicas.
	Indexingdeletetotal = CatIndicesColumn{"indexing.delete_total"}

	// Priindexingdeletetotal The number of delete operations for primaries.
	Priindexingdeletetotal = CatIndicesColumn{"pri.indexing.delete_total"}

	// Indexingindexcurrent The number of current indexing operations for primaries and replicas.
	Indexingindexcurrent = CatIndicesColumn{"indexing.index_current"}

	// Priindexingindexcurrent The number of current indexing operations for primaries.
	Priindexingindexcurrent = CatIndicesColumn{"pri.indexing.index_current"}

	// Indexingindextime The time spent in indexing for primaries and replicas.
	Indexingindextime = CatIndicesColumn{"indexing.index_time"}

	// Priindexingindextime The time spent in indexing for primaries.
	Priindexingindextime = CatIndicesColumn{"pri.indexing.index_time"}

	// Indexingindextotal The number of indexing operations for primaries and replicas.
	Indexingindextotal = CatIndicesColumn{"indexing.index_total"}

	// Priindexingindextotal The number of indexing operations for primaries.
	Priindexingindextotal = CatIndicesColumn{"pri.indexing.index_total"}

	// Indexingindexfailed The number of failed indexing operations for primaries and replicas.
	Indexingindexfailed = CatIndicesColumn{"indexing.index_failed"}

	// Priindexingindexfailed The number of failed indexing operations for primaries.
	Priindexingindexfailed = CatIndicesColumn{"pri.indexing.index_failed"}

	// Indexingindexfailedduetoversionconflict The number of failed indexing operations due to version conflict for
	// primaries and replicas.
	Indexingindexfailedduetoversionconflict = CatIndicesColumn{"indexing.index_failed_due_to_version_conflict"}

	// Priindexingindexfailedduetoversionconflict The number of failed indexing operations due to version conflict for
	// primaries.
	Priindexingindexfailedduetoversionconflict = CatIndicesColumn{"pri.indexing.index_failed_due_to_version_conflict"}

	// Mergescurrent The number of current merges for primaries and replicas.
	Mergescurrent = CatIndicesColumn{"merges.current"}

	// Primergescurrent The number of current merges for primaries.
	Primergescurrent = CatIndicesColumn{"pri.merges.current"}

	// Mergescurrentdocs The number of current merging documents for primaries and replicas.
	Mergescurrentdocs = CatIndicesColumn{"merges.current_docs"}

	// Primergescurrentdocs The number of current merging documents for primaries.
	Primergescurrentdocs = CatIndicesColumn{"pri.merges.current_docs"}

	// Mergescurrentsize The size of current merges for primaries and replicas.
	Mergescurrentsize = CatIndicesColumn{"merges.current_size"}

	// Primergescurrentsize The size of current merges for primaries.
	Primergescurrentsize = CatIndicesColumn{"pri.merges.current_size"}

	// Mergestotal The number of completed merge operations for primaries and replicas.
	Mergestotal = CatIndicesColumn{"merges.total"}

	// Primergestotal The number of completed merge operations for primaries.
	Primergestotal = CatIndicesColumn{"pri.merges.total"}

	// Mergestotaldocs The number of merged documents for primaries and replicas.
	Mergestotaldocs = CatIndicesColumn{"merges.total_docs"}

	// Primergestotaldocs The number of merged documents for primaries.
	Primergestotaldocs = CatIndicesColumn{"pri.merges.total_docs"}

	// Mergestotalsize The merged size for primaries and replicas.
	Mergestotalsize = CatIndicesColumn{"merges.total_size"}

	// Primergestotalsize The merged size for primaries.
	Primergestotalsize = CatIndicesColumn{"pri.merges.total_size"}

	// Mergestotaltime The time spent in merges for primaries and replicas.
	Mergestotaltime = CatIndicesColumn{"merges.total_time"}

	// Primergestotaltime The time spent in merges for primaries.
	Primergestotaltime = CatIndicesColumn{"pri.merges.total_time"}

	// Refreshtotal The total refreshes for primaries and replicas.
	Refreshtotal = CatIndicesColumn{"refresh.total"}

	// Prirefreshtotal The total refreshes for primaries.
	Prirefreshtotal = CatIndicesColumn{"pri.refresh.total"}

	// Refreshtime The time spent in refreshes for primaries and replicas.
	Refreshtime = CatIndicesColumn{"refresh.time"}

	// Prirefreshtime The time spent in refreshes for primaries.
	Prirefreshtime = CatIndicesColumn{"pri.refresh.time"}

	// Refreshexternaltotal The total external refreshes for primaries and replicas.
	Refreshexternaltotal = CatIndicesColumn{"refresh.external_total"}

	// Prirefreshexternaltotal The total external refreshes for primaries.
	Prirefreshexternaltotal = CatIndicesColumn{"pri.refresh.external_total"}

	// Refreshexternaltime The time spent in external refreshes for primaries and replicas.
	Refreshexternaltime = CatIndicesColumn{"refresh.external_time"}

	// Prirefreshexternaltime The time spent in external refreshes for primaries.
	Prirefreshexternaltime = CatIndicesColumn{"pri.refresh.external_time"}

	// Refreshlisteners The number of pending refresh listeners for primaries and replicas.
	Refreshlisteners = CatIndicesColumn{"refresh.listeners"}

	// Prirefreshlisteners The number of pending refresh listeners for primaries.
	Prirefreshlisteners = CatIndicesColumn{"pri.refresh.listeners"}

	// Searchfetchcurrent The current fetch phase operations for primaries and replicas.
	Searchfetchcurrent = CatIndicesColumn{"search.fetch_current"}

	// Prisearchfetchcurrent The current fetch phase operations for primaries.
	Prisearchfetchcurrent = CatIndicesColumn{"pri.search.fetch_current"}

	// Searchfetchtime The time spent in fetch phase for primaries and replicas.
	Searchfetchtime = CatIndicesColumn{"search.fetch_time"}

	// Prisearchfetchtime The time spent in fetch phase for primaries.
	Prisearchfetchtime = CatIndicesColumn{"pri.search.fetch_time"}

	// Searchfetchtotal The total fetch operations for primaries and replicas.
	Searchfetchtotal = CatIndicesColumn{"search.fetch_total"}

	// Prisearchfetchtotal The total fetch operations for primaries.
	Prisearchfetchtotal = CatIndicesColumn{"pri.search.fetch_total"}

	// Searchopencontexts The open search contexts for primaries and replicas.
	Searchopencontexts = CatIndicesColumn{"search.open_contexts"}

	// Prisearchopencontexts The open search contexts for primaries.
	Prisearchopencontexts = CatIndicesColumn{"pri.search.open_contexts"}

	// Searchquerycurrent The current query phase operations for primaries and replicas.
	Searchquerycurrent = CatIndicesColumn{"search.query_current"}

	// Prisearchquerycurrent The current query phase operations for primaries.
	Prisearchquerycurrent = CatIndicesColumn{"pri.search.query_current"}

	// Searchquerytime The time spent in query phase for primaries and replicas.
	Searchquerytime = CatIndicesColumn{"search.query_time"}

	// Prisearchquerytime The time spent in query phase for primaries.
	Prisearchquerytime = CatIndicesColumn{"pri.search.query_time"}

	// Searchquerytotal The total query phase operations for primaries and replicas.
	Searchquerytotal = CatIndicesColumn{"search.query_total"}

	// Prisearchquerytotal The total query phase operations for primaries.
	Prisearchquerytotal = CatIndicesColumn{"pri.search.query_total"}

	// Searchscrollcurrent The open scroll contexts for primaries and replicas.
	Searchscrollcurrent = CatIndicesColumn{"search.scroll_current"}

	// Prisearchscrollcurrent The open scroll contexts for primaries.
	Prisearchscrollcurrent = CatIndicesColumn{"pri.search.scroll_current"}

	// Searchscrolltime The time scroll contexts held open for primaries and replicas.
	Searchscrolltime = CatIndicesColumn{"search.scroll_time"}

	// Prisearchscrolltime The time scroll contexts held open for primaries.
	Prisearchscrolltime = CatIndicesColumn{"pri.search.scroll_time"}

	// Searchscrolltotal The completed scroll contexts for primaries and replicas.
	Searchscrolltotal = CatIndicesColumn{"search.scroll_total"}

	// Prisearchscrolltotal The completed scroll contexts for primaries.
	Prisearchscrolltotal = CatIndicesColumn{"pri.search.scroll_total"}

	// Segmentscount The number of segments for primaries and replicas.
	Segmentscount = CatIndicesColumn{"segments.count"}

	// Prisegmentscount The number of segments for primaries.
	Prisegmentscount = CatIndicesColumn{"pri.segments.count"}

	// Segmentsmemory The memory used by segments for primaries and replicas.
	Segmentsmemory = CatIndicesColumn{"segments.memory"}

	// Prisegmentsmemory The memory used by segments for primaries.
	Prisegmentsmemory = CatIndicesColumn{"pri.segments.memory"}

	// Segmentsindexwritermemory The memory used by index writer for primaries and replicas.
	Segmentsindexwritermemory = CatIndicesColumn{"segments.index_writer_memory"}

	// Prisegmentsindexwritermemory The memory used by index writer for primaries.
	Prisegmentsindexwritermemory = CatIndicesColumn{"pri.segments.index_writer_memory"}

	// Segmentsversionmapmemory The memory used by version map for primaries and replicas.
	Segmentsversionmapmemory = CatIndicesColumn{"segments.version_map_memory"}

	// Prisegmentsversionmapmemory The memory used by version map for primaries.
	Prisegmentsversionmapmemory = CatIndicesColumn{"pri.segments.version_map_memory"}

	// Segmentsfixedbitsetmemory The memory used by fixed bit sets for nested object field types and type
	// filters for types referred in _parent fields. Applicable for primaries and
	// replicas.
	Segmentsfixedbitsetmemory = CatIndicesColumn{"segments.fixed_bitset_memory"}

	// Prisegmentsfixedbitsetmemory The memory used by fixed bit sets for nested object field types and type
	// filters for types referred in _parent fields. Applicable for primaries.
	Prisegmentsfixedbitsetmemory = CatIndicesColumn{"pri.segments.fixed_bitset_memory"}

	// Warmercurrent The current warmer operations for primaries and replicas.
	Warmercurrent = CatIndicesColumn{"warmer.current"}

	// Priwarmercurrent The current warmer operations for primaries.
	Priwarmercurrent = CatIndicesColumn{"pri.warmer.current"}

	// Warmertotal The total warmer operations for primaries and replicas.
	Warmertotal = CatIndicesColumn{"warmer.total"}

	// Priwarmertotal The total warmer operations for primaries.
	Priwarmertotal = CatIndicesColumn{"pri.warmer.total"}

	// Warmertotaltime The time spent in warmers for primaries and replicas.
	Warmertotaltime = CatIndicesColumn{"warmer.total_time"}

	// Priwarmertotaltime The time spent in warmers for primaries.
	Priwarmertotaltime = CatIndicesColumn{"pri.warmer.total_time"}

	// Suggestcurrent The current suggest operations for primaries and replicas.
	Suggestcurrent = CatIndicesColumn{"suggest.current"}

	// Prisuggestcurrent The current suggest operations for primaries.
	Prisuggestcurrent = CatIndicesColumn{"pri.suggest.current"}

	// Suggesttime The time spent in suggest for primaries and replicas.
	Suggesttime = CatIndicesColumn{"suggest.time"}

	// Prisuggesttime The time spent in suggest for primaries.
	Prisuggesttime = CatIndicesColumn{"pri.suggest.time"}

	// Suggesttotal The number of suggest operations for primaries and replicas.
	Suggesttotal = CatIndicesColumn{"suggest.total"}

	// Prisuggesttotal The number of suggest operations for primaries.
	Prisuggesttotal = CatIndicesColumn{"pri.suggest.total"}

	// Memorytotal The total used memory for primaries and replicas.
	Memorytotal = CatIndicesColumn{"memory.total"}

	// Primemorytotal The total used memory for primaries.
	Primemorytotal = CatIndicesColumn{"pri.memory.total"}

	// Bulktotaloperations The number of bulk shard operations for primaries and replicas.
	Bulktotaloperations = CatIndicesColumn{"bulk.total_operations"}

	// Pribulktotaloperations The number of bulk shard operations for primaries.
	Pribulktotaloperations = CatIndicesColumn{"pri.bulk.total_operations"}

	// Bulktotaltime The time spent in shard bulk for primaries and replicas.
	Bulktotaltime = CatIndicesColumn{"bulk.total_time"}

	// Pribulktotaltime The time spent in shard bulk for primaries.
	Pribulktotaltime = CatIndicesColumn{"pri.bulk.total_time"}

	// Bulktotalsizeinbytes The total size in bytes of shard bulk for primaries and replicas.
	Bulktotalsizeinbytes = CatIndicesColumn{"bulk.total_size_in_bytes"}

	// Pribulktotalsizeinbytes The total size in bytes of shard bulk for primaries.
	Pribulktotalsizeinbytes = CatIndicesColumn{"pri.bulk.total_size_in_bytes"}

	// Bulkavgtime The average time spent in shard bulk for primaries and replicas.
	Bulkavgtime = CatIndicesColumn{"bulk.avg_time"}

	// Pribulkavgtime The average time spent in shard bulk for primaries.
	Pribulkavgtime = CatIndicesColumn{"pri.bulk.avg_time"}

	// Bulkavgsizeinbytes The average size in bytes of shard bulk for primaries and replicas.
	Bulkavgsizeinbytes = CatIndicesColumn{"bulk.avg_size_in_bytes"}

	// Pribulkavgsizeinbytes The average size in bytes of shard bulk for primaries.
	Pribulkavgsizeinbytes = CatIndicesColumn{"pri.bulk.avg_size_in_bytes"}

	// Densevectorvaluecount The total count of indexed dense vectors for primaries and replicas.
	Densevectorvaluecount = CatIndicesColumn{"dense_vector.value_count"}

	// Pridensevectorvaluecount The total count of indexed dense vectors for primaries.
	Pridensevectorvaluecount = CatIndicesColumn{"pri.dense_vector.value_count"}

	// Sparsevectorvaluecount The total count of indexed sparse vectors for primaries and replicas.
	Sparsevectorvaluecount = CatIndicesColumn{"sparse_vector.value_count"}

	// Prisparsevectorvaluecount The total count of indexed sparse vectors for primaries.
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
