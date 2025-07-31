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
// https://github.com/elastic/elasticsearch-specification/tree/86f41834c7bb975159a38a73be8a9d930010d673

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _longRangeProperty struct {
	v *types.LongRangeProperty
}

func NewLongRangeProperty() *_longRangeProperty {

	return &_longRangeProperty{v: types.NewLongRangeProperty()}

}

func (s *_longRangeProperty) Boost(boost types.Float64) *_longRangeProperty {

	s.v.Boost = &boost

	return s
}

func (s *_longRangeProperty) Coerce(coerce bool) *_longRangeProperty {

	s.v.Coerce = &coerce

	return s
}

func (s *_longRangeProperty) CopyTo(fields ...string) *_longRangeProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_longRangeProperty) DocValues(docvalues bool) *_longRangeProperty {

	s.v.DocValues = &docvalues

	return s
}

func (s *_longRangeProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_longRangeProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_longRangeProperty) Fields(fields map[string]types.Property) *_longRangeProperty {

	s.v.Fields = fields
	return s
}

func (s *_longRangeProperty) AddField(key string, value types.PropertyVariant) *_longRangeProperty {

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

func (s *_longRangeProperty) IgnoreAbove(ignoreabove int) *_longRangeProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_longRangeProperty) Index(index bool) *_longRangeProperty {

	s.v.Index = &index

	return s
}

func (s *_longRangeProperty) Meta(meta map[string]string) *_longRangeProperty {

	s.v.Meta = meta
	return s
}

func (s *_longRangeProperty) AddMeta(key string, value string) *_longRangeProperty {

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

func (s *_longRangeProperty) Properties(properties map[string]types.Property) *_longRangeProperty {

	s.v.Properties = properties
	return s
}

func (s *_longRangeProperty) AddProperty(key string, value types.PropertyVariant) *_longRangeProperty {

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

func (s *_longRangeProperty) Store(store bool) *_longRangeProperty {

	s.v.Store = &store

	return s
}

func (s *_longRangeProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_longRangeProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_longRangeProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_longRangeProperty) LongRangePropertyCaster() *types.LongRangeProperty {
	return s.v
}
