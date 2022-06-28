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

// GetStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/_types/Stats.ts#L87-L98
type GetStats struct {
	Current             int64   `json:"current"`
	ExistsTime          *string `json:"exists_time,omitempty"`
	ExistsTimeInMillis  int64   `json:"exists_time_in_millis"`
	ExistsTotal         int64   `json:"exists_total"`
	MissingTime         *string `json:"missing_time,omitempty"`
	MissingTimeInMillis int64   `json:"missing_time_in_millis"`
	MissingTotal        int64   `json:"missing_total"`
	Time                *string `json:"time,omitempty"`
	TimeInMillis        int64   `json:"time_in_millis"`
	Total               int64   `json:"total"`
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

func (rb *GetStatsBuilder) ExistsTime(existstime string) *GetStatsBuilder {
	rb.v.ExistsTime = &existstime
	return rb
}

func (rb *GetStatsBuilder) ExistsTimeInMillis(existstimeinmillis int64) *GetStatsBuilder {
	rb.v.ExistsTimeInMillis = existstimeinmillis
	return rb
}

func (rb *GetStatsBuilder) ExistsTotal(existstotal int64) *GetStatsBuilder {
	rb.v.ExistsTotal = existstotal
	return rb
}

func (rb *GetStatsBuilder) MissingTime(missingtime string) *GetStatsBuilder {
	rb.v.MissingTime = &missingtime
	return rb
}

func (rb *GetStatsBuilder) MissingTimeInMillis(missingtimeinmillis int64) *GetStatsBuilder {
	rb.v.MissingTimeInMillis = missingtimeinmillis
	return rb
}

func (rb *GetStatsBuilder) MissingTotal(missingtotal int64) *GetStatsBuilder {
	rb.v.MissingTotal = missingtotal
	return rb
}

func (rb *GetStatsBuilder) Time(time string) *GetStatsBuilder {
	rb.v.Time = &time
	return rb
}

func (rb *GetStatsBuilder) TimeInMillis(timeinmillis int64) *GetStatsBuilder {
	rb.v.TimeInMillis = timeinmillis
	return rb
}

func (rb *GetStatsBuilder) Total(total int64) *GetStatsBuilder {
	rb.v.Total = total
	return rb
}
