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

package deletebyqueryrethrottle

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package deletebyqueryrethrottle
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_global/delete_by_query_rethrottle/DeleteByQueryRethrottleResponse.ts#L22-L24
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
			if s.Nodes == nil {
				s.Nodes = make(map[string]types.NodeTasks, 0)
			}
			if err := dec.Decode(&s.Nodes); err != nil {
				return err
			}

		case "task_failures":
			if err := dec.Decode(&s.TaskFailures); err != nil {
				return err
			}

		case "tasks":

			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			source := bytes.NewReader(rawMsg)
			localDec := json.NewDecoder(source)
			switch rawMsg[0] {
			case '{':
				o := make(map[string]types.ParentTaskInfo, 0)
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Tasks = o
			case '[':
				o := []types.TaskInfo{}
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Tasks = o
			}

		}
	}
	return nil
}
