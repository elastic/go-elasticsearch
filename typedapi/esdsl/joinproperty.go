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
// https://github.com/elastic/elasticsearch-specification/tree/beeb1dc688bcc058488dcc45d9cbd2cd364e9943

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _joinProperty struct {
	v *types.JoinProperty
}

func NewJoinProperty() *_joinProperty {

	return &_joinProperty{v: types.NewJoinProperty()}

}

func (s *_joinProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_joinProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_joinProperty) EagerGlobalOrdinals(eagerglobalordinals bool) *_joinProperty {

	s.v.EagerGlobalOrdinals = &eagerglobalordinals

	return s
}

func (s *_joinProperty) Fields(fields map[string]types.Property) *_joinProperty {

	s.v.Fields = fields
	return s
}

func (s *_joinProperty) AddField(key string, value types.PropertyVariant) *_joinProperty {

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

func (s *_joinProperty) IgnoreAbove(ignoreabove int) *_joinProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_joinProperty) Meta(meta map[string]string) *_joinProperty {

	s.v.Meta = meta
	return s
}

func (s *_joinProperty) AddMeta(key string, value string) *_joinProperty {

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

func (s *_joinProperty) Properties(properties map[string]types.Property) *_joinProperty {

	s.v.Properties = properties
	return s
}

func (s *_joinProperty) AddProperty(key string, value types.PropertyVariant) *_joinProperty {

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

func (s *_joinProperty) Relations(relations map[string][]string) *_joinProperty {

	s.v.Relations = relations
	return s
}

func (s *_joinProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_joinProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_joinProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_joinProperty) JoinPropertyCaster() *types.JoinProperty {
	return s.v
}
