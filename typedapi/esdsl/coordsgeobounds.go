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

type _coordsGeoBounds struct {
	v *types.CoordsGeoBounds
}

func NewCoordsGeoBounds(bottom types.Float64, left types.Float64, right types.Float64, top types.Float64) *_coordsGeoBounds {

	tmp := &_coordsGeoBounds{v: types.NewCoordsGeoBounds()}

	tmp.Bottom(bottom)

	tmp.Left(left)

	tmp.Right(right)

	tmp.Top(top)

	return tmp

}

func (s *_coordsGeoBounds) Bottom(bottom types.Float64) *_coordsGeoBounds {

	s.v.Bottom = bottom

	return s
}

func (s *_coordsGeoBounds) Left(left types.Float64) *_coordsGeoBounds {

	s.v.Left = left

	return s
}

func (s *_coordsGeoBounds) Right(right types.Float64) *_coordsGeoBounds {

	s.v.Right = right

	return s
}

func (s *_coordsGeoBounds) Top(top types.Float64) *_coordsGeoBounds {

	s.v.Top = top

	return s
}

func (s *_coordsGeoBounds) CoordsGeoBoundsCaster() *types.CoordsGeoBounds {
	return s.v
}
