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
// https://github.com/elastic/elasticsearch-specification/tree/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27

// Package cattaskscolumn
package cattaskscolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27/specification/cat/_types/CatBase.ts#L2216-L2292
type CatTasksColumn struct {
	Name string
}

var (
	Id = CatTasksColumn{"id"}

	Action = CatTasksColumn{"action"}

	Taskid = CatTasksColumn{"task_id"}

	Parenttaskid = CatTasksColumn{"parent_task_id"}

	Type = CatTasksColumn{"type"}

	Starttime = CatTasksColumn{"start_time"}

	Timestamp = CatTasksColumn{"timestamp"}

	Runningtimens = CatTasksColumn{"running_time_ns"}

	Runningtime = CatTasksColumn{"running_time"}

	Nodeid = CatTasksColumn{"node_id"}

	Ip = CatTasksColumn{"ip"}

	Port = CatTasksColumn{"port"}

	Node = CatTasksColumn{"node"}

	Version = CatTasksColumn{"version"}

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
