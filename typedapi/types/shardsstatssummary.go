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

// ShardsStatsSummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/snapshot/_types/SnapshotShardsStatus.ts#L28-L33
type ShardsStatsSummary struct {
	Incremental       ShardsStatsSummaryItem `json:"incremental"`
	StartTimeInMillis int64                  `json:"start_time_in_millis"`
	TimeInMillis      int64                  `json:"time_in_millis"`
	Total             ShardsStatsSummaryItem `json:"total"`
}

// ShardsStatsSummaryBuilder holds ShardsStatsSummary struct and provides a builder API.
type ShardsStatsSummaryBuilder struct {
	v *ShardsStatsSummary
}

// NewShardsStatsSummary provides a builder for the ShardsStatsSummary struct.
func NewShardsStatsSummaryBuilder() *ShardsStatsSummaryBuilder {
	r := ShardsStatsSummaryBuilder{
		&ShardsStatsSummary{},
	}

	return &r
}

// Build finalize the chain and returns the ShardsStatsSummary struct
func (rb *ShardsStatsSummaryBuilder) Build() ShardsStatsSummary {
	return *rb.v
}

func (rb *ShardsStatsSummaryBuilder) Incremental(incremental *ShardsStatsSummaryItemBuilder) *ShardsStatsSummaryBuilder {
	v := incremental.Build()
	rb.v.Incremental = v
	return rb
}

func (rb *ShardsStatsSummaryBuilder) StartTimeInMillis(starttimeinmillis int64) *ShardsStatsSummaryBuilder {
	rb.v.StartTimeInMillis = starttimeinmillis
	return rb
}

func (rb *ShardsStatsSummaryBuilder) TimeInMillis(timeinmillis int64) *ShardsStatsSummaryBuilder {
	rb.v.TimeInMillis = timeinmillis
	return rb
}

func (rb *ShardsStatsSummaryBuilder) Total(total *ShardsStatsSummaryItemBuilder) *ShardsStatsSummaryBuilder {
	v := total.Build()
	rb.v.Total = v
	return rb
}
