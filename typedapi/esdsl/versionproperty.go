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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/syntheticsourcekeepenum"
)

type _versionProperty struct {
	v *types.VersionProperty
}

func NewVersionProperty() *_versionProperty {

	return &_versionProperty{v: types.NewVersionProperty()}

}

func (s *_versionProperty) CopyTo(fields ...string) *_versionProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_versionProperty) DocValues(docvalues bool) *_versionProperty {

	s.v.DocValues = &docvalues

	return s
}

func (s *_versionProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_versionProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_versionProperty) Fields(fields map[string]types.Property) *_versionProperty {

	s.v.Fields = fields
	return s
}

func (s *_versionProperty) AddField(key string, value types.PropertyVariant) *_versionProperty {

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

func (s *_versionProperty) IgnoreAbove(ignoreabove int) *_versionProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

// Metadata about the field.
func (s *_versionProperty) Meta(meta map[string]string) *_versionProperty {

	s.v.Meta = meta
	return s
}

func (s *_versionProperty) AddMeta(key string, value string) *_versionProperty {

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

func (s *_versionProperty) Properties(properties map[string]types.Property) *_versionProperty {

	s.v.Properties = properties
	return s
}

func (s *_versionProperty) AddProperty(key string, value types.PropertyVariant) *_versionProperty {

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

func (s *_versionProperty) Store(store bool) *_versionProperty {

	s.v.Store = &store

	return s
}

func (s *_versionProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_versionProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_versionProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_versionProperty) VersionPropertyCaster() *types.VersionProperty {
	return s.v
}
