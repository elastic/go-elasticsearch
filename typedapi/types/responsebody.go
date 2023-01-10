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

// ResponseBody type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/_global/search/SearchResponse.ts#L38-L54
type ResponseBody struct {
	Aggregations    map[string]Aggregate   `json:"aggregations,omitempty"`
	Clusters_       *ClusterStatistics     `json:"_clusters,omitempty"`
	Fields          map[string]interface{} `json:"fields,omitempty"`
	Hits            HitsMetadata           `json:"hits"`
	MaxScore        *float64               `json:"max_score,omitempty"`
	NumReducePhases *int64                 `json:"num_reduce_phases,omitempty"`
	PitId           *string                `json:"pit_id,omitempty"`
	Profile         *Profile               `json:"profile,omitempty"`
	ScrollId_       *string                `json:"_scroll_id,omitempty"`
	Shards_         ShardStatistics        `json:"_shards"`
	Suggest         map[string][]Suggest   `json:"suggest,omitempty"`
	TerminatedEarly *bool                  `json:"terminated_early,omitempty"`
	TimedOut        bool                   `json:"timed_out"`
	Took            int64                  `json:"took"`
}

// NewResponseBody returns a ResponseBody.
func NewResponseBody() *ResponseBody {
	r := &ResponseBody{
		Aggregations: make(map[string]Aggregate, 0),
		Fields:       make(map[string]interface{}, 0),
		Suggest:      make(map[string][]Suggest, 0),
	}

	return r
}
