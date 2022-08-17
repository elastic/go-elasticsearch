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

// NestedQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/joining.ts#L63-L71
type NestedQuery struct {
	Boost          *float32                       `json:"boost,omitempty"`
	IgnoreUnmapped *bool                          `json:"ignore_unmapped,omitempty"`
	InnerHits      *InnerHits                     `json:"inner_hits,omitempty"`
	Path           Field                          `json:"path"`
	Query          *QueryContainer                `json:"query,omitempty"`
	QueryName_     *string                        `json:"_name,omitempty"`
	ScoreMode      *childscoremode.ChildScoreMode `json:"score_mode,omitempty"`
}

// NestedQueryBuilder holds NestedQuery struct and provides a builder API.
type NestedQueryBuilder struct {
	v *NestedQuery
}

// NewNestedQuery provides a builder for the NestedQuery struct.
func NewNestedQueryBuilder() *NestedQueryBuilder {
	r := NestedQueryBuilder{
		&NestedQuery{},
	}

	return &r
}

// Build finalize the chain and returns the NestedQuery struct
func (rb *NestedQueryBuilder) Build() NestedQuery {
	return *rb.v
}

func (rb *NestedQueryBuilder) Boost(boost float32) *NestedQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *NestedQueryBuilder) IgnoreUnmapped(ignoreunmapped bool) *NestedQueryBuilder {
	rb.v.IgnoreUnmapped = &ignoreunmapped
	return rb
}

func (rb *NestedQueryBuilder) InnerHits(innerhits *InnerHitsBuilder) *NestedQueryBuilder {
	v := innerhits.Build()
	rb.v.InnerHits = &v
	return rb
}

func (rb *NestedQueryBuilder) Path(path Field) *NestedQueryBuilder {
	rb.v.Path = path
	return rb
}

func (rb *NestedQueryBuilder) Query(query *QueryContainerBuilder) *NestedQueryBuilder {
	v := query.Build()
	rb.v.Query = &v
	return rb
}

func (rb *NestedQueryBuilder) QueryName_(queryname_ string) *NestedQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

func (rb *NestedQueryBuilder) ScoreMode(scoremode childscoremode.ChildScoreMode) *NestedQueryBuilder {
	rb.v.ScoreMode = &scoremode
	return rb
}
