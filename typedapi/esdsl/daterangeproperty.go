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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/syntheticsourcekeepenum"
)

type _dateRangeProperty struct {
	v *types.DateRangeProperty
}

func NewDateRangeProperty() *_dateRangeProperty {

	return &_dateRangeProperty{v: types.NewDateRangeProperty()}

}

func (s *_dateRangeProperty) Boost(boost types.Float64) *_dateRangeProperty {

	s.v.Boost = &boost

	return s
}

func (s *_dateRangeProperty) Coerce(coerce bool) *_dateRangeProperty {

	s.v.Coerce = &coerce

	return s
}

func (s *_dateRangeProperty) CopyTo(fields ...string) *_dateRangeProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_dateRangeProperty) DocValues(docvalues bool) *_dateRangeProperty {

	s.v.DocValues = &docvalues

	return s
}

func (s *_dateRangeProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_dateRangeProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_dateRangeProperty) Fields(fields map[string]types.Property) *_dateRangeProperty {

	s.v.Fields = fields
	return s
}

func (s *_dateRangeProperty) AddField(key string, value types.PropertyVariant) *_dateRangeProperty {

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

func (s *_dateRangeProperty) Format(format string) *_dateRangeProperty {

	s.v.Format = &format

	return s
}

func (s *_dateRangeProperty) IgnoreAbove(ignoreabove int) *_dateRangeProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_dateRangeProperty) Index(index bool) *_dateRangeProperty {

	s.v.Index = &index

	return s
}

// Metadata about the field.
func (s *_dateRangeProperty) Meta(meta map[string]string) *_dateRangeProperty {

	s.v.Meta = meta
	return s
}

func (s *_dateRangeProperty) AddMeta(key string, value string) *_dateRangeProperty {

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

func (s *_dateRangeProperty) Properties(properties map[string]types.Property) *_dateRangeProperty {

	s.v.Properties = properties
	return s
}

func (s *_dateRangeProperty) AddProperty(key string, value types.PropertyVariant) *_dateRangeProperty {

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

func (s *_dateRangeProperty) Store(store bool) *_dateRangeProperty {

	s.v.Store = &store

	return s
}

func (s *_dateRangeProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_dateRangeProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_dateRangeProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_dateRangeProperty) DateRangePropertyCaster() *types.DateRangeProperty {
	return s.v
}
