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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _intervalsQuery struct {
	v *types.IntervalsQuery
}

func NewIntervalsQuery() *_intervalsQuery {
	return &_intervalsQuery{v: types.NewIntervalsQuery()}
}

// AdditionalIntervalsQueryProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_intervalsQuery) AdditionalIntervalsQueryProperty(key string, value json.RawMessage) *_intervalsQuery {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalIntervalsQueryProperty = tmp
	return s
}

func (s *_intervalsQuery) AllOf(allof types.IntervalsAllOfVariant) *_intervalsQuery {

	s.v.AllOf = allof.IntervalsAllOfCaster()

	return s
}

func (s *_intervalsQuery) AnyOf(anyof types.IntervalsAnyOfVariant) *_intervalsQuery {

	s.v.AnyOf = anyof.IntervalsAnyOfCaster()

	return s
}

func (s *_intervalsQuery) Boost(boost float32) *_intervalsQuery {

	s.v.Boost = &boost

	return s
}

func (s *_intervalsQuery) Fuzzy(fuzzy types.IntervalsFuzzyVariant) *_intervalsQuery {

	s.v.Fuzzy = fuzzy.IntervalsFuzzyCaster()

	return s
}

func (s *_intervalsQuery) Match(match types.IntervalsMatchVariant) *_intervalsQuery {

	s.v.Match = match.IntervalsMatchCaster()

	return s
}

func (s *_intervalsQuery) Prefix(prefix types.IntervalsPrefixVariant) *_intervalsQuery {

	s.v.Prefix = prefix.IntervalsPrefixCaster()

	return s
}

func (s *_intervalsQuery) QueryName_(queryname_ string) *_intervalsQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_intervalsQuery) Wildcard(wildcard types.IntervalsWildcardVariant) *_intervalsQuery {

	s.v.Wildcard = wildcard.IntervalsWildcardCaster()

	return s
}

func (s *_intervalsQuery) IntervalsQueryCaster() *types.IntervalsQuery {
	return s.v
}
