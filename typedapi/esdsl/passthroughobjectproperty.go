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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _passthroughObjectProperty struct {
	v *types.PassthroughObjectProperty
}

func NewPassthroughObjectProperty() *_passthroughObjectProperty {

	return &_passthroughObjectProperty{v: types.NewPassthroughObjectProperty()}

}

func (s *_passthroughObjectProperty) Enabled(enabled bool) *_passthroughObjectProperty {

	s.v.Enabled = &enabled

	return s
}

func (s *_passthroughObjectProperty) Priority(priority int) *_passthroughObjectProperty {

	s.v.Priority = &priority

	return s
}

func (s *_passthroughObjectProperty) TimeSeriesDimension(timeseriesdimension bool) *_passthroughObjectProperty {

	s.v.TimeSeriesDimension = &timeseriesdimension

	return s
}

func (s *_passthroughObjectProperty) CopyTo(fields ...string) *_passthroughObjectProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_passthroughObjectProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_passthroughObjectProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_passthroughObjectProperty) Fields(fields map[string]types.Property) *_passthroughObjectProperty {

	s.v.Fields = fields
	return s
}

func (s *_passthroughObjectProperty) AddField(key string, value types.PropertyVariant) *_passthroughObjectProperty {

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

func (s *_passthroughObjectProperty) IgnoreAbove(ignoreabove int) *_passthroughObjectProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_passthroughObjectProperty) Meta(meta map[string]string) *_passthroughObjectProperty {

	s.v.Meta = meta
	return s
}

func (s *_passthroughObjectProperty) AddMeta(key string, value string) *_passthroughObjectProperty {

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

func (s *_passthroughObjectProperty) Properties(properties map[string]types.Property) *_passthroughObjectProperty {

	s.v.Properties = properties
	return s
}

func (s *_passthroughObjectProperty) AddProperty(key string, value types.PropertyVariant) *_passthroughObjectProperty {

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

func (s *_passthroughObjectProperty) Store(store bool) *_passthroughObjectProperty {

	s.v.Store = &store

	return s
}

func (s *_passthroughObjectProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_passthroughObjectProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_passthroughObjectProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_passthroughObjectProperty) PassthroughObjectPropertyCaster() *types.PassthroughObjectProperty {
	return s.v
}
