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

package cancel

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package cancel
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/tasks/cancel/CancelTasksResponse.ts#L22-L24

type Response struct {
	NodeFailures []types.ErrorCause `json:"node_failures,omitempty"`
	// Nodes Task information grouped by node, if `group_by` was set to `node` (the
	// default).
	Nodes        map[string]types.NodeTasks `json:"nodes,omitempty"`
	TaskFailures []types.TaskFailure        `json:"task_failures,omitempty"`
	// Tasks Either a flat list of tasks if `group_by` was set to `none`, or grouped by
	// parents if
	// `group_by` was set to `parents`.
	Tasks types.TaskInfos `json:"tasks,omitempty"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{
		Nodes: make(map[string]types.NodeTasks, 0),
	}
	return r
}

func (s *Response) UnmarshalJSON(data []byte) error {
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

		case "node_failures":
			if err := dec.Decode(&s.NodeFailures); err != nil {
				return err
			}

		case "nodes":
			if err := dec.Decode(&s.Nodes); err != nil {
				return err
			}

		case "task_failures":
			if err := dec.Decode(&s.TaskFailures); err != nil {
				return err
			}

		case "tasks":
			if err := dec.Decode(&s.Tasks); err != nil {
				return err
			}

		}
	}
	return nil
}
