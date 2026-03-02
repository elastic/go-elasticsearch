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
// https://github.com/elastic/elasticsearch-specification/tree/e196f9953fa743572ee46884835f1934bce9a16b

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _suggester struct {
	v *types.Suggester
}

func NewSuggester() *_suggester {

	return &_suggester{v: types.NewSuggester()}

}

func (s *_suggester) Suggesters(suggesters map[string]types.FieldSuggester) *_suggester {

	s.v.Suggesters = suggesters
	return s
}

func (s *_suggester) AddSuggester(key string, value types.FieldSuggesterVariant) *_suggester {

	var tmp map[string]types.FieldSuggester
	if s.v.Suggesters == nil {
		s.v.Suggesters = make(map[string]types.FieldSuggester)
	} else {
		tmp = s.v.Suggesters
	}

	tmp[key] = *value.FieldSuggesterCaster()

	s.v.Suggesters = tmp
	return s
}

func (s *_suggester) Text(text string) *_suggester {

	s.v.Text = &text

	return s
}

func (s *_suggester) SuggesterCaster() *types.Suggester {
	return s.v
}
