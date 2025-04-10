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

type _intervals struct {
	v *types.Intervals
}

func NewIntervals() *_intervals {
	return &_intervals{v: types.NewIntervals()}
}

// AdditionalIntervalsProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_intervals) AdditionalIntervalsProperty(key string, value json.RawMessage) *_intervals {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalIntervalsProperty = tmp
	return s
}

func (s *_intervals) AllOf(allof types.IntervalsAllOfVariant) *_intervals {

	s.v.AllOf = allof.IntervalsAllOfCaster()

	return s
}

func (s *_intervals) AnyOf(anyof types.IntervalsAnyOfVariant) *_intervals {

	s.v.AnyOf = anyof.IntervalsAnyOfCaster()

	return s
}

func (s *_intervals) Fuzzy(fuzzy types.IntervalsFuzzyVariant) *_intervals {

	s.v.Fuzzy = fuzzy.IntervalsFuzzyCaster()

	return s
}

func (s *_intervals) Match(match types.IntervalsMatchVariant) *_intervals {

	s.v.Match = match.IntervalsMatchCaster()

	return s
}

func (s *_intervals) Prefix(prefix types.IntervalsPrefixVariant) *_intervals {

	s.v.Prefix = prefix.IntervalsPrefixCaster()

	return s
}

func (s *_intervals) Wildcard(wildcard types.IntervalsWildcardVariant) *_intervals {

	s.v.Wildcard = wildcard.IntervalsWildcardCaster()

	return s
}

func (s *_intervals) IntervalsCaster() *types.Intervals {
	return s.v
}
