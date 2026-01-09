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
// https://github.com/elastic/elasticsearch-specification/tree/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27

// Package nodestatsmetric
package nodestatsmetric

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27/specification/nodes/stats/NodesStatsRequest.ts#L112-L133
type NodeStatsMetric struct {
	Name string
}

var (
	All = NodeStatsMetric{"_all"}

	None = NodeStatsMetric{"_none"}

	Indices = NodeStatsMetric{"indices"}

	Os = NodeStatsMetric{"os"}

	Process = NodeStatsMetric{"process"}

	Jvm = NodeStatsMetric{"jvm"}

	Threadpool = NodeStatsMetric{"thread_pool"}

	Fs = NodeStatsMetric{"fs"}

	Transport = NodeStatsMetric{"transport"}

	Http = NodeStatsMetric{"http"}

	Breaker = NodeStatsMetric{"breaker"}

	Script = NodeStatsMetric{"script"}

	Discovery = NodeStatsMetric{"discovery"}

	Ingest = NodeStatsMetric{"ingest"}

	Adaptiveselection = NodeStatsMetric{"adaptive_selection"}

	Scriptcache = NodeStatsMetric{"script_cache"}

	Indexingpressure = NodeStatsMetric{"indexing_pressure"}

	Repositories = NodeStatsMetric{"repositories"}

	Allocations = NodeStatsMetric{"allocations"}
)

func (n NodeStatsMetric) MarshalText() (text []byte, err error) {
	return []byte(n.String()), nil
}

func (n *NodeStatsMetric) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "_all":
		*n = All
	case "_none":
		*n = None
	case "indices":
		*n = Indices
	case "os":
		*n = Os
	case "process":
		*n = Process
	case "jvm":
		*n = Jvm
	case "thread_pool":
		*n = Threadpool
	case "fs":
		*n = Fs
	case "transport":
		*n = Transport
	case "http":
		*n = Http
	case "breaker":
		*n = Breaker
	case "script":
		*n = Script
	case "discovery":
		*n = Discovery
	case "ingest":
		*n = Ingest
	case "adaptive_selection":
		*n = Adaptiveselection
	case "script_cache":
		*n = Scriptcache
	case "indexing_pressure":
		*n = Indexingpressure
	case "repositories":
		*n = Repositories
	case "allocations":
		*n = Allocations
	default:
		*n = NodeStatsMetric{string(text)}
	}

	return nil
}

func (n NodeStatsMetric) String() string {
	return n.Name
}
