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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/syntheticsourcekeepenum"
)

type _ipRangeProperty struct {
	v *types.IpRangeProperty
}

func NewIpRangeProperty() *_ipRangeProperty {

	return &_ipRangeProperty{v: types.NewIpRangeProperty()}

}

func (s *_ipRangeProperty) Boost(boost types.Float64) *_ipRangeProperty {

	s.v.Boost = &boost

	return s
}

func (s *_ipRangeProperty) Coerce(coerce bool) *_ipRangeProperty {

	s.v.Coerce = &coerce

	return s
}

func (s *_ipRangeProperty) CopyTo(fields ...string) *_ipRangeProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_ipRangeProperty) DocValues(docvalues bool) *_ipRangeProperty {

	s.v.DocValues = &docvalues

	return s
}

func (s *_ipRangeProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_ipRangeProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_ipRangeProperty) Fields(fields map[string]types.Property) *_ipRangeProperty {

	s.v.Fields = fields
	return s
}

func (s *_ipRangeProperty) AddField(key string, value types.PropertyVariant) *_ipRangeProperty {

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

func (s *_ipRangeProperty) IgnoreAbove(ignoreabove int) *_ipRangeProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_ipRangeProperty) Index(index bool) *_ipRangeProperty {

	s.v.Index = &index

	return s
}

// Metadata about the field.
func (s *_ipRangeProperty) Meta(meta map[string]string) *_ipRangeProperty {

	s.v.Meta = meta
	return s
}

func (s *_ipRangeProperty) AddMeta(key string, value string) *_ipRangeProperty {

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

func (s *_ipRangeProperty) Properties(properties map[string]types.Property) *_ipRangeProperty {

	s.v.Properties = properties
	return s
}

func (s *_ipRangeProperty) AddProperty(key string, value types.PropertyVariant) *_ipRangeProperty {

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

func (s *_ipRangeProperty) Store(store bool) *_ipRangeProperty {

	s.v.Store = &store

	return s
}

func (s *_ipRangeProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_ipRangeProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_ipRangeProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_ipRangeProperty) IpRangePropertyCaster() *types.IpRangeProperty {
	return s.v
}
