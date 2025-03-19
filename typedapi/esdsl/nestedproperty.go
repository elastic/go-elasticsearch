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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/syntheticsourcekeepenum"
)

type _nestedProperty struct {
	v *types.NestedProperty
}

func NewNestedProperty() *_nestedProperty {

	return &_nestedProperty{v: types.NewNestedProperty()}

}

func (s *_nestedProperty) CopyTo(fields ...string) *_nestedProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_nestedProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_nestedProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_nestedProperty) Enabled(enabled bool) *_nestedProperty {

	s.v.Enabled = &enabled

	return s
}

func (s *_nestedProperty) Fields(fields map[string]types.Property) *_nestedProperty {

	s.v.Fields = fields
	return s
}

func (s *_nestedProperty) AddField(key string, value types.PropertyVariant) *_nestedProperty {

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

func (s *_nestedProperty) IgnoreAbove(ignoreabove int) *_nestedProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_nestedProperty) IncludeInParent(includeinparent bool) *_nestedProperty {

	s.v.IncludeInParent = &includeinparent

	return s
}

func (s *_nestedProperty) IncludeInRoot(includeinroot bool) *_nestedProperty {

	s.v.IncludeInRoot = &includeinroot

	return s
}

// Metadata about the field.
func (s *_nestedProperty) Meta(meta map[string]string) *_nestedProperty {

	s.v.Meta = meta
	return s
}

func (s *_nestedProperty) AddMeta(key string, value string) *_nestedProperty {

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

func (s *_nestedProperty) Properties(properties map[string]types.Property) *_nestedProperty {

	s.v.Properties = properties
	return s
}

func (s *_nestedProperty) AddProperty(key string, value types.PropertyVariant) *_nestedProperty {

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

func (s *_nestedProperty) Store(store bool) *_nestedProperty {

	s.v.Store = &store

	return s
}

func (s *_nestedProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_nestedProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_nestedProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_nestedProperty) NestedPropertyCaster() *types.NestedProperty {
	return s.v
}
