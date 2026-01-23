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

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _requestChatCompletion struct {
	v *types.RequestChatCompletion
}

func NewRequestChatCompletion() *_requestChatCompletion {

	return &_requestChatCompletion{v: types.NewRequestChatCompletion()}

}

func (s *_requestChatCompletion) MaxCompletionTokens(maxcompletiontokens int64) *_requestChatCompletion {

	s.v.MaxCompletionTokens = &maxcompletiontokens

	return s
}

func (s *_requestChatCompletion) Messages(messages ...types.MessageVariant) *_requestChatCompletion {

	for _, v := range messages {

		s.v.Messages = append(s.v.Messages, *v.MessageCaster())

	}
	return s
}

func (s *_requestChatCompletion) Model(model string) *_requestChatCompletion {

	s.v.Model = &model

	return s
}

func (s *_requestChatCompletion) Stop(stops ...string) *_requestChatCompletion {

	for _, v := range stops {

		s.v.Stop = append(s.v.Stop, v)

	}
	return s
}

func (s *_requestChatCompletion) Temperature(temperature float32) *_requestChatCompletion {

	s.v.Temperature = &temperature

	return s
}

func (s *_requestChatCompletion) ToolChoice(completiontooltype types.CompletionToolTypeVariant) *_requestChatCompletion {

	s.v.ToolChoice = *completiontooltype.CompletionToolTypeCaster()

	return s
}

func (s *_requestChatCompletion) Tools(tools ...types.CompletionToolVariant) *_requestChatCompletion {

	for _, v := range tools {

		s.v.Tools = append(s.v.Tools, *v.CompletionToolCaster())

	}
	return s
}

func (s *_requestChatCompletion) TopP(topp float32) *_requestChatCompletion {

	s.v.TopP = &topp

	return s
}

func (s *_requestChatCompletion) RequestChatCompletionCaster() *types.RequestChatCompletion {
	return s.v
}
