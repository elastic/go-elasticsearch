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

// GetStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/Stats.ts#L88-L99
type GetStats struct {
	Current             int64                   `json:"current"`
	ExistsTime          *Duration               `json:"exists_time,omitempty"`
	ExistsTimeInMillis  DurationValueUnitMillis `json:"exists_time_in_millis"`
	ExistsTotal         int64                   `json:"exists_total"`
	MissingTime         *Duration               `json:"missing_time,omitempty"`
	MissingTimeInMillis DurationValueUnitMillis `json:"missing_time_in_millis"`
	MissingTotal        int64                   `json:"missing_total"`
	Time                *Duration               `json:"time,omitempty"`
	TimeInMillis        DurationValueUnitMillis `json:"time_in_millis"`
	Total               int64                   `json:"total"`
}

// GetStatsBuilder holds GetStats struct and provides a builder API.
type GetStatsBuilder struct {
	v *GetStats
}

// NewGetStats provides a builder for the GetStats struct.
func NewGetStatsBuilder() *GetStatsBuilder {
	r := GetStatsBuilder{
		&GetStats{},
	}

	return &r
}

// Build finalize the chain and returns the GetStats struct
func (rb *GetStatsBuilder) Build() GetStats {
	return *rb.v
}

func (rb *GetStatsBuilder) Current(current int64) *GetStatsBuilder {
	rb.v.Current = current
	return rb
}

func (rb *GetStatsBuilder) ExistsTime(existstime *DurationBuilder) *GetStatsBuilder {
	v := existstime.Build()
	rb.v.ExistsTime = &v
	return rb
}

func (rb *GetStatsBuilder) ExistsTimeInMillis(existstimeinmillis *DurationValueUnitMillisBuilder) *GetStatsBuilder {
	v := existstimeinmillis.Build()
	rb.v.ExistsTimeInMillis = v
	return rb
}

func (rb *GetStatsBuilder) ExistsTotal(existstotal int64) *GetStatsBuilder {
	rb.v.ExistsTotal = existstotal
	return rb
}

func (rb *GetStatsBuilder) MissingTime(missingtime *DurationBuilder) *GetStatsBuilder {
	v := missingtime.Build()
	rb.v.MissingTime = &v
	return rb
}

func (rb *GetStatsBuilder) MissingTimeInMillis(missingtimeinmillis *DurationValueUnitMillisBuilder) *GetStatsBuilder {
	v := missingtimeinmillis.Build()
	rb.v.MissingTimeInMillis = v
	return rb
}

func (rb *GetStatsBuilder) MissingTotal(missingtotal int64) *GetStatsBuilder {
	rb.v.MissingTotal = missingtotal
	return rb
}

func (rb *GetStatsBuilder) Time(time *DurationBuilder) *GetStatsBuilder {
	v := time.Build()
	rb.v.Time = &v
	return rb
}

func (rb *GetStatsBuilder) TimeInMillis(timeinmillis *DurationValueUnitMillisBuilder) *GetStatsBuilder {
	v := timeinmillis.Build()
	rb.v.TimeInMillis = v
	return rb
}

func (rb *GetStatsBuilder) Total(total int64) *GetStatsBuilder {
	rb.v.Total = total
	return rb
}
