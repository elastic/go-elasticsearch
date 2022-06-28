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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/calendarinterval"
)

// DateHistogramAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/_types/aggregations/bucket.ts#L90-L106
type DateHistogramAggregation struct {
	CalendarInterval *calendarinterval.CalendarInterval `json:"calendar_interval,omitempty"`
	ExtendedBounds   *ExtendedBoundsFieldDateMath       `json:"extended_bounds,omitempty"`
	Field            *Field                             `json:"field,omitempty"`
	FixedInterval    *Time                              `json:"fixed_interval,omitempty"`
	Format           *string                            `json:"format,omitempty"`
	HardBounds       *ExtendedBoundsFieldDateMath       `json:"hard_bounds,omitempty"`
	Interval         *Time                              `json:"interval,omitempty"`
	Keyed            *bool                              `json:"keyed,omitempty"`
	Meta             *Metadata                          `json:"meta,omitempty"`
	MinDocCount      *int                               `json:"min_doc_count,omitempty"`
	Missing          *DateString                        `json:"missing,omitempty"`
	Name             *string                            `json:"name,omitempty"`
	Offset           *Time                              `json:"offset,omitempty"`
	Order            *HistogramOrder                    `json:"order,omitempty"`
	Params           map[string]interface{}             `json:"params,omitempty"`
	Script           *Script                            `json:"script,omitempty"`
	TimeZone         *string                            `json:"time_zone,omitempty"`
}

// DateHistogramAggregationBuilder holds DateHistogramAggregation struct and provides a builder API.
type DateHistogramAggregationBuilder struct {
	v *DateHistogramAggregation
}

// NewDateHistogramAggregation provides a builder for the DateHistogramAggregation struct.
func NewDateHistogramAggregationBuilder() *DateHistogramAggregationBuilder {
	r := DateHistogramAggregationBuilder{
		&DateHistogramAggregation{
			Params: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the DateHistogramAggregation struct
func (rb *DateHistogramAggregationBuilder) Build() DateHistogramAggregation {
	return *rb.v
}

func (rb *DateHistogramAggregationBuilder) CalendarInterval(calendarinterval calendarinterval.CalendarInterval) *DateHistogramAggregationBuilder {
	rb.v.CalendarInterval = &calendarinterval
	return rb
}

func (rb *DateHistogramAggregationBuilder) ExtendedBounds(extendedbounds *ExtendedBoundsFieldDateMathBuilder) *DateHistogramAggregationBuilder {
	v := extendedbounds.Build()
	rb.v.ExtendedBounds = &v
	return rb
}

func (rb *DateHistogramAggregationBuilder) Field(field Field) *DateHistogramAggregationBuilder {
	rb.v.Field = &field
	return rb
}

func (rb *DateHistogramAggregationBuilder) FixedInterval(fixedinterval *TimeBuilder) *DateHistogramAggregationBuilder {
	v := fixedinterval.Build()
	rb.v.FixedInterval = &v
	return rb
}

func (rb *DateHistogramAggregationBuilder) Format(format string) *DateHistogramAggregationBuilder {
	rb.v.Format = &format
	return rb
}

func (rb *DateHistogramAggregationBuilder) HardBounds(hardbounds *ExtendedBoundsFieldDateMathBuilder) *DateHistogramAggregationBuilder {
	v := hardbounds.Build()
	rb.v.HardBounds = &v
	return rb
}

func (rb *DateHistogramAggregationBuilder) Interval(interval *TimeBuilder) *DateHistogramAggregationBuilder {
	v := interval.Build()
	rb.v.Interval = &v
	return rb
}

func (rb *DateHistogramAggregationBuilder) Keyed(keyed bool) *DateHistogramAggregationBuilder {
	rb.v.Keyed = &keyed
	return rb
}

func (rb *DateHistogramAggregationBuilder) Meta(meta *MetadataBuilder) *DateHistogramAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *DateHistogramAggregationBuilder) MinDocCount(mindoccount int) *DateHistogramAggregationBuilder {
	rb.v.MinDocCount = &mindoccount
	return rb
}

func (rb *DateHistogramAggregationBuilder) Missing(missing DateString) *DateHistogramAggregationBuilder {
	rb.v.Missing = &missing
	return rb
}

func (rb *DateHistogramAggregationBuilder) Name(name string) *DateHistogramAggregationBuilder {
	rb.v.Name = &name
	return rb
}

func (rb *DateHistogramAggregationBuilder) Offset(offset *TimeBuilder) *DateHistogramAggregationBuilder {
	v := offset.Build()
	rb.v.Offset = &v
	return rb
}

func (rb *DateHistogramAggregationBuilder) Order(order *HistogramOrderBuilder) *DateHistogramAggregationBuilder {
	v := order.Build()
	rb.v.Order = &v
	return rb
}

func (rb *DateHistogramAggregationBuilder) Params(value map[string]interface{}) *DateHistogramAggregationBuilder {
	rb.v.Params = value
	return rb
}

func (rb *DateHistogramAggregationBuilder) Script(script *ScriptBuilder) *DateHistogramAggregationBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}

func (rb *DateHistogramAggregationBuilder) TimeZone(timezone string) *DateHistogramAggregationBuilder {
	rb.v.TimeZone = &timezone
	return rb
}
