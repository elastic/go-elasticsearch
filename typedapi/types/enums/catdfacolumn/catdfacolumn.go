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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


// Package catdfacolumn
package catdfacolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/cat/_types/CatBase.ts#L472-L557
type CatDfaColumn struct {
	Name string
}

var (
	Assignmentexplanation = CatDfaColumn{"assignment_explanation"}

	Createtime = CatDfaColumn{"create_time"}

	Description = CatDfaColumn{"description"}

	Destindex = CatDfaColumn{"dest_index"}

	Failurereason = CatDfaColumn{"failure_reason"}

	Id = CatDfaColumn{"id"}

	Modelmemorylimit = CatDfaColumn{"model_memory_limit"}

	NodeAddress = CatDfaColumn{"node.address"}

	NodeEphemeralid = CatDfaColumn{"node.ephemeral_id"}

	NodeId = CatDfaColumn{"node.id"}

	NodeName = CatDfaColumn{"node.name"}

	Progress = CatDfaColumn{"progress"}

	Sourceindex = CatDfaColumn{"source_index"}

	State = CatDfaColumn{"state"}

	Type = CatDfaColumn{"type"}

	Version = CatDfaColumn{"version"}
)

func (c CatDfaColumn) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CatDfaColumn) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "assignment_explanation":
		*c = Assignmentexplanation
	case "create_time":
		*c = Createtime
	case "description":
		*c = Description
	case "dest_index":
		*c = Destindex
	case "failure_reason":
		*c = Failurereason
	case "id":
		*c = Id
	case "model_memory_limit":
		*c = Modelmemorylimit
	case "node.address":
		*c = NodeAddress
	case "node.ephemeral_id":
		*c = NodeEphemeralid
	case "node.id":
		*c = NodeId
	case "node.name":
		*c = NodeName
	case "progress":
		*c = Progress
	case "source_index":
		*c = Sourceindex
	case "state":
		*c = State
	case "type":
		*c = Type
	case "version":
		*c = Version
	default:
		*c = CatDfaColumn{string(text)}
	}

	return nil
}

func (c CatDfaColumn) String() string {
	return c.Name
}
