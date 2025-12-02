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
// https://github.com/elastic/elasticsearch-specification/tree/aa1459fbdcaf57c653729142b3b6e9982373bb1c

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _commandCancelAction struct {
	v *types.CommandCancelAction
}

func NewCommandCancelAction(node string, shard int) *_commandCancelAction {

	tmp := &_commandCancelAction{v: types.NewCommandCancelAction()}

	tmp.Node(node)

	tmp.Shard(shard)

	return tmp

}

func (s *_commandCancelAction) AllowPrimary(allowprimary bool) *_commandCancelAction {

	s.v.AllowPrimary = &allowprimary

	return s
}

func (s *_commandCancelAction) Index(indexname string) *_commandCancelAction {

	s.v.Index = indexname

	return s
}

func (s *_commandCancelAction) Node(node string) *_commandCancelAction {

	s.v.Node = node

	return s
}

func (s *_commandCancelAction) Shard(shard int) *_commandCancelAction {

	s.v.Shard = shard

	return s
}

func (s *_commandCancelAction) CommandCancelActionCaster() *types.CommandCancelAction {
	return s.v
}
