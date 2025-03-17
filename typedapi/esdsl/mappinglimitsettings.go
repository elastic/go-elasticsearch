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

type _mappingLimitSettings struct {
	v *types.MappingLimitSettings
}

func NewMappingLimitSettings() *_mappingLimitSettings {

	return &_mappingLimitSettings{v: types.NewMappingLimitSettings()}

}

func (s *_mappingLimitSettings) Coerce(coerce bool) *_mappingLimitSettings {

	s.v.Coerce = &coerce

	return s
}

func (s *_mappingLimitSettings) Depth(depth types.MappingLimitSettingsDepthVariant) *_mappingLimitSettings {

	s.v.Depth = depth.MappingLimitSettingsDepthCaster()

	return s
}

func (s *_mappingLimitSettings) DimensionFields(dimensionfields types.MappingLimitSettingsDimensionFieldsVariant) *_mappingLimitSettings {

	s.v.DimensionFields = dimensionfields.MappingLimitSettingsDimensionFieldsCaster()

	return s
}

func (s *_mappingLimitSettings) FieldNameLength(fieldnamelength types.MappingLimitSettingsFieldNameLengthVariant) *_mappingLimitSettings {

	s.v.FieldNameLength = fieldnamelength.MappingLimitSettingsFieldNameLengthCaster()

	return s
}

func (s *_mappingLimitSettings) IgnoreMalformed(ignoremalformed string) *_mappingLimitSettings {

	s.v.IgnoreMalformed = ignoremalformed

	return s
}

func (s *_mappingLimitSettings) NestedFields(nestedfields types.MappingLimitSettingsNestedFieldsVariant) *_mappingLimitSettings {

	s.v.NestedFields = nestedfields.MappingLimitSettingsNestedFieldsCaster()

	return s
}

func (s *_mappingLimitSettings) NestedObjects(nestedobjects types.MappingLimitSettingsNestedObjectsVariant) *_mappingLimitSettings {

	s.v.NestedObjects = nestedobjects.MappingLimitSettingsNestedObjectsCaster()

	return s
}

func (s *_mappingLimitSettings) Source(source types.MappingLimitSettingsSourceFieldsVariant) *_mappingLimitSettings {

	s.v.Source = source.MappingLimitSettingsSourceFieldsCaster()

	return s
}

func (s *_mappingLimitSettings) TotalFields(totalfields types.MappingLimitSettingsTotalFieldsVariant) *_mappingLimitSettings {

	s.v.TotalFields = totalfields.MappingLimitSettingsTotalFieldsCaster()

	return s
}

func (s *_mappingLimitSettings) MappingLimitSettingsCaster() *types.MappingLimitSettings {
	return s.v
}
