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

// MatrixAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/matrix.ts#L26-L29
type MatrixAggregation struct {
	Fields  *Fields           `json:"fields,omitempty"`
	Meta    *Metadata         `json:"meta,omitempty"`
	Missing map[Field]float64 `json:"missing,omitempty"`
	Name    *string           `json:"name,omitempty"`
}

// MatrixAggregationBuilder holds MatrixAggregation struct and provides a builder API.
type MatrixAggregationBuilder struct {
	v *MatrixAggregation
}

// NewMatrixAggregation provides a builder for the MatrixAggregation struct.
func NewMatrixAggregationBuilder() *MatrixAggregationBuilder {
	r := MatrixAggregationBuilder{
		&MatrixAggregation{
			Missing: make(map[Field]float64, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the MatrixAggregation struct
func (rb *MatrixAggregationBuilder) Build() MatrixAggregation {
	return *rb.v
}

func (rb *MatrixAggregationBuilder) Fields(fields *FieldsBuilder) *MatrixAggregationBuilder {
	v := fields.Build()
	rb.v.Fields = &v
	return rb
}

func (rb *MatrixAggregationBuilder) Meta(meta *MetadataBuilder) *MatrixAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *MatrixAggregationBuilder) Missing(value map[Field]float64) *MatrixAggregationBuilder {
	rb.v.Missing = value
	return rb
}

func (rb *MatrixAggregationBuilder) Name(name string) *MatrixAggregationBuilder {
	rb.v.Name = &name
	return rb
}
