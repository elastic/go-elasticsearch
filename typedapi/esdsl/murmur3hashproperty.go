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
// https://github.com/elastic/elasticsearch-specification/tree/b1811e10a0722431d79d1c234dd412ff47d8656f

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _murmur3HashProperty struct {
	v *types.Murmur3HashProperty
}

func NewMurmur3HashProperty() *_murmur3HashProperty {

	return &_murmur3HashProperty{v: types.NewMurmur3HashProperty()}

}

func (s *_murmur3HashProperty) CopyTo(fields ...string) *_murmur3HashProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_murmur3HashProperty) DocValues(docvalues bool) *_murmur3HashProperty {

	s.v.DocValues = &docvalues

	return s
}

func (s *_murmur3HashProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_murmur3HashProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_murmur3HashProperty) Fields(fields map[string]types.Property) *_murmur3HashProperty {

	s.v.Fields = fields
	return s
}

func (s *_murmur3HashProperty) AddField(key string, value types.PropertyVariant) *_murmur3HashProperty {

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

func (s *_murmur3HashProperty) IgnoreAbove(ignoreabove int) *_murmur3HashProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_murmur3HashProperty) Meta(meta map[string]string) *_murmur3HashProperty {

	s.v.Meta = meta
	return s
}

func (s *_murmur3HashProperty) AddMeta(key string, value string) *_murmur3HashProperty {

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

func (s *_murmur3HashProperty) Properties(properties map[string]types.Property) *_murmur3HashProperty {

	s.v.Properties = properties
	return s
}

func (s *_murmur3HashProperty) AddProperty(key string, value types.PropertyVariant) *_murmur3HashProperty {

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

func (s *_murmur3HashProperty) Store(store bool) *_murmur3HashProperty {

	s.v.Store = &store

	return s
}

func (s *_murmur3HashProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_murmur3HashProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_murmur3HashProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_murmur3HashProperty) Murmur3HashPropertyCaster() *types.Murmur3HashProperty {
	return s.v
}
