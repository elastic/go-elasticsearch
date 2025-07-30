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
// https://github.com/elastic/elasticsearch-specification/tree/de4ff9ec1f716256f521d9e30011ad9c284b0dcc

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _removeIndexAction struct {
	v *types.RemoveIndexAction
}

// Deletes an index.
// You cannot use this action on aliases or data streams.
func NewRemoveIndexAction() *_removeIndexAction {

	return &_removeIndexAction{v: types.NewRemoveIndexAction()}

}

func (s *_removeIndexAction) Index(indexname string) *_removeIndexAction {

	s.v.Index = &indexname

	return s
}

func (s *_removeIndexAction) Indices(indices ...string) *_removeIndexAction {

	s.v.Indices = indices

	return s
}

func (s *_removeIndexAction) MustExist(mustexist bool) *_removeIndexAction {

	s.v.MustExist = &mustexist

	return s
}

func (s *_removeIndexAction) IndicesActionCaster() *types.IndicesAction {
	container := types.NewIndicesAction()

	container.RemoveIndex = s.v

	return container
}

func (s *_removeIndexAction) RemoveIndexActionCaster() *types.RemoveIndexAction {
	return s.v
}
