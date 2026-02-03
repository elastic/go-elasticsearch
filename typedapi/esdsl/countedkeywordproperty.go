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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _countedKeywordProperty struct {
	v *types.CountedKeywordProperty
}

func NewCountedKeywordProperty() *_countedKeywordProperty {

	return &_countedKeywordProperty{v: types.NewCountedKeywordProperty()}

}

func (s *_countedKeywordProperty) Index(index bool) *_countedKeywordProperty {

	s.v.Index = &index

	return s
}

func (s *_countedKeywordProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_countedKeywordProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_countedKeywordProperty) Fields(fields map[string]types.Property) *_countedKeywordProperty {

	s.v.Fields = fields
	return s
}

func (s *_countedKeywordProperty) AddField(key string, value types.PropertyVariant) *_countedKeywordProperty {

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

func (s *_countedKeywordProperty) IgnoreAbove(ignoreabove int) *_countedKeywordProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_countedKeywordProperty) Meta(meta map[string]string) *_countedKeywordProperty {

	s.v.Meta = meta
	return s
}

func (s *_countedKeywordProperty) AddMeta(key string, value string) *_countedKeywordProperty {

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

func (s *_countedKeywordProperty) Properties(properties map[string]types.Property) *_countedKeywordProperty {

	s.v.Properties = properties
	return s
}

func (s *_countedKeywordProperty) AddProperty(key string, value types.PropertyVariant) *_countedKeywordProperty {

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

func (s *_countedKeywordProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_countedKeywordProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_countedKeywordProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_countedKeywordProperty) CountedKeywordPropertyCaster() *types.CountedKeywordProperty {
	return s.v
}
