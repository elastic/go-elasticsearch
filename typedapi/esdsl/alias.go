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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _alias struct {
	v *types.Alias
}

func NewAlias() *_alias {

	return &_alias{v: types.NewAlias()}

}

// Query used to limit documents the alias can access.
func (s *_alias) Filter(filter types.QueryVariant) *_alias {

	s.v.Filter = filter.QueryCaster()

	return s
}

// Value used to route indexing operations to a specific shard.
// If specified, this overwrites the `routing` value for indexing operations.
func (s *_alias) IndexRouting(routing string) *_alias {

	s.v.IndexRouting = &routing

	return s
}

// If `true`, the alias is hidden.
// All indices for the alias must have the same `is_hidden` value.
func (s *_alias) IsHidden(ishidden bool) *_alias {

	s.v.IsHidden = &ishidden

	return s
}

// If `true`, the index is the write index for the alias.
func (s *_alias) IsWriteIndex(iswriteindex bool) *_alias {

	s.v.IsWriteIndex = &iswriteindex

	return s
}

// Value used to route indexing and search operations to a specific shard.
func (s *_alias) Routing(routing string) *_alias {

	s.v.Routing = &routing

	return s
}

// Value used to route search operations to a specific shard.
// If specified, this overwrites the `routing` value for search operations.
func (s *_alias) SearchRouting(routing string) *_alias {

	s.v.SearchRouting = &routing

	return s
}

func (s *_alias) AliasCaster() *types.Alias {
	return s.v
}
