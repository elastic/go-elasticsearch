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
// https://github.com/elastic/elasticsearch-specification/tree/de4ff9ec1f716256f521d9e30011ad9c284b0dcc

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _binaryProperty struct {
	v *types.BinaryProperty
}

func NewBinaryProperty() *_binaryProperty {

	return &_binaryProperty{v: types.NewBinaryProperty()}

}

func (s *_binaryProperty) CopyTo(fields ...string) *_binaryProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_binaryProperty) DocValues(docvalues bool) *_binaryProperty {

	s.v.DocValues = &docvalues

	return s
}

func (s *_binaryProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_binaryProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_binaryProperty) Fields(fields map[string]types.Property) *_binaryProperty {

	s.v.Fields = fields
	return s
}

func (s *_binaryProperty) AddField(key string, value types.PropertyVariant) *_binaryProperty {

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

func (s *_binaryProperty) IgnoreAbove(ignoreabove int) *_binaryProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_binaryProperty) Meta(meta map[string]string) *_binaryProperty {

	s.v.Meta = meta
	return s
}

func (s *_binaryProperty) AddMeta(key string, value string) *_binaryProperty {

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

func (s *_binaryProperty) Properties(properties map[string]types.Property) *_binaryProperty {

	s.v.Properties = properties
	return s
}

func (s *_binaryProperty) AddProperty(key string, value types.PropertyVariant) *_binaryProperty {

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

func (s *_binaryProperty) Store(store bool) *_binaryProperty {

	s.v.Store = &store

	return s
}

func (s *_binaryProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_binaryProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_binaryProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_binaryProperty) BinaryPropertyCaster() *types.BinaryProperty {
	return s.v
}
