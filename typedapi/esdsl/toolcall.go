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
// https://github.com/elastic/elasticsearch-specification/tree/86f41834c7bb975159a38a73be8a9d930010d673

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _toolCall struct {
	v *types.ToolCall
}

func NewToolCall(function types.ToolCallFunctionVariant, type_ string) *_toolCall {

	tmp := &_toolCall{v: types.NewToolCall()}

	tmp.Function(function)

	tmp.Type(type_)

	return tmp

}

func (s *_toolCall) Function(function types.ToolCallFunctionVariant) *_toolCall {

	s.v.Function = *function.ToolCallFunctionCaster()

	return s
}

func (s *_toolCall) Id(id string) *_toolCall {

	s.v.Id = id

	return s
}

func (s *_toolCall) Type(type_ string) *_toolCall {

	s.v.Type = type_

	return s
}

func (s *_toolCall) ToolCallCaster() *types.ToolCall {
	return s.v
}
