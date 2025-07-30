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
// https://github.com/elastic/elasticsearch-specification/tree/e585438d116b00ff34643179e6286e402c0bcaaf

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _completionProperty struct {
	v *types.CompletionProperty
}

func NewCompletionProperty() *_completionProperty {

	return &_completionProperty{v: types.NewCompletionProperty()}

}

func (s *_completionProperty) Analyzer(analyzer string) *_completionProperty {

	s.v.Analyzer = &analyzer

	return s
}

func (s *_completionProperty) Contexts(contexts ...types.SuggestContextVariant) *_completionProperty {

	for _, v := range contexts {

		s.v.Contexts = append(s.v.Contexts, *v.SuggestContextCaster())

	}
	return s
}

func (s *_completionProperty) CopyTo(fields ...string) *_completionProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_completionProperty) DocValues(docvalues bool) *_completionProperty {

	s.v.DocValues = &docvalues

	return s
}

func (s *_completionProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_completionProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_completionProperty) Fields(fields map[string]types.Property) *_completionProperty {

	s.v.Fields = fields
	return s
}

func (s *_completionProperty) AddField(key string, value types.PropertyVariant) *_completionProperty {

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

func (s *_completionProperty) IgnoreAbove(ignoreabove int) *_completionProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_completionProperty) MaxInputLength(maxinputlength int) *_completionProperty {

	s.v.MaxInputLength = &maxinputlength

	return s
}

func (s *_completionProperty) Meta(meta map[string]string) *_completionProperty {

	s.v.Meta = meta
	return s
}

func (s *_completionProperty) AddMeta(key string, value string) *_completionProperty {

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

func (s *_completionProperty) PreservePositionIncrements(preservepositionincrements bool) *_completionProperty {

	s.v.PreservePositionIncrements = &preservepositionincrements

	return s
}

func (s *_completionProperty) PreserveSeparators(preserveseparators bool) *_completionProperty {

	s.v.PreserveSeparators = &preserveseparators

	return s
}

func (s *_completionProperty) Properties(properties map[string]types.Property) *_completionProperty {

	s.v.Properties = properties
	return s
}

func (s *_completionProperty) AddProperty(key string, value types.PropertyVariant) *_completionProperty {

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

func (s *_completionProperty) SearchAnalyzer(searchanalyzer string) *_completionProperty {

	s.v.SearchAnalyzer = &searchanalyzer

	return s
}

func (s *_completionProperty) Store(store bool) *_completionProperty {

	s.v.Store = &store

	return s
}

func (s *_completionProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_completionProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_completionProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_completionProperty) CompletionPropertyCaster() *types.CompletionProperty {
	return s.v
}
