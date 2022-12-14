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
// https://github.com/elastic/elasticsearch-specification/tree/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33


package types

// Status type.
//
// https://github.com/elastic/elasticsearch-specification/blob/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33/specification/snapshot/_types/SnapshotStatus.ts#L26-L35
type Status struct {
	IncludeGlobalState bool                          `json:"include_global_state"`
	Indices            map[string]SnapshotIndexStats `json:"indices"`
	Repository         string                        `json:"repository"`
	ShardsStats        SnapshotShardsStats           `json:"shards_stats"`
	Snapshot           string                        `json:"snapshot"`
	State              string                        `json:"state"`
	Stats              SnapshotStats                 `json:"stats"`
	Uuid               string                        `json:"uuid"`
}

// NewStatus returns a Status.
func NewStatus() *Status {
	r := &Status{
		Indices: make(map[string]SnapshotIndexStats, 0),
	}

	return r
}
