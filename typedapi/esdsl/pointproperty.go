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

type _pointProperty struct {
	v *types.PointProperty
}

func NewPointProperty() *_pointProperty {

	return &_pointProperty{v: types.NewPointProperty()}

}

func (s *_pointProperty) IgnoreMalformed(ignoremalformed bool) *_pointProperty {

	s.v.IgnoreMalformed = &ignoremalformed

	return s
}

func (s *_pointProperty) IgnoreZValue(ignorezvalue bool) *_pointProperty {

	s.v.IgnoreZValue = &ignorezvalue

	return s
}

func (s *_pointProperty) NullValue(nullvalue string) *_pointProperty {

	s.v.NullValue = &nullvalue

	return s
}

func (s *_pointProperty) CopyTo(fields ...string) *_pointProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_pointProperty) DocValues(docvalues bool) *_pointProperty {

	s.v.DocValues = &docvalues

	return s
}

func (s *_pointProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_pointProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_pointProperty) Fields(fields map[string]types.Property) *_pointProperty {

	s.v.Fields = fields
	return s
}

func (s *_pointProperty) AddField(key string, value types.PropertyVariant) *_pointProperty {

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

func (s *_pointProperty) IgnoreAbove(ignoreabove int) *_pointProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_pointProperty) Meta(meta map[string]string) *_pointProperty {

	s.v.Meta = meta
	return s
}

func (s *_pointProperty) AddMeta(key string, value string) *_pointProperty {

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

func (s *_pointProperty) Properties(properties map[string]types.Property) *_pointProperty {

	s.v.Properties = properties
	return s
}

func (s *_pointProperty) AddProperty(key string, value types.PropertyVariant) *_pointProperty {

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

func (s *_pointProperty) Store(store bool) *_pointProperty {

	s.v.Store = &store

	return s
}

func (s *_pointProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_pointProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_pointProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_pointProperty) PointPropertyCaster() *types.PointProperty {
	return s.v
}
