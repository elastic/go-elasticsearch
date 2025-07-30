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

type _settingsSimilarityLmd struct {
	v *types.SettingsSimilarityLmd
}

func NewSettingsSimilarityLmd() *_settingsSimilarityLmd {

	return &_settingsSimilarityLmd{v: types.NewSettingsSimilarityLmd()}

}

func (s *_settingsSimilarityLmd) Mu(mu types.Float64) *_settingsSimilarityLmd {

	s.v.Mu = &mu

	return s
}

func (s *_settingsSimilarityLmd) SettingsSimilarityLmdCaster() *types.SettingsSimilarityLmd {
	return s.v
}
