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
// https://github.com/elastic/elasticsearch-specification/tree/37285cbd3fd155f913b50d880b40ec45f9df64b3

package reindexrethrottle

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

// Response holds the response body struct for the package reindexrethrottle
//
// https://github.com/elastic/elasticsearch-specification/blob/37285cbd3fd155f913b50d880b40ec45f9df64b3/specification/_global/reindex_rethrottle/ReindexRethrottleResponse.ts#L24-L54
type Response struct {
	// NodeFailures Node-level failures encountered while applying the rethrottle request. Will
	// return a `failed_node_exception` wrapping a `no_such_node_exception`, if a
	// node handling the task either never existed, or has left the cluster, and one
	// of the following is true: 1. The task has completed. 2. The task cannot be
	// found.
	//
	// Note: Rethrottle handles relocations, so it should succeed if the task can be
	// found and has not completed.
	NodeFailures []types.ErrorCause `json:"node_failures,omitempty"`
	// Nodes Tasks grouped by node, returned only when `group_by=nodes` (the default).
	Nodes map[string]types.ReindexNode `json:"nodes,omitempty"`
	// TaskFailures Per-task failures encountered while applying the rethrottle. If a rethrottle
	// is attempted during a relocation handoff, the failure object reports `status:
	// SERVICE_UNAVAILABLE` (the HTTP response itself is still `200 OK`). In this
	// case, the request can be retried until success.
	TaskFailures []types.TaskFailure `json:"task_failures,omitempty"`
	// Tasks The tasks that were successfully rethrottled. Always returned in serverless.
	// Returned with `group_by=none` or `group_by=parents` in stack.
	Tasks types.ReindexTasks `json:"tasks,omitempty"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{
		Nodes: make(map[string]types.ReindexNode, 0),
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
				return fmt.Errorf("%s | %w", "NodeFailures", err)
			}

		case "nodes":
			if s.Nodes == nil {
				s.Nodes = make(map[string]types.ReindexNode, 0)
			}
			if err := dec.Decode(&s.Nodes); err != nil {
				return fmt.Errorf("%s | %w", "Nodes", err)
			}

		case "task_failures":
			if err := dec.Decode(&s.TaskFailures); err != nil {
				return fmt.Errorf("%s | %w", "TaskFailures", err)
			}

		case "tasks":

			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			source := bytes.NewReader(rawMsg)
			localDec := json.NewDecoder(source)
			switch rawMsg[0] {
			case '{':
				o := make(map[string]types.ParentReindexTask, 0)
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Tasks", err)
				}
				s.Tasks = o
			case '[':
				o := []types.ReindexTask{}
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Tasks", err)
				}
				s.Tasks = o
			}

		}
	}
	return nil
}
