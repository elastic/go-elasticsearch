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

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_regexpQuery) Boost(boost float32) *_regexpQuery {

	s.v.Boost = &boost

	return s
}

// Allows case insensitive matching of the regular expression value with the
// indexed field values when set to `true`.
// When `false`, case sensitivity of matching depends on the underlying field’s
// mapping.
func (s *_regexpQuery) CaseInsensitive(caseinsensitive bool) *_regexpQuery {

	s.v.CaseInsensitive = &caseinsensitive

	return s
}

// Enables optional operators for the regular expression.
func (s *_regexpQuery) Flags(flags string) *_regexpQuery {

	s.v.Flags = &flags

	return s
}

// Maximum number of automaton states required for the query.
func (s *_regexpQuery) MaxDeterminizedStates(maxdeterminizedstates int) *_regexpQuery {

	s.v.MaxDeterminizedStates = &maxdeterminizedstates

	return s
}

func (s *_regexpQuery) QueryName_(queryname_ string) *_regexpQuery {

	s.v.QueryName_ = &queryname_

	return s
}

// Method used to rewrite the query.
func (s *_regexpQuery) Rewrite(multitermqueryrewrite string) *_regexpQuery {

	s.v.Rewrite = &multitermqueryrewrite

	return s
}

// Regular expression for terms you wish to find in the provided field.
func (s *_regexpQuery) Value(value string) *_regexpQuery {

	s.v.Value = value

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
