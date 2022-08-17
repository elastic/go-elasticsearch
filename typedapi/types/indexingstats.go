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

// IndexingStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/Stats.ts#L101-L116
type IndexingStats struct {
	DeleteCurrent        int64                    `json:"delete_current"`
	DeleteTime           *Duration                `json:"delete_time,omitempty"`
	DeleteTimeInMillis   DurationValueUnitMillis  `json:"delete_time_in_millis"`
	DeleteTotal          int64                    `json:"delete_total"`
	IndexCurrent         int64                    `json:"index_current"`
	IndexFailed          int64                    `json:"index_failed"`
	IndexTime            *Duration                `json:"index_time,omitempty"`
	IndexTimeInMillis    DurationValueUnitMillis  `json:"index_time_in_millis"`
	IndexTotal           int64                    `json:"index_total"`
	IsThrottled          bool                     `json:"is_throttled"`
	NoopUpdateTotal      int64                    `json:"noop_update_total"`
	ThrottleTime         *Duration                `json:"throttle_time,omitempty"`
	ThrottleTimeInMillis DurationValueUnitMillis  `json:"throttle_time_in_millis"`
	Types                map[string]IndexingStats `json:"types,omitempty"`
}

// IndexingStatsBuilder holds IndexingStats struct and provides a builder API.
type IndexingStatsBuilder struct {
	v *IndexingStats
}

// NewIndexingStats provides a builder for the IndexingStats struct.
func NewIndexingStatsBuilder() *IndexingStatsBuilder {
	r := IndexingStatsBuilder{
		&IndexingStats{
			Types: make(map[string]IndexingStats, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the IndexingStats struct
func (rb *IndexingStatsBuilder) Build() IndexingStats {
	return *rb.v
}

func (rb *IndexingStatsBuilder) DeleteCurrent(deletecurrent int64) *IndexingStatsBuilder {
	rb.v.DeleteCurrent = deletecurrent
	return rb
}

func (rb *IndexingStatsBuilder) DeleteTime(deletetime *DurationBuilder) *IndexingStatsBuilder {
	v := deletetime.Build()
	rb.v.DeleteTime = &v
	return rb
}

func (rb *IndexingStatsBuilder) DeleteTimeInMillis(deletetimeinmillis *DurationValueUnitMillisBuilder) *IndexingStatsBuilder {
	v := deletetimeinmillis.Build()
	rb.v.DeleteTimeInMillis = v
	return rb
}

func (rb *IndexingStatsBuilder) DeleteTotal(deletetotal int64) *IndexingStatsBuilder {
	rb.v.DeleteTotal = deletetotal
	return rb
}

func (rb *IndexingStatsBuilder) IndexCurrent(indexcurrent int64) *IndexingStatsBuilder {
	rb.v.IndexCurrent = indexcurrent
	return rb
}

func (rb *IndexingStatsBuilder) IndexFailed(indexfailed int64) *IndexingStatsBuilder {
	rb.v.IndexFailed = indexfailed
	return rb
}

func (rb *IndexingStatsBuilder) IndexTime(indextime *DurationBuilder) *IndexingStatsBuilder {
	v := indextime.Build()
	rb.v.IndexTime = &v
	return rb
}

func (rb *IndexingStatsBuilder) IndexTimeInMillis(indextimeinmillis *DurationValueUnitMillisBuilder) *IndexingStatsBuilder {
	v := indextimeinmillis.Build()
	rb.v.IndexTimeInMillis = v
	return rb
}

func (rb *IndexingStatsBuilder) IndexTotal(indextotal int64) *IndexingStatsBuilder {
	rb.v.IndexTotal = indextotal
	return rb
}

func (rb *IndexingStatsBuilder) IsThrottled(isthrottled bool) *IndexingStatsBuilder {
	rb.v.IsThrottled = isthrottled
	return rb
}

func (rb *IndexingStatsBuilder) NoopUpdateTotal(noopupdatetotal int64) *IndexingStatsBuilder {
	rb.v.NoopUpdateTotal = noopupdatetotal
	return rb
}

func (rb *IndexingStatsBuilder) ThrottleTime(throttletime *DurationBuilder) *IndexingStatsBuilder {
	v := throttletime.Build()
	rb.v.ThrottleTime = &v
	return rb
}

func (rb *IndexingStatsBuilder) ThrottleTimeInMillis(throttletimeinmillis *DurationValueUnitMillisBuilder) *IndexingStatsBuilder {
	v := throttletimeinmillis.Build()
	rb.v.ThrottleTimeInMillis = v
	return rb
}

func (rb *IndexingStatsBuilder) Types(values map[string]*IndexingStatsBuilder) *IndexingStatsBuilder {
	tmp := make(map[string]IndexingStats, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Types = tmp
	return rb
}
