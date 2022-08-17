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

// ShardsStatsSummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/snapshot/_types/SnapshotShardsStatus.ts#L29-L35
type ShardsStatsSummary struct {
	Incremental       ShardsStatsSummaryItem  `json:"incremental"`
	StartTimeInMillis EpochTimeUnitMillis     `json:"start_time_in_millis"`
	Time              *Duration               `json:"time,omitempty"`
	TimeInMillis      DurationValueUnitMillis `json:"time_in_millis"`
	Total             ShardsStatsSummaryItem  `json:"total"`
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

func (rb *ShardsStatsSummaryBuilder) StartTimeInMillis(starttimeinmillis *EpochTimeUnitMillisBuilder) *ShardsStatsSummaryBuilder {
	v := starttimeinmillis.Build()
	rb.v.StartTimeInMillis = v
	return rb
}

func (rb *ShardsStatsSummaryBuilder) Time(time *DurationBuilder) *ShardsStatsSummaryBuilder {
	v := time.Build()
	rb.v.Time = &v
	return rb
}

func (rb *ShardsStatsSummaryBuilder) TimeInMillis(timeinmillis *DurationValueUnitMillisBuilder) *ShardsStatsSummaryBuilder {
	v := timeinmillis.Build()
	rb.v.TimeInMillis = v
	return rb
}

func (rb *ShardsStatsSummaryBuilder) Total(total *ShardsStatsSummaryItemBuilder) *ShardsStatsSummaryBuilder {
	v := total.Build()
	rb.v.Total = v
	return rb
}
