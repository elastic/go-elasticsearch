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

type _semanticTextProperty struct {
	v *types.SemanticTextProperty
}

func NewSemanticTextProperty() *_semanticTextProperty {

	return &_semanticTextProperty{v: types.NewSemanticTextProperty()}

}

func (s *_semanticTextProperty) ChunkingSettings(chunkingsettings types.ChunkingSettingsVariant) *_semanticTextProperty {

	s.v.ChunkingSettings = chunkingsettings.ChunkingSettingsCaster()

	return s
}

func (s *_semanticTextProperty) Fields(fields map[string]types.Property) *_semanticTextProperty {

	s.v.Fields = fields
	return s
}

func (s *_semanticTextProperty) AddField(key string, value types.PropertyVariant) *_semanticTextProperty {

	var tmp map[string]types.Property
	if s.v.Fields == nil {
		s.v.Fields = make(map[string]types.Property)
	} else {
		tmp = s.v.Fields
	}

	tmp[key] = *value.PropertyCaster()

	s.v.Fields = tmp
	return s
}

func (s *_semanticTextProperty) IndexOptions(indexoptions types.SemanticTextIndexOptionsVariant) *_semanticTextProperty {

	s.v.IndexOptions = indexoptions.SemanticTextIndexOptionsCaster()

	return s
}

func (s *_semanticTextProperty) InferenceId(id string) *_semanticTextProperty {

	s.v.InferenceId = &id

	return s
}

func (s *_semanticTextProperty) Meta(meta map[string]string) *_semanticTextProperty {

	s.v.Meta = meta
	return s
}

func (s *_semanticTextProperty) AddMeta(key string, value string) *_semanticTextProperty {

	var tmp map[string]string
	if s.v.Meta == nil {
		s.v.Meta = make(map[string]string)
	} else {
		tmp = s.v.Meta
	}

	tmp[key] = value

	s.v.Meta = tmp
	return s
}

func (s *_semanticTextProperty) SearchInferenceId(id string) *_semanticTextProperty {

	s.v.SearchInferenceId = &id

	return s
}

func (s *_semanticTextProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_semanticTextProperty) SemanticTextPropertyCaster() *types.SemanticTextProperty {
	return s.v
}
