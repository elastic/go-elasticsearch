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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/scoremode"
)

// RescoreQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/search/_types/rescoring.ts#L28-L34
type RescoreQuery struct {
	Query              QueryContainer       `json:"rescore_query"`
	QueryWeight        *float64             `json:"query_weight,omitempty"`
	RescoreQueryWeight *float64             `json:"rescore_query_weight,omitempty"`
	ScoreMode          *scoremode.ScoreMode `json:"score_mode,omitempty"`
}

// RescoreQueryBuilder holds RescoreQuery struct and provides a builder API.
type RescoreQueryBuilder struct {
	v *RescoreQuery
}

// NewRescoreQuery provides a builder for the RescoreQuery struct.
func NewRescoreQueryBuilder() *RescoreQueryBuilder {
	r := RescoreQueryBuilder{
		&RescoreQuery{},
	}

	return &r
}

// Build finalize the chain and returns the RescoreQuery struct
func (rb *RescoreQueryBuilder) Build() RescoreQuery {
	return *rb.v
}

func (rb *RescoreQueryBuilder) Query(query *QueryContainerBuilder) *RescoreQueryBuilder {
	v := query.Build()
	rb.v.Query = v
	return rb
}

func (rb *RescoreQueryBuilder) QueryWeight(queryweight float64) *RescoreQueryBuilder {
	rb.v.QueryWeight = &queryweight
	return rb
}

func (rb *RescoreQueryBuilder) RescoreQueryWeight(rescorequeryweight float64) *RescoreQueryBuilder {
	rb.v.RescoreQueryWeight = &rescorequeryweight
	return rb
}

func (rb *RescoreQueryBuilder) ScoreMode(scoremode scoremode.ScoreMode) *RescoreQueryBuilder {
	rb.v.ScoreMode = &scoremode
	return rb
}
