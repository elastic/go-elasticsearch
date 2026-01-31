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
// https://github.com/elastic/elasticsearch-specification/tree/6785a6caa1fa3ca5ab3308963d79dce923a3469f

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _indexTemplateSummary struct {
	v *types.IndexTemplateSummary
}

func NewIndexTemplateSummary() *_indexTemplateSummary {

	return &_indexTemplateSummary{v: types.NewIndexTemplateSummary()}

}

func (s *_indexTemplateSummary) Aliases(aliases map[string]types.Alias) *_indexTemplateSummary {

	s.v.Aliases = aliases
	return s
}

func (s *_indexTemplateSummary) AddAlias(key string, value types.AliasVariant) *_indexTemplateSummary {

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

func (s *_indexTemplateSummary) DataStreamOptions(datastreamoptions types.DataStreamOptionsVariant) *_indexTemplateSummary {

	s.v.DataStreamOptions = datastreamoptions.DataStreamOptionsCaster()

	return s
}

func (s *_indexTemplateSummary) Lifecycle(lifecycle types.DataStreamLifecycleWithRolloverVariant) *_indexTemplateSummary {

	s.v.Lifecycle = lifecycle.DataStreamLifecycleWithRolloverCaster()

	return s
}

func (s *_indexTemplateSummary) Mappings(mappings types.TypeMappingVariant) *_indexTemplateSummary {

	s.v.Mappings = mappings.TypeMappingCaster()

	return s
}

func (s *_indexTemplateSummary) Settings(settings types.IndexSettingsVariant) *_indexTemplateSummary {

	s.v.Settings = settings.IndexSettingsCaster()

	return s
}

func (s *_indexTemplateSummary) IndexTemplateSummaryCaster() *types.IndexTemplateSummary {
	return s.v
}
