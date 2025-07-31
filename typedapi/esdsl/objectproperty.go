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
// https://github.com/elastic/elasticsearch-specification/tree/86f41834c7bb975159a38a73be8a9d930010d673

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/subobjects"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _objectProperty struct {
	v *types.ObjectProperty
}

func NewObjectProperty() *_objectProperty {

	return &_objectProperty{v: types.NewObjectProperty()}

}

func (s *_objectProperty) CopyTo(fields ...string) *_objectProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_objectProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_objectProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_objectProperty) Enabled(enabled bool) *_objectProperty {

	s.v.Enabled = &enabled

	return s
}

func (s *_objectProperty) Fields(fields map[string]types.Property) *_objectProperty {

	s.v.Fields = fields
	return s
}

func (s *_objectProperty) AddField(key string, value types.PropertyVariant) *_objectProperty {

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

func (s *_objectProperty) IgnoreAbove(ignoreabove int) *_objectProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_objectProperty) Meta(meta map[string]string) *_objectProperty {

	s.v.Meta = meta
	return s
}

func (s *_objectProperty) AddMeta(key string, value string) *_objectProperty {

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

func (s *_objectProperty) Properties(properties map[string]types.Property) *_objectProperty {

	s.v.Properties = properties
	return s
}

func (s *_objectProperty) AddProperty(key string, value types.PropertyVariant) *_objectProperty {

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

func (s *_objectProperty) Store(store bool) *_objectProperty {

	s.v.Store = &store

	return s
}

func (s *_objectProperty) Subobjects(subobjects subobjects.Subobjects) *_objectProperty {

	s.v.Subobjects = &subobjects
	return s
}

func (s *_objectProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_objectProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_objectProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_objectProperty) ObjectPropertyCaster() *types.ObjectProperty {
	return s.v
}
