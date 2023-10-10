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
// https://github.com/elastic/elasticsearch-specification/tree/24afbdf78c21fde141eb2cad34491d952bd6daa8

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// TaskFailure type.
//
// https://github.com/elastic/elasticsearch-specification/blob/24afbdf78c21fde141eb2cad34491d952bd6daa8/specification/_types/Errors.ts#L66-L71
type TaskFailure struct {
	NodeId string     `json:"node_id"`
	Reason ErrorCause `json:"reason"`
	Status string     `json:"status"`
	TaskId int64      `json:"task_id"`
}

func (s *TaskFailure) UnmarshalJSON(data []byte) error {

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

		case "node_id":
			if err := dec.Decode(&s.NodeId); err != nil {
				return err
			}

		case "reason":
			if err := dec.Decode(&s.Reason); err != nil {
				return err
			}

		case "status":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Status = o

		case "task_id":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.TaskId = value
			case float64:
				f := int64(v)
				s.TaskId = f
			}

		}
	}
	return nil
}

// NewTaskFailure returns a TaskFailure.
func NewTaskFailure() *TaskFailure {
	r := &TaskFailure{}

	return r
}
