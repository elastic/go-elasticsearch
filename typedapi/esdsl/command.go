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

type _command struct {
	v *types.Command
}

func NewCommand() *_command {

	return &_command{v: types.NewCommand()}

}

func (s *_command) AllocateEmptyPrimary(allocateemptyprimary types.CommandAllocatePrimaryActionVariant) *_command {

	s.v.AllocateEmptyPrimary = allocateemptyprimary.CommandAllocatePrimaryActionCaster()

	return s
}

func (s *_command) AllocateReplica(allocatereplica types.CommandAllocateReplicaActionVariant) *_command {

	s.v.AllocateReplica = allocatereplica.CommandAllocateReplicaActionCaster()

	return s
}

func (s *_command) AllocateStalePrimary(allocatestaleprimary types.CommandAllocatePrimaryActionVariant) *_command {

	s.v.AllocateStalePrimary = allocatestaleprimary.CommandAllocatePrimaryActionCaster()

	return s
}

func (s *_command) Cancel(cancel types.CommandCancelActionVariant) *_command {

	s.v.Cancel = cancel.CommandCancelActionCaster()

	return s
}

func (s *_command) Move(move types.CommandMoveActionVariant) *_command {

	s.v.Move = move.CommandMoveActionCaster()

	return s
}

func (s *_command) CommandCaster() *types.Command {
	return s.v
}
