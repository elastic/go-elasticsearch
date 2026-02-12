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
// https://github.com/elastic/elasticsearch-specification/tree/55f8d05b44cea956ae5ceddfcb02770ea2a24ff6

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _scriptQuery struct {
	v *types.ScriptQuery
}

// Filters documents based on a provided script.
// The script query is typically used in a filter context.
func NewScriptQuery(script types.ScriptVariant) *_scriptQuery {

	tmp := &_scriptQuery{v: types.NewScriptQuery()}

	tmp.Script(script)

	return tmp

}

func (s *_scriptQuery) Script(script types.ScriptVariant) *_scriptQuery {

	s.v.Script = *script.ScriptCaster()

	return s
}

func (s *_scriptQuery) Boost(boost float32) *_scriptQuery {

	s.v.Boost = &boost

	return s
}

func (s *_scriptQuery) QueryName_(queryname_ string) *_scriptQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_scriptQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.Script = s.v

	return container
}

func (s *_scriptQuery) ScriptQueryCaster() *types.ScriptQuery {
	return s.v
}
