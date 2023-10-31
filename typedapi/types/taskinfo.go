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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// TaskInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/tasks/_types/TaskInfo.ts#L32-L47
type TaskInfo struct {
	Action             string            `json:"action"`
	Cancellable        bool              `json:"cancellable"`
	Cancelled          *bool             `json:"cancelled,omitempty"`
	Description        *string           `json:"description,omitempty"`
	Headers            map[string]string `json:"headers"`
	Id                 int64             `json:"id"`
	Node               string            `json:"node"`
	ParentTaskId       TaskId            `json:"parent_task_id,omitempty"`
	RunningTime        Duration          `json:"running_time,omitempty"`
	RunningTimeInNanos int64             `json:"running_time_in_nanos"`
	StartTimeInMillis  int64             `json:"start_time_in_millis"`
	// Status Task status information can vary wildly from task to task.
	Status json.RawMessage `json:"status,omitempty"`
	Type   string          `json:"type"`
}

func (s *TaskInfo) UnmarshalJSON(data []byte) error {

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

		case "action":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Action = o

		case "cancellable":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Cancellable = value
			case bool:
				s.Cancellable = v
			}

		case "cancelled":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Cancelled = &value
			case bool:
				s.Cancelled = &v
			}

		case "description":
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

		case "headers":
			if s.Headers == nil {
				s.Headers = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.Headers); err != nil {
				return err
			}

		case "id":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Id = value
			case float64:
				f := int64(v)
				s.Id = f
			}

		case "node":
			if err := dec.Decode(&s.Node); err != nil {
				return err
			}

		case "parent_task_id":
			if err := dec.Decode(&s.ParentTaskId); err != nil {
				return err
			}

		case "running_time":
			if err := dec.Decode(&s.RunningTime); err != nil {
				return err
			}

		case "running_time_in_nanos":
			if err := dec.Decode(&s.RunningTimeInNanos); err != nil {
				return err
			}

		case "start_time_in_millis":
			if err := dec.Decode(&s.StartTimeInMillis); err != nil {
				return err
			}

		case "status":
			if err := dec.Decode(&s.Status); err != nil {
				return err
			}

		case "type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Type = o

		}
	}
	return nil
}

// NewTaskInfo returns a TaskInfo.
func NewTaskInfo() *TaskInfo {
	r := &TaskInfo{
		Headers: make(map[string]string, 0),
	}

	return r
}
