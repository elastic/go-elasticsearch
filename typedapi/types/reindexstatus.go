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

// ReindexStatus type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/_global/reindex_rethrottle/types.ts#L37-L51
type ReindexStatus struct {
	Batches              int64     `json:"batches"`
	Created              int64     `json:"created"`
	Deleted              int64     `json:"deleted"`
	Noops                int64     `json:"noops"`
	RequestsPerSecond    float32   `json:"requests_per_second"`
	Retries              Retries   `json:"retries"`
	Throttled            *Duration `json:"throttled,omitempty"`
	ThrottledMillis      int64     `json:"throttled_millis"`
	ThrottledUntil       *Duration `json:"throttled_until,omitempty"`
	ThrottledUntilMillis int64     `json:"throttled_until_millis"`
	Total                int64     `json:"total"`
	Updated              int64     `json:"updated"`
	VersionConflicts     int64     `json:"version_conflicts"`
}

// NewReindexStatus returns a ReindexStatus.
func NewReindexStatus() *ReindexStatus {
	r := &ReindexStatus{}

	return r
}
