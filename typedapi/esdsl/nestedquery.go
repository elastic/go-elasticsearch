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

type _nestedQuery struct {
	v *types.NestedQuery
}

// Wraps another query to search nested fields.
// If an object matches the search, the nested query returns the root parent
// document.
func NewNestedQuery(query types.QueryVariant) *_nestedQuery {

	tmp := &_nestedQuery{v: types.NewNestedQuery()}

	tmp.Query(query)

	return tmp

}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_nestedQuery) Boost(boost float32) *_nestedQuery {

	s.v.Boost = &boost

	return s
}

// Indicates whether to ignore an unmapped path and not return any documents
// instead of an error.
func (s *_nestedQuery) IgnoreUnmapped(ignoreunmapped bool) *_nestedQuery {

	s.v.IgnoreUnmapped = &ignoreunmapped

	return s
}

// If defined, each search hit will contain inner hits.
func (s *_nestedQuery) InnerHits(innerhits types.InnerHitsVariant) *_nestedQuery {

	s.v.InnerHits = innerhits.InnerHitsCaster()

	return s
}

// Path to the nested object you wish to search.
func (s *_nestedQuery) Path(field string) *_nestedQuery {

	s.v.Path = field

	return s
}

// Query you wish to run on nested objects in the path.
func (s *_nestedQuery) Query(query types.QueryVariant) *_nestedQuery {

	s.v.Query = *query.QueryCaster()

	return s
}

func (s *_nestedQuery) QueryName_(queryname_ string) *_nestedQuery {

	s.v.QueryName_ = &queryname_

	return s
}

// How scores for matching child objects affect the root parent document’s
// relevance score.
func (s *_nestedQuery) ScoreMode(scoremode childscoremode.ChildScoreMode) *_nestedQuery {

	s.v.ScoreMode = &scoremode
	return s
}

func (s *_nestedQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.Nested = s.v

	return container
}

func (s *_nestedQuery) NestedQueryCaster() *types.NestedQuery {
	return s.v
}
