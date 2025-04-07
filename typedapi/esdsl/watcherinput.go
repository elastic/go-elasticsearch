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

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _watcherInput struct {
	v *types.WatcherInput
}

func NewWatcherInput() *_watcherInput {
	return &_watcherInput{v: types.NewWatcherInput()}
}

// AdditionalWatcherInputProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_watcherInput) AdditionalWatcherInputProperty(key string, value json.RawMessage) *_watcherInput {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalWatcherInputProperty = tmp
	return s
}

func (s *_watcherInput) Chain(chain types.ChainInputVariant) *_watcherInput {

	s.v.Chain = chain.ChainInputCaster()

	return s
}

func (s *_watcherInput) Http(http types.HttpInputVariant) *_watcherInput {

	s.v.Http = http.HttpInputCaster()

	return s
}

func (s *_watcherInput) Search(search types.SearchInputVariant) *_watcherInput {

	s.v.Search = search.SearchInputCaster()

	return s
}

func (s *_watcherInput) Simple(simple map[string]json.RawMessage) *_watcherInput {

	s.v.Simple = simple
	return s
}

func (s *_watcherInput) AddSimple(key string, value json.RawMessage) *_watcherInput {

	var tmp map[string]json.RawMessage
	if s.v.Simple == nil {
		s.v.Simple = make(map[string]json.RawMessage)
	} else {
		tmp = s.v.Simple
	}

	tmp[key] = value

	s.v.Simple = tmp
	return s
}

func (s *_watcherInput) WatcherInputCaster() *types.WatcherInput {
	return s.v
}
