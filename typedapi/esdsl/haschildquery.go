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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/childscoremode"
)

type _hasChildQuery struct {
	v *types.HasChildQuery
}

// Returns parent documents whose joined child documents match a provided query.
func NewHasChildQuery(query types.QueryVariant) *_hasChildQuery {

	tmp := &_hasChildQuery{v: types.NewHasChildQuery()}

	tmp.Query(query)

	return tmp

}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_hasChildQuery) Boost(boost float32) *_hasChildQuery {

	s.v.Boost = &boost

	return s
}

// Indicates whether to ignore an unmapped `type` and not return any documents
// instead of an error.
func (s *_hasChildQuery) IgnoreUnmapped(ignoreunmapped bool) *_hasChildQuery {

	s.v.IgnoreUnmapped = &ignoreunmapped

	return s
}

// If defined, each search hit will contain inner hits.
func (s *_hasChildQuery) InnerHits(innerhits types.InnerHitsVariant) *_hasChildQuery {

	s.v.InnerHits = innerhits.InnerHitsCaster()

	return s
}

// Maximum number of child documents that match the query allowed for a returned
// parent document.
// If the parent document exceeds this limit, it is excluded from the search
// results.
func (s *_hasChildQuery) MaxChildren(maxchildren int) *_hasChildQuery {

	s.v.MaxChildren = &maxchildren

	return s
}

// Minimum number of child documents that match the query required to match the
// query for a returned parent document.
// If the parent document does not meet this limit, it is excluded from the
// search results.
func (s *_hasChildQuery) MinChildren(minchildren int) *_hasChildQuery {

	s.v.MinChildren = &minchildren

	return s
}

// Query you wish to run on child documents of the `type` field.
// If a child document matches the search, the query returns the parent
// document.
func (s *_hasChildQuery) Query(query types.QueryVariant) *_hasChildQuery {

	s.v.Query = *query.QueryCaster()

	return s
}

func (s *_hasChildQuery) QueryName_(queryname_ string) *_hasChildQuery {

	s.v.QueryName_ = &queryname_

	return s
}

// Indicates how scores for matching child documents affect the root parent
// document’s relevance score.
func (s *_hasChildQuery) ScoreMode(scoremode childscoremode.ChildScoreMode) *_hasChildQuery {

	s.v.ScoreMode = &scoremode
	return s
}

// Name of the child relationship mapped for the `join` field.
func (s *_hasChildQuery) Type(relationname string) *_hasChildQuery {

	s.v.Type = relationname

	return s
}

func (s *_hasChildQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.HasChild = s.v

	return container
}

func (s *_hasChildQuery) HasChildQueryCaster() *types.HasChildQuery {
	return s.v
}
