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
// https://github.com/elastic/elasticsearch-specification/tree/2514615770f18dbb4e3887cc1a279995dbfd0724

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _scriptScoreFunction struct {
	v *types.ScriptScoreFunction
}

// Enables you to wrap another query and customize the scoring of it optionally
// with a computation derived from other numeric field values in the doc using a
// script expression.
func NewScriptScoreFunction(script types.ScriptVariant) *_scriptScoreFunction {

	tmp := &_scriptScoreFunction{v: types.NewScriptScoreFunction()}

	tmp.Script(script)

	return tmp

}

func (s *_scriptScoreFunction) Script(script types.ScriptVariant) *_scriptScoreFunction {

	s.v.Script = *script.ScriptCaster()

	return s
}

func (s *_scriptScoreFunction) FunctionScoreCaster() *types.FunctionScore {
	container := types.NewFunctionScore()

	container.ScriptScore = s.v

	return container
}

func (s *_scriptScoreFunction) ScriptScoreFunctionCaster() *types.ScriptScoreFunction {
	return s.v
}
