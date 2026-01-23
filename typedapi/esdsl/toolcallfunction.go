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

type _toolCallFunction struct {
	v *types.ToolCallFunction
}

func NewToolCallFunction(arguments string, name string) *_toolCallFunction {

	tmp := &_toolCallFunction{v: types.NewToolCallFunction()}

	tmp.Arguments(arguments)

	tmp.Name(name)

	return tmp

}

func (s *_toolCallFunction) Arguments(arguments string) *_toolCallFunction {

	s.v.Arguments = arguments

	return s
}

func (s *_toolCallFunction) Name(name string) *_toolCallFunction {

	s.v.Name = name

	return s
}

func (s *_toolCallFunction) ToolCallFunctionCaster() *types.ToolCallFunction {
	return s.v
}
