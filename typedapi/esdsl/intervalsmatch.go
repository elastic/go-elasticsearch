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

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _intervalsMatch struct {
	v *types.IntervalsMatch
}

// Matches analyzed text.
func NewIntervalsMatch(query string) *_intervalsMatch {

	tmp := &_intervalsMatch{v: types.NewIntervalsMatch()}

	tmp.Query(query)

	return tmp

}

func (s *_intervalsMatch) Analyzer(analyzer string) *_intervalsMatch {

	s.v.Analyzer = &analyzer

	return s
}

func (s *_intervalsMatch) Filter(filter types.IntervalsFilterVariant) *_intervalsMatch {

	s.v.Filter = filter.IntervalsFilterCaster()

	return s
}

func (s *_intervalsMatch) MaxGaps(maxgaps int) *_intervalsMatch {

	s.v.MaxGaps = &maxgaps

	return s
}

func (s *_intervalsMatch) Ordered(ordered bool) *_intervalsMatch {

	s.v.Ordered = &ordered

	return s
}

func (s *_intervalsMatch) Query(query string) *_intervalsMatch {

	s.v.Query = query

	return s
}

func (s *_intervalsMatch) UseField(field string) *_intervalsMatch {

	s.v.UseField = &field

	return s
}

func (s *_intervalsMatch) IntervalsCaster() *types.Intervals {
	container := types.NewIntervals()

	container.Match = s.v

	return container
}

func (s *_intervalsMatch) IntervalsQueryCaster() *types.IntervalsQuery {
	container := types.NewIntervalsQuery()

	container.Match = s.v

	return container
}

func (s *_intervalsMatch) IntervalsMatchCaster() *types.IntervalsMatch {
	return s.v
}
