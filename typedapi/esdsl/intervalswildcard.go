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

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _intervalsWildcard struct {
	v *types.IntervalsWildcard
}

// Matches terms using a wildcard pattern.
func NewIntervalsWildcard(pattern string) *_intervalsWildcard {

	tmp := &_intervalsWildcard{v: types.NewIntervalsWildcard()}

	tmp.Pattern(pattern)

	return tmp

}

// Analyzer used to analyze the `pattern`.
// Defaults to the top-level field's analyzer.
func (s *_intervalsWildcard) Analyzer(analyzer string) *_intervalsWildcard {

	s.v.Analyzer = &analyzer

	return s
}

// Wildcard pattern used to find matching terms.
func (s *_intervalsWildcard) Pattern(pattern string) *_intervalsWildcard {

	s.v.Pattern = pattern

	return s
}

// If specified, match intervals from this field rather than the top-level
// field.
// The `pattern` is normalized using the search analyzer from this field, unless
// `analyzer` is specified separately.
func (s *_intervalsWildcard) UseField(field string) *_intervalsWildcard {

	s.v.UseField = &field

	return s
}

func (s *_intervalsWildcard) IntervalsCaster() *types.Intervals {
	container := types.NewIntervals()

	container.Wildcard = s.v

	return container
}

func (s *_intervalsWildcard) IntervalsQueryCaster() *types.IntervalsQuery {
	container := types.NewIntervalsQuery()

	container.Wildcard = s.v

	return container
}

func (s *_intervalsWildcard) IntervalsWildcardCaster() *types.IntervalsWildcard {
	return s.v
}
