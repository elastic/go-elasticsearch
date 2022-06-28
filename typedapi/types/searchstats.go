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
// https://github.com/elastic/elasticsearch-specification/tree/135ae054e304239743b5777ad8d41cb2c9091d35


package types

// SearchStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/_types/Stats.ts#L184-L199
type SearchStats struct {
	FetchCurrent        int64                  `json:"fetch_current"`
	FetchTimeInMillis   int64                  `json:"fetch_time_in_millis"`
	FetchTotal          int64                  `json:"fetch_total"`
	Groups              map[string]SearchStats `json:"groups,omitempty"`
	OpenContexts        *int64                 `json:"open_contexts,omitempty"`
	QueryCurrent        int64                  `json:"query_current"`
	QueryTimeInMillis   int64                  `json:"query_time_in_millis"`
	QueryTotal          int64                  `json:"query_total"`
	ScrollCurrent       int64                  `json:"scroll_current"`
	ScrollTimeInMillis  int64                  `json:"scroll_time_in_millis"`
	ScrollTotal         int64                  `json:"scroll_total"`
	SuggestCurrent      int64                  `json:"suggest_current"`
	SuggestTimeInMillis int64                  `json:"suggest_time_in_millis"`
	SuggestTotal        int64                  `json:"suggest_total"`
}

// SearchStatsBuilder holds SearchStats struct and provides a builder API.
type SearchStatsBuilder struct {
	v *SearchStats
}

// NewSearchStats provides a builder for the SearchStats struct.
func NewSearchStatsBuilder() *SearchStatsBuilder {
	r := SearchStatsBuilder{
		&SearchStats{
			Groups: make(map[string]SearchStats, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the SearchStats struct
func (rb *SearchStatsBuilder) Build() SearchStats {
	return *rb.v
}

func (rb *SearchStatsBuilder) FetchCurrent(fetchcurrent int64) *SearchStatsBuilder {
	rb.v.FetchCurrent = fetchcurrent
	return rb
}

func (rb *SearchStatsBuilder) FetchTimeInMillis(fetchtimeinmillis int64) *SearchStatsBuilder {
	rb.v.FetchTimeInMillis = fetchtimeinmillis
	return rb
}

func (rb *SearchStatsBuilder) FetchTotal(fetchtotal int64) *SearchStatsBuilder {
	rb.v.FetchTotal = fetchtotal
	return rb
}

func (rb *SearchStatsBuilder) Groups(values map[string]*SearchStatsBuilder) *SearchStatsBuilder {
	tmp := make(map[string]SearchStats, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Groups = tmp
	return rb
}

func (rb *SearchStatsBuilder) OpenContexts(opencontexts int64) *SearchStatsBuilder {
	rb.v.OpenContexts = &opencontexts
	return rb
}

func (rb *SearchStatsBuilder) QueryCurrent(querycurrent int64) *SearchStatsBuilder {
	rb.v.QueryCurrent = querycurrent
	return rb
}

func (rb *SearchStatsBuilder) QueryTimeInMillis(querytimeinmillis int64) *SearchStatsBuilder {
	rb.v.QueryTimeInMillis = querytimeinmillis
	return rb
}

func (rb *SearchStatsBuilder) QueryTotal(querytotal int64) *SearchStatsBuilder {
	rb.v.QueryTotal = querytotal
	return rb
}

func (rb *SearchStatsBuilder) ScrollCurrent(scrollcurrent int64) *SearchStatsBuilder {
	rb.v.ScrollCurrent = scrollcurrent
	return rb
}

func (rb *SearchStatsBuilder) ScrollTimeInMillis(scrolltimeinmillis int64) *SearchStatsBuilder {
	rb.v.ScrollTimeInMillis = scrolltimeinmillis
	return rb
}

func (rb *SearchStatsBuilder) ScrollTotal(scrolltotal int64) *SearchStatsBuilder {
	rb.v.ScrollTotal = scrolltotal
	return rb
}

func (rb *SearchStatsBuilder) SuggestCurrent(suggestcurrent int64) *SearchStatsBuilder {
	rb.v.SuggestCurrent = suggestcurrent
	return rb
}

func (rb *SearchStatsBuilder) SuggestTimeInMillis(suggesttimeinmillis int64) *SearchStatsBuilder {
	rb.v.SuggestTimeInMillis = suggesttimeinmillis
	return rb
}

func (rb *SearchStatsBuilder) SuggestTotal(suggesttotal int64) *SearchStatsBuilder {
	rb.v.SuggestTotal = suggesttotal
	return rb
}
