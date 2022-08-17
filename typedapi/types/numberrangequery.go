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

// NumberRangeQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/term.ts#L83-L90
type NumberRangeQuery struct {
	Boost      *float32                     `json:"boost,omitempty"`
	From       float64                      `json:"from,omitempty"`
	Gt         *float64                     `json:"gt,omitempty"`
	Gte        *float64                     `json:"gte,omitempty"`
	Lt         *float64                     `json:"lt,omitempty"`
	Lte        *float64                     `json:"lte,omitempty"`
	QueryName_ *string                      `json:"_name,omitempty"`
	Relation   *rangerelation.RangeRelation `json:"relation,omitempty"`
	To         float64                      `json:"to,omitempty"`
}

// NumberRangeQueryBuilder holds NumberRangeQuery struct and provides a builder API.
type NumberRangeQueryBuilder struct {
	v *NumberRangeQuery
}

// NewNumberRangeQuery provides a builder for the NumberRangeQuery struct.
func NewNumberRangeQueryBuilder() *NumberRangeQueryBuilder {
	r := NumberRangeQueryBuilder{
		&NumberRangeQuery{},
	}

	return &r
}

// Build finalize the chain and returns the NumberRangeQuery struct
func (rb *NumberRangeQueryBuilder) Build() NumberRangeQuery {
	return *rb.v
}

func (rb *NumberRangeQueryBuilder) Boost(boost float32) *NumberRangeQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *NumberRangeQueryBuilder) From(from float64) *NumberRangeQueryBuilder {
	rb.v.From = from
	return rb
}

func (rb *NumberRangeQueryBuilder) Gt(gt float64) *NumberRangeQueryBuilder {
	rb.v.Gt = &gt
	return rb
}

func (rb *NumberRangeQueryBuilder) Gte(gte float64) *NumberRangeQueryBuilder {
	rb.v.Gte = &gte
	return rb
}

func (rb *NumberRangeQueryBuilder) Lt(lt float64) *NumberRangeQueryBuilder {
	rb.v.Lt = &lt
	return rb
}

func (rb *NumberRangeQueryBuilder) Lte(lte float64) *NumberRangeQueryBuilder {
	rb.v.Lte = &lte
	return rb
}

func (rb *NumberRangeQueryBuilder) QueryName_(queryname_ string) *NumberRangeQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

func (rb *NumberRangeQueryBuilder) Relation(relation rangerelation.RangeRelation) *NumberRangeQueryBuilder {
	rb.v.Relation = &relation
	return rb
}

func (rb *NumberRangeQueryBuilder) To(to float64) *NumberRangeQueryBuilder {
	rb.v.To = to
	return rb
}
