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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/ttesttype"
)

// TTestAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/metric.ts#L153-L157
type TTestAggregation struct {
	A    *TestPopulation      `json:"a,omitempty"`
	B    *TestPopulation      `json:"b,omitempty"`
	Meta *Metadata            `json:"meta,omitempty"`
	Name *string              `json:"name,omitempty"`
	Type *ttesttype.TTestType `json:"type,omitempty"`
}

// TTestAggregationBuilder holds TTestAggregation struct and provides a builder API.
type TTestAggregationBuilder struct {
	v *TTestAggregation
}

// NewTTestAggregation provides a builder for the TTestAggregation struct.
func NewTTestAggregationBuilder() *TTestAggregationBuilder {
	r := TTestAggregationBuilder{
		&TTestAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the TTestAggregation struct
func (rb *TTestAggregationBuilder) Build() TTestAggregation {
	return *rb.v
}

func (rb *TTestAggregationBuilder) A(a *TestPopulationBuilder) *TTestAggregationBuilder {
	v := a.Build()
	rb.v.A = &v
	return rb
}

func (rb *TTestAggregationBuilder) B(b *TestPopulationBuilder) *TTestAggregationBuilder {
	v := b.Build()
	rb.v.B = &v
	return rb
}

func (rb *TTestAggregationBuilder) Meta(meta *MetadataBuilder) *TTestAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *TTestAggregationBuilder) Name(name string) *TTestAggregationBuilder {
	rb.v.Name = &name
	return rb
}

func (rb *TTestAggregationBuilder) Type_(type_ ttesttype.TTestType) *TTestAggregationBuilder {
	rb.v.Type = &type_
	return rb
}
