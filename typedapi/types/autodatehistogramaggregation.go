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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/minimuminterval"
)

// AutoDateHistogramAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/bucket.ts#L52-L62
type AutoDateHistogramAggregation struct {
	Buckets         *int                             `json:"buckets,omitempty"`
	Field           *Field                           `json:"field,omitempty"`
	Format          *string                          `json:"format,omitempty"`
	Meta            *Metadata                        `json:"meta,omitempty"`
	MinimumInterval *minimuminterval.MinimumInterval `json:"minimum_interval,omitempty"`
	Missing         *DateTime                        `json:"missing,omitempty"`
	Name            *string                          `json:"name,omitempty"`
	Offset          *string                          `json:"offset,omitempty"`
	Params          map[string]interface{}           `json:"params,omitempty"`
	Script          *Script                          `json:"script,omitempty"`
	TimeZone        *TimeZone                        `json:"time_zone,omitempty"`
}

// AutoDateHistogramAggregationBuilder holds AutoDateHistogramAggregation struct and provides a builder API.
type AutoDateHistogramAggregationBuilder struct {
	v *AutoDateHistogramAggregation
}

// NewAutoDateHistogramAggregation provides a builder for the AutoDateHistogramAggregation struct.
func NewAutoDateHistogramAggregationBuilder() *AutoDateHistogramAggregationBuilder {
	r := AutoDateHistogramAggregationBuilder{
		&AutoDateHistogramAggregation{
			Params: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the AutoDateHistogramAggregation struct
func (rb *AutoDateHistogramAggregationBuilder) Build() AutoDateHistogramAggregation {
	return *rb.v
}

func (rb *AutoDateHistogramAggregationBuilder) Buckets(buckets int) *AutoDateHistogramAggregationBuilder {
	rb.v.Buckets = &buckets
	return rb
}

func (rb *AutoDateHistogramAggregationBuilder) Field(field Field) *AutoDateHistogramAggregationBuilder {
	rb.v.Field = &field
	return rb
}

func (rb *AutoDateHistogramAggregationBuilder) Format(format string) *AutoDateHistogramAggregationBuilder {
	rb.v.Format = &format
	return rb
}

func (rb *AutoDateHistogramAggregationBuilder) Meta(meta *MetadataBuilder) *AutoDateHistogramAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *AutoDateHistogramAggregationBuilder) MinimumInterval(minimuminterval minimuminterval.MinimumInterval) *AutoDateHistogramAggregationBuilder {
	rb.v.MinimumInterval = &minimuminterval
	return rb
}

func (rb *AutoDateHistogramAggregationBuilder) Missing(missing *DateTimeBuilder) *AutoDateHistogramAggregationBuilder {
	v := missing.Build()
	rb.v.Missing = &v
	return rb
}

func (rb *AutoDateHistogramAggregationBuilder) Name(name string) *AutoDateHistogramAggregationBuilder {
	rb.v.Name = &name
	return rb
}

func (rb *AutoDateHistogramAggregationBuilder) Offset(offset string) *AutoDateHistogramAggregationBuilder {
	rb.v.Offset = &offset
	return rb
}

func (rb *AutoDateHistogramAggregationBuilder) Params(value map[string]interface{}) *AutoDateHistogramAggregationBuilder {
	rb.v.Params = value
	return rb
}

func (rb *AutoDateHistogramAggregationBuilder) Script(script *ScriptBuilder) *AutoDateHistogramAggregationBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}

func (rb *AutoDateHistogramAggregationBuilder) TimeZone(timezone TimeZone) *AutoDateHistogramAggregationBuilder {
	rb.v.TimeZone = &timezone
	return rb
}
