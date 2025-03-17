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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _indexTemplateMapping struct {
	v *types.IndexTemplateMapping
}

func NewIndexTemplateMapping() *_indexTemplateMapping {

	return &_indexTemplateMapping{v: types.NewIndexTemplateMapping()}

}

// Aliases to add.
// If the index template includes a `data_stream` object, these are data stream
// aliases.
// Otherwise, these are index aliases.
// Data stream aliases ignore the `index_routing`, `routing`, and
// `search_routing` options.
func (s *_indexTemplateMapping) Aliases(aliases map[string]types.Alias) *_indexTemplateMapping {

	s.v.Aliases = aliases
	return s
}

func (s *_indexTemplateMapping) AddAlias(key string, value types.AliasVariant) *_indexTemplateMapping {

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

func (s *_indexTemplateMapping) Lifecycle(lifecycle types.DataStreamLifecycleVariant) *_indexTemplateMapping {

	s.v.Lifecycle = lifecycle.DataStreamLifecycleCaster()

	return s
}

// Mapping for fields in the index.
// If specified, this mapping can include field names, field data types, and
// mapping parameters.
func (s *_indexTemplateMapping) Mappings(mappings types.TypeMappingVariant) *_indexTemplateMapping {

	s.v.Mappings = mappings.TypeMappingCaster()

	return s
}

// Configuration options for the index.
func (s *_indexTemplateMapping) Settings(settings types.IndexSettingsVariant) *_indexTemplateMapping {

	s.v.Settings = settings.IndexSettingsCaster()

	return s
}

func (s *_indexTemplateMapping) IndexTemplateMappingCaster() *types.IndexTemplateMapping {
	return s.v
}
