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
// https://github.com/elastic/elasticsearch-specification/tree/555082f38110f65b60d470107d211fc354a5c55a


package types

// ShardsSegment type.
//
// https://github.com/elastic/elasticsearch-specification/blob/555082f38110f65b60d470107d211fc354a5c55a/specification/indices/segments/types.ts#L47-L52
type ShardsSegment struct {
	NumCommittedSegments int                 `json:"num_committed_segments"`
	NumSearchSegments    int                 `json:"num_search_segments"`
	Routing              ShardSegmentRouting `json:"routing"`
	Segments             map[string]Segment  `json:"segments"`
}

// NewShardsSegment returns a ShardsSegment.
func NewShardsSegment() *ShardsSegment {
	r := &ShardsSegment{
		Segments: make(map[string]Segment, 0),
	}

	return r
}
