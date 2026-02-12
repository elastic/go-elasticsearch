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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _sparseVectorProperty struct {
	v *types.SparseVectorProperty
}

func NewSparseVectorProperty() *_sparseVectorProperty {

	return &_sparseVectorProperty{v: types.NewSparseVectorProperty()}

}

func (s *_sparseVectorProperty) IndexOptions(indexoptions types.SparseVectorIndexOptionsVariant) *_sparseVectorProperty {

	s.v.IndexOptions = indexoptions.SparseVectorIndexOptionsCaster()

	return s
}

func (s *_sparseVectorProperty) Store(store bool) *_sparseVectorProperty {

	s.v.Store = &store

	return s
}

func (s *_sparseVectorProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_sparseVectorProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_sparseVectorProperty) Fields(fields map[string]types.Property) *_sparseVectorProperty {

	s.v.Fields = fields
	return s
}

func (s *_sparseVectorProperty) AddField(key string, value types.PropertyVariant) *_sparseVectorProperty {

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

func (s *_sparseVectorProperty) IgnoreAbove(ignoreabove int) *_sparseVectorProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_sparseVectorProperty) Meta(meta map[string]string) *_sparseVectorProperty {

	s.v.Meta = meta
	return s
}

func (s *_sparseVectorProperty) AddMeta(key string, value string) *_sparseVectorProperty {

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

func (s *_sparseVectorProperty) Properties(properties map[string]types.Property) *_sparseVectorProperty {

	s.v.Properties = properties
	return s
}

func (s *_sparseVectorProperty) AddProperty(key string, value types.PropertyVariant) *_sparseVectorProperty {

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

func (s *_sparseVectorProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_sparseVectorProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_sparseVectorProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_sparseVectorProperty) SparseVectorPropertyCaster() *types.SparseVectorProperty {
	return s.v
}
