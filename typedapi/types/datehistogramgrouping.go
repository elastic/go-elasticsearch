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

// DateHistogramGrouping type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/rollup/_types/Groupings.ts#L30-L38
type DateHistogramGrouping struct {
	CalendarInterval *Time   `json:"calendar_interval,omitempty"`
	Delay            *Time   `json:"delay,omitempty"`
	Field            Field   `json:"field"`
	FixedInterval    *Time   `json:"fixed_interval,omitempty"`
	Format           *string `json:"format,omitempty"`
	Interval         *Time   `json:"interval,omitempty"`
	TimeZone         *string `json:"time_zone,omitempty"`
}

// DateHistogramGroupingBuilder holds DateHistogramGrouping struct and provides a builder API.
type DateHistogramGroupingBuilder struct {
	v *DateHistogramGrouping
}

// NewDateHistogramGrouping provides a builder for the DateHistogramGrouping struct.
func NewDateHistogramGroupingBuilder() *DateHistogramGroupingBuilder {
	r := DateHistogramGroupingBuilder{
		&DateHistogramGrouping{},
	}

	return &r
}

// Build finalize the chain and returns the DateHistogramGrouping struct
func (rb *DateHistogramGroupingBuilder) Build() DateHistogramGrouping {
	return *rb.v
}

func (rb *DateHistogramGroupingBuilder) CalendarInterval(calendarinterval *TimeBuilder) *DateHistogramGroupingBuilder {
	v := calendarinterval.Build()
	rb.v.CalendarInterval = &v
	return rb
}

func (rb *DateHistogramGroupingBuilder) Delay(delay *TimeBuilder) *DateHistogramGroupingBuilder {
	v := delay.Build()
	rb.v.Delay = &v
	return rb
}

func (rb *DateHistogramGroupingBuilder) Field(field Field) *DateHistogramGroupingBuilder {
	rb.v.Field = field
	return rb
}

func (rb *DateHistogramGroupingBuilder) FixedInterval(fixedinterval *TimeBuilder) *DateHistogramGroupingBuilder {
	v := fixedinterval.Build()
	rb.v.FixedInterval = &v
	return rb
}

func (rb *DateHistogramGroupingBuilder) Format(format string) *DateHistogramGroupingBuilder {
	rb.v.Format = &format
	return rb
}

func (rb *DateHistogramGroupingBuilder) Interval(interval *TimeBuilder) *DateHistogramGroupingBuilder {
	v := interval.Build()
	rb.v.Interval = &v
	return rb
}

func (rb *DateHistogramGroupingBuilder) TimeZone(timezone string) *DateHistogramGroupingBuilder {
	rb.v.TimeZone = &timezone
	return rb
}
