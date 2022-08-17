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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/childscoremode"
)

// HasChildQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/joining.ts#L41-L51
type HasChildQuery struct {
	Boost          *float32                       `json:"boost,omitempty"`
	IgnoreUnmapped *bool                          `json:"ignore_unmapped,omitempty"`
	InnerHits      *InnerHits                     `json:"inner_hits,omitempty"`
	MaxChildren    *int                           `json:"max_children,omitempty"`
	MinChildren    *int                           `json:"min_children,omitempty"`
	Query          *QueryContainer                `json:"query,omitempty"`
	QueryName_     *string                        `json:"_name,omitempty"`
	ScoreMode      *childscoremode.ChildScoreMode `json:"score_mode,omitempty"`
	Type           RelationName                   `json:"type"`
}

// HasChildQueryBuilder holds HasChildQuery struct and provides a builder API.
type HasChildQueryBuilder struct {
	v *HasChildQuery
}

// NewHasChildQuery provides a builder for the HasChildQuery struct.
func NewHasChildQueryBuilder() *HasChildQueryBuilder {
	r := HasChildQueryBuilder{
		&HasChildQuery{},
	}

	return &r
}

// Build finalize the chain and returns the HasChildQuery struct
func (rb *HasChildQueryBuilder) Build() HasChildQuery {
	return *rb.v
}

func (rb *HasChildQueryBuilder) Boost(boost float32) *HasChildQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *HasChildQueryBuilder) IgnoreUnmapped(ignoreunmapped bool) *HasChildQueryBuilder {
	rb.v.IgnoreUnmapped = &ignoreunmapped
	return rb
}

func (rb *HasChildQueryBuilder) InnerHits(innerhits *InnerHitsBuilder) *HasChildQueryBuilder {
	v := innerhits.Build()
	rb.v.InnerHits = &v
	return rb
}

func (rb *HasChildQueryBuilder) MaxChildren(maxchildren int) *HasChildQueryBuilder {
	rb.v.MaxChildren = &maxchildren
	return rb
}

func (rb *HasChildQueryBuilder) MinChildren(minchildren int) *HasChildQueryBuilder {
	rb.v.MinChildren = &minchildren
	return rb
}

func (rb *HasChildQueryBuilder) Query(query *QueryContainerBuilder) *HasChildQueryBuilder {
	v := query.Build()
	rb.v.Query = &v
	return rb
}

func (rb *HasChildQueryBuilder) QueryName_(queryname_ string) *HasChildQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

func (rb *HasChildQueryBuilder) ScoreMode(scoremode childscoremode.ChildScoreMode) *HasChildQueryBuilder {
	rb.v.ScoreMode = &scoremode
	return rb
}

func (rb *HasChildQueryBuilder) Type_(type_ RelationName) *HasChildQueryBuilder {
	rb.v.Type = type_
	return rb
}
