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
// https://github.com/elastic/elasticsearch-specification/tree/b1811e10a0722431d79d1c234dd412ff47d8656f

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icucollationalternate"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icucollationcasefirst"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icucollationdecomposition"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icucollationstrength"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexoptions"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _icuCollationProperty struct {
	v *types.IcuCollationProperty
}

func NewIcuCollationProperty() *_icuCollationProperty {

	return &_icuCollationProperty{v: types.NewIcuCollationProperty()}

}

func (s *_icuCollationProperty) Alternate(alternate icucollationalternate.IcuCollationAlternate) *_icuCollationProperty {

	s.v.Alternate = &alternate
	return s
}

func (s *_icuCollationProperty) CaseFirst(casefirst icucollationcasefirst.IcuCollationCaseFirst) *_icuCollationProperty {

	s.v.CaseFirst = &casefirst
	return s
}

func (s *_icuCollationProperty) CaseLevel(caselevel bool) *_icuCollationProperty {

	s.v.CaseLevel = &caselevel

	return s
}

func (s *_icuCollationProperty) Country(country string) *_icuCollationProperty {

	s.v.Country = &country

	return s
}

func (s *_icuCollationProperty) Decomposition(decomposition icucollationdecomposition.IcuCollationDecomposition) *_icuCollationProperty {

	s.v.Decomposition = &decomposition
	return s
}

func (s *_icuCollationProperty) HiraganaQuaternaryMode(hiraganaquaternarymode bool) *_icuCollationProperty {

	s.v.HiraganaQuaternaryMode = &hiraganaquaternarymode

	return s
}

func (s *_icuCollationProperty) Index(index bool) *_icuCollationProperty {

	s.v.Index = &index

	return s
}

func (s *_icuCollationProperty) IndexOptions(indexoptions indexoptions.IndexOptions) *_icuCollationProperty {

	s.v.IndexOptions = &indexoptions
	return s
}

func (s *_icuCollationProperty) Language(language string) *_icuCollationProperty {

	s.v.Language = &language

	return s
}

func (s *_icuCollationProperty) Norms(norms bool) *_icuCollationProperty {

	s.v.Norms = &norms

	return s
}

func (s *_icuCollationProperty) NullValue(nullvalue string) *_icuCollationProperty {

	s.v.NullValue = &nullvalue

	return s
}

func (s *_icuCollationProperty) Numeric(numeric bool) *_icuCollationProperty {

	s.v.Numeric = &numeric

	return s
}

func (s *_icuCollationProperty) Rules(rules string) *_icuCollationProperty {

	s.v.Rules = &rules

	return s
}

func (s *_icuCollationProperty) Strength(strength icucollationstrength.IcuCollationStrength) *_icuCollationProperty {

	s.v.Strength = &strength
	return s
}

func (s *_icuCollationProperty) VariableTop(variabletop string) *_icuCollationProperty {

	s.v.VariableTop = &variabletop

	return s
}

func (s *_icuCollationProperty) Variant(variant string) *_icuCollationProperty {

	s.v.Variant = &variant

	return s
}

func (s *_icuCollationProperty) CopyTo(fields ...string) *_icuCollationProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_icuCollationProperty) DocValues(docvalues bool) *_icuCollationProperty {

	s.v.DocValues = &docvalues

	return s
}

func (s *_icuCollationProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_icuCollationProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_icuCollationProperty) Fields(fields map[string]types.Property) *_icuCollationProperty {

	s.v.Fields = fields
	return s
}

func (s *_icuCollationProperty) AddField(key string, value types.PropertyVariant) *_icuCollationProperty {

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

func (s *_icuCollationProperty) IgnoreAbove(ignoreabove int) *_icuCollationProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_icuCollationProperty) Meta(meta map[string]string) *_icuCollationProperty {

	s.v.Meta = meta
	return s
}

func (s *_icuCollationProperty) AddMeta(key string, value string) *_icuCollationProperty {

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

func (s *_icuCollationProperty) Properties(properties map[string]types.Property) *_icuCollationProperty {

	s.v.Properties = properties
	return s
}

func (s *_icuCollationProperty) AddProperty(key string, value types.PropertyVariant) *_icuCollationProperty {

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

func (s *_icuCollationProperty) Store(store bool) *_icuCollationProperty {

	s.v.Store = &store

	return s
}

func (s *_icuCollationProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_icuCollationProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_icuCollationProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_icuCollationProperty) IcuCollationPropertyCaster() *types.IcuCollationProperty {
	return s.v
}
