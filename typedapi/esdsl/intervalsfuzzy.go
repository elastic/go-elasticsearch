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

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _intervalsFuzzy struct {
	v *types.IntervalsFuzzy
}

// Matches analyzed text.
func NewIntervalsFuzzy(term string) *_intervalsFuzzy {

	tmp := &_intervalsFuzzy{v: types.NewIntervalsFuzzy()}

	tmp.Term(term)

	return tmp

}

// Analyzer used to normalize the term.
func (s *_intervalsFuzzy) Analyzer(analyzer string) *_intervalsFuzzy {

	s.v.Analyzer = &analyzer

	return s
}

// Maximum edit distance allowed for matching.
func (s *_intervalsFuzzy) Fuzziness(fuzziness types.FuzzinessVariant) *_intervalsFuzzy {

	s.v.Fuzziness = *fuzziness.FuzzinessCaster()

	return s
}

// Number of beginning characters left unchanged when creating expansions.
func (s *_intervalsFuzzy) PrefixLength(prefixlength int) *_intervalsFuzzy {

	s.v.PrefixLength = &prefixlength

	return s
}

// The term to match.
func (s *_intervalsFuzzy) Term(term string) *_intervalsFuzzy {

	s.v.Term = term

	return s
}

// Indicates whether edits include transpositions of two adjacent characters
// (for example, `ab` to `ba`).
func (s *_intervalsFuzzy) Transpositions(transpositions bool) *_intervalsFuzzy {

	s.v.Transpositions = &transpositions

	return s
}

// If specified, match intervals from this field rather than the top-level
// field.
// The `term` is normalized using the search analyzer from this field, unless
// `analyzer` is specified separately.
func (s *_intervalsFuzzy) UseField(field string) *_intervalsFuzzy {

	s.v.UseField = &field

	return s
}

func (s *_intervalsFuzzy) IntervalsCaster() *types.Intervals {
	container := types.NewIntervals()

	container.Fuzzy = s.v

	return container
}

func (s *_intervalsFuzzy) IntervalsQueryCaster() *types.IntervalsQuery {
	container := types.NewIntervalsQuery()

	container.Fuzzy = s.v

	return container
}

func (s *_intervalsFuzzy) IntervalsFuzzyCaster() *types.IntervalsFuzzy {
	return s.v
}
