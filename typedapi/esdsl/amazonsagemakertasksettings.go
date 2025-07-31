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

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _amazonSageMakerTaskSettings struct {
	v *types.AmazonSageMakerTaskSettings
}

func NewAmazonSageMakerTaskSettings() *_amazonSageMakerTaskSettings {

	return &_amazonSageMakerTaskSettings{v: types.NewAmazonSageMakerTaskSettings()}

}

func (s *_amazonSageMakerTaskSettings) CustomAttributes(customattributes string) *_amazonSageMakerTaskSettings {

	s.v.CustomAttributes = &customattributes

	return s
}

func (s *_amazonSageMakerTaskSettings) EnableExplanations(enableexplanations string) *_amazonSageMakerTaskSettings {

	s.v.EnableExplanations = &enableexplanations

	return s
}

func (s *_amazonSageMakerTaskSettings) InferenceId(inferenceid string) *_amazonSageMakerTaskSettings {

	s.v.InferenceId = &inferenceid

	return s
}

func (s *_amazonSageMakerTaskSettings) SessionId(sessionid string) *_amazonSageMakerTaskSettings {

	s.v.SessionId = &sessionid

	return s
}

func (s *_amazonSageMakerTaskSettings) TargetVariant(targetvariant string) *_amazonSageMakerTaskSettings {

	s.v.TargetVariant = &targetvariant

	return s
}

func (s *_amazonSageMakerTaskSettings) AmazonSageMakerTaskSettingsCaster() *types.AmazonSageMakerTaskSettings {
	return s.v
}
