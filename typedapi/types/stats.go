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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

import "github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/noderole"

// Stats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/nodes/_types/Stats.ts#L30-L53
type Stats struct {
	AdaptiveSelection map[string]AdaptiveSelection `json:"adaptive_selection,omitempty"`
	Attributes        map[string]string            `json:"attributes,omitempty"`
	Breakers          map[string]Breaker           `json:"breakers,omitempty"`
	Discovery         *Discovery                   `json:"discovery,omitempty"`
	Fs                *FileSystem                  `json:"fs,omitempty"`
	Host              *string                      `json:"host,omitempty"`
	Http              *Http                        `json:"http,omitempty"`
	IndexingPressure  *NodesIndexingPressure       `json:"indexing_pressure,omitempty"`
	Indices           *IndicesShardStats           `json:"indices,omitempty"`
	Ingest            *NodesIngest                 `json:"ingest,omitempty"`
	Ip                []string                     `json:"ip,omitempty"`
	Jvm               *Jvm                         `json:"jvm,omitempty"`
	Name              *string                      `json:"name,omitempty"`
	Os                *OperatingSystem             `json:"os,omitempty"`
	Process           *Process                     `json:"process,omitempty"`
	Roles             []noderole.NodeRole          `json:"roles,omitempty"`
	Script            *Scripting                   `json:"script,omitempty"`
	ScriptCache       map[string][]ScriptCache     `json:"script_cache,omitempty"`
	ThreadPool        map[string]ThreadCount       `json:"thread_pool,omitempty"`
	Timestamp         *int64                       `json:"timestamp,omitempty"`
	Transport         *Transport                   `json:"transport,omitempty"`
	TransportAddress  *string                      `json:"transport_address,omitempty"`
}

// NewStats returns a Stats.
func NewStats() *Stats {
	r := &Stats{
		AdaptiveSelection: make(map[string]AdaptiveSelection, 0),
		Attributes:        make(map[string]string, 0),
		Breakers:          make(map[string]Breaker, 0),
		ScriptCache:       make(map[string][]ScriptCache, 0),
		ThreadPool:        make(map[string]ThreadCount, 0),
	}

	return r
}
