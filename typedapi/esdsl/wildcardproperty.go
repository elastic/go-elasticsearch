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

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _wildcardProperty struct {
	v *types.WildcardProperty
}

func NewWildcardProperty() *_wildcardProperty {

	return &_wildcardProperty{v: types.NewWildcardProperty()}

}

func (s *_wildcardProperty) NullValue(nullvalue string) *_wildcardProperty {

	s.v.NullValue = &nullvalue

	return s
}

func (s *_wildcardProperty) CopyTo(fields ...string) *_wildcardProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_wildcardProperty) DocValues(docvalues bool) *_wildcardProperty {

	s.v.DocValues = &docvalues

	return s
}

func (s *_wildcardProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_wildcardProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_wildcardProperty) Fields(fields map[string]types.Property) *_wildcardProperty {

	s.v.Fields = fields
	return s
}

func (s *_wildcardProperty) AddField(key string, value types.PropertyVariant) *_wildcardProperty {

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

func (s *_wildcardProperty) IgnoreAbove(ignoreabove int) *_wildcardProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_wildcardProperty) Meta(meta map[string]string) *_wildcardProperty {

	s.v.Meta = meta
	return s
}

func (s *_wildcardProperty) AddMeta(key string, value string) *_wildcardProperty {

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

func (s *_wildcardProperty) Properties(properties map[string]types.Property) *_wildcardProperty {

	s.v.Properties = properties
	return s
}

func (s *_wildcardProperty) AddProperty(key string, value types.PropertyVariant) *_wildcardProperty {

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

func (s *_wildcardProperty) Store(store bool) *_wildcardProperty {

	s.v.Store = &store

	return s
}

func (s *_wildcardProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_wildcardProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_wildcardProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_wildcardProperty) WildcardPropertyCaster() *types.WildcardProperty {
	return s.v
}
