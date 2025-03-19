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

type _apiKeyQueryContainer struct {
	v *types.ApiKeyQueryContainer
}

func NewApiKeyQueryContainer() *_apiKeyQueryContainer {
	return &_apiKeyQueryContainer{v: types.NewApiKeyQueryContainer()}
}

// AdditionalApiKeyQueryContainerProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_apiKeyQueryContainer) AdditionalApiKeyQueryContainerProperty(key string, value json.RawMessage) *_apiKeyQueryContainer {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalApiKeyQueryContainerProperty = tmp
	return s
}

// Matches documents matching boolean combinations of other queries.
func (s *_apiKeyQueryContainer) Bool(bool types.BoolQueryVariant) *_apiKeyQueryContainer {

	s.v.Bool = bool.BoolQueryCaster()

	return s
}

// Returns documents that contain an indexed value for a field.
func (s *_apiKeyQueryContainer) Exists(exists types.ExistsQueryVariant) *_apiKeyQueryContainer {

	s.v.Exists = exists.ExistsQueryCaster()

	return s
}

// Returns documents based on their IDs.
// This query uses document IDs stored in the `_id` field.
func (s *_apiKeyQueryContainer) Ids(ids types.IdsQueryVariant) *_apiKeyQueryContainer {

	s.v.Ids = ids.IdsQueryCaster()

	return s
}

// Returns documents that match a provided text, number, date or boolean value.
// The provided text is analyzed before matching.
// Match is a single key dictionnary.
// It will replace the current value on each call.
func (s *_apiKeyQueryContainer) Match(key string, value types.MatchQueryVariant) *_apiKeyQueryContainer {

	tmp := make(map[string]types.MatchQuery)

	tmp[key] = *value.MatchQueryCaster()

	s.v.Match = tmp
	return s
}

// Matches all documents, giving them all a `_score` of 1.0.
func (s *_apiKeyQueryContainer) MatchAll(matchall types.MatchAllQueryVariant) *_apiKeyQueryContainer {

	s.v.MatchAll = matchall.MatchAllQueryCaster()

	return s
}

// Returns documents that contain a specific prefix in a provided field.
// Prefix is a single key dictionnary.
// It will replace the current value on each call.
func (s *_apiKeyQueryContainer) Prefix(key string, value types.PrefixQueryVariant) *_apiKeyQueryContainer {

	tmp := make(map[string]types.PrefixQuery)

	tmp[key] = *value.PrefixQueryCaster()

	s.v.Prefix = tmp
	return s
}

// Returns documents that contain terms within a provided range.
// Range is a single key dictionnary.
// It will replace the current value on each call.
func (s *_apiKeyQueryContainer) Range(key string, value types.RangeQueryVariant) *_apiKeyQueryContainer {

	tmp := make(map[string]types.RangeQuery)

	tmp[key] = *value.RangeQueryCaster()

	s.v.Range = tmp
	return s
}

// Returns documents based on a provided query string, using a parser with a
// limited but fault-tolerant syntax.
func (s *_apiKeyQueryContainer) SimpleQueryString(simplequerystring types.SimpleQueryStringQueryVariant) *_apiKeyQueryContainer {

	s.v.SimpleQueryString = simplequerystring.SimpleQueryStringQueryCaster()

	return s
}

// Returns documents that contain an exact term in a provided field.
// To return a document, the query term must exactly match the queried field's
// value, including whitespace and capitalization.
// Term is a single key dictionnary.
// It will replace the current value on each call.
func (s *_apiKeyQueryContainer) Term(key string, value types.TermQueryVariant) *_apiKeyQueryContainer {

	tmp := make(map[string]types.TermQuery)

	tmp[key] = *value.TermQueryCaster()

	s.v.Term = tmp
	return s
}

// Returns documents that contain one or more exact terms in a provided field.
// To return a document, one or more terms must exactly match a field value,
// including whitespace and capitalization.
func (s *_apiKeyQueryContainer) Terms(terms types.TermsQueryVariant) *_apiKeyQueryContainer {

	s.v.Terms = terms.TermsQueryCaster()

	return s
}

// Returns documents that contain terms matching a wildcard pattern.
// Wildcard is a single key dictionnary.
// It will replace the current value on each call.
func (s *_apiKeyQueryContainer) Wildcard(key string, value types.WildcardQueryVariant) *_apiKeyQueryContainer {

	tmp := make(map[string]types.WildcardQuery)

	tmp[key] = *value.WildcardQueryCaster()

	s.v.Wildcard = tmp
	return s
}

func (s *_apiKeyQueryContainer) ApiKeyQueryContainerCaster() *types.ApiKeyQueryContainer {
	return s.v
}
