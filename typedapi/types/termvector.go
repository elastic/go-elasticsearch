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

// TermVector type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/termvectors/types.ts#L23-L26
type TermVector struct {
	FieldStatistics FieldStatistics `json:"field_statistics"`
	Terms           map[string]Term `json:"terms"`
}

// TermVectorBuilder holds TermVector struct and provides a builder API.
type TermVectorBuilder struct {
	v *TermVector
}

// NewTermVector provides a builder for the TermVector struct.
func NewTermVectorBuilder() *TermVectorBuilder {
	r := TermVectorBuilder{
		&TermVector{
			Terms: make(map[string]Term, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the TermVector struct
func (rb *TermVectorBuilder) Build() TermVector {
	return *rb.v
}

func (rb *TermVectorBuilder) FieldStatistics(fieldstatistics *FieldStatisticsBuilder) *TermVectorBuilder {
	v := fieldstatistics.Build()
	rb.v.FieldStatistics = v
	return rb
}

func (rb *TermVectorBuilder) Terms(values map[string]*TermBuilder) *TermVectorBuilder {
	tmp := make(map[string]Term, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Terms = tmp
	return rb
}
