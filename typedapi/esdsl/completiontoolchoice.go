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

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _completionToolChoice struct {
	v *types.CompletionToolChoice
}

func NewCompletionToolChoice(function types.CompletionToolChoiceFunctionVariant, type_ string) *_completionToolChoice {

	tmp := &_completionToolChoice{v: types.NewCompletionToolChoice()}

	tmp.Function(function)

	tmp.Type(type_)

	return tmp

}

// The tool choice function.
func (s *_completionToolChoice) Function(function types.CompletionToolChoiceFunctionVariant) *_completionToolChoice {

	s.v.Function = *function.CompletionToolChoiceFunctionCaster()

	return s
}

// The type of the tool.
func (s *_completionToolChoice) Type(type_ string) *_completionToolChoice {

	s.v.Type = type_

	return s
}

func (s *_completionToolChoice) CompletionToolChoiceCaster() *types.CompletionToolChoice {
	return s.v
}
