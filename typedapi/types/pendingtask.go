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

// PendingTask type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/cluster/pending_tasks/types.ts#L23-L47
type PendingTask struct {
	// Executing Indicates whether the pending tasks are currently executing or not.
	Executing bool `json:"executing"`
	// InsertOrder The number that represents when the task has been inserted into the task
	// queue.
	InsertOrder int `json:"insert_order"`
	// Priority The priority of the pending task.
	// The valid priorities in descending priority order are: `IMMEDIATE` > `URGENT`
	// > `HIGH` > `NORMAL` > `LOW` > `LANGUID`.
	Priority string `json:"priority"`
	// Source A general description of the cluster task that may include a reason and
	// origin.
	Source string `json:"source"`
	// TimeInQueue The time since the task is waiting for being performed.
	TimeInQueue Duration `json:"time_in_queue,omitempty"`
	// TimeInQueueMillis The time expressed in milliseconds since the task is waiting for being
	// performed.
	TimeInQueueMillis int64 `json:"time_in_queue_millis"`
}

func (s *PendingTask) UnmarshalJSON(data []byte) error {

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

		case "executing":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Executing = value
			case bool:
				s.Executing = v
			}

		case "insert_order":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.InsertOrder = value
			case float64:
				f := int(v)
				s.InsertOrder = f
			}

		case "priority":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Priority = o

		case "source":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Source = o

		case "time_in_queue":
			if err := dec.Decode(&s.TimeInQueue); err != nil {
				return err
			}

		case "time_in_queue_millis":
			if err := dec.Decode(&s.TimeInQueueMillis); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewPendingTask returns a PendingTask.
func NewPendingTask() *PendingTask {
	r := &PendingTask{}

	return r
}
