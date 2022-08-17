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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/functionboostmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/functionscoremode"
)

// FunctionScoreQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/compound.ts#L52-L59
type FunctionScoreQuery struct {
	Boost      *float32                             `json:"boost,omitempty"`
	BoostMode  *functionboostmode.FunctionBoostMode `json:"boost_mode,omitempty"`
	Functions  []FunctionScoreContainer             `json:"functions,omitempty"`
	MaxBoost   *float64                             `json:"max_boost,omitempty"`
	MinScore   *float64                             `json:"min_score,omitempty"`
	Query      *QueryContainer                      `json:"query,omitempty"`
	QueryName_ *string                              `json:"_name,omitempty"`
	ScoreMode  *functionscoremode.FunctionScoreMode `json:"score_mode,omitempty"`
}

// FunctionScoreQueryBuilder holds FunctionScoreQuery struct and provides a builder API.
type FunctionScoreQueryBuilder struct {
	v *FunctionScoreQuery
}

// NewFunctionScoreQuery provides a builder for the FunctionScoreQuery struct.
func NewFunctionScoreQueryBuilder() *FunctionScoreQueryBuilder {
	r := FunctionScoreQueryBuilder{
		&FunctionScoreQuery{},
	}

	return &r
}

// Build finalize the chain and returns the FunctionScoreQuery struct
func (rb *FunctionScoreQueryBuilder) Build() FunctionScoreQuery {
	return *rb.v
}

func (rb *FunctionScoreQueryBuilder) Boost(boost float32) *FunctionScoreQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *FunctionScoreQueryBuilder) BoostMode(boostmode functionboostmode.FunctionBoostMode) *FunctionScoreQueryBuilder {
	rb.v.BoostMode = &boostmode
	return rb
}

func (rb *FunctionScoreQueryBuilder) Functions(functions []FunctionScoreContainerBuilder) *FunctionScoreQueryBuilder {
	tmp := make([]FunctionScoreContainer, len(functions))
	for _, value := range functions {
		tmp = append(tmp, value.Build())
	}
	rb.v.Functions = tmp
	return rb
}

func (rb *FunctionScoreQueryBuilder) MaxBoost(maxboost float64) *FunctionScoreQueryBuilder {
	rb.v.MaxBoost = &maxboost
	return rb
}

func (rb *FunctionScoreQueryBuilder) MinScore(minscore float64) *FunctionScoreQueryBuilder {
	rb.v.MinScore = &minscore
	return rb
}

func (rb *FunctionScoreQueryBuilder) Query(query *QueryContainerBuilder) *FunctionScoreQueryBuilder {
	v := query.Build()
	rb.v.Query = &v
	return rb
}

func (rb *FunctionScoreQueryBuilder) QueryName_(queryname_ string) *FunctionScoreQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

func (rb *FunctionScoreQueryBuilder) ScoreMode(scoremode functionscoremode.FunctionScoreMode) *FunctionScoreQueryBuilder {
	rb.v.ScoreMode = &scoremode
	return rb
}
