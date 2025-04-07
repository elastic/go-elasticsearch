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
// https://github.com/elastic/elasticsearch-specification/tree/60a81659be928bfe6cec53708c7f7613555a5eaf

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/syntheticsourcekeepenum"
)

type _constantKeywordProperty struct {
	v *types.ConstantKeywordProperty
}

func NewConstantKeywordProperty() *_constantKeywordProperty {

	return &_constantKeywordProperty{v: types.NewConstantKeywordProperty()}

}

func (s *_constantKeywordProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_constantKeywordProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_constantKeywordProperty) Fields(fields map[string]types.Property) *_constantKeywordProperty {

	s.v.Fields = fields
	return s
}

func (s *_constantKeywordProperty) AddField(key string, value types.PropertyVariant) *_constantKeywordProperty {

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

func (s *_constantKeywordProperty) IgnoreAbove(ignoreabove int) *_constantKeywordProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_constantKeywordProperty) Meta(meta map[string]string) *_constantKeywordProperty {

	s.v.Meta = meta
	return s
}

func (s *_constantKeywordProperty) AddMeta(key string, value string) *_constantKeywordProperty {

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

func (s *_constantKeywordProperty) Properties(properties map[string]types.Property) *_constantKeywordProperty {

	s.v.Properties = properties
	return s
}

func (s *_constantKeywordProperty) AddProperty(key string, value types.PropertyVariant) *_constantKeywordProperty {

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

func (s *_constantKeywordProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_constantKeywordProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_constantKeywordProperty) Value(value json.RawMessage) *_constantKeywordProperty {

	s.v.Value = value

	return s
}

func (s *_constantKeywordProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_constantKeywordProperty) ConstantKeywordPropertyCaster() *types.ConstantKeywordProperty {
	return s.v
}
