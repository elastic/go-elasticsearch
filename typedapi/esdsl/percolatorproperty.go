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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/syntheticsourcekeepenum"
)

type _percolatorProperty struct {
	v *types.PercolatorProperty
}

func NewPercolatorProperty() *_percolatorProperty {

	return &_percolatorProperty{v: types.NewPercolatorProperty()}

}

func (s *_percolatorProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_percolatorProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_percolatorProperty) Fields(fields map[string]types.Property) *_percolatorProperty {

	s.v.Fields = fields
	return s
}

func (s *_percolatorProperty) AddField(key string, value types.PropertyVariant) *_percolatorProperty {

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

func (s *_percolatorProperty) IgnoreAbove(ignoreabove int) *_percolatorProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_percolatorProperty) Meta(meta map[string]string) *_percolatorProperty {

	s.v.Meta = meta
	return s
}

func (s *_percolatorProperty) AddMeta(key string, value string) *_percolatorProperty {

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

func (s *_percolatorProperty) Properties(properties map[string]types.Property) *_percolatorProperty {

	s.v.Properties = properties
	return s
}

func (s *_percolatorProperty) AddProperty(key string, value types.PropertyVariant) *_percolatorProperty {

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

func (s *_percolatorProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_percolatorProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_percolatorProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_percolatorProperty) PercolatorPropertyCaster() *types.PercolatorProperty {
	return s.v
}
