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

// Package catnodecolumn
package catnodecolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/cat/_types/CatBase.ts#L567-L1024
type CatNodeColumn struct {
	Name string
}

var (

	// Build The Elasticsearch build hash. For example: `5c03844`.
	Build = CatNodeColumn{"build"}

	// Completionsize The size of completion. For example: `0b`.
	Completionsize = CatNodeColumn{"completion.size"}

	// Cpu The percentage of recent system CPU used.
	Cpu = CatNodeColumn{"cpu"}

	// Diskavail The available disk space. For example: `198.4gb`.
	Diskavail = CatNodeColumn{"disk.avail"}

	// Disktotal The total disk space. For example: `458.3gb`.
	Disktotal = CatNodeColumn{"disk.total"}

	// Diskused The used disk space. For example: `259.8gb`.
	Diskused = CatNodeColumn{"disk.used"}

	// Diskusedpercent The percentage of disk space used.
	Diskusedpercent = CatNodeColumn{"disk.used_percent"}

	// Fielddataevictions The number of fielddata cache evictions.
	Fielddataevictions = CatNodeColumn{"fielddata.evictions"}

	// Fielddatamemorysize The fielddata cache memory used. For example: `0b`.
	Fielddatamemorysize = CatNodeColumn{"fielddata.memory_size"}

	// Filedesccurrent The number of file descriptors used.
	Filedesccurrent = CatNodeColumn{"file_desc.current"}

	// Filedescmax The maximum number of file descriptors.
	Filedescmax = CatNodeColumn{"file_desc.max"}

	// Filedescpercent The percentage of file descriptors used.
	Filedescpercent = CatNodeColumn{"file_desc.percent"}

	// Flushtotal The number of flushes.
	Flushtotal = CatNodeColumn{"flush.total"}

	// Flushtotaltime The amount of time spent in flush.
	Flushtotaltime = CatNodeColumn{"flush.total_time"}

	// Getcurrent The number of current get operations.
	Getcurrent = CatNodeColumn{"get.current"}

	// Getexiststime The time spent in successful get operations. For example: `14ms`.
	Getexiststime = CatNodeColumn{"get.exists_time"}

	// Getexiststotal The number of successful get operations.
	Getexiststotal = CatNodeColumn{"get.exists_total"}

	// Getmissingtime The time spent in failed get operations. For example: `0s`.
	Getmissingtime = CatNodeColumn{"get.missing_time"}

	// Getmissingtotal The number of failed get operations.
	Getmissingtotal = CatNodeColumn{"get.missing_total"}

	// Gettime The amount of time spent in get operations. For example: `14ms`.
	Gettime = CatNodeColumn{"get.time"}

	// Gettotal The number of get operations.
	Gettotal = CatNodeColumn{"get.total"}

	// Heapcurrent The used heap size. For example: `311.2mb`.
	Heapcurrent = CatNodeColumn{"heap.current"}

	// Heapmax The total heap size. For example: `4gb`.
	Heapmax = CatNodeColumn{"heap.max"}

	// Heappercent The used percentage of total allocated Elasticsearch JVM heap. This value
	// reflects only the Elasticsearch process running within the operating system
	// and is the most direct indicator of its JVM, heap, or memory resource
	// performance.
	Heappercent = CatNodeColumn{"heap.percent"}

	// Httpaddress The bound HTTP address.
	Httpaddress = CatNodeColumn{"http_address"}

	// Id The identifier for the node.
	Id = CatNodeColumn{"id"}

	// Indexingdeletecurrent The number of current deletion operations.
	Indexingdeletecurrent = CatNodeColumn{"indexing.delete_current"}

	// Indexingdeletetime The time spent in deletion operations. For example: `2ms`.
	Indexingdeletetime = CatNodeColumn{"indexing.delete_time"}

	// Indexingdeletetotal The number of deletion operations.
	Indexingdeletetotal = CatNodeColumn{"indexing.delete_total"}

	// Indexingindexcurrent The number of current indexing operations.
	Indexingindexcurrent = CatNodeColumn{"indexing.index_current"}

	// Indexingindexfailed The number of failed indexing operations.
	Indexingindexfailed = CatNodeColumn{"indexing.index_failed"}

	// Indexingindexfailedduetoversionconflict The number of indexing operations that failed due to version conflict.
	Indexingindexfailedduetoversionconflict = CatNodeColumn{"indexing.index_failed_due_to_version_conflict"}

	// Indexingindextime The time spent in indexing operations. For example: `134ms`.
	Indexingindextime = CatNodeColumn{"indexing.index_time"}

	// Indexingindextotal The number of indexing operations.
	Indexingindextotal = CatNodeColumn{"indexing.index_total"}

	// Ip The IP address.
	Ip = CatNodeColumn{"ip"}

	// Jdk The Java version. For example: `1.8.0`.
	Jdk = CatNodeColumn{"jdk"}

	// Load1m The most recent load average. For example: `0.22`.
	Load1m = CatNodeColumn{"load_1m"}

	// Load5m The load average for the last five minutes. For example: `0.78`.
	Load5m = CatNodeColumn{"load_5m"}

	// Load15m The load average for the last fifteen minutes. For example: `1.24`.
	Load15m = CatNodeColumn{"load_15m"}

	// Availableprocessors The number of available processors (logical CPU cores available to the JVM).
	Availableprocessors = CatNodeColumn{"available_processors"}

	// Mappingstotalcount The number of mappings, including runtime and object fields.
	Mappingstotalcount = CatNodeColumn{"mappings.total_count"}

	// Mappingstotalestimatedoverheadinbytes The estimated heap overhead, in bytes, of mappings on this node, which allows
	// for 1KiB of heap for every mapped field.
	Mappingstotalestimatedoverheadinbytes = CatNodeColumn{"mappings.total_estimated_overhead_in_bytes"}

	// Master Indicates whether the node is the elected master node. Returned values
	// include `*` (elected master) and `-` (not elected master).
	Master = CatNodeColumn{"master"}

	// Mergescurrent The number of current merge operations.
	Mergescurrent = CatNodeColumn{"merges.current"}

	// Mergescurrentdocs The number of current merging documents.
	Mergescurrentdocs = CatNodeColumn{"merges.current_docs"}

	// Mergescurrentsize The size of current merges. For example: `0b`.
	Mergescurrentsize = CatNodeColumn{"merges.current_size"}

	// Mergestotal The number of completed merge operations.
	Mergestotal = CatNodeColumn{"merges.total"}

	// Mergestotaldocs The number of merged documents.
	Mergestotaldocs = CatNodeColumn{"merges.total_docs"}

	// Mergestotalsize The total size of merges. For example: `0b`.
	Mergestotalsize = CatNodeColumn{"merges.total_size"}

	// Mergestotaltime The time spent merging documents. For example: `0s`.
	Mergestotaltime = CatNodeColumn{"merges.total_time"}

	// Name The node name.
	Name = CatNodeColumn{"name"}

	// Noderole The roles of the node. Returned values include `c` (cold node), `d` (data
	// node), `f` (frozen node), `h` (hot node), `i` (ingest node), `l` (machine
	// learning node), `m` (master-eligible node), `r` (remote cluster client node),
	// `s` (content node), `t` (transform node), `v` (voting-only node), `w` (warm
	// node), and `-` (coordinating node only). For example, `dim` indicates a
	// master-eligible data and ingest node.
	Noderole = CatNodeColumn{"node.role"}

	// Pid The process identifier.
	Pid = CatNodeColumn{"pid"}

	// Port The bound transport port number.
	Port = CatNodeColumn{"port"}

	// Querycachememorysize The used query cache memory. For example: `0b`.
	Querycachememorysize = CatNodeColumn{"query_cache.memory_size"}

	// Querycacheevictions The number of query cache evictions.
	Querycacheevictions = CatNodeColumn{"query_cache.evictions"}

	// Querycachehitcount The query cache hit count.
	Querycachehitcount = CatNodeColumn{"query_cache.hit_count"}

	// Querycachemisscount The query cache miss count.
	Querycachemisscount = CatNodeColumn{"query_cache.miss_count"}

	// Ramcurrent The used total memory. For example: `513.4mb`.
	Ramcurrent = CatNodeColumn{"ram.current"}

	// Rammax The total memory. For example: `2.9gb`.
	Rammax = CatNodeColumn{"ram.max"}

	// Rampercent The used percentage of the total operating system memory. This reflects all
	// processes running on the operating system instead of only Elasticsearch and
	// is not guaranteed to correlate to its performance.
	Rampercent = CatNodeColumn{"ram.percent"}

	// Refreshtotal The number of refresh operations.
	Refreshtotal = CatNodeColumn{"refresh.total"}

	// Refreshtime The time spent in refresh operations. For example: `91ms`.
	Refreshtime = CatNodeColumn{"refresh.time"}

	// Requestcachememorysize The used request cache memory. For example: `0b`.
	Requestcachememorysize = CatNodeColumn{"request_cache.memory_size"}

	// Requestcacheevictions The number of request cache evictions.
	Requestcacheevictions = CatNodeColumn{"request_cache.evictions"}

	// Requestcachehitcount The request cache hit count.
	Requestcachehitcount = CatNodeColumn{"request_cache.hit_count"}

	// Requestcachemisscount The request cache miss count.
	Requestcachemisscount = CatNodeColumn{"request_cache.miss_count"}

	// Scriptcompilations The number of total script compilations.
	Scriptcompilations = CatNodeColumn{"script.compilations"}

	// Scriptcacheevictions The number of total compiled scripts evicted from cache.
	Scriptcacheevictions = CatNodeColumn{"script.cache_evictions"}

	// Searchfetchcurrent The number of current fetch phase operations.
	Searchfetchcurrent = CatNodeColumn{"search.fetch_current"}

	// Searchfetchtime The time spent in fetch phase. For example: `37ms`.
	Searchfetchtime = CatNodeColumn{"search.fetch_time"}

	// Searchfetchtotal The number of fetch operations.
	Searchfetchtotal = CatNodeColumn{"search.fetch_total"}

	// Searchopencontexts The number of open search contexts.
	Searchopencontexts = CatNodeColumn{"search.open_contexts"}

	// Searchquerycurrent The number of current query phase operations.
	Searchquerycurrent = CatNodeColumn{"search.query_current"}

	// Searchquerytime The time spent in query phase. For example: `43ms`.
	Searchquerytime = CatNodeColumn{"search.query_time"}

	// Searchquerytotal The number of query operations.
	Searchquerytotal = CatNodeColumn{"search.query_total"}

	// Searchscrollcurrent The number of open scroll contexts.
	Searchscrollcurrent = CatNodeColumn{"search.scroll_current"}

	// Searchscrolltime The amount of time scroll contexts were held open. For example: `2m`.
	Searchscrolltime = CatNodeColumn{"search.scroll_time"}

	// Searchscrolltotal The number of completed scroll contexts.
	Searchscrolltotal = CatNodeColumn{"search.scroll_total"}

	// Segmentscount The number of segments.
	Segmentscount = CatNodeColumn{"segments.count"}

	// Segmentsfixedbitsetmemory The memory used by fixed bit sets for nested object field types and type
	// filters for types referred in join fields. For example: `1.0kb`.
	Segmentsfixedbitsetmemory = CatNodeColumn{"segments.fixed_bitset_memory"}

	// Segmentsindexwritermemory The memory used by the index writer. For example: `18mb`.
	Segmentsindexwritermemory = CatNodeColumn{"segments.index_writer_memory"}

	// Segmentsmemory The memory used by segments. For example: `1.4kb`.
	Segmentsmemory = CatNodeColumn{"segments.memory"}

	// Segmentsversionmapmemory The memory used by the version map. For example: `1.0kb`.
	Segmentsversionmapmemory = CatNodeColumn{"segments.version_map_memory"}

	// Shardstatstotalcount The number of shards assigned.
	Shardstatstotalcount = CatNodeColumn{"shard_stats.total_count"}

	// Suggestcurrent The number of current suggest operations.
	Suggestcurrent = CatNodeColumn{"suggest.current"}

	// Suggesttime The time spent in suggest operations.
	Suggesttime = CatNodeColumn{"suggest.time"}

	// Suggesttotal The number of suggest operations.
	Suggesttotal = CatNodeColumn{"suggest.total"}

	// Uptime The amount of node uptime. For example: `17.3m`.
	Uptime = CatNodeColumn{"uptime"}

	// Version The Elasticsearch version. For example: `9.0.0`.
	Version = CatNodeColumn{"version"}
)

