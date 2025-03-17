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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _createFrom struct {
	v *types.CreateFrom
}

func NewCreateFrom() *_createFrom {

	return &_createFrom{v: types.NewCreateFrom()}

}

// Mappings overrides to be applied to the destination index (optional)
func (s *_createFrom) MappingsOverride(mappingsoverride types.TypeMappingVariant) *_createFrom {

	s.v.MappingsOverride = mappingsoverride.TypeMappingCaster()

	return s
}

// If index blocks should be removed when creating destination index (optional)
func (s *_createFrom) RemoveIndexBlocks(removeindexblocks bool) *_createFrom {

	s.v.RemoveIndexBlocks = &removeindexblocks

	return s
}

// Settings overrides to be applied to the destination index (optional)
func (s *_createFrom) SettingsOverride(settingsoverride types.IndexSettingsVariant) *_createFrom {

	s.v.SettingsOverride = settingsoverride.IndexSettingsCaster()

	return s
}

func (s *_createFrom) CreateFromCaster() *types.CreateFrom {
	return s.v
}
