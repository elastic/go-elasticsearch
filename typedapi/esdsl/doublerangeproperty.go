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

type _doubleRangeProperty struct {
	v *types.DoubleRangeProperty
}

func NewDoubleRangeProperty() *_doubleRangeProperty {

	return &_doubleRangeProperty{v: types.NewDoubleRangeProperty()}

}

func (s *_doubleRangeProperty) Boost(boost types.Float64) *_doubleRangeProperty {

	s.v.Boost = &boost

	return s
}

func (s *_doubleRangeProperty) Coerce(coerce bool) *_doubleRangeProperty {

	s.v.Coerce = &coerce

	return s
}

func (s *_doubleRangeProperty) CopyTo(fields ...string) *_doubleRangeProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_doubleRangeProperty) DocValues(docvalues bool) *_doubleRangeProperty {

	s.v.DocValues = &docvalues

	return s
}

func (s *_doubleRangeProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_doubleRangeProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_doubleRangeProperty) Fields(fields map[string]types.Property) *_doubleRangeProperty {

	s.v.Fields = fields
	return s
}

func (s *_doubleRangeProperty) AddField(key string, value types.PropertyVariant) *_doubleRangeProperty {

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

func (s *_doubleRangeProperty) IgnoreAbove(ignoreabove int) *_doubleRangeProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_doubleRangeProperty) Index(index bool) *_doubleRangeProperty {

	s.v.Index = &index

	return s
}

func (s *_doubleRangeProperty) Meta(meta map[string]string) *_doubleRangeProperty {

	s.v.Meta = meta
	return s
}

func (s *_doubleRangeProperty) AddMeta(key string, value string) *_doubleRangeProperty {

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

func (s *_doubleRangeProperty) Properties(properties map[string]types.Property) *_doubleRangeProperty {

	s.v.Properties = properties
	return s
}

func (s *_doubleRangeProperty) AddProperty(key string, value types.PropertyVariant) *_doubleRangeProperty {

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

func (s *_doubleRangeProperty) Store(store bool) *_doubleRangeProperty {

	s.v.Store = &store

	return s
}

func (s *_doubleRangeProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_doubleRangeProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_doubleRangeProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_doubleRangeProperty) DoubleRangePropertyCaster() *types.DoubleRangeProperty {
	return s.v
}
