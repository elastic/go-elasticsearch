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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _textExpansionQuery struct {
	k string
	v *types.TextExpansionQuery
}

// Uses a natural language processing model to convert the query text into a
// list of token-weight pairs which are then used in a query against a sparse
// vector or rank features field.
func NewTextExpansionQuery(key string) *_textExpansionQuery {
	return &_textExpansionQuery{
		k: key,
		v: types.NewTextExpansionQuery(),
	}
}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_textExpansionQuery) Boost(boost float32) *_textExpansionQuery {

	s.v.Boost = &boost

	return s
}

// The text expansion NLP model to use
func (s *_textExpansionQuery) ModelId(modelid string) *_textExpansionQuery {

	s.v.ModelId = modelid

	return s
}

// The query text
func (s *_textExpansionQuery) ModelText(modeltext string) *_textExpansionQuery {

	s.v.ModelText = modeltext

	return s
}

// Token pruning configurations
func (s *_textExpansionQuery) PruningConfig(pruningconfig types.TokenPruningConfigVariant) *_textExpansionQuery {

	s.v.PruningConfig = pruningconfig.TokenPruningConfigCaster()

	return s
}

func (s *_textExpansionQuery) QueryName_(queryname_ string) *_textExpansionQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_textExpansionQuery) QueryCaster() *types.Query {
	container := types.NewQuery()
	container.TextExpansion = map[string]types.TextExpansionQuery{
		s.k: *s.v,
	}
	return container
}

// NewSingleTextExpansionQuery should be used when you want to
// create a single key dictionary without specifying the key in the
// constructor. Usually key is already defined within the parent container.
func NewSingleTextExpansionQuery() *_textExpansionQuery {
	return &_textExpansionQuery{
		k: "",
		v: types.NewTextExpansionQuery(),
	}
}

func (s *_textExpansionQuery) TextExpansionQueryCaster() *types.TextExpansionQuery {
	return s.v.TextExpansionQueryCaster()
}
