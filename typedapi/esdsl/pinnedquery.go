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
// https://github.com/elastic/elasticsearch-specification/tree/beeb1dc688bcc058488dcc45d9cbd2cd364e9943

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _pinnedQuery struct {
	v *types.PinnedQuery
}

func NewPinnedQuery() *_pinnedQuery {
	return &_pinnedQuery{v: types.NewPinnedQuery()}
}

// AdditionalPinnedQueryProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_pinnedQuery) AdditionalPinnedQueryProperty(key string, value json.RawMessage) *_pinnedQuery {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalPinnedQueryProperty = tmp
	return s
}

func (s *_pinnedQuery) Boost(boost float32) *_pinnedQuery {

	s.v.Boost = &boost

	return s
}

func (s *_pinnedQuery) Docs(docs ...types.PinnedDocVariant) *_pinnedQuery {

	for _, v := range docs {

		s.v.Docs = append(s.v.Docs, *v.PinnedDocCaster())

	}
	return s
}

func (s *_pinnedQuery) Ids(ids ...string) *_pinnedQuery {

	for _, v := range ids {

		s.v.Ids = append(s.v.Ids, v)

	}
	return s
}

func (s *_pinnedQuery) Organic(organic types.QueryVariant) *_pinnedQuery {

	s.v.Organic = *organic.QueryCaster()

	return s
}

func (s *_pinnedQuery) QueryName_(queryname_ string) *_pinnedQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_pinnedQuery) PinnedQueryCaster() *types.PinnedQuery {
	return s.v
}
