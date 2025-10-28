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
// https://github.com/elastic/elasticsearch-specification/tree/d520d9e8cf14cad487de5e0654007686c395b494

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _scriptScoreQuery struct {
	v *types.ScriptScoreQuery
}

// Uses a script to provide a custom score for returned documents.
func NewScriptScoreQuery(query types.QueryVariant, script types.ScriptVariant) *_scriptScoreQuery {

	tmp := &_scriptScoreQuery{v: types.NewScriptScoreQuery()}

	tmp.Query(query)

	tmp.Script(script)

	return tmp

}

func (s *_scriptScoreQuery) Boost(boost float32) *_scriptScoreQuery {

	s.v.Boost = &boost

	return s
}

func (s *_scriptScoreQuery) MinScore(minscore float32) *_scriptScoreQuery {

	s.v.MinScore = &minscore

	return s
}

func (s *_scriptScoreQuery) Query(query types.QueryVariant) *_scriptScoreQuery {

	s.v.Query = *query.QueryCaster()

	return s
}

func (s *_scriptScoreQuery) QueryName_(queryname_ string) *_scriptScoreQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_scriptScoreQuery) Script(script types.ScriptVariant) *_scriptScoreQuery {

	s.v.Script = *script.ScriptCaster()

	return s
}

func (s *_scriptScoreQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.ScriptScore = s.v

	return container
}

func (s *_scriptScoreQuery) ScriptScoreQueryCaster() *types.ScriptScoreQuery {
	return s.v
}
