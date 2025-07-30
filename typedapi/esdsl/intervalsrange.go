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
// https://github.com/elastic/elasticsearch-specification/tree/e585438d116b00ff34643179e6286e402c0bcaaf

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _intervalsRange struct {
	v *types.IntervalsRange
}

func NewIntervalsRange() *_intervalsRange {

	return &_intervalsRange{v: types.NewIntervalsRange()}

}

func (s *_intervalsRange) Analyzer(analyzer string) *_intervalsRange {

	s.v.Analyzer = &analyzer

	return s
}

func (s *_intervalsRange) Gt(gt string) *_intervalsRange {

	s.v.Gt = &gt

	return s
}

func (s *_intervalsRange) Gte(gte string) *_intervalsRange {

	s.v.Gte = &gte

	return s
}

func (s *_intervalsRange) Lt(lt string) *_intervalsRange {

	s.v.Lt = &lt

	return s
}

func (s *_intervalsRange) Lte(lte string) *_intervalsRange {

	s.v.Lte = &lte

	return s
}

func (s *_intervalsRange) UseField(field string) *_intervalsRange {

	s.v.UseField = &field

	return s
}

func (s *_intervalsRange) IntervalsCaster() *types.Intervals {
	container := types.NewIntervals()

	container.Range = s.v

	return container
}

func (s *_intervalsRange) IntervalsQueryCaster() *types.IntervalsQuery {
	container := types.NewIntervalsQuery()

	container.Range = s.v

	return container
}

func (s *_intervalsRange) IntervalsRangeCaster() *types.IntervalsRange {
	return s.v
}
