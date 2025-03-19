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
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _phraseSuggestCollate struct {
	v *types.PhraseSuggestCollate
}

func NewPhraseSuggestCollate(query types.PhraseSuggestCollateQueryVariant) *_phraseSuggestCollate {

	tmp := &_phraseSuggestCollate{v: types.NewPhraseSuggestCollate()}

	tmp.Query(query)

	return tmp

}

// Parameters to use if the query is templated.
func (s *_phraseSuggestCollate) Params(params map[string]json.RawMessage) *_phraseSuggestCollate {

	s.v.Params = params
	return s
}

func (s *_phraseSuggestCollate) AddParam(key string, value json.RawMessage) *_phraseSuggestCollate {

	var tmp map[string]json.RawMessage
	if s.v.Params == nil {
		s.v.Params = make(map[string]json.RawMessage)
	} else {
		tmp = s.v.Params
	}

	tmp[key] = value

	s.v.Params = tmp
	return s
}

// Returns all suggestions with an extra `collate_match` option indicating
// whether the generated phrase matched any document.
func (s *_phraseSuggestCollate) Prune(prune bool) *_phraseSuggestCollate {

	s.v.Prune = &prune

	return s
}

// A collate query that is run once for every suggestion.
func (s *_phraseSuggestCollate) Query(query types.PhraseSuggestCollateQueryVariant) *_phraseSuggestCollate {

	s.v.Query = *query.PhraseSuggestCollateQueryCaster()

	return s
}

func (s *_phraseSuggestCollate) PhraseSuggestCollateCaster() *types.PhraseSuggestCollate {
	return s.v
}
