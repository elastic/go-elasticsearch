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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _rankFeaturesProperty struct {
	v *types.RankFeaturesProperty
}

func NewRankFeaturesProperty() *_rankFeaturesProperty {

	return &_rankFeaturesProperty{v: types.NewRankFeaturesProperty()}

}

func (s *_rankFeaturesProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_rankFeaturesProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_rankFeaturesProperty) Fields(fields map[string]types.Property) *_rankFeaturesProperty {

	s.v.Fields = fields
	return s
}

func (s *_rankFeaturesProperty) AddField(key string, value types.PropertyVariant) *_rankFeaturesProperty {

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

func (s *_rankFeaturesProperty) IgnoreAbove(ignoreabove int) *_rankFeaturesProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_rankFeaturesProperty) Meta(meta map[string]string) *_rankFeaturesProperty {

	s.v.Meta = meta
	return s
}

func (s *_rankFeaturesProperty) AddMeta(key string, value string) *_rankFeaturesProperty {

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

func (s *_rankFeaturesProperty) PositiveScoreImpact(positivescoreimpact bool) *_rankFeaturesProperty {

	s.v.PositiveScoreImpact = &positivescoreimpact

	return s
}

func (s *_rankFeaturesProperty) Properties(properties map[string]types.Property) *_rankFeaturesProperty {

	s.v.Properties = properties
	return s
}

func (s *_rankFeaturesProperty) AddProperty(key string, value types.PropertyVariant) *_rankFeaturesProperty {

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

func (s *_rankFeaturesProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_rankFeaturesProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_rankFeaturesProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_rankFeaturesProperty) RankFeaturesPropertyCaster() *types.RankFeaturesProperty {
	return s.v
}
