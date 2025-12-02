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

// Package commonstatsflag
package commonstatsflag

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/aa1459fbdcaf57c653729142b3b6e9982373bb1c/specification/_types/common.ts#L392-L415
type CommonStatsFlag struct {
	Name string
}

var (
	All = CommonStatsFlag{"_all"}

	Store = CommonStatsFlag{"store"}

	Indexing = CommonStatsFlag{"indexing"}

	Get = CommonStatsFlag{"get"}

	Search = CommonStatsFlag{"search"}

	Merge = CommonStatsFlag{"merge"}

	Flush = CommonStatsFlag{"flush"}

	Refresh = CommonStatsFlag{"refresh"}

	Querycache = CommonStatsFlag{"query_cache"}

	Fielddata = CommonStatsFlag{"fielddata"}

	Docs = CommonStatsFlag{"docs"}

	Warmer = CommonStatsFlag{"warmer"}

	Completion = CommonStatsFlag{"completion"}

	Segments = CommonStatsFlag{"segments"}

	Translog = CommonStatsFlag{"translog"}

	Requestcache = CommonStatsFlag{"request_cache"}

	Recovery = CommonStatsFlag{"recovery"}

	Bulk = CommonStatsFlag{"bulk"}

	Shardstats = CommonStatsFlag{"shard_stats"}

	Mappings = CommonStatsFlag{"mappings"}

	Densevector = CommonStatsFlag{"dense_vector"}

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
