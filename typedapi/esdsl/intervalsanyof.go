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

type _intervalsAnyOf struct {
	v *types.IntervalsAnyOf
}

// Returns intervals produced by any of its sub-rules.
func NewIntervalsAnyOf() *_intervalsAnyOf {

	return &_intervalsAnyOf{v: types.NewIntervalsAnyOf()}

}

// Rule used to filter returned intervals.
func (s *_intervalsAnyOf) Filter(filter types.IntervalsFilterVariant) *_intervalsAnyOf {

	s.v.Filter = filter.IntervalsFilterCaster()

	return s
}

// An array of rules to match.
func (s *_intervalsAnyOf) Intervals(intervals ...types.IntervalsVariant) *_intervalsAnyOf {

	for _, v := range intervals {

		s.v.Intervals = append(s.v.Intervals, *v.IntervalsCaster())

	}
	return s
}

func (s *_intervalsAnyOf) IntervalsCaster() *types.Intervals {
	container := types.NewIntervals()

	container.AnyOf = s.v

	return container
}

func (s *_intervalsAnyOf) IntervalsQueryCaster() *types.IntervalsQuery {
	container := types.NewIntervalsQuery()

	container.AnyOf = s.v

	return container
}

func (s *_intervalsAnyOf) IntervalsAnyOfCaster() *types.IntervalsAnyOf {
	return s.v
}
