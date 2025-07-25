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
// https://github.com/elastic/elasticsearch-specification/tree/cf6914e80d9c586e872b7d5e9e74ca34905dcf5f

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/optype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/refresh"
)

type _indexAction struct {
	v *types.IndexAction
}

func NewIndexAction() *_indexAction {

	return &_indexAction{v: types.NewIndexAction()}

}

func (s *_indexAction) DocId(id string) *_indexAction {

	s.v.DocId = &id

	return s
}

func (s *_indexAction) ExecutionTimeField(field string) *_indexAction {

	s.v.ExecutionTimeField = &field

	return s
}

func (s *_indexAction) Index(indexname string) *_indexAction {

	s.v.Index = indexname

	return s
}

func (s *_indexAction) OpType(optype optype.OpType) *_indexAction {

	s.v.OpType = &optype
	return s
}

func (s *_indexAction) Refresh(refresh refresh.Refresh) *_indexAction {

	s.v.Refresh = &refresh
	return s
}

func (s *_indexAction) Timeout(duration types.DurationVariant) *_indexAction {

	s.v.Timeout = *duration.DurationCaster()

	return s
}

func (s *_indexAction) IndexActionCaster() *types.IndexAction {
	return s.v
}
