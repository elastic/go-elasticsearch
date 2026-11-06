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

type _choiceDecisionQuestion struct {
	v *types.ChoiceDecisionQuestion
}

func NewChoiceDecisionQuestion(instructions string) *_choiceDecisionQuestion {

	tmp := &_choiceDecisionQuestion{v: types.NewChoiceDecisionQuestion()}

	tmp.Instructions(instructions)

	return tmp

}

func (s *_choiceDecisionQuestion) Choices(choices map[string]string) *_choiceDecisionQuestion {

	s.v.Choices = choices
	return s
}

func (s *_choiceDecisionQuestion) AddChoice(key string, value string) *_choiceDecisionQuestion {

	var tmp map[string]string
	if s.v.Choices == nil {
		s.v.Choices = make(map[string]string)
	} else {
		tmp = s.v.Choices
	}

	tmp[key] = value

	s.v.Choices = tmp
	return s
}

func (s *_choiceDecisionQuestion) Instructions(instructions string) *_choiceDecisionQuestion {

	s.v.Instructions = instructions

	return s
}

func (s *_choiceDecisionQuestion) ChoiceDecisionQuestionCaster() *types.ChoiceDecisionQuestion {
	return s.v
}
