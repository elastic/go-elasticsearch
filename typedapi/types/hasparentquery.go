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

// HasParentQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/joining.ts#L53-L61
type HasParentQuery struct {
	Boost          *float32        `json:"boost,omitempty"`
	IgnoreUnmapped *bool           `json:"ignore_unmapped,omitempty"`
	InnerHits      *InnerHits      `json:"inner_hits,omitempty"`
	ParentType     RelationName    `json:"parent_type"`
	Query          *QueryContainer `json:"query,omitempty"`
	QueryName_     *string         `json:"_name,omitempty"`
	Score          *bool           `json:"score,omitempty"`
}

// HasParentQueryBuilder holds HasParentQuery struct and provides a builder API.
type HasParentQueryBuilder struct {
	v *HasParentQuery
}

// NewHasParentQuery provides a builder for the HasParentQuery struct.
func NewHasParentQueryBuilder() *HasParentQueryBuilder {
	r := HasParentQueryBuilder{
		&HasParentQuery{},
	}

	return &r
}

// Build finalize the chain and returns the HasParentQuery struct
func (rb *HasParentQueryBuilder) Build() HasParentQuery {
	return *rb.v
}

func (rb *HasParentQueryBuilder) Boost(boost float32) *HasParentQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *HasParentQueryBuilder) IgnoreUnmapped(ignoreunmapped bool) *HasParentQueryBuilder {
	rb.v.IgnoreUnmapped = &ignoreunmapped
	return rb
}

func (rb *HasParentQueryBuilder) InnerHits(innerhits *InnerHitsBuilder) *HasParentQueryBuilder {
	v := innerhits.Build()
	rb.v.InnerHits = &v
	return rb
}

func (rb *HasParentQueryBuilder) ParentType(parenttype RelationName) *HasParentQueryBuilder {
	rb.v.ParentType = parenttype
	return rb
}

func (rb *HasParentQueryBuilder) Query(query *QueryContainerBuilder) *HasParentQueryBuilder {
	v := query.Build()
	rb.v.Query = &v
	return rb
}

func (rb *HasParentQueryBuilder) QueryName_(queryname_ string) *HasParentQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

func (rb *HasParentQueryBuilder) Score(score bool) *HasParentQueryBuilder {
	rb.v.Score = &score
	return rb
}
