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
// https://github.com/elastic/elasticsearch-specification/tree/60a81659be928bfe6cec53708c7f7613555a5eaf

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _googleVertexAITaskSettings struct {
	v *types.GoogleVertexAITaskSettings
}

func NewGoogleVertexAITaskSettings() *_googleVertexAITaskSettings {

	return &_googleVertexAITaskSettings{v: types.NewGoogleVertexAITaskSettings()}

}

func (s *_googleVertexAITaskSettings) AutoTruncate(autotruncate bool) *_googleVertexAITaskSettings {

	s.v.AutoTruncate = &autotruncate

	return s
}

func (s *_googleVertexAITaskSettings) TopN(topn int) *_googleVertexAITaskSettings {

	s.v.TopN = &topn

	return s
}

func (s *_googleVertexAITaskSettings) GoogleVertexAITaskSettingsCaster() *types.GoogleVertexAITaskSettings {
	return s.v
}
