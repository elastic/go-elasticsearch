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
// https://github.com/elastic/elasticsearch-specification/tree/6785a6caa1fa3ca5ab3308963d79dce923a3469f

// Package nodesinfometric
package nodesinfometric

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/6785a6caa1fa3ca5ab3308963d79dce923a3469f/specification/nodes/info/NodesInfoRequest.ts#L74-L89
type NodesInfoMetric struct {
	Name string
}

var (
	All = NodesInfoMetric{"_all"}

	None = NodesInfoMetric{"_none"}

	Settings = NodesInfoMetric{"settings"}

	Os = NodesInfoMetric{"os"}

	Process = NodesInfoMetric{"process"}

	Jvm = NodesInfoMetric{"jvm"}

	Threadpool = NodesInfoMetric{"thread_pool"}

	Transport = NodesInfoMetric{"transport"}

	Http = NodesInfoMetric{"http"}

	Remoteclusterserver = NodesInfoMetric{"remote_cluster_server"}

	Plugins = NodesInfoMetric{"plugins"}

	Ingest = NodesInfoMetric{"ingest"}

	Aggregations = NodesInfoMetric{"aggregations"}

	Indices = NodesInfoMetric{"indices"}
)

func (n NodesInfoMetric) MarshalText() (text []byte, err error) {
	return []byte(n.String()), nil
}

func (n *NodesInfoMetric) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "_all":
		*n = All
	case "_none":
		*n = None
	case "settings":
		*n = Settings
	case "os":
		*n = Os
	case "process":
		*n = Process
	case "jvm":
		*n = Jvm
	case "thread_pool":
		*n = Threadpool
	case "transport":
		*n = Transport
	case "http":
		*n = Http
	case "remote_cluster_server":
		*n = Remoteclusterserver
	case "plugins":
		*n = Plugins
	case "ingest":
		*n = Ingest
	case "aggregations":
		*n = Aggregations
	case "indices":
		*n = Indices
	default:
		*n = NodesInfoMetric{string(text)}
	}

	return nil
}

func (n NodesInfoMetric) String() string {
	return n.Name
}
