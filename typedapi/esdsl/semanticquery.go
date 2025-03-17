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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _semanticQuery struct {
	v *types.SemanticQuery
}

// A semantic query to semantic_text field types
func NewSemanticQuery(field string, query string) *_semanticQuery {

	tmp := &_semanticQuery{v: types.NewSemanticQuery()}

	tmp.Field(field)

	tmp.Query(query)

	return tmp

}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_semanticQuery) Boost(boost float32) *_semanticQuery {

	s.v.Boost = &boost

	return s
}

// The field to query, which must be a semantic_text field type
func (s *_semanticQuery) Field(field string) *_semanticQuery {

	s.v.Field = field

	return s
}

// The query text
func (s *_semanticQuery) Query(query string) *_semanticQuery {

	s.v.Query = query

	return s
}

func (s *_semanticQuery) QueryName_(queryname_ string) *_semanticQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_semanticQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.Semantic = s.v

	return container
}

func (s *_semanticQuery) SemanticQueryCaster() *types.SemanticQuery {
	return s.v
}
