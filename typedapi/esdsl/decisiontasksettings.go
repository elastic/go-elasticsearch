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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _decisionTaskSettings struct {
	v *types.DecisionTaskSettings
}

func NewDecisionTaskSettings() *_decisionTaskSettings {

	return &_decisionTaskSettings{v: types.NewDecisionTaskSettings()}

}

func (s *_decisionTaskSettings) Questions(questions map[string]types.DecisionQuestion) *_decisionTaskSettings {

	s.v.Questions = questions
	return s
}

func (s *_decisionTaskSettings) AddQuestion(key string, value types.DecisionQuestionVariant) *_decisionTaskSettings {

	var tmp map[string]types.DecisionQuestion
	if s.v.Questions == nil {
		s.v.Questions = make(map[string]types.DecisionQuestion)
	} else {
		tmp = s.v.Questions
	}

	tmp[key] = *value.DecisionQuestionCaster()

	s.v.Questions = tmp
	return s
}

func (s *_decisionTaskSettings) DecisionTaskSettingsCaster() *types.DecisionTaskSettings {
	return s.v
}
