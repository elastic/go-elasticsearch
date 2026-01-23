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
// https://github.com/elastic/elasticsearch-specification/tree/b1811e10a0722431d79d1c234dd412ff47d8656f

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _learningToRank struct {
	v *types.LearningToRank
}

func NewLearningToRank(modelid string) *_learningToRank {

	tmp := &_learningToRank{v: types.NewLearningToRank()}

	tmp.ModelId(modelid)

	return tmp

}

func (s *_learningToRank) ModelId(modelid string) *_learningToRank {

	s.v.ModelId = modelid

	return s
}

func (s *_learningToRank) Params(params map[string]json.RawMessage) *_learningToRank {

	s.v.Params = params
	return s
}

func (s *_learningToRank) AddParam(key string, value json.RawMessage) *_learningToRank {

	var tmp map[string]json.RawMessage
	if s.v.Params == nil {
		s.v.Params = make(map[string]json.RawMessage)
	} else {
		tmp = s.v.Params
	}

	tmp[key] = value

	s.v.Params = tmp
	return s
}

func (s *_learningToRank) RescoreCaster() *types.Rescore {
	container := types.NewRescore()

	container.LearningToRank = s.v

	return container
}

func (s *_learningToRank) LearningToRankCaster() *types.LearningToRank {
	return s.v
}
