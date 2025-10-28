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
// https://github.com/elastic/elasticsearch-specification/tree/d520d9e8cf14cad487de5e0654007686c395b494

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _rankContainer struct {
	v *types.RankContainer
}

func NewRankContainer() *_rankContainer {
	return &_rankContainer{v: types.NewRankContainer()}
}

// AdditionalRankContainerProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_rankContainer) AdditionalRankContainerProperty(key string, value json.RawMessage) *_rankContainer {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalRankContainerProperty = tmp
	return s
}

func (s *_rankContainer) Rrf(rrf types.RrfRankVariant) *_rankContainer {

	s.v.Rrf = rrf.RrfRankCaster()

	return s
}

func (s *_rankContainer) RankContainerCaster() *types.RankContainer {
	return s.v
}
