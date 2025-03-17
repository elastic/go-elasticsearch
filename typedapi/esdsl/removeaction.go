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

type _removeAction struct {
	v *types.RemoveAction
}

// Removes a data stream or index from an alias.
func NewRemoveAction() *_removeAction {

	return &_removeAction{v: types.NewRemoveAction()}

}

// Alias for the action.
// Index alias names support date math.
func (s *_removeAction) Alias(indexalias string) *_removeAction {

	s.v.Alias = &indexalias

	return s
}

// Aliases for the action.
// Index alias names support date math.
func (s *_removeAction) Aliases(aliases ...string) *_removeAction {

	s.v.Aliases = make([]string, len(aliases))
	s.v.Aliases = aliases

	return s
}

// Data stream or index for the action.
// Supports wildcards (`*`).
func (s *_removeAction) Index(indexname string) *_removeAction {

	s.v.Index = &indexname

	return s
}

// Data streams or indices for the action.
// Supports wildcards (`*`).
func (s *_removeAction) Indices(indices ...string) *_removeAction {

	s.v.Indices = indices

	return s
}

// If `true`, the alias must exist to perform the action.
func (s *_removeAction) MustExist(mustexist bool) *_removeAction {

	s.v.MustExist = &mustexist

	return s
}

func (s *_removeAction) IndicesActionCaster() *types.IndicesAction {
	container := types.NewIndicesAction()

	container.Remove = s.v

	return container
}

func (s *_removeAction) RemoveActionCaster() *types.RemoveAction {
	return s.v
}
