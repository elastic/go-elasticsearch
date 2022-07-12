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
// https://github.com/elastic/elasticsearch-specification/tree/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741

package types

// WarmerStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/_types/Stats.ts#L252-L257
type WarmerStats struct {
	Current           int64                   `json:"current"`
	Total             int64                   `json:"total"`
	TotalTime         *Duration               `json:"total_time,omitempty"`
	TotalTimeInMillis DurationValueUnitMillis `json:"total_time_in_millis"`
}

// WarmerStatsBuilder holds WarmerStats struct and provides a builder API.
type WarmerStatsBuilder struct {
	v *WarmerStats
}

// NewWarmerStats provides a builder for the WarmerStats struct.
func NewWarmerStatsBuilder() *WarmerStatsBuilder {
	r := WarmerStatsBuilder{
		&WarmerStats{},
	}

	return &r
}

// Build finalize the chain and returns the WarmerStats struct
func (rb *WarmerStatsBuilder) Build() WarmerStats {
	return *rb.v
}

func (rb *WarmerStatsBuilder) Current(current int64) *WarmerStatsBuilder {
	rb.v.Current = current
	return rb
}

func (rb *WarmerStatsBuilder) Total(total int64) *WarmerStatsBuilder {
	rb.v.Total = total
	return rb
}

func (rb *WarmerStatsBuilder) TotalTime(totaltime *DurationBuilder) *WarmerStatsBuilder {
	v := totaltime.Build()
	rb.v.TotalTime = &v
	return rb
}

func (rb *WarmerStatsBuilder) TotalTimeInMillis(totaltimeinmillis *DurationValueUnitMillisBuilder) *WarmerStatsBuilder {
	v := totaltimeinmillis.Build()
	rb.v.TotalTimeInMillis = v
	return rb
}
