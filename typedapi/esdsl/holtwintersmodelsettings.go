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
// https://github.com/elastic/elasticsearch-specification/tree/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/holtwinterstype"
)

type _holtWintersModelSettings struct {
	v *types.HoltWintersModelSettings
}

func NewHoltWintersModelSettings() *_holtWintersModelSettings {

	return &_holtWintersModelSettings{v: types.NewHoltWintersModelSettings()}

}

func (s *_holtWintersModelSettings) Alpha(alpha float32) *_holtWintersModelSettings {

	s.v.Alpha = &alpha

	return s
}

func (s *_holtWintersModelSettings) Beta(beta float32) *_holtWintersModelSettings {

	s.v.Beta = &beta

	return s
}

func (s *_holtWintersModelSettings) Gamma(gamma float32) *_holtWintersModelSettings {

	s.v.Gamma = &gamma

	return s
}

func (s *_holtWintersModelSettings) Pad(pad bool) *_holtWintersModelSettings {

	s.v.Pad = &pad

	return s
}

func (s *_holtWintersModelSettings) Period(period int) *_holtWintersModelSettings {

	s.v.Period = &period

	return s
}

func (s *_holtWintersModelSettings) Type(type_ holtwinterstype.HoltWintersType) *_holtWintersModelSettings {

	s.v.Type = &type_
	return s
}

func (s *_holtWintersModelSettings) HoltWintersModelSettingsCaster() *types.HoltWintersModelSettings {
	return s.v
}
