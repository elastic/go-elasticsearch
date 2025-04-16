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
// https://github.com/elastic/elasticsearch-specification/tree/cbfcc73d01310bed2a480ec35aaef98138b598e5

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dfraftereffect"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dfrbasicmodel"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/normalization"
)

type _settingsSimilarityDfr struct {
	v *types.SettingsSimilarityDfr
}

func NewSettingsSimilarityDfr(aftereffect dfraftereffect.DFRAfterEffect, basicmodel dfrbasicmodel.DFRBasicModel, normalization normalization.Normalization) *_settingsSimilarityDfr {

	tmp := &_settingsSimilarityDfr{v: types.NewSettingsSimilarityDfr()}

	tmp.AfterEffect(aftereffect)

	tmp.BasicModel(basicmodel)

	tmp.Normalization(normalization)

	return tmp

}

func (s *_settingsSimilarityDfr) AfterEffect(aftereffect dfraftereffect.DFRAfterEffect) *_settingsSimilarityDfr {

	s.v.AfterEffect = aftereffect
	return s
}

func (s *_settingsSimilarityDfr) BasicModel(basicmodel dfrbasicmodel.DFRBasicModel) *_settingsSimilarityDfr {

	s.v.BasicModel = basicmodel
	return s
}

func (s *_settingsSimilarityDfr) Normalization(normalization normalization.Normalization) *_settingsSimilarityDfr {

	s.v.Normalization = normalization
	return s
}

func (s *_settingsSimilarityDfr) SettingsSimilarityDfrCaster() *types.SettingsSimilarityDfr {
	return s.v
}
