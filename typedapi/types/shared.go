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

// Shared type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/searchable_snapshots/cache_stats/Response.ts#L34-L43
type Shared struct {
	BytesReadInBytes    ByteSize `json:"bytes_read_in_bytes"`
	BytesWrittenInBytes ByteSize `json:"bytes_written_in_bytes"`
	Evictions           int64    `json:"evictions"`
	NumRegions          int      `json:"num_regions"`
	Reads               int64    `json:"reads"`
	RegionSizeInBytes   ByteSize `json:"region_size_in_bytes"`
	SizeInBytes         ByteSize `json:"size_in_bytes"`
	Writes              int64    `json:"writes"`
}

// NewShared returns a Shared.
func NewShared() *Shared {
	r := &Shared{}

	return r
}
