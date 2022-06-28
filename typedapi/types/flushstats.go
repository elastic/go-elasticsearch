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

// FlushStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/_types/Stats.ts#L80-L85
type FlushStats struct {
	Periodic          int64   `json:"periodic"`
	Total             int64   `json:"total"`
	TotalTime         *string `json:"total_time,omitempty"`
	TotalTimeInMillis int64   `json:"total_time_in_millis"`
}

// FlushStatsBuilder holds FlushStats struct and provides a builder API.
type FlushStatsBuilder struct {
	v *FlushStats
}

// NewFlushStats provides a builder for the FlushStats struct.
func NewFlushStatsBuilder() *FlushStatsBuilder {
	r := FlushStatsBuilder{
		&FlushStats{},
	}

	return &r
}

// Build finalize the chain and returns the FlushStats struct
func (rb *FlushStatsBuilder) Build() FlushStats {
	return *rb.v
}

func (rb *FlushStatsBuilder) Periodic(periodic int64) *FlushStatsBuilder {
	rb.v.Periodic = periodic
	return rb
}

func (rb *FlushStatsBuilder) Total(total int64) *FlushStatsBuilder {
	rb.v.Total = total
	return rb
}

func (rb *FlushStatsBuilder) TotalTime(totaltime string) *FlushStatsBuilder {
	rb.v.TotalTime = &totaltime
	return rb
}

func (rb *FlushStatsBuilder) TotalTimeInMillis(totaltimeinmillis int64) *FlushStatsBuilder {
	rb.v.TotalTimeInMillis = totaltimeinmillis
	return rb
}
