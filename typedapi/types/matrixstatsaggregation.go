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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortmode"
)

// MatrixStatsAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/matrix.ts#L31-L33
type MatrixStatsAggregation struct {
	Fields  *Fields            `json:"fields,omitempty"`
	Meta    *Metadata          `json:"meta,omitempty"`
	Missing map[Field]float64  `json:"missing,omitempty"`
	Mode    *sortmode.SortMode `json:"mode,omitempty"`
	Name    *string            `json:"name,omitempty"`
}

// MatrixStatsAggregationBuilder holds MatrixStatsAggregation struct and provides a builder API.
type MatrixStatsAggregationBuilder struct {
	v *MatrixStatsAggregation
}

// NewMatrixStatsAggregation provides a builder for the MatrixStatsAggregation struct.
func NewMatrixStatsAggregationBuilder() *MatrixStatsAggregationBuilder {
	r := MatrixStatsAggregationBuilder{
		&MatrixStatsAggregation{
			Missing: make(map[Field]float64, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the MatrixStatsAggregation struct
func (rb *MatrixStatsAggregationBuilder) Build() MatrixStatsAggregation {
	return *rb.v
}

func (rb *MatrixStatsAggregationBuilder) Fields(fields *FieldsBuilder) *MatrixStatsAggregationBuilder {
	v := fields.Build()
	rb.v.Fields = &v
	return rb
}

func (rb *MatrixStatsAggregationBuilder) Meta(meta *MetadataBuilder) *MatrixStatsAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *MatrixStatsAggregationBuilder) Missing(value map[Field]float64) *MatrixStatsAggregationBuilder {
	rb.v.Missing = value
	return rb
}

func (rb *MatrixStatsAggregationBuilder) Mode(mode sortmode.SortMode) *MatrixStatsAggregationBuilder {
	rb.v.Mode = &mode
	return rb
}

func (rb *MatrixStatsAggregationBuilder) Name(name string) *MatrixStatsAggregationBuilder {
	rb.v.Name = &name
	return rb
}
