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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _userQueryContainer struct {
	v *types.UserQueryContainer
}

func NewUserQueryContainer() *_userQueryContainer {
	return &_userQueryContainer{v: types.NewUserQueryContainer()}
}

// AdditionalUserQueryContainerProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_userQueryContainer) AdditionalUserQueryContainerProperty(key string, value json.RawMessage) *_userQueryContainer {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalUserQueryContainerProperty = tmp
	return s
}

// matches users matching boolean combinations of other queries.
func (s *_userQueryContainer) Bool(bool types.BoolQueryVariant) *_userQueryContainer {

	s.v.Bool = bool.BoolQueryCaster()

	return s
}

// Returns users that contain an indexed value for a field.
func (s *_userQueryContainer) Exists(exists types.ExistsQueryVariant) *_userQueryContainer {

	s.v.Exists = exists.ExistsQueryCaster()

	return s
}

// Returns users based on their IDs.
// This query uses the user document IDs stored in the `_id` field.
func (s *_userQueryContainer) Ids(ids types.IdsQueryVariant) *_userQueryContainer {

	s.v.Ids = ids.IdsQueryCaster()

	return s
}

// Returns users that match a provided text, number, date or boolean value.
// The provided text is analyzed before matching.
// Match is a single key dictionnary.
// It will replace the current value on each call.
func (s *_userQueryContainer) Match(key string, value types.MatchQueryVariant) *_userQueryContainer {

	tmp := make(map[string]types.MatchQuery)

	tmp[key] = *value.MatchQueryCaster()

	s.v.Match = tmp
	return s
}

// Matches all users, giving them all a `_score` of 1.0.
func (s *_userQueryContainer) MatchAll(matchall types.MatchAllQueryVariant) *_userQueryContainer {

	s.v.MatchAll = matchall.MatchAllQueryCaster()

	return s
}

// Returns users that contain a specific prefix in a provided field.
// Prefix is a single key dictionnary.
// It will replace the current value on each call.
func (s *_userQueryContainer) Prefix(key string, value types.PrefixQueryVariant) *_userQueryContainer {

	tmp := make(map[string]types.PrefixQuery)

	tmp[key] = *value.PrefixQueryCaster()

	s.v.Prefix = tmp
	return s
}

// Returns users that contain terms within a provided range.
// Range is a single key dictionnary.
// It will replace the current value on each call.
func (s *_userQueryContainer) Range(key string, value types.RangeQueryVariant) *_userQueryContainer {

	tmp := make(map[string]types.RangeQuery)

	tmp[key] = *value.RangeQueryCaster()

	s.v.Range = tmp
	return s
}

// Returns users based on a provided query string, using a parser with a limited
// but fault-tolerant syntax.
func (s *_userQueryContainer) SimpleQueryString(simplequerystring types.SimpleQueryStringQueryVariant) *_userQueryContainer {

	s.v.SimpleQueryString = simplequerystring.SimpleQueryStringQueryCaster()

	return s
}

// Returns users that contain an exact term in a provided field.
// To return a document, the query term must exactly match the queried field's
// value, including whitespace and capitalization.
// Term is a single key dictionnary.
// It will replace the current value on each call.
func (s *_userQueryContainer) Term(key string, value types.TermQueryVariant) *_userQueryContainer {

	tmp := make(map[string]types.TermQuery)

	tmp[key] = *value.TermQueryCaster()

	s.v.Term = tmp
	return s
}

// Returns users that contain one or more exact terms in a provided field.
// To return a document, one or more terms must exactly match a field value,
// including whitespace and capitalization.
func (s *_userQueryContainer) Terms(terms types.TermsQueryVariant) *_userQueryContainer {

	s.v.Terms = terms.TermsQueryCaster()

	return s
}

// Returns users that contain terms matching a wildcard pattern.
// Wildcard is a single key dictionnary.
// It will replace the current value on each call.
func (s *_userQueryContainer) Wildcard(key string, value types.WildcardQueryVariant) *_userQueryContainer {

	tmp := make(map[string]types.WildcardQuery)

	tmp[key] = *value.WildcardQueryCaster()

	s.v.Wildcard = tmp
	return s
}

func (s *_userQueryContainer) UserQueryContainerCaster() *types.UserQueryContainer {
	return s.v
}
