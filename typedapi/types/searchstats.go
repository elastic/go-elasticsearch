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

// SearchStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/_types/Stats.ts#L185-L204
type SearchStats struct {
	FetchCurrent        int64                  `json:"fetch_current"`
	FetchTime           *Duration              `json:"fetch_time,omitempty"`
	FetchTimeInMillis   int64                  `json:"fetch_time_in_millis"`
	FetchTotal          int64                  `json:"fetch_total"`
	Groups              map[string]SearchStats `json:"groups,omitempty"`
	OpenContexts        *int64                 `json:"open_contexts,omitempty"`
	QueryCurrent        int64                  `json:"query_current"`
	QueryTime           *Duration              `json:"query_time,omitempty"`
	QueryTimeInMillis   int64                  `json:"query_time_in_millis"`
	QueryTotal          int64                  `json:"query_total"`
	ScrollCurrent       int64                  `json:"scroll_current"`
	ScrollTime          *Duration              `json:"scroll_time,omitempty"`
	ScrollTimeInMillis  int64                  `json:"scroll_time_in_millis"`
	ScrollTotal         int64                  `json:"scroll_total"`
	SuggestCurrent      int64                  `json:"suggest_current"`
	SuggestTime         *Duration              `json:"suggest_time,omitempty"`
	SuggestTimeInMillis int64                  `json:"suggest_time_in_millis"`
	SuggestTotal        int64                  `json:"suggest_total"`
}

// NewSearchStats returns a SearchStats.
func NewSearchStats() *SearchStats {
	r := &SearchStats{
		Groups: make(map[string]SearchStats, 0),
	}

	return r
}
