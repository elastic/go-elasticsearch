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
// https://github.com/elastic/elasticsearch-specification/tree/52c473efb1fb5320a5bac12572d0b285882862fb

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _indexAndDataStreamAction struct {
	v *types.IndexAndDataStreamAction
}

// Removes a backing index from a data stream.
// The index is unhidden as part of this operation.
// A data stream’s write index cannot be removed.
func NewIndexAndDataStreamAction() *_indexAndDataStreamAction {

	return &_indexAndDataStreamAction{v: types.NewIndexAndDataStreamAction()}

}

func (s *_indexAndDataStreamAction) DataStream(datastreamname string) *_indexAndDataStreamAction {

	s.v.DataStream = datastreamname

	return s
}

func (s *_indexAndDataStreamAction) Index(indexname string) *_indexAndDataStreamAction {

	s.v.Index = indexname

	return s
}

func (s *_indexAndDataStreamAction) IndicesModifyActionCaster() *types.IndicesModifyAction {
	container := types.NewIndicesModifyAction()

	container.RemoveBackingIndex = s.v

	return container
}

func (s *_indexAndDataStreamAction) IndexAndDataStreamActionCaster() *types.IndexAndDataStreamAction {
	return s.v
}
