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
// https://github.com/elastic/elasticsearch-specification/tree/de4ff9ec1f716256f521d9e30011ad9c284b0dcc

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _intervalsPrefix struct {
	v *types.IntervalsPrefix
}

// Matches terms that start with a specified set of characters.
func NewIntervalsPrefix(prefix string) *_intervalsPrefix {

	tmp := &_intervalsPrefix{v: types.NewIntervalsPrefix()}

	tmp.Prefix(prefix)

	return tmp

}

func (s *_intervalsPrefix) Analyzer(analyzer string) *_intervalsPrefix {

	s.v.Analyzer = &analyzer

	return s
}

func (s *_intervalsPrefix) Prefix(prefix string) *_intervalsPrefix {

	s.v.Prefix = prefix

	return s
}

func (s *_intervalsPrefix) UseField(field string) *_intervalsPrefix {

	s.v.UseField = &field

	return s
}

func (s *_intervalsPrefix) IntervalsCaster() *types.Intervals {
	container := types.NewIntervals()

	container.Prefix = s.v

	return container
}

func (s *_intervalsPrefix) IntervalsQueryCaster() *types.IntervalsQuery {
	container := types.NewIntervalsQuery()

	container.Prefix = s.v

	return container
}

func (s *_intervalsPrefix) IntervalsPrefixCaster() *types.IntervalsPrefix {
	return s.v
}
