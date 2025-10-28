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

type _addAction struct {
	v *types.AddAction
}

// Adds a data stream or index to an alias.
// If the alias doesn’t exist, the `add` action creates it.
func NewAddAction() *_addAction {

	return &_addAction{v: types.NewAddAction()}

}

func (s *_addAction) Alias(indexalias string) *_addAction {

	s.v.Alias = &indexalias

	return s
}

func (s *_addAction) Aliases(aliases ...string) *_addAction {

	s.v.Aliases = make([]string, len(aliases))
	s.v.Aliases = aliases

	return s
}

func (s *_addAction) Filter(filter types.QueryVariant) *_addAction {

	s.v.Filter = filter.QueryCaster()

	return s
}

func (s *_addAction) Index(indexname string) *_addAction {

	s.v.Index = &indexname

	return s
}

func (s *_addAction) IndexRouting(routing string) *_addAction {

	s.v.IndexRouting = &routing

	return s
}

func (s *_addAction) Indices(indices ...string) *_addAction {

	s.v.Indices = indices

	return s
}

func (s *_addAction) IsHidden(ishidden bool) *_addAction {

	s.v.IsHidden = &ishidden

	return s
}

func (s *_addAction) IsWriteIndex(iswriteindex bool) *_addAction {

	s.v.IsWriteIndex = &iswriteindex

	return s
}

func (s *_addAction) MustExist(mustexist bool) *_addAction {

	s.v.MustExist = &mustexist

	return s
}

func (s *_addAction) Routing(routing string) *_addAction {

	s.v.Routing = &routing

	return s
}

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
