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

type _intervalsAllOf struct {
	v *types.IntervalsAllOf
}

// Returns matches that span a combination of other rules.
func NewIntervalsAllOf() *_intervalsAllOf {

	return &_intervalsAllOf{v: types.NewIntervalsAllOf()}

}

func (s *_intervalsAllOf) Filter(filter types.IntervalsFilterVariant) *_intervalsAllOf {

	s.v.Filter = filter.IntervalsFilterCaster()

	return s
}

func (s *_intervalsAllOf) Intervals(intervals ...types.IntervalsVariant) *_intervalsAllOf {

	for _, v := range intervals {

		s.v.Intervals = append(s.v.Intervals, *v.IntervalsCaster())

	}
	return s
}

func (s *_intervalsAllOf) MaxGaps(maxgaps int) *_intervalsAllOf {

	s.v.MaxGaps = &maxgaps

	return s
}

func (s *_intervalsAllOf) Ordered(ordered bool) *_intervalsAllOf {

	s.v.Ordered = &ordered

	return s
}

func (s *_intervalsAllOf) IntervalsCaster() *types.Intervals {
	container := types.NewIntervals()

	container.AllOf = s.v

	return container
}

func (s *_intervalsAllOf) IntervalsQueryCaster() *types.IntervalsQuery {
	container := types.NewIntervalsQuery()

	container.AllOf = s.v

	return container
}

func (s *_intervalsAllOf) IntervalsAllOfCaster() *types.IntervalsAllOf {
	return s.v
}
