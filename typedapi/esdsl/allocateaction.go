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
// https://github.com/elastic/elasticsearch-specification/tree/86f41834c7bb975159a38a73be8a9d930010d673

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _allocateAction struct {
	v *types.AllocateAction
}

func NewAllocateAction() *_allocateAction {

	return &_allocateAction{v: types.NewAllocateAction()}

}

func (s *_allocateAction) Exclude(exclude map[string]string) *_allocateAction {

	s.v.Exclude = exclude
	return s
}

func (s *_allocateAction) AddExclude(key string, value string) *_allocateAction {

	var tmp map[string]string
	if s.v.Exclude == nil {
		s.v.Exclude = make(map[string]string)
	} else {
		tmp = s.v.Exclude
	}

	tmp[key] = value

	s.v.Exclude = tmp
	return s
}

func (s *_allocateAction) Include(include map[string]string) *_allocateAction {

	s.v.Include = include
	return s
}

func (s *_allocateAction) AddInclude(key string, value string) *_allocateAction {

	var tmp map[string]string
	if s.v.Include == nil {
		s.v.Include = make(map[string]string)
	} else {
		tmp = s.v.Include
	}

	tmp[key] = value

	s.v.Include = tmp
	return s
}

func (s *_allocateAction) NumberOfReplicas(numberofreplicas int) *_allocateAction {

	s.v.NumberOfReplicas = &numberofreplicas

	return s
}

func (s *_allocateAction) Require(require map[string]string) *_allocateAction {

	s.v.Require = require
	return s
}

func (s *_allocateAction) AddRequire(key string, value string) *_allocateAction {

	var tmp map[string]string
	if s.v.Require == nil {
		s.v.Require = make(map[string]string)
	} else {
		tmp = s.v.Require
	}

	tmp[key] = value

	s.v.Require = tmp
	return s
}

func (s *_allocateAction) TotalShardsPerNode(totalshardspernode int) *_allocateAction {

	s.v.TotalShardsPerNode = &totalshardspernode

	return s
}

func (s *_allocateAction) AllocateActionCaster() *types.AllocateAction {
	return s.v
}
