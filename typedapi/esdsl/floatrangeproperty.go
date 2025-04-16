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
// https://github.com/elastic/elasticsearch-specification/tree/52c473efb1fb5320a5bac12572d0b285882862fb

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _floatRangeProperty struct {
	v *types.FloatRangeProperty
}

func NewFloatRangeProperty() *_floatRangeProperty {

	return &_floatRangeProperty{v: types.NewFloatRangeProperty()}

}

func (s *_floatRangeProperty) Boost(boost types.Float64) *_floatRangeProperty {

	s.v.Boost = &boost

	return s
}

func (s *_floatRangeProperty) Coerce(coerce bool) *_floatRangeProperty {

	s.v.Coerce = &coerce

	return s
}

func (s *_floatRangeProperty) CopyTo(fields ...string) *_floatRangeProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_floatRangeProperty) DocValues(docvalues bool) *_floatRangeProperty {

	s.v.DocValues = &docvalues

	return s
}

func (s *_floatRangeProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_floatRangeProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_floatRangeProperty) Fields(fields map[string]types.Property) *_floatRangeProperty {

	s.v.Fields = fields
	return s
}

func (s *_floatRangeProperty) AddField(key string, value types.PropertyVariant) *_floatRangeProperty {

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

func (s *_floatRangeProperty) IgnoreAbove(ignoreabove int) *_floatRangeProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_floatRangeProperty) Index(index bool) *_floatRangeProperty {

	s.v.Index = &index

	return s
}

func (s *_floatRangeProperty) Meta(meta map[string]string) *_floatRangeProperty {

	s.v.Meta = meta
	return s
}

func (s *_floatRangeProperty) AddMeta(key string, value string) *_floatRangeProperty {

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

func (s *_floatRangeProperty) Properties(properties map[string]types.Property) *_floatRangeProperty {

	s.v.Properties = properties
	return s
}

func (s *_floatRangeProperty) AddProperty(key string, value types.PropertyVariant) *_floatRangeProperty {

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

func (s *_floatRangeProperty) Store(store bool) *_floatRangeProperty {

	s.v.Store = &store

	return s
}

func (s *_floatRangeProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_floatRangeProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_floatRangeProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_floatRangeProperty) FloatRangePropertyCaster() *types.FloatRangeProperty {
	return s.v
}
