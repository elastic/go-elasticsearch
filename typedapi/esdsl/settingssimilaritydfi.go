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
// https://github.com/elastic/elasticsearch-specification/tree/86f41834c7bb975159a38a73be8a9d930010d673

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dfiindependencemeasure"
)

type _settingsSimilarityDfi struct {
	v *types.SettingsSimilarityDfi
}

func NewSettingsSimilarityDfi(independencemeasure dfiindependencemeasure.DFIIndependenceMeasure) *_settingsSimilarityDfi {

	tmp := &_settingsSimilarityDfi{v: types.NewSettingsSimilarityDfi()}

	tmp.IndependenceMeasure(independencemeasure)

	return tmp

}

func (s *_settingsSimilarityDfi) IndependenceMeasure(independencemeasure dfiindependencemeasure.DFIIndependenceMeasure) *_settingsSimilarityDfi {

	s.v.IndependenceMeasure = independencemeasure
	return s
}

func (s *_settingsSimilarityDfi) SettingsSimilarityDfiCaster() *types.SettingsSimilarityDfi {
	return s.v
}
