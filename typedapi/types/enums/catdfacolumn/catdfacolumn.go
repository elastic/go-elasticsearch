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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

// Package catdfacolumn
package catdfacolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/cat/_types/CatBase.ts#L480-L565
type CatDfaColumn struct {
	Name string
}

var (

	// Assignmentexplanation Contains messages relating to the selection of a node.
	Assignmentexplanation = CatDfaColumn{"assignment_explanation"}

	// Createtime The time when the data frame analytics job was created.
	Createtime = CatDfaColumn{"create_time"}

	// Description A description of a job.
	Description = CatDfaColumn{"description"}

	// Destindex Name of the destination index.
	Destindex = CatDfaColumn{"dest_index"}

	// Failurereason Contains messages about the reason why a data frame analytics job failed.
	Failurereason = CatDfaColumn{"failure_reason"}

	// Id Identifier for the data frame analytics job.
	Id = CatDfaColumn{"id"}

	// Modelmemorylimit The approximate maximum amount of memory resources that are permitted for the
	// data frame analytics job.
	Modelmemorylimit = CatDfaColumn{"model_memory_limit"}

	// Nodeaddress The network address of the node that the data frame analytics job is assigned
	// to.
	Nodeaddress = CatDfaColumn{"node.address"}

	// Nodeephemeralid The ephemeral ID of the node that the data frame analytics job is assigned
	// to.
	Nodeephemeralid = CatDfaColumn{"node.ephemeral_id"}

	// Nodeid The unique identifier of the node that the data frame analytics job is
	// assigned to.
	Nodeid = CatDfaColumn{"node.id"}

	// Nodename The name of the node that the data frame analytics job is assigned to.
	Nodename = CatDfaColumn{"node.name"}

	// Progress The progress report of the data frame analytics job by phase.
	Progress = CatDfaColumn{"progress"}

	// Sourceindex Name of the source index.
	Sourceindex = CatDfaColumn{"source_index"}

	// State Current state of the data frame analytics job.
	State = CatDfaColumn{"state"}

	// Type The type of analysis that the data frame analytics job performs.
	Type = CatDfaColumn{"type"}

	// Version The Elasticsearch version number in which the data frame analytics job was
	// created.
	Version = CatDfaColumn{"version"}
)

func (c CatDfaColumn) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CatDfaColumn) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

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
