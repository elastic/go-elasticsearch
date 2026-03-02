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
// https://github.com/elastic/elasticsearch-specification/tree/e196f9953fa743572ee46884835f1934bce9a16b

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _componentTemplateSummary struct {
	v *types.ComponentTemplateSummary
}

func NewComponentTemplateSummary() *_componentTemplateSummary {

	return &_componentTemplateSummary{v: types.NewComponentTemplateSummary()}

}

func (s *_componentTemplateSummary) Aliases(aliases map[string]types.AliasDefinition) *_componentTemplateSummary {

	s.v.Aliases = aliases
	return s
}

func (s *_componentTemplateSummary) AddAlias(key string, value types.AliasDefinitionVariant) *_componentTemplateSummary {

	var tmp map[string]types.AliasDefinition
	if s.v.Aliases == nil {
		s.v.Aliases = make(map[string]types.AliasDefinition)
	} else {
		tmp = s.v.Aliases
	}

	tmp[key] = *value.AliasDefinitionCaster()

	s.v.Aliases = tmp
	return s
}

func (s *_componentTemplateSummary) DataStreamOptions(datastreamoptions types.DataStreamOptionsVariant) *_componentTemplateSummary {

	s.v.DataStreamOptions = datastreamoptions.DataStreamOptionsCaster()

	return s
}

func (s *_componentTemplateSummary) Lifecycle(lifecycle types.DataStreamLifecycleWithRolloverVariant) *_componentTemplateSummary {

	s.v.Lifecycle = lifecycle.DataStreamLifecycleWithRolloverCaster()

	return s
}

func (s *_componentTemplateSummary) Mappings(mappings types.TypeMappingVariant) *_componentTemplateSummary {

	s.v.Mappings = mappings.TypeMappingCaster()

	return s
}

func (s *_componentTemplateSummary) Meta_(metadata types.MetadataVariant) *_componentTemplateSummary {

	s.v.Meta_ = *metadata.MetadataCaster()

	return s
}

func (s *_componentTemplateSummary) Settings(settings map[string]types.IndexSettings) *_componentTemplateSummary {

	s.v.Settings = settings
	return s
}

func (s *_componentTemplateSummary) AddSetting(key string, value types.IndexSettingsVariant) *_componentTemplateSummary {

	var tmp map[string]types.IndexSettings
	if s.v.Settings == nil {
		s.v.Settings = make(map[string]types.IndexSettings)
	} else {
		tmp = s.v.Settings
	}

	tmp[key] = *value.IndexSettingsCaster()

	s.v.Settings = tmp
	return s
}

func (s *_componentTemplateSummary) Version(versionnumber int64) *_componentTemplateSummary {

	s.v.Version = &versionnumber

	return s
}

func (s *_componentTemplateSummary) ComponentTemplateSummaryCaster() *types.ComponentTemplateSummary {
	return s.v
}
