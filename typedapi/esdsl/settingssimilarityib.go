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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/ibdistribution"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/iblambda"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/normalization"
)

type _settingsSimilarityIb struct {
	v *types.SettingsSimilarityIb
}

func NewSettingsSimilarityIb(distribution ibdistribution.IBDistribution, lambda iblambda.IBLambda, normalization normalization.Normalization) *_settingsSimilarityIb {

	tmp := &_settingsSimilarityIb{v: types.NewSettingsSimilarityIb()}

	tmp.Distribution(distribution)

	tmp.Lambda(lambda)

	tmp.Normalization(normalization)

	return tmp

}

func (s *_settingsSimilarityIb) Distribution(distribution ibdistribution.IBDistribution) *_settingsSimilarityIb {

	s.v.Distribution = distribution
	return s
}

func (s *_settingsSimilarityIb) Lambda(lambda iblambda.IBLambda) *_settingsSimilarityIb {

	s.v.Lambda = lambda
	return s
}

func (s *_settingsSimilarityIb) Normalization(normalization normalization.Normalization) *_settingsSimilarityIb {

	s.v.Normalization = normalization
	return s
}

func (s *_settingsSimilarityIb) SettingsSimilarityIbCaster() *types.SettingsSimilarityIb {
	return s.v
}
