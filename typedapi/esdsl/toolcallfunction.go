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

type _toolCallFunction struct {
	v *types.ToolCallFunction
}

func NewToolCallFunction(arguments string, name string) *_toolCallFunction {

	tmp := &_toolCallFunction{v: types.NewToolCallFunction()}

	tmp.Arguments(arguments)

	tmp.Name(name)

	return tmp

}

// The arguments to call the function with in JSON format.
func (s *_toolCallFunction) Arguments(arguments string) *_toolCallFunction {

	s.v.Arguments = arguments

	return s
}

// The name of the function to call.
func (s *_toolCallFunction) Name(name string) *_toolCallFunction {

	s.v.Name = name

	return s
}

func (s *_toolCallFunction) ToolCallFunctionCaster() *types.ToolCallFunction {
	return s.v
}
