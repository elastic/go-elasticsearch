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

// RefreshStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/Stats.ts#L168-L175
type RefreshStats struct {
	ExternalTotal             int64                   `json:"external_total"`
	ExternalTotalTimeInMillis DurationValueUnitMillis `json:"external_total_time_in_millis"`
	Listeners                 int64                   `json:"listeners"`
	Total                     int64                   `json:"total"`
	TotalTime                 *Duration               `json:"total_time,omitempty"`
	TotalTimeInMillis         DurationValueUnitMillis `json:"total_time_in_millis"`
}

// RefreshStatsBuilder holds RefreshStats struct and provides a builder API.
type RefreshStatsBuilder struct {
	v *RefreshStats
}

// NewRefreshStats provides a builder for the RefreshStats struct.
func NewRefreshStatsBuilder() *RefreshStatsBuilder {
	r := RefreshStatsBuilder{
		&RefreshStats{},
	}

	return &r
}

// Build finalize the chain and returns the RefreshStats struct
func (rb *RefreshStatsBuilder) Build() RefreshStats {
	return *rb.v
}

func (rb *RefreshStatsBuilder) ExternalTotal(externaltotal int64) *RefreshStatsBuilder {
	rb.v.ExternalTotal = externaltotal
	return rb
}

func (rb *RefreshStatsBuilder) ExternalTotalTimeInMillis(externaltotaltimeinmillis *DurationValueUnitMillisBuilder) *RefreshStatsBuilder {
	v := externaltotaltimeinmillis.Build()
	rb.v.ExternalTotalTimeInMillis = v
	return rb
}

func (rb *RefreshStatsBuilder) Listeners(listeners int64) *RefreshStatsBuilder {
	rb.v.Listeners = listeners
	return rb
}

func (rb *RefreshStatsBuilder) Total(total int64) *RefreshStatsBuilder {
	rb.v.Total = total
	return rb
}

func (rb *RefreshStatsBuilder) TotalTime(totaltime *DurationBuilder) *RefreshStatsBuilder {
	v := totaltime.Build()
	rb.v.TotalTime = &v
	return rb
}

func (rb *RefreshStatsBuilder) TotalTimeInMillis(totaltimeinmillis *DurationValueUnitMillisBuilder) *RefreshStatsBuilder {
	v := totaltimeinmillis.Build()
	rb.v.TotalTimeInMillis = v
	return rb
}
