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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _filteringRules struct {
	v *types.FilteringRules
}

func NewFilteringRules(advancedsnippet types.FilteringAdvancedSnippetVariant, validation types.FilteringRulesValidationVariant) *_filteringRules {

	tmp := &_filteringRules{v: types.NewFilteringRules()}

	tmp.AdvancedSnippet(advancedsnippet)

	tmp.Validation(validation)

	return tmp

}

func (s *_filteringRules) AdvancedSnippet(advancedsnippet types.FilteringAdvancedSnippetVariant) *_filteringRules {

	s.v.AdvancedSnippet = *advancedsnippet.FilteringAdvancedSnippetCaster()

	return s
}

func (s *_filteringRules) Rules(rules ...types.FilteringRuleVariant) *_filteringRules {

	for _, v := range rules {

		s.v.Rules = append(s.v.Rules, *v.FilteringRuleCaster())

	}
	return s
}

func (s *_filteringRules) Validation(validation types.FilteringRulesValidationVariant) *_filteringRules {

	s.v.Validation = *validation.FilteringRulesValidationCaster()

	return s
}

func (s *_filteringRules) FilteringRulesCaster() *types.FilteringRules {
	return s.v
}
