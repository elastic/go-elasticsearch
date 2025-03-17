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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/day"
)

type _timeOfWeek struct {
	v *types.TimeOfWeek
}

func NewTimeOfWeek() *_timeOfWeek {

	return &_timeOfWeek{v: types.NewTimeOfWeek()}

}

func (s *_timeOfWeek) At(ats ...string) *_timeOfWeek {

	for _, v := range ats {

		s.v.At = append(s.v.At, v)

	}
	return s
}

func (s *_timeOfWeek) On(ons ...day.Day) *_timeOfWeek {

	for _, v := range ons {

		s.v.On = append(s.v.On, v)

	}
	return s
}

func (s *_timeOfWeek) TimeOfWeekCaster() *types.TimeOfWeek {
	return s.v
}
