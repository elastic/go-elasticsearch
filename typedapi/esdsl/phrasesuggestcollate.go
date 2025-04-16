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
// https://github.com/elastic/elasticsearch-specification/tree/cbfcc73d01310bed2a480ec35aaef98138b598e5

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _phraseSuggestCollate struct {
	v *types.PhraseSuggestCollate
}

func NewPhraseSuggestCollate(query types.PhraseSuggestCollateQueryVariant) *_phraseSuggestCollate {

	tmp := &_phraseSuggestCollate{v: types.NewPhraseSuggestCollate()}

	tmp.Query(query)

	return tmp

}

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

func (s *_phraseSuggestCollate) Prune(prune bool) *_phraseSuggestCollate {

	s.v.Prune = &prune

	return s
}

func (s *_phraseSuggestCollate) Query(query types.PhraseSuggestCollateQueryVariant) *_phraseSuggestCollate {

	s.v.Query = *query.PhraseSuggestCollateQueryCaster()

	return s
}

func (s *_phraseSuggestCollate) PhraseSuggestCollateCaster() *types.PhraseSuggestCollate {
	return s.v
}
