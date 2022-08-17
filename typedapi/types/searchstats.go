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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// SearchStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/Stats.ts#L185-L204
type SearchStats struct {
	FetchCurrent        int64                   `json:"fetch_current"`
	FetchTime           *Duration               `json:"fetch_time,omitempty"`
	FetchTimeInMillis   DurationValueUnitMillis `json:"fetch_time_in_millis"`
	FetchTotal          int64                   `json:"fetch_total"`
	Groups              map[string]SearchStats  `json:"groups,omitempty"`
	OpenContexts        *int64                  `json:"open_contexts,omitempty"`
	QueryCurrent        int64                   `json:"query_current"`
	QueryTime           *Duration               `json:"query_time,omitempty"`
	QueryTimeInMillis   DurationValueUnitMillis `json:"query_time_in_millis"`
	QueryTotal          int64                   `json:"query_total"`
	ScrollCurrent       int64                   `json:"scroll_current"`
	ScrollTime          *Duration               `json:"scroll_time,omitempty"`
	ScrollTimeInMillis  DurationValueUnitMillis `json:"scroll_time_in_millis"`
	ScrollTotal         int64                   `json:"scroll_total"`
	SuggestCurrent      int64                   `json:"suggest_current"`
	SuggestTime         *Duration               `json:"suggest_time,omitempty"`
	SuggestTimeInMillis DurationValueUnitMillis `json:"suggest_time_in_millis"`
	SuggestTotal        int64                   `json:"suggest_total"`
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

func (rb *SearchStatsBuilder) FetchTime(fetchtime *DurationBuilder) *SearchStatsBuilder {
	v := fetchtime.Build()
	rb.v.FetchTime = &v
	return rb
}

func (rb *SearchStatsBuilder) FetchTimeInMillis(fetchtimeinmillis *DurationValueUnitMillisBuilder) *SearchStatsBuilder {
	v := fetchtimeinmillis.Build()
	rb.v.FetchTimeInMillis = v
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

func (rb *SearchStatsBuilder) QueryTime(querytime *DurationBuilder) *SearchStatsBuilder {
	v := querytime.Build()
	rb.v.QueryTime = &v
	return rb
}

func (rb *SearchStatsBuilder) QueryTimeInMillis(querytimeinmillis *DurationValueUnitMillisBuilder) *SearchStatsBuilder {
	v := querytimeinmillis.Build()
	rb.v.QueryTimeInMillis = v
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

func (rb *SearchStatsBuilder) ScrollTime(scrolltime *DurationBuilder) *SearchStatsBuilder {
	v := scrolltime.Build()
	rb.v.ScrollTime = &v
	return rb
}

func (rb *SearchStatsBuilder) ScrollTimeInMillis(scrolltimeinmillis *DurationValueUnitMillisBuilder) *SearchStatsBuilder {
	v := scrolltimeinmillis.Build()
	rb.v.ScrollTimeInMillis = v
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

func (rb *SearchStatsBuilder) SuggestTime(suggesttime *DurationBuilder) *SearchStatsBuilder {
	v := suggesttime.Build()
	rb.v.SuggestTime = &v
	return rb
}

func (rb *SearchStatsBuilder) SuggestTimeInMillis(suggesttimeinmillis *DurationValueUnitMillisBuilder) *SearchStatsBuilder {
	v := suggesttimeinmillis.Build()
	rb.v.SuggestTimeInMillis = v
	return rb
}

func (rb *SearchStatsBuilder) SuggestTotal(suggesttotal int64) *SearchStatsBuilder {
	rb.v.SuggestTotal = suggesttotal
	return rb
}
