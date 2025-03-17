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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/indexoptions"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/termvectoroption"
)

type _searchAsYouTypeProperty struct {
	v *types.SearchAsYouTypeProperty
}

func NewSearchAsYouTypeProperty() *_searchAsYouTypeProperty {

	return &_searchAsYouTypeProperty{v: types.NewSearchAsYouTypeProperty()}

}

func (s *_searchAsYouTypeProperty) Analyzer(analyzer string) *_searchAsYouTypeProperty {

	s.v.Analyzer = &analyzer

	return s
}

func (s *_searchAsYouTypeProperty) CopyTo(fields ...string) *_searchAsYouTypeProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_searchAsYouTypeProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_searchAsYouTypeProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_searchAsYouTypeProperty) Fields(fields map[string]types.Property) *_searchAsYouTypeProperty {

	s.v.Fields = fields
	return s
}

func (s *_searchAsYouTypeProperty) AddField(key string, value types.PropertyVariant) *_searchAsYouTypeProperty {

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

func (s *_searchAsYouTypeProperty) IgnoreAbove(ignoreabove int) *_searchAsYouTypeProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_searchAsYouTypeProperty) Index(index bool) *_searchAsYouTypeProperty {

	s.v.Index = &index

	return s
}

func (s *_searchAsYouTypeProperty) IndexOptions(indexoptions indexoptions.IndexOptions) *_searchAsYouTypeProperty {

	s.v.IndexOptions = &indexoptions
	return s
}

func (s *_searchAsYouTypeProperty) MaxShingleSize(maxshinglesize int) *_searchAsYouTypeProperty {

	s.v.MaxShingleSize = &maxshinglesize

	return s
}

// Metadata about the field.
func (s *_searchAsYouTypeProperty) Meta(meta map[string]string) *_searchAsYouTypeProperty {

	s.v.Meta = meta
	return s
}

func (s *_searchAsYouTypeProperty) AddMeta(key string, value string) *_searchAsYouTypeProperty {

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

func (s *_searchAsYouTypeProperty) Norms(norms bool) *_searchAsYouTypeProperty {

	s.v.Norms = &norms

	return s
}

func (s *_searchAsYouTypeProperty) Properties(properties map[string]types.Property) *_searchAsYouTypeProperty {

	s.v.Properties = properties
	return s
}

func (s *_searchAsYouTypeProperty) AddProperty(key string, value types.PropertyVariant) *_searchAsYouTypeProperty {

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

func (s *_searchAsYouTypeProperty) SearchAnalyzer(searchanalyzer string) *_searchAsYouTypeProperty {

	s.v.SearchAnalyzer = &searchanalyzer

	return s
}

func (s *_searchAsYouTypeProperty) SearchQuoteAnalyzer(searchquoteanalyzer string) *_searchAsYouTypeProperty {

	s.v.SearchQuoteAnalyzer = &searchquoteanalyzer

	return s
}

func (s *_searchAsYouTypeProperty) Similarity(similarity string) *_searchAsYouTypeProperty {

	s.v.Similarity = &similarity

	return s
}

func (s *_searchAsYouTypeProperty) Store(store bool) *_searchAsYouTypeProperty {

	s.v.Store = &store

	return s
}

func (s *_searchAsYouTypeProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_searchAsYouTypeProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_searchAsYouTypeProperty) TermVector(termvector termvectoroption.TermVectorOption) *_searchAsYouTypeProperty {

	s.v.TermVector = &termvector
	return s
}

func (s *_searchAsYouTypeProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_searchAsYouTypeProperty) SearchAsYouTypePropertyCaster() *types.SearchAsYouTypeProperty {
	return s.v
}
