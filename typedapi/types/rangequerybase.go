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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/rangerelation"
)

// RangeQueryBase type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/term.ts#L68-L70
type RangeQueryBase struct {
	Boost      *float32                     `json:"boost,omitempty"`
	QueryName_ *string                      `json:"_name,omitempty"`
	Relation   *rangerelation.RangeRelation `json:"relation,omitempty"`
}

// RangeQueryBaseBuilder holds RangeQueryBase struct and provides a builder API.
type RangeQueryBaseBuilder struct {
	v *RangeQueryBase
}

// NewRangeQueryBase provides a builder for the RangeQueryBase struct.
func NewRangeQueryBaseBuilder() *RangeQueryBaseBuilder {
	r := RangeQueryBaseBuilder{
		&RangeQueryBase{},
	}

	return &r
}

// Build finalize the chain and returns the RangeQueryBase struct
func (rb *RangeQueryBaseBuilder) Build() RangeQueryBase {
	return *rb.v
}

func (rb *RangeQueryBaseBuilder) Boost(boost float32) *RangeQueryBaseBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *RangeQueryBaseBuilder) QueryName_(queryname_ string) *RangeQueryBaseBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

func (rb *RangeQueryBaseBuilder) Relation(relation rangerelation.RangeRelation) *RangeQueryBaseBuilder {
	rb.v.Relation = &relation
	return rb
}
