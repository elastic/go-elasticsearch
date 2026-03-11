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
// https://github.com/elastic/elasticsearch-specification/tree/55f8d05b44cea956ae5ceddfcb02770ea2a24ff6

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/rankvectorelementtype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _rankVectorProperty struct {
	v *types.RankVectorProperty
}

func NewRankVectorProperty() *_rankVectorProperty {

	return &_rankVectorProperty{v: types.NewRankVectorProperty()}

}

func (s *_rankVectorProperty) Dims(dims int) *_rankVectorProperty {

	s.v.Dims = &dims

	return s
}

func (s *_rankVectorProperty) ElementType(elementtype rankvectorelementtype.RankVectorElementType) *_rankVectorProperty {

	s.v.ElementType = &elementtype
	return s
}

func (s *_rankVectorProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_rankVectorProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_rankVectorProperty) Fields(fields map[string]types.Property) *_rankVectorProperty {

	s.v.Fields = fields
	return s
}

func (s *_rankVectorProperty) AddField(key string, value types.PropertyVariant) *_rankVectorProperty {

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

func (s *_rankVectorProperty) IgnoreAbove(ignoreabove int) *_rankVectorProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_rankVectorProperty) Meta(meta map[string]string) *_rankVectorProperty {

	s.v.Meta = meta
	return s
}

func (s *_rankVectorProperty) AddMeta(key string, value string) *_rankVectorProperty {

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

func (s *_rankVectorProperty) Properties(properties map[string]types.Property) *_rankVectorProperty {

	s.v.Properties = properties
	return s
}

func (s *_rankVectorProperty) AddProperty(key string, value types.PropertyVariant) *_rankVectorProperty {

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

func (s *_rankVectorProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_rankVectorProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_rankVectorProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_rankVectorProperty) RankVectorPropertyCaster() *types.RankVectorProperty {
	return s.v
}
