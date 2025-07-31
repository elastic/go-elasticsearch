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

// Package catnodecolumn
package catnodecolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/cat/_types/CatBase.ts#L560-L1012
type CatNodeColumn struct {
	Name string
}

var (
	Build = CatNodeColumn{"build"}

	Completionsize = CatNodeColumn{"completion.size"}

	Cpu = CatNodeColumn{"cpu"}

	Diskavail = CatNodeColumn{"disk.avail"}

	Disktotal = CatNodeColumn{"disk.total"}

	Diskused = CatNodeColumn{"disk.used"}

	Diskusedpercent = CatNodeColumn{"disk.used_percent"}

	Fielddataevictions = CatNodeColumn{"fielddata.evictions"}

	Fielddatamemorysize = CatNodeColumn{"fielddata.memory_size"}

	Filedesccurrent = CatNodeColumn{"file_desc.current"}

	Filedescmax = CatNodeColumn{"file_desc.max"}

	Filedescpercent = CatNodeColumn{"file_desc.percent"}

	Flushtotal = CatNodeColumn{"flush.total"}

	Flushtotaltime = CatNodeColumn{"flush.total_time"}

	Getcurrent = CatNodeColumn{"get.current"}

	Getexiststime = CatNodeColumn{"get.exists_time"}

	Getexiststotal = CatNodeColumn{"get.exists_total"}

	Getmissingtime = CatNodeColumn{"get.missing_time"}

	Getmissingtotal = CatNodeColumn{"get.missing_total"}

	Gettime = CatNodeColumn{"get.time"}

	Gettotal = CatNodeColumn{"get.total"}

	Heapcurrent = CatNodeColumn{"heap.current"}

	Heapmax = CatNodeColumn{"heap.max"}

	Heappercent = CatNodeColumn{"heap.percent"}

	Httpaddress = CatNodeColumn{"http_address"}

	Id = CatNodeColumn{"id"}

	Indexingdeletecurrent = CatNodeColumn{"indexing.delete_current"}

	Indexingdeletetime = CatNodeColumn{"indexing.delete_time"}

	Indexingdeletetotal = CatNodeColumn{"indexing.delete_total"}

	Indexingindexcurrent = CatNodeColumn{"indexing.index_current"}

	Indexingindexfailed = CatNodeColumn{"indexing.index_failed"}

	Indexingindexfailedduetoversionconflict = CatNodeColumn{"indexing.index_failed_due_to_version_conflict"}

	Indexingindextime = CatNodeColumn{"indexing.index_time"}

	Indexingindextotal = CatNodeColumn{"indexing.index_total"}

	Ip = CatNodeColumn{"ip"}

	Jdk = CatNodeColumn{"jdk"}

	Load1m = CatNodeColumn{"load_1m"}

	Load5m = CatNodeColumn{"load_5m"}

	Load15m = CatNodeColumn{"load_15m"}

	Mappingstotalcount = CatNodeColumn{"mappings.total_count"}

	Mappingstotalestimatedoverheadinbytes = CatNodeColumn{"mappings.total_estimated_overhead_in_bytes"}

	Master = CatNodeColumn{"master"}

	Mergescurrent = CatNodeColumn{"merges.current"}

	Mergescurrentdocs = CatNodeColumn{"merges.current_docs"}

	Mergescurrentsize = CatNodeColumn{"merges.current_size"}

	Mergestotal = CatNodeColumn{"merges.total"}

	Mergestotaldocs = CatNodeColumn{"merges.total_docs"}

	Mergestotalsize = CatNodeColumn{"merges.total_size"}

	Mergestotaltime = CatNodeColumn{"merges.total_time"}

	Name = CatNodeColumn{"name"}

	Noderole = CatNodeColumn{"node.role"}

	Pid = CatNodeColumn{"pid"}

	Port = CatNodeColumn{"port"}

	Querycachememorysize = CatNodeColumn{"query_cache.memory_size"}

	Querycacheevictions = CatNodeColumn{"query_cache.evictions"}

	Querycachehitcount = CatNodeColumn{"query_cache.hit_count"}

	Querycachemisscount = CatNodeColumn{"query_cache.miss_count"}

	Ramcurrent = CatNodeColumn{"ram.current"}

	Rammax = CatNodeColumn{"ram.max"}

	Rampercent = CatNodeColumn{"ram.percent"}

	Refreshtotal = CatNodeColumn{"refresh.total"}

	Refreshtime = CatNodeColumn{"refresh.time"}

	Requestcachememorysize = CatNodeColumn{"request_cache.memory_size"}

	Requestcacheevictions = CatNodeColumn{"request_cache.evictions"}

	Requestcachehitcount = CatNodeColumn{"request_cache.hit_count"}

	Requestcachemisscount = CatNodeColumn{"request_cache.miss_count"}

	Scriptcompilations = CatNodeColumn{"script.compilations"}

	Scriptcacheevictions = CatNodeColumn{"script.cache_evictions"}

	Searchfetchcurrent = CatNodeColumn{"search.fetch_current"}

	Searchfetchtime = CatNodeColumn{"search.fetch_time"}

	Searchfetchtotal = CatNodeColumn{"search.fetch_total"}

	Searchopencontexts = CatNodeColumn{"search.open_contexts"}

	Searchquerycurrent = CatNodeColumn{"search.query_current"}

	Searchquerytime = CatNodeColumn{"search.query_time"}

	Searchquerytotal = CatNodeColumn{"search.query_total"}

	Searchscrollcurrent = CatNodeColumn{"search.scroll_current"}

	Searchscrolltime = CatNodeColumn{"search.scroll_time"}

	Searchscrolltotal = CatNodeColumn{"search.scroll_total"}

	Segmentscount = CatNodeColumn{"segments.count"}

	Segmentsfixedbitsetmemory = CatNodeColumn{"segments.fixed_bitset_memory"}

	Segmentsindexwritermemory = CatNodeColumn{"segments.index_writer_memory"}

	Segmentsmemory = CatNodeColumn{"segments.memory"}

	Segmentsversionmapmemory = CatNodeColumn{"segments.version_map_memory"}

	Shardstatstotalcount = CatNodeColumn{"shard_stats.total_count"}

	Suggestcurrent = CatNodeColumn{"suggest.current"}

	Suggesttime = CatNodeColumn{"suggest.time"}

	Suggesttotal = CatNodeColumn{"suggest.total"}

	Uptime = CatNodeColumn{"uptime"}

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
