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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _regexpQuery struct {
	k string
	v *types.RegexpQuery
}

// Returns documents that contain terms matching a regular expression.
func NewRegexpQuery(field string, value string) *_regexpQuery {
	tmp := &_regexpQuery{
		k: field,
		v: types.NewRegexpQuery(),
	}

	tmp.Value(value)
	return tmp
}

func (s *_regexpQuery) CaseInsensitive(caseinsensitive bool) *_regexpQuery {

	s.v.CaseInsensitive = &caseinsensitive

	return s
}

func (s *_regexpQuery) Flags(flags string) *_regexpQuery {

	s.v.Flags = &flags

	return s
}

func (s *_regexpQuery) MaxDeterminizedStates(maxdeterminizedstates int) *_regexpQuery {

	s.v.MaxDeterminizedStates = &maxdeterminizedstates

	return s
}

func (s *_regexpQuery) Rewrite(multitermqueryrewrite string) *_regexpQuery {

	s.v.Rewrite = &multitermqueryrewrite

	return s
}

func (s *_regexpQuery) Value(value string) *_regexpQuery {

	s.v.Value = value

	return s
}

func (s *_regexpQuery) Boost(boost float32) *_regexpQuery {

	s.v.Boost = &boost

	return s
}

func (s *_regexpQuery) QueryName_(queryname_ string) *_regexpQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_regexpQuery) QueryCaster() *types.Query {
	container := types.NewQuery()
	container.Regexp = map[string]types.RegexpQuery{
		s.k: *s.v,
	}
	return container
}

// NewSingleRegexpQuery should be used when you want to
// create a single key dictionary without specifying the key in the
// constructor. Usually key is already defined within the parent container.
func NewSingleRegexpQuery() *_regexpQuery {
	return &_regexpQuery{
		k: "",
		v: types.NewRegexpQuery(),
	}
}

func (s *_regexpQuery) RegexpQueryCaster() *types.RegexpQuery {
	return s.v.RegexpQueryCaster()
}
