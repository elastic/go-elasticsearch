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
// https://github.com/elastic/elasticsearch-specification/tree/e585438d116b00ff34643179e6286e402c0bcaaf

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _commandAllocatePrimaryAction struct {
	v *types.CommandAllocatePrimaryAction
}

func NewCommandAllocatePrimaryAction(acceptdataloss bool, node string, shard int) *_commandAllocatePrimaryAction {

	tmp := &_commandAllocatePrimaryAction{v: types.NewCommandAllocatePrimaryAction()}

	tmp.AcceptDataLoss(acceptdataloss)

	tmp.Node(node)

	tmp.Shard(shard)

	return tmp

}

func (s *_commandAllocatePrimaryAction) AcceptDataLoss(acceptdataloss bool) *_commandAllocatePrimaryAction {

	s.v.AcceptDataLoss = acceptdataloss

	return s
}

func (s *_commandAllocatePrimaryAction) Index(indexname string) *_commandAllocatePrimaryAction {

	s.v.Index = indexname

	return s
}

func (s *_commandAllocatePrimaryAction) Node(node string) *_commandAllocatePrimaryAction {

	s.v.Node = node

	return s
}

func (s *_commandAllocatePrimaryAction) Shard(shard int) *_commandAllocatePrimaryAction {

	s.v.Shard = shard

	return s
}

func (s *_commandAllocatePrimaryAction) CommandAllocatePrimaryActionCaster() *types.CommandAllocatePrimaryAction {
	return s.v
}
