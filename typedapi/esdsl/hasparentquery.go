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

type _hasParentQuery struct {
	v *types.HasParentQuery
}

// Returns child documents whose joined parent document matches a provided
// query.
func NewHasParentQuery(query types.QueryVariant) *_hasParentQuery {

	tmp := &_hasParentQuery{v: types.NewHasParentQuery()}

	tmp.Query(query)

	return tmp

}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_hasParentQuery) Boost(boost float32) *_hasParentQuery {

	s.v.Boost = &boost

	return s
}

// Indicates whether to ignore an unmapped `parent_type` and not return any
// documents instead of an error.
// You can use this parameter to query multiple indices that may not contain the
// `parent_type`.
func (s *_hasParentQuery) IgnoreUnmapped(ignoreunmapped bool) *_hasParentQuery {

	s.v.IgnoreUnmapped = &ignoreunmapped

	return s
}

// If defined, each search hit will contain inner hits.
func (s *_hasParentQuery) InnerHits(innerhits types.InnerHitsVariant) *_hasParentQuery {

	s.v.InnerHits = innerhits.InnerHitsCaster()

	return s
}

// Name of the parent relationship mapped for the `join` field.
func (s *_hasParentQuery) ParentType(relationname string) *_hasParentQuery {

	s.v.ParentType = relationname

	return s
}

// Query you wish to run on parent documents of the `parent_type` field.
// If a parent document matches the search, the query returns its child
// documents.
func (s *_hasParentQuery) Query(query types.QueryVariant) *_hasParentQuery {

	s.v.Query = *query.QueryCaster()

	return s
}

func (s *_hasParentQuery) QueryName_(queryname_ string) *_hasParentQuery {

	s.v.QueryName_ = &queryname_

	return s
}

// Indicates whether the relevance score of a matching parent document is
// aggregated into its child documents.
func (s *_hasParentQuery) Score(score bool) *_hasParentQuery {

	s.v.Score = &score

	return s
}

func (s *_hasParentQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.HasParent = s.v

	return container
}

func (s *_hasParentQuery) HasParentQueryCaster() *types.HasParentQuery {
	return s.v
}
