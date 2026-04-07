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
// https://github.com/elastic/elasticsearch-specification/tree/836fca874204ca4173ae5c36fb6b5107d28d2fc0

// Package commonstatsflag
package commonstatsflag

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/836fca874204ca4173ae5c36fb6b5107d28d2fc0/specification/_types/common.ts#L418-L472
type CommonStatsFlag struct {
	Name string
}

var (

	// All Return all statistics.
	All = CommonStatsFlag{"_all"}

	// Store Size of the index in byte units.
	Store = CommonStatsFlag{"store"}

	// Indexing Indexing statistics.
	Indexing = CommonStatsFlag{"indexing"}

	// Get Get statistics, including missing stats.
	Get = CommonStatsFlag{"get"}

	// Search Search statistics including suggest statistics. You can include statistics
	// for custom groups by adding an extra `groups` parameter (search operations
	// can be associated with one or more groups). The `groups` parameter accepts a
	// comma-separated list of group names. Use `_all` to return statistics for all
	// groups.
	Search = CommonStatsFlag{"search"}

	// Merge Merge statistics.
	Merge = CommonStatsFlag{"merge"}

	// Flush Flush statistics.
	Flush = CommonStatsFlag{"flush"}

	// Refresh Refresh statistics.
	Refresh = CommonStatsFlag{"refresh"}

	// Querycache Query cache statistics.
	Querycache = CommonStatsFlag{"query_cache"}

	// Fielddata Fielddata statistics.
	Fielddata = CommonStatsFlag{"fielddata"}

	// Docs Number of documents and deleted docs not yet merged out. Index refreshes can
	// affect this statistic.
	Docs = CommonStatsFlag{"docs"}

	// Warmer Index warming statistics.
	Warmer = CommonStatsFlag{"warmer"}

	// Completion Completion suggester statistics.
	Completion = CommonStatsFlag{"completion"}

	// Segments Memory use of all open segments. If the `include_segment_file_sizes`
	// parameter is `true`, this metric includes the aggregated disk usage of each
	// Lucene index file.
	Segments = CommonStatsFlag{"segments"}

	// Translog Translog statistics.
	Translog = CommonStatsFlag{"translog"}

	// Requestcache Shard request cache statistics.
	Requestcache = CommonStatsFlag{"request_cache"}

	// Recovery Recovery statistics.
	Recovery = CommonStatsFlag{"recovery"}

	// Bulk Bulk operations statistics.
	Bulk = CommonStatsFlag{"bulk"}

	// Shardstats Shard statistics, including the total number of shards.
	Shardstats = CommonStatsFlag{"shard_stats"}

	// Mappings Mapping statistics, including the total count and estimated overhead.
	Mappings = CommonStatsFlag{"mappings"}

	// Densevector Total number of dense vectors indexed. Index refreshes can affect this
	// statistic.
	Densevector = CommonStatsFlag{"dense_vector"}

	// Sparsevector Total number of sparse vectors indexed. Index refreshes can affect this
	// statistic.
	Sparsevector = CommonStatsFlag{"sparse_vector"}
)

func (c CommonStatsFlag) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CommonStatsFlag) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "_all":
		*c = All
	case "store":
		*c = Store
	case "indexing":
		*c = Indexing
	case "get":
		*c = Get
	case "search":
		*c = Search
	case "merge":
		*c = Merge
	case "flush":
		*c = Flush
	case "refresh":
		*c = Refresh
	case "query_cache":
		*c = Querycache
	case "fielddata":
		*c = Fielddata
	case "docs":
		*c = Docs
	case "warmer":
		*c = Warmer
	case "completion":
		*c = Completion
	case "segments":
		*c = Segments
	case "translog":
		*c = Translog
	case "request_cache":
		*c = Requestcache
	case "recovery":
		*c = Recovery
	case "bulk":
		*c = Bulk
	case "shard_stats":
		*c = Shardstats
	case "mappings":
		*c = Mappings
	case "dense_vector":
		*c = Densevector
	case "sparse_vector":
		*c = Sparsevector
	default:
		*c = CommonStatsFlag{string(text)}
	}

	return nil
}

func (c CommonStatsFlag) String() string {
	return c.Name
}
