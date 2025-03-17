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

type _addAction struct {
	v *types.AddAction
}

// Adds a data stream or index to an alias.
// If the alias doesn’t exist, the `add` action creates it.
func NewAddAction() *_addAction {

	return &_addAction{v: types.NewAddAction()}

}

// Alias for the action.
// Index alias names support date math.
func (s *_addAction) Alias(indexalias string) *_addAction {

	s.v.Alias = &indexalias

	return s
}

// Aliases for the action.
// Index alias names support date math.
func (s *_addAction) Aliases(aliases ...string) *_addAction {

	s.v.Aliases = make([]string, len(aliases))
	s.v.Aliases = aliases

	return s
}

// Query used to limit documents the alias can access.
func (s *_addAction) Filter(filter types.QueryVariant) *_addAction {

	s.v.Filter = filter.QueryCaster()

	return s
}

// Data stream or index for the action.
// Supports wildcards (`*`).
func (s *_addAction) Index(indexname string) *_addAction {

	s.v.Index = &indexname

	return s
}

// Value used to route indexing operations to a specific shard.
// If specified, this overwrites the `routing` value for indexing operations.
// Data stream aliases don’t support this parameter.
func (s *_addAction) IndexRouting(routing string) *_addAction {

	s.v.IndexRouting = &routing

	return s
}

// Data streams or indices for the action.
// Supports wildcards (`*`).
func (s *_addAction) Indices(indices ...string) *_addAction {

	s.v.Indices = indices

	return s
}

// If `true`, the alias is hidden.
func (s *_addAction) IsHidden(ishidden bool) *_addAction {

	s.v.IsHidden = &ishidden

	return s
}

// If `true`, sets the write index or data stream for the alias.
func (s *_addAction) IsWriteIndex(iswriteindex bool) *_addAction {

	s.v.IsWriteIndex = &iswriteindex

	return s
}

// If `true`, the alias must exist to perform the action.
func (s *_addAction) MustExist(mustexist bool) *_addAction {

	s.v.MustExist = &mustexist

	return s
}

// Value used to route indexing and search operations to a specific shard.
// Data stream aliases don’t support this parameter.
func (s *_addAction) Routing(routing string) *_addAction {

	s.v.Routing = &routing

	return s
}

// Value used to route search operations to a specific shard.
// If specified, this overwrites the `routing` value for search operations.
// Data stream aliases don’t support this parameter.
func (s *_addAction) SearchRouting(routing string) *_addAction {

	s.v.SearchRouting = &routing

	return s
}

func (s *_addAction) IndicesActionCaster() *types.IndicesAction {
	container := types.NewIndicesAction()

	container.Add = s.v

	return container
}

func (s *_addAction) AddActionCaster() *types.AddAction {
	return s.v
}
