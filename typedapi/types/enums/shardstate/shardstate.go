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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

// Package shardstate
package shardstate

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/cluster/stats/types.ts#L689-L698
type ShardState struct {
	Name string
}

var (
	INIT = ShardState{"INIT"}

	SUCCESS = ShardState{"SUCCESS"}

	FAILED = ShardState{"FAILED"}

	ABORTED = ShardState{"ABORTED"}

	MISSING = ShardState{"MISSING"}

	WAITING = ShardState{"WAITING"}

	QUEUED = ShardState{"QUEUED"}

	PAUSEDFORNODEREMOVAL = ShardState{"PAUSED_FOR_NODE_REMOVAL"}
)

func (s ShardState) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}

func (s *ShardState) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "init":
		*s = INIT
	case "success":
		*s = SUCCESS
	case "failed":
		*s = FAILED
	case "aborted":
		*s = ABORTED
	case "missing":
		*s = MISSING
	case "waiting":
		*s = WAITING
	case "queued":
		*s = QUEUED
	case "paused_for_node_removal":
		*s = PAUSEDFORNODEREMOVAL
	default:
		*s = ShardState{string(text)}
	}

	return nil
}

func (s ShardState) String() string {
	return s.Name
}
