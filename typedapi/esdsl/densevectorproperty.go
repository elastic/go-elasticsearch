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
// https://github.com/elastic/elasticsearch-specification/tree/a0b0db20330063a6d11f7997ff443fd2a1a827d1

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/densevectorelementtype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/densevectorsimilarity"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _denseVectorProperty struct {
	v *types.DenseVectorProperty
}

func NewDenseVectorProperty() *_denseVectorProperty {

	return &_denseVectorProperty{v: types.NewDenseVectorProperty()}

}

func (s *_denseVectorProperty) Dims(dims int) *_denseVectorProperty {

	s.v.Dims = &dims

	return s
}

func (s *_denseVectorProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_denseVectorProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_denseVectorProperty) ElementType(elementtype densevectorelementtype.DenseVectorElementType) *_denseVectorProperty {

	s.v.ElementType = &elementtype
	return s
}

func (s *_denseVectorProperty) Fields(fields map[string]types.Property) *_denseVectorProperty {

	s.v.Fields = fields
	return s
}

func (s *_denseVectorProperty) AddField(key string, value types.PropertyVariant) *_denseVectorProperty {

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

func (s *_denseVectorProperty) IgnoreAbove(ignoreabove int) *_denseVectorProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_denseVectorProperty) Index(index bool) *_denseVectorProperty {

	s.v.Index = &index

	return s
}

func (s *_denseVectorProperty) IndexOptions(indexoptions types.DenseVectorIndexOptionsVariant) *_denseVectorProperty {

	s.v.IndexOptions = indexoptions.DenseVectorIndexOptionsCaster()

	return s
}

func (s *_denseVectorProperty) Meta(meta map[string]string) *_denseVectorProperty {

	s.v.Meta = meta
	return s
}

func (s *_denseVectorProperty) AddMeta(key string, value string) *_denseVectorProperty {

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

func (s *_denseVectorProperty) Properties(properties map[string]types.Property) *_denseVectorProperty {

	s.v.Properties = properties
	return s
}

func (s *_denseVectorProperty) AddProperty(key string, value types.PropertyVariant) *_denseVectorProperty {

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

func (s *_denseVectorProperty) Similarity(similarity densevectorsimilarity.DenseVectorSimilarity) *_denseVectorProperty {

	s.v.Similarity = &similarity
	return s
}

func (s *_denseVectorProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_denseVectorProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_denseVectorProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_denseVectorProperty) DenseVectorPropertyCaster() *types.DenseVectorProperty {
	return s.v
}
