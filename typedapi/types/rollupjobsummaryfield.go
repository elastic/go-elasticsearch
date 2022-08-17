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

// RollupJobSummaryField type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/rollup/get_rollup_index_caps/types.ts#L35-L39
type RollupJobSummaryField struct {
	Agg              string    `json:"agg"`
	CalendarInterval *Duration `json:"calendar_interval,omitempty"`
	TimeZone         *TimeZone `json:"time_zone,omitempty"`
}

// RollupJobSummaryFieldBuilder holds RollupJobSummaryField struct and provides a builder API.
type RollupJobSummaryFieldBuilder struct {
	v *RollupJobSummaryField
}

// NewRollupJobSummaryField provides a builder for the RollupJobSummaryField struct.
func NewRollupJobSummaryFieldBuilder() *RollupJobSummaryFieldBuilder {
	r := RollupJobSummaryFieldBuilder{
		&RollupJobSummaryField{},
	}

	return &r
}

// Build finalize the chain and returns the RollupJobSummaryField struct
func (rb *RollupJobSummaryFieldBuilder) Build() RollupJobSummaryField {
	return *rb.v
}

func (rb *RollupJobSummaryFieldBuilder) Agg(agg string) *RollupJobSummaryFieldBuilder {
	rb.v.Agg = agg
	return rb
}

func (rb *RollupJobSummaryFieldBuilder) CalendarInterval(calendarinterval *DurationBuilder) *RollupJobSummaryFieldBuilder {
	v := calendarinterval.Build()
	rb.v.CalendarInterval = &v
	return rb
}

func (rb *RollupJobSummaryFieldBuilder) TimeZone(timezone TimeZone) *RollupJobSummaryFieldBuilder {
	rb.v.TimeZone = &timezone
	return rb
}
