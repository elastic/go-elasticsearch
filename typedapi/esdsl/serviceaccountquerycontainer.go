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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _serviceAccountQueryContainer struct {
	v *types.ServiceAccountQueryContainer
}

func NewServiceAccountQueryContainer() *_serviceAccountQueryContainer {
	return &_serviceAccountQueryContainer{v: types.NewServiceAccountQueryContainer()}
}

// AdditionalServiceAccountQueryContainerProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_serviceAccountQueryContainer) AdditionalServiceAccountQueryContainerProperty(key string, value json.RawMessage) *_serviceAccountQueryContainer {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalServiceAccountQueryContainerProperty = tmp
	return s
}

func (s *_serviceAccountQueryContainer) Bool(bool types.BoolQueryVariant) *_serviceAccountQueryContainer {

	s.v.Bool = bool.BoolQueryCaster()

	return s
}

func (s *_serviceAccountQueryContainer) Exists(exists types.ExistsQueryVariant) *_serviceAccountQueryContainer {

	s.v.Exists = exists.ExistsQueryCaster()

	return s
}

func (s *_serviceAccountQueryContainer) Ids(ids types.IdsQueryVariant) *_serviceAccountQueryContainer {

	s.v.Ids = ids.IdsQueryCaster()

	return s
}

// Match is a single key dictionnary.
// It will replace the current value on each call.
func (s *_serviceAccountQueryContainer) Match(key string, value types.MatchQueryVariant) *_serviceAccountQueryContainer {

	tmp := make(map[string]types.MatchQuery)

	tmp[key] = *value.MatchQueryCaster()

	s.v.Match = tmp
	return s
}

func (s *_serviceAccountQueryContainer) MatchAll(matchall types.MatchAllQueryVariant) *_serviceAccountQueryContainer {

	s.v.MatchAll = matchall.MatchAllQueryCaster()

	return s
}

// Prefix is a single key dictionnary.
// It will replace the current value on each call.
func (s *_serviceAccountQueryContainer) Prefix(key string, value types.PrefixQueryVariant) *_serviceAccountQueryContainer {

	tmp := make(map[string]types.PrefixQuery)

	tmp[key] = *value.PrefixQueryCaster()

	s.v.Prefix = tmp
	return s
}

// Range is a single key dictionnary.
// It will replace the current value on each call.
func (s *_serviceAccountQueryContainer) Range(key string, value types.RangeQueryVariant) *_serviceAccountQueryContainer {

	tmp := make(map[string]types.RangeQuery)

	tmp[key] = *value.RangeQueryCaster()

	s.v.Range = tmp
	return s
}

func (s *_serviceAccountQueryContainer) SimpleQueryString(simplequerystring types.SimpleQueryStringQueryVariant) *_serviceAccountQueryContainer {

	s.v.SimpleQueryString = simplequerystring.SimpleQueryStringQueryCaster()

	return s
}

// Term is a single key dictionnary.
// It will replace the current value on each call.
func (s *_serviceAccountQueryContainer) Term(key string, value types.TermQueryVariant) *_serviceAccountQueryContainer {

	tmp := make(map[string]types.TermQuery)

	tmp[key] = *value.TermQueryCaster()

	s.v.Term = tmp
	return s
}

func (s *_serviceAccountQueryContainer) Terms(terms types.TermsQueryVariant) *_serviceAccountQueryContainer {

	s.v.Terms = terms.TermsQueryCaster()

	return s
}

// Wildcard is a single key dictionnary.
// It will replace the current value on each call.
func (s *_serviceAccountQueryContainer) Wildcard(key string, value types.WildcardQueryVariant) *_serviceAccountQueryContainer {

	tmp := make(map[string]types.WildcardQuery)

	tmp[key] = *value.WildcardQueryCaster()

	s.v.Wildcard = tmp
	return s
}

func (s *_serviceAccountQueryContainer) ServiceAccountQueryContainerCaster() *types.ServiceAccountQueryContainer {
	return s.v
}
