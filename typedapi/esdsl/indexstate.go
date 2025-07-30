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
// https://github.com/elastic/elasticsearch-specification/tree/e585438d116b00ff34643179e6286e402c0bcaaf

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _indexState struct {
	v *types.IndexState
}

func NewIndexState() *_indexState {

	return &_indexState{v: types.NewIndexState()}

}

func (s *_indexState) Aliases(aliases map[string]types.Alias) *_indexState {

	s.v.Aliases = aliases
	return s
}

func (s *_indexState) AddAlias(key string, value types.AliasVariant) *_indexState {

	var tmp map[string]types.Alias
	if s.v.Aliases == nil {
		s.v.Aliases = make(map[string]types.Alias)
	} else {
		tmp = s.v.Aliases
	}

	tmp[key] = *value.AliasCaster()

	s.v.Aliases = tmp
	return s
}

func (s *_indexState) DataStream(datastreamname string) *_indexState {

	s.v.DataStream = &datastreamname

	return s
}

func (s *_indexState) Defaults(defaults types.IndexSettingsVariant) *_indexState {

	s.v.Defaults = defaults.IndexSettingsCaster()

	return s
}

func (s *_indexState) Lifecycle(lifecycle types.DataStreamLifecycleVariant) *_indexState {

	s.v.Lifecycle = lifecycle.DataStreamLifecycleCaster()

	return s
}

func (s *_indexState) Mappings(mappings types.TypeMappingVariant) *_indexState {

	s.v.Mappings = mappings.TypeMappingCaster()

	return s
}

func (s *_indexState) Settings(settings types.IndexSettingsVariant) *_indexState {

	s.v.Settings = settings.IndexSettingsCaster()

	return s
}

func (s *_indexState) IndexStateCaster() *types.IndexState {
	return s.v
}
