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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// ParentTaskInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/tasks/_types/TaskListResponseBase.ts#L45-L47
type ParentTaskInfo struct {
	Action      string     `json:"action"`
	Cancellable bool       `json:"cancellable"`
	Cancelled   *bool      `json:"cancelled,omitempty"`
	Children    []TaskInfo `json:"children,omitempty"`
	// Description Human readable text that identifies the particular request that the task is
	// performing.
	// For example, it might identify the search request being performed by a search
	// task.
	// Other kinds of tasks have different descriptions, like `_reindex` which has
	// the source and the destination, or `_bulk` which just has the number of
	// requests and the destination indices.
	// Many requests will have only an empty description because more detailed
	// information about the request is not easily available or particularly helpful
	// in identifying the request.
	Description        *string           `json:"description,omitempty"`
	Headers            map[string]string `json:"headers"`
	Id                 int64             `json:"id"`
	Node               string            `json:"node"`
	ParentTaskId       TaskId            `json:"parent_task_id,omitempty"`
	RunningTime        Duration          `json:"running_time,omitempty"`
	RunningTimeInNanos int64             `json:"running_time_in_nanos"`
	StartTimeInMillis  int64             `json:"start_time_in_millis"`
	// Status The internal status of the task, which varies from task to task.
	// The format also varies.
	// While the goal is to keep the status for a particular task consistent from
	// version to version, this is not always possible because sometimes the
	// implementation changes.
	// Fields might be removed from the status for a particular request so any
	// parsing you do of the status might break in minor releases.
	Status json.RawMessage `json:"status,omitempty"`
	Type   string          `json:"type"`
}

func (s *ParentTaskInfo) UnmarshalJSON(data []byte) error {

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
				return fmt.Errorf("%s | %w", "Action", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Action = o

		case "cancellable":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Cancellable", err)
				}
				s.Cancellable = value
			case bool:
				s.Cancellable = v
			}

		case "cancelled":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Cancelled", err)
				}
				s.Cancelled = &value
			case bool:
				s.Cancelled = &v
			}

		case "children":
			if err := dec.Decode(&s.Children); err != nil {
				return fmt.Errorf("%s | %w", "Children", err)
			}

		case "description":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Description", err)
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
				return fmt.Errorf("%s | %w", "Headers", err)
			}

		case "id":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Id", err)
				}
				s.Id = value
			case float64:
				f := int64(v)
				s.Id = f
			}

		case "node":
			if err := dec.Decode(&s.Node); err != nil {
				return fmt.Errorf("%s | %w", "Node", err)
			}

		case "parent_task_id":
			if err := dec.Decode(&s.ParentTaskId); err != nil {
				return fmt.Errorf("%s | %w", "ParentTaskId", err)
			}

		case "running_time":
			if err := dec.Decode(&s.RunningTime); err != nil {
				return fmt.Errorf("%s | %w", "RunningTime", err)
			}

		case "running_time_in_nanos":
			if err := dec.Decode(&s.RunningTimeInNanos); err != nil {
				return fmt.Errorf("%s | %w", "RunningTimeInNanos", err)
			}

		case "start_time_in_millis":
			if err := dec.Decode(&s.StartTimeInMillis); err != nil {
				return fmt.Errorf("%s | %w", "StartTimeInMillis", err)
			}

		case "status":
			if err := dec.Decode(&s.Status); err != nil {
				return fmt.Errorf("%s | %w", "Status", err)
			}

		case "type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
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

// NewParentTaskInfo returns a ParentTaskInfo.
func NewParentTaskInfo() *ParentTaskInfo {
	r := &ParentTaskInfo{
		Headers: make(map[string]string),
	}

	return r
}
