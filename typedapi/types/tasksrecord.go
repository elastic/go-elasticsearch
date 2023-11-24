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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// TasksRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/cat/tasks/types.ts#L22-L101
type TasksRecord struct {
	// Action The task action.
	Action *string `json:"action,omitempty"`
	// Description The task action description.
	Description *string `json:"description,omitempty"`
	// Id The identifier of the task with the node.
	Id *string `json:"id,omitempty"`
	// Ip The IP address for the node.
	Ip *string `json:"ip,omitempty"`
	// Node The node name.
	Node *string `json:"node,omitempty"`
	// NodeId The unique node identifier.
	NodeId *string `json:"node_id,omitempty"`
	// ParentTaskId The parent task identifier.
	ParentTaskId *string `json:"parent_task_id,omitempty"`
	// Port The bound transport port for the node.
	Port *string `json:"port,omitempty"`
	// RunningTime The running time.
	RunningTime *string `json:"running_time,omitempty"`
	// RunningTimeNs The running time in nanoseconds.
	RunningTimeNs *string `json:"running_time_ns,omitempty"`
	// StartTime The start time in milliseconds.
	StartTime *string `json:"start_time,omitempty"`
	// TaskId The unique task identifier.
	TaskId *string `json:"task_id,omitempty"`
	// Timestamp The start time in `HH:MM:SS` format.
	Timestamp *string `json:"timestamp,omitempty"`
	// Type The task type.
	Type *string `json:"type,omitempty"`
	// Version The Elasticsearch version.
	Version *string `json:"version,omitempty"`
	// XOpaqueId The X-Opaque-ID header.
	XOpaqueId *string `json:"x_opaque_id,omitempty"`
}

func (s *TasksRecord) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "action", "ac":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Action = &o

		case "description", "desc":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Description = &o

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return err
			}

		case "ip", "i":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Ip = &o

		case "node", "n":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Node = &o

		case "node_id", "ni":
			if err := dec.Decode(&s.NodeId); err != nil {
				return err
			}

		case "parent_task_id", "pti":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ParentTaskId = &o

		case "port", "po":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Port = &o

		case "running_time", "time":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.RunningTime = &o

		case "running_time_ns":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.RunningTimeNs = &o

		case "start_time", "start":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.StartTime = &o

		case "task_id", "ti":
			if err := dec.Decode(&s.TaskId); err != nil {
				return err
			}

		case "timestamp", "ts", "hms", "hhmmss":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Timestamp = &o

		case "type", "ty":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Type = &o

		case "version", "v":
			if err := dec.Decode(&s.Version); err != nil {
				return err
			}

		case "x_opaque_id", "x":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.XOpaqueId = &o

		}
	}
	return nil
}

// NewTasksRecord returns a TasksRecord.
func NewTasksRecord() *TasksRecord {
	r := &TasksRecord{}

	return r
}