func (c CatNodeColumn) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CatNodeColumn) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "build":
		*c = Build
	case "completion.size":
		*c = Completionsize
	case "cpu":
		*c = Cpu
	case "disk.avail":
		*c = Diskavail
	case "disk.total":
		*c = Disktotal
	case "disk.used":
		*c = Diskused
	case "disk.used_percent":
		*c = Diskusedpercent
	case "fielddata.evictions":
		*c = Fielddataevictions
	case "fielddata.memory_size":
		*c = Fielddatamemorysize
	case "file_desc.current":
		*c = Filedesccurrent
	case "file_desc.max":
		*c = Filedescmax
	case "file_desc.percent":
		*c = Filedescpercent
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
	case "heap.current":
		*c = Heapcurrent
	case "heap.max":
		*c = Heapmax
	case "heap.percent":
		*c = Heappercent
	case "http_address":
		*c = Httpaddress
	case "id":
		*c = Id
	case "indexing.delete_current":
		*c = Indexingdeletecurrent
	case "indexing.delete_time":
		*c = Indexingdeletetime
	case "indexing.delete_total":
		*c = Indexingdeletetotal
	case "indexing.index_current":
		*c = Indexingindexcurrent
	case "indexing.index_failed":
		*c = Indexingindexfailed
	case "indexing.index_failed_due_to_version_conflict":
		*c = Indexingindexfailedduetoversionconflict
	case "indexing.index_time":
		*c = Indexingindextime
	case "indexing.index_total":
		*c = Indexingindextotal
	case "ip":
		*c = Ip
	case "jdk":
		*c = Jdk
	case "load_1m":
		*c = Load1m
	case "load_5m":
		*c = Load5m
	case "load_15m":
		*c = Load15m
	case "available_processors":
		*c = Availableprocessors
	case "mappings.total_count":
		*c = Mappingstotalcount
	case "mappings.total_estimated_overhead_in_bytes":
		*c = Mappingstotalestimatedoverheadinbytes
	case "master":
		*c = Master
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
	case "name":
		*c = Name
	case "node.role":
		*c = Noderole
	case "pid":
		*c = Pid
	case "port":
		*c = Port
	case "query_cache.memory_size":
		*c = Querycachememorysize
	case "query_cache.evictions":
		*c = Querycacheevictions
	case "query_cache.hit_count":
		*c = Querycachehitcount
	case "query_cache.miss_count":
		*c = Querycachemisscount
	case "ram.current":
		*c = Ramcurrent
	case "ram.max":
		*c = Rammax
	case "ram.percent":
		*c = Rampercent
	case "refresh.total":
		*c = Refreshtotal
	case "refresh.time":
		*c = Refreshtime
	case "request_cache.memory_size":
		*c = Requestcachememorysize
	case "request_cache.evictions":
		*c = Requestcacheevictions
	case "request_cache.hit_count":
		*c = Requestcachehitcount
	case "request_cache.miss_count":
		*c = Requestcachemisscount
	case "script.compilations":
		*c = Scriptcompilations
	case "script.cache_evictions":
		*c = Scriptcacheevictions
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
	case "shard_stats.total_count":
		*c = Shardstatstotalcount
	case "suggest.current":
		*c = Suggestcurrent
	case "suggest.time":
		*c = Suggesttime
	case "suggest.total":
		*c = Suggesttotal
	case "uptime":
		*c = Uptime
	case "version":
		*c = Version
	default:
		*c = CatNodeColumn{string(text)}
	}

	return nil
}

func (c CatNodeColumn) String() string {
	return c.Name
}
