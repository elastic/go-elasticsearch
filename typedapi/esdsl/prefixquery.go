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

type _prefixQuery struct {
	k string
	v *types.PrefixQuery
}

// Returns documents that contain a specific prefix in a provided field.
func NewPrefixQuery(field string, value string) *_prefixQuery {
	tmp := &_prefixQuery{
		k: field,
		v: types.NewPrefixQuery(),
	}

	tmp.Value(value)
	return tmp
}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_prefixQuery) Boost(boost float32) *_prefixQuery {

	s.v.Boost = &boost

	return s
}

// Allows ASCII case insensitive matching of the value with the indexed field
// values when set to `true`.
// Default is `false` which means the case sensitivity of matching depends on
// the underlying field’s mapping.
func (s *_prefixQuery) CaseInsensitive(caseinsensitive bool) *_prefixQuery {

	s.v.CaseInsensitive = &caseinsensitive

	return s
}

func (s *_prefixQuery) QueryName_(queryname_ string) *_prefixQuery {

	s.v.QueryName_ = &queryname_

	return s
}

// Method used to rewrite the query.
func (s *_prefixQuery) Rewrite(multitermqueryrewrite string) *_prefixQuery {

	s.v.Rewrite = &multitermqueryrewrite

	return s
}

// Beginning characters of terms you wish to find in the provided field.
func (s *_prefixQuery) Value(value string) *_prefixQuery {

	s.v.Value = value

	return s
}

func (s *_prefixQuery) QueryCaster() *types.Query {
	container := types.NewQuery()
	container.Prefix = map[string]types.PrefixQuery{
		s.k: *s.v,
	}
	return container
}

func (s *_prefixQuery) ApiKeyQueryContainerCaster() *types.ApiKeyQueryContainer {
	container := types.NewApiKeyQueryContainer()
	container.Prefix = map[string]types.PrefixQuery{
		s.k: *s.v,
	}
	return container
}

func (s *_prefixQuery) RoleQueryContainerCaster() *types.RoleQueryContainer {
	container := types.NewRoleQueryContainer()
	container.Prefix = map[string]types.PrefixQuery{
		s.k: *s.v,
	}
	return container
}

func (s *_prefixQuery) UserQueryContainerCaster() *types.UserQueryContainer {
	container := types.NewUserQueryContainer()
	container.Prefix = map[string]types.PrefixQuery{
		s.k: *s.v,
	}
	return container
}

// NewSinglePrefixQuery should be used when you want to
// create a single key dictionary without specifying the key in the
// constructor. Usually key is already defined within the parent container.
func NewSinglePrefixQuery() *_prefixQuery {
	return &_prefixQuery{
		k: "",
		v: types.NewPrefixQuery(),
	}
}

func (s *_prefixQuery) PrefixQueryCaster() *types.PrefixQuery {
	return s.v.PrefixQueryCaster()
}
