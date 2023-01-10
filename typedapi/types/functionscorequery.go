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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/functionboostmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/functionscoremode"
)

// FunctionScoreQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/_types/query_dsl/compound.ts#L52-L59
type FunctionScoreQuery struct {
	Boost      *float32                             `json:"boost,omitempty"`
	BoostMode  *functionboostmode.FunctionBoostMode `json:"boost_mode,omitempty"`
	Functions  []FunctionScore                      `json:"functions,omitempty"`
	MaxBoost   *float64                             `json:"max_boost,omitempty"`
	MinScore   *float64                             `json:"min_score,omitempty"`
	Query      *Query                               `json:"query,omitempty"`
	QueryName_ *string                              `json:"_name,omitempty"`
	ScoreMode  *functionscoremode.FunctionScoreMode `json:"score_mode,omitempty"`
}

// NewFunctionScoreQuery returns a FunctionScoreQuery.
func NewFunctionScoreQuery() *FunctionScoreQuery {
	r := &FunctionScoreQuery{}

	return r
}
