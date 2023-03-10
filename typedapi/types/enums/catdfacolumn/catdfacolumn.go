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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

// Package catdfacolumn
package catdfacolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/cat/_types/CatBase.ts#L472-L557
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

	Nodeaddress = CatDfaColumn{"node.address"}

	Nodeephemeralid = CatDfaColumn{"node.ephemeral_id"}

	Nodeid = CatDfaColumn{"node.id"}

	Nodename = CatDfaColumn{"node.name"}

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
		*c = Nodeaddress
	case "node.ephemeral_id":
		*c = Nodeephemeralid
	case "node.id":
		*c = Nodeid
	case "node.name":
		*c = Nodename
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
