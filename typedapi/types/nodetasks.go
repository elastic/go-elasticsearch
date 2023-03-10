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

package types

// NodeTasks type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/tasks/_types/TaskListResponseBase.ts#L49-L57
type NodeTasks struct {
	Attributes       map[string]string   `json:"attributes,omitempty"`
	Host             *string             `json:"host,omitempty"`
	Ip               *string             `json:"ip,omitempty"`
	Name             *string             `json:"name,omitempty"`
	Roles            []string            `json:"roles,omitempty"`
	Tasks            map[TaskId]TaskInfo `json:"tasks"`
	TransportAddress *string             `json:"transport_address,omitempty"`
}

// NewNodeTasks returns a NodeTasks.
func NewNodeTasks() *NodeTasks {
	r := &NodeTasks{
		Attributes: make(map[string]string, 0),
		Tasks:      make(map[TaskId]TaskInfo, 0),
	}

	return r
}
