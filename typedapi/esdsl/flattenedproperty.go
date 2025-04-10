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
// https://github.com/elastic/elasticsearch-specification/tree/beeb1dc688bcc058488dcc45d9cbd2cd364e9943

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexoptions"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _flattenedProperty struct {
	v *types.FlattenedProperty
}

func NewFlattenedProperty() *_flattenedProperty {

	return &_flattenedProperty{v: types.NewFlattenedProperty()}

}

func (s *_flattenedProperty) Boost(boost types.Float64) *_flattenedProperty {

	s.v.Boost = &boost

	return s
}

func (s *_flattenedProperty) DepthLimit(depthlimit int) *_flattenedProperty {

	s.v.DepthLimit = &depthlimit

	return s
}

func (s *_flattenedProperty) DocValues(docvalues bool) *_flattenedProperty {

	s.v.DocValues = &docvalues

	return s
}

func (s *_flattenedProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_flattenedProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_flattenedProperty) EagerGlobalOrdinals(eagerglobalordinals bool) *_flattenedProperty {

	s.v.EagerGlobalOrdinals = &eagerglobalordinals

	return s
}

func (s *_flattenedProperty) Fields(fields map[string]types.Property) *_flattenedProperty {

	s.v.Fields = fields
	return s
}

func (s *_flattenedProperty) AddField(key string, value types.PropertyVariant) *_flattenedProperty {

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

func (s *_flattenedProperty) IgnoreAbove(ignoreabove int) *_flattenedProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_flattenedProperty) Index(index bool) *_flattenedProperty {

	s.v.Index = &index

	return s
}

func (s *_flattenedProperty) IndexOptions(indexoptions indexoptions.IndexOptions) *_flattenedProperty {

	s.v.IndexOptions = &indexoptions
	return s
}

func (s *_flattenedProperty) Meta(meta map[string]string) *_flattenedProperty {

	s.v.Meta = meta
	return s
}

func (s *_flattenedProperty) AddMeta(key string, value string) *_flattenedProperty {

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

func (s *_flattenedProperty) NullValue(nullvalue string) *_flattenedProperty {

	s.v.NullValue = &nullvalue

	return s
}

func (s *_flattenedProperty) Properties(properties map[string]types.Property) *_flattenedProperty {

	s.v.Properties = properties
	return s
}

func (s *_flattenedProperty) AddProperty(key string, value types.PropertyVariant) *_flattenedProperty {

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

func (s *_flattenedProperty) Similarity(similarity string) *_flattenedProperty {

	s.v.Similarity = &similarity

	return s
}

func (s *_flattenedProperty) SplitQueriesOnWhitespace(splitqueriesonwhitespace bool) *_flattenedProperty {

	s.v.SplitQueriesOnWhitespace = &splitqueriesonwhitespace

	return s
}

func (s *_flattenedProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_flattenedProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_flattenedProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_flattenedProperty) FlattenedPropertyCaster() *types.FlattenedProperty {
	return s.v
}
