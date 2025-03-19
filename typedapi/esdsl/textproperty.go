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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/indexoptions"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/termvectoroption"
)

type _textProperty struct {
	v *types.TextProperty
}

func NewTextProperty() *_textProperty {

	return &_textProperty{v: types.NewTextProperty()}

}

func (s *_textProperty) Analyzer(analyzer string) *_textProperty {

	s.v.Analyzer = &analyzer

	return s
}

func (s *_textProperty) Boost(boost types.Float64) *_textProperty {

	s.v.Boost = &boost

	return s
}

func (s *_textProperty) CopyTo(fields ...string) *_textProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_textProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_textProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_textProperty) EagerGlobalOrdinals(eagerglobalordinals bool) *_textProperty {

	s.v.EagerGlobalOrdinals = &eagerglobalordinals

	return s
}

func (s *_textProperty) Fielddata(fielddata bool) *_textProperty {

	s.v.Fielddata = &fielddata

	return s
}

func (s *_textProperty) FielddataFrequencyFilter(fielddatafrequencyfilter types.FielddataFrequencyFilterVariant) *_textProperty {

	s.v.FielddataFrequencyFilter = fielddatafrequencyfilter.FielddataFrequencyFilterCaster()

	return s
}

func (s *_textProperty) Fields(fields map[string]types.Property) *_textProperty {

	s.v.Fields = fields
	return s
}

func (s *_textProperty) AddField(key string, value types.PropertyVariant) *_textProperty {

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

func (s *_textProperty) IgnoreAbove(ignoreabove int) *_textProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_textProperty) Index(index bool) *_textProperty {

	s.v.Index = &index

	return s
}

func (s *_textProperty) IndexOptions(indexoptions indexoptions.IndexOptions) *_textProperty {

	s.v.IndexOptions = &indexoptions
	return s
}

func (s *_textProperty) IndexPhrases(indexphrases bool) *_textProperty {

	s.v.IndexPhrases = &indexphrases

	return s
}

func (s *_textProperty) IndexPrefixes(indexprefixes types.TextIndexPrefixesVariant) *_textProperty {

	s.v.IndexPrefixes = indexprefixes.TextIndexPrefixesCaster()

	return s
}

// Metadata about the field.
func (s *_textProperty) Meta(meta map[string]string) *_textProperty {

	s.v.Meta = meta
	return s
}

func (s *_textProperty) AddMeta(key string, value string) *_textProperty {

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

func (s *_textProperty) Norms(norms bool) *_textProperty {

	s.v.Norms = &norms

	return s
}

func (s *_textProperty) PositionIncrementGap(positionincrementgap int) *_textProperty {

	s.v.PositionIncrementGap = &positionincrementgap

	return s
}

func (s *_textProperty) Properties(properties map[string]types.Property) *_textProperty {

	s.v.Properties = properties
	return s
}

func (s *_textProperty) AddProperty(key string, value types.PropertyVariant) *_textProperty {

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

func (s *_textProperty) SearchAnalyzer(searchanalyzer string) *_textProperty {

	s.v.SearchAnalyzer = &searchanalyzer

	return s
}

func (s *_textProperty) SearchQuoteAnalyzer(searchquoteanalyzer string) *_textProperty {

	s.v.SearchQuoteAnalyzer = &searchquoteanalyzer

	return s
}

func (s *_textProperty) Similarity(similarity string) *_textProperty {

	s.v.Similarity = &similarity

	return s
}

func (s *_textProperty) Store(store bool) *_textProperty {

	s.v.Store = &store

	return s
}

func (s *_textProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_textProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_textProperty) TermVector(termvector termvectoroption.TermVectorOption) *_textProperty {

	s.v.TermVector = &termvector
	return s
}

func (s *_textProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_textProperty) TextPropertyCaster() *types.TextProperty {
	return s.v
}
