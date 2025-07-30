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
// https://github.com/elastic/elasticsearch-specification/tree/de4ff9ec1f716256f521d9e30011ad9c284b0dcc

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _rescore struct {
	v *types.Rescore
}

func NewRescore() *_rescore {
	return &_rescore{v: types.NewRescore()}
}

// AdditionalRescoreProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_rescore) AdditionalRescoreProperty(key string, value json.RawMessage) *_rescore {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalRescoreProperty = tmp
	return s
}

func (s *_rescore) LearningToRank(learningtorank types.LearningToRankVariant) *_rescore {

	s.v.LearningToRank = learningtorank.LearningToRankCaster()

	return s
}

func (s *_rescore) Query(query types.RescoreQueryVariant) *_rescore {

	s.v.Query = query.RescoreQueryCaster()

	return s
}

func (s *_rescore) WindowSize(windowsize int) *_rescore {

	s.v.WindowSize = &windowsize

	return s
}

func (s *_rescore) RescoreCaster() *types.Rescore {
	return s.v
}
