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

// Returns matches that span a combination of other rules.
func (s *_intervalsQuery) AllOf(allof types.IntervalsAllOfVariant) *_intervalsQuery {

	s.v.AllOf = allof.IntervalsAllOfCaster()

	return s
}

// Returns intervals produced by any of its sub-rules.
func (s *_intervalsQuery) AnyOf(anyof types.IntervalsAnyOfVariant) *_intervalsQuery {

	s.v.AnyOf = anyof.IntervalsAnyOfCaster()

	return s
}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_intervalsQuery) Boost(boost float32) *_intervalsQuery {

	s.v.Boost = &boost

	return s
}

// Matches terms that are similar to the provided term, within an edit distance
// defined by `fuzziness`.
func (s *_intervalsQuery) Fuzzy(fuzzy types.IntervalsFuzzyVariant) *_intervalsQuery {

	s.v.Fuzzy = fuzzy.IntervalsFuzzyCaster()

	return s
}

// Matches analyzed text.
func (s *_intervalsQuery) Match(match types.IntervalsMatchVariant) *_intervalsQuery {

	s.v.Match = match.IntervalsMatchCaster()

	return s
}

// Matches terms that start with a specified set of characters.
func (s *_intervalsQuery) Prefix(prefix types.IntervalsPrefixVariant) *_intervalsQuery {

	s.v.Prefix = prefix.IntervalsPrefixCaster()

	return s
}

func (s *_intervalsQuery) QueryName_(queryname_ string) *_intervalsQuery {

	s.v.QueryName_ = &queryname_

	return s
}

// Matches terms using a wildcard pattern.
func (s *_intervalsQuery) Wildcard(wildcard types.IntervalsWildcardVariant) *_intervalsQuery {

	s.v.Wildcard = wildcard.IntervalsWildcardCaster()

	return s
}

func (s *_intervalsQuery) IntervalsQueryCaster() *types.IntervalsQuery {
	return s.v
}
