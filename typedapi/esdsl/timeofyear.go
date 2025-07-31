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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/month"
)

type _timeOfYear struct {
	v *types.TimeOfYear
}

func NewTimeOfYear() *_timeOfYear {

	return &_timeOfYear{v: types.NewTimeOfYear()}

}

func (s *_timeOfYear) At(ats ...string) *_timeOfYear {

	for _, v := range ats {

		s.v.At = append(s.v.At, v)

	}
	return s
}

func (s *_timeOfYear) Int(ints ...month.Month) *_timeOfYear {

	for _, v := range ints {

		s.v.Int = append(s.v.Int, v)

	}
	return s
}

func (s *_timeOfYear) On(ons ...int) *_timeOfYear {

	for _, v := range ons {

		s.v.On = append(s.v.On, v)

	}
	return s
}

func (s *_timeOfYear) TimeOfYearCaster() *types.TimeOfYear {
	return s.v
}
