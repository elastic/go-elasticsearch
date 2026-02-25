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

// Package cattaskscolumn
package cattaskscolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/cat/_types/CatBase.ts#L2216-L2292
type CatTasksColumn struct {
	Name string
}

var (

	// Id The ID of the task with the node.
	Id = CatTasksColumn{"id"}

	// Action The task action.
	Action = CatTasksColumn{"action"}

	// Taskid The unique task ID.
	Taskid = CatTasksColumn{"task_id"}

	// Parenttaskid The parent task ID.
	Parenttaskid = CatTasksColumn{"parent_task_id"}

	// Type The task type.
	Type = CatTasksColumn{"type"}

	// Starttime The start time in milliseconds.
	Starttime = CatTasksColumn{"start_time"}

	// Timestamp The start time in HH:MM:SS.
	Timestamp = CatTasksColumn{"timestamp"}

	// Runningtimens The running time in nanoseconds.
	Runningtimens = CatTasksColumn{"running_time_ns"}

	// Runningtime The running time.
	Runningtime = CatTasksColumn{"running_time"}

	// Nodeid The unique node ID.
	Nodeid = CatTasksColumn{"node_id"}

	// Ip The IP address.
	Ip = CatTasksColumn{"ip"}

	// Port The bound transport port.
	Port = CatTasksColumn{"port"}

	// Node The node name.
	Node = CatTasksColumn{"node"}

	// Version The Elasticsearch version.
	Version = CatTasksColumn{"version"}

	// Xopaqueid The X-Opaque-ID header.
	Xopaqueid = CatTasksColumn{"x_opaque_id"}
)

func (c CatTasksColumn) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CatTasksColumn) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "id":
		*c = Id
	case "action":
		*c = Action
	case "task_id":
		*c = Taskid
	case "parent_task_id":
		*c = Parenttaskid
	case "type":
		*c = Type
	case "start_time":
		*c = Starttime
	case "timestamp":
		*c = Timestamp
	case "running_time_ns":
		*c = Runningtimens
	case "running_time":
		*c = Runningtime
	case "node_id":
		*c = Nodeid
	case "ip":
		*c = Ip
	case "port":
		*c = Port
	case "node":
		*c = Node
	case "version":
		*c = Version
	case "x_opaque_id":
		*c = Xopaqueid
	default:
		*c = CatTasksColumn{string(text)}
	}

	return nil
}

func (c CatTasksColumn) String() string {
	return c.Name
}
