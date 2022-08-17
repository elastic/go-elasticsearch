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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/calendarinterval"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/ratemode"
)

// RateAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/metric.ts#L127-L130
type RateAggregation struct {
	Field   *Field                             `json:"field,omitempty"`
	Format  *string                            `json:"format,omitempty"`
	Missing *Missing                           `json:"missing,omitempty"`
	Mode    *ratemode.RateMode                 `json:"mode,omitempty"`
	Script  *Script                            `json:"script,omitempty"`
	Unit    *calendarinterval.CalendarInterval `json:"unit,omitempty"`
}

// RateAggregationBuilder holds RateAggregation struct and provides a builder API.
type RateAggregationBuilder struct {
	v *RateAggregation
}

// NewRateAggregation provides a builder for the RateAggregation struct.
func NewRateAggregationBuilder() *RateAggregationBuilder {
	r := RateAggregationBuilder{
		&RateAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the RateAggregation struct
func (rb *RateAggregationBuilder) Build() RateAggregation {
	return *rb.v
}

func (rb *RateAggregationBuilder) Field(field Field) *RateAggregationBuilder {
	rb.v.Field = &field
	return rb
}

func (rb *RateAggregationBuilder) Format(format string) *RateAggregationBuilder {
	rb.v.Format = &format
	return rb
}

func (rb *RateAggregationBuilder) Missing(missing *MissingBuilder) *RateAggregationBuilder {
	v := missing.Build()
	rb.v.Missing = &v
	return rb
}

func (rb *RateAggregationBuilder) Mode(mode ratemode.RateMode) *RateAggregationBuilder {
	rb.v.Mode = &mode
	return rb
}

func (rb *RateAggregationBuilder) Script(script *ScriptBuilder) *RateAggregationBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}

func (rb *RateAggregationBuilder) Unit(unit calendarinterval.CalendarInterval) *RateAggregationBuilder {
	rb.v.Unit = &unit
	return rb
}
