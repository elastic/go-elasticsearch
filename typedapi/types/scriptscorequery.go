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

// ScriptScoreQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/specialized.ts#L168-L172
type ScriptScoreQuery struct {
	Boost      *float32        `json:"boost,omitempty"`
	MinScore   *float32        `json:"min_score,omitempty"`
	Query      *QueryContainer `json:"query,omitempty"`
	QueryName_ *string         `json:"_name,omitempty"`
	Script     Script          `json:"script"`
}

// ScriptScoreQueryBuilder holds ScriptScoreQuery struct and provides a builder API.
type ScriptScoreQueryBuilder struct {
	v *ScriptScoreQuery
}

// NewScriptScoreQuery provides a builder for the ScriptScoreQuery struct.
func NewScriptScoreQueryBuilder() *ScriptScoreQueryBuilder {
	r := ScriptScoreQueryBuilder{
		&ScriptScoreQuery{},
	}

	return &r
}

// Build finalize the chain and returns the ScriptScoreQuery struct
func (rb *ScriptScoreQueryBuilder) Build() ScriptScoreQuery {
	return *rb.v
}

func (rb *ScriptScoreQueryBuilder) Boost(boost float32) *ScriptScoreQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *ScriptScoreQueryBuilder) MinScore(minscore float32) *ScriptScoreQueryBuilder {
	rb.v.MinScore = &minscore
	return rb
}

func (rb *ScriptScoreQueryBuilder) Query(query *QueryContainerBuilder) *ScriptScoreQueryBuilder {
	v := query.Build()
	rb.v.Query = &v
	return rb
}

func (rb *ScriptScoreQueryBuilder) QueryName_(queryname_ string) *ScriptScoreQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

func (rb *ScriptScoreQueryBuilder) Script(script *ScriptBuilder) *ScriptScoreQueryBuilder {
	v := script.Build()
	rb.v.Script = v
	return rb
}
