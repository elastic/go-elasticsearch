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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/subobjects"
)

type _typeMapping struct {
	v *types.TypeMapping
}

func NewTypeMapping() *_typeMapping {

	return &_typeMapping{v: types.NewTypeMapping()}

}

func (s *_typeMapping) AllField(allfield types.AllFieldVariant) *_typeMapping {

	s.v.AllField = allfield.AllFieldCaster()

	return s
}

func (s *_typeMapping) DataStreamTimestamp_(datastreamtimestamp_ types.DataStreamTimestampVariant) *_typeMapping {

	s.v.DataStreamTimestamp_ = datastreamtimestamp_.DataStreamTimestampCaster()

	return s
}

func (s *_typeMapping) DateDetection(datedetection bool) *_typeMapping {

	s.v.DateDetection = &datedetection

	return s
}

func (s *_typeMapping) Dynamic(dynamic dynamicmapping.DynamicMapping) *_typeMapping {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_typeMapping) DynamicDateFormats(dynamicdateformats ...string) *_typeMapping {

	for _, v := range dynamicdateformats {

		s.v.DynamicDateFormats = append(s.v.DynamicDateFormats, v)

	}
	return s
}

func (s *_typeMapping) DynamicTemplates(dynamictemplates []map[string]types.DynamicTemplate) *_typeMapping {

	s.v.DynamicTemplates = dynamictemplates

	return s
}

func (s *_typeMapping) Enabled(enabled bool) *_typeMapping {

	s.v.Enabled = &enabled

	return s
}

func (s *_typeMapping) FieldNames_(fieldnames_ types.FieldNamesFieldVariant) *_typeMapping {

	s.v.FieldNames_ = fieldnames_.FieldNamesFieldCaster()

	return s
}

func (s *_typeMapping) IndexField(indexfield types.IndexFieldVariant) *_typeMapping {

	s.v.IndexField = indexfield.IndexFieldCaster()

	return s
}

func (s *_typeMapping) Meta_(metadata types.MetadataVariant) *_typeMapping {

	s.v.Meta_ = *metadata.MetadataCaster()

	return s
}

func (s *_typeMapping) NumericDetection(numericdetection bool) *_typeMapping {

	s.v.NumericDetection = &numericdetection

	return s
}

func (s *_typeMapping) Properties(properties map[string]types.Property) *_typeMapping {

	s.v.Properties = properties
	return s
}

func (s *_typeMapping) AddProperty(key string, value types.PropertyVariant) *_typeMapping {

	var tmp map[string]types.Property
	if s.v.Properties == nil {
		s.v.Properties = make(map[string]types.Property)
	} else {
		tmp = s.v.Properties
	}

	tmp[key] = *value.PropertyCaster()

	s.v.Properties = tmp
	return s
}

func (s *_typeMapping) Routing_(routing_ types.RoutingFieldVariant) *_typeMapping {

	s.v.Routing_ = routing_.RoutingFieldCaster()

	return s
}

func (s *_typeMapping) Runtime(runtime map[string]types.RuntimeField) *_typeMapping {

	s.v.Runtime = runtime
	return s
}

func (s *_typeMapping) AddRuntime(key string, value types.RuntimeFieldVariant) *_typeMapping {

	var tmp map[string]types.RuntimeField
	if s.v.Runtime == nil {
		s.v.Runtime = make(map[string]types.RuntimeField)
	} else {
		tmp = s.v.Runtime
	}

	tmp[key] = *value.RuntimeFieldCaster()

	s.v.Runtime = tmp
	return s
}

func (s *_typeMapping) Size_(size_ types.SizeFieldVariant) *_typeMapping {

	s.v.Size_ = size_.SizeFieldCaster()

	return s
}

func (s *_typeMapping) Source_(source_ types.SourceFieldVariant) *_typeMapping {

	s.v.Source_ = source_.SourceFieldCaster()

	return s
}

func (s *_typeMapping) Subobjects(subobjects subobjects.Subobjects) *_typeMapping {

	s.v.Subobjects = &subobjects
	return s
}

func (s *_typeMapping) TypeMappingCaster() *types.TypeMapping {
	return s.v
}
