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
// https://github.com/elastic/elasticsearch-specification/tree/d520d9e8cf14cad487de5e0654007686c395b494

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _integerRangeProperty struct {
	v *types.IntegerRangeProperty
}

func NewIntegerRangeProperty() *_integerRangeProperty {

	return &_integerRangeProperty{v: types.NewIntegerRangeProperty()}

}

func (s *_integerRangeProperty) Boost(boost types.Float64) *_integerRangeProperty {

	s.v.Boost = &boost

	return s
}

func (s *_integerRangeProperty) Coerce(coerce bool) *_integerRangeProperty {

	s.v.Coerce = &coerce

	return s
}

func (s *_integerRangeProperty) CopyTo(fields ...string) *_integerRangeProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_integerRangeProperty) DocValues(docvalues bool) *_integerRangeProperty {

	s.v.DocValues = &docvalues

	return s
}

func (s *_integerRangeProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_integerRangeProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_integerRangeProperty) Fields(fields map[string]types.Property) *_integerRangeProperty {

	s.v.Fields = fields
	return s
}

func (s *_integerRangeProperty) AddField(key string, value types.PropertyVariant) *_integerRangeProperty {

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

func (s *_integerRangeProperty) IgnoreAbove(ignoreabove int) *_integerRangeProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_integerRangeProperty) Index(index bool) *_integerRangeProperty {

	s.v.Index = &index

	return s
}

func (s *_integerRangeProperty) Meta(meta map[string]string) *_integerRangeProperty {

	s.v.Meta = meta
	return s
}

func (s *_integerRangeProperty) AddMeta(key string, value string) *_integerRangeProperty {

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

func (s *_integerRangeProperty) Properties(properties map[string]types.Property) *_integerRangeProperty {

	s.v.Properties = properties
	return s
}

func (s *_integerRangeProperty) AddProperty(key string, value types.PropertyVariant) *_integerRangeProperty {

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

func (s *_integerRangeProperty) Store(store bool) *_integerRangeProperty {

	s.v.Store = &store

	return s
}

func (s *_integerRangeProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_integerRangeProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_integerRangeProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_integerRangeProperty) IntegerRangePropertyCaster() *types.IntegerRangeProperty {
	return s.v
}
