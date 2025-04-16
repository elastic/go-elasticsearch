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

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _voyageAITaskSettings struct {
	v *types.VoyageAITaskSettings
}

func NewVoyageAITaskSettings() *_voyageAITaskSettings {

	return &_voyageAITaskSettings{v: types.NewVoyageAITaskSettings()}

}

func (s *_voyageAITaskSettings) InputType(inputtype string) *_voyageAITaskSettings {

	s.v.InputType = &inputtype

	return s
}

func (s *_voyageAITaskSettings) ReturnDocuments(returndocuments bool) *_voyageAITaskSettings {

	s.v.ReturnDocuments = &returndocuments

	return s
}

func (s *_voyageAITaskSettings) TopK(topk int) *_voyageAITaskSettings {

	s.v.TopK = &topk

	return s
}

func (s *_voyageAITaskSettings) Truncation(truncation bool) *_voyageAITaskSettings {

	s.v.Truncation = &truncation

	return s
}

func (s *_voyageAITaskSettings) VoyageAITaskSettingsCaster() *types.VoyageAITaskSettings {
	return s.v
}
