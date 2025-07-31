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

// ThreadCount type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/nodes/_types/Stats.ts#L1082-L1107
type ThreadCount struct {
	// Active Number of active threads in the thread pool.
	Active *int64 `json:"active,omitempty"`
	// Completed Number of tasks completed by the thread pool executor.
	Completed *int64 `json:"completed,omitempty"`
	// Largest Highest number of active threads in the thread pool.
	Largest *int64 `json:"largest,omitempty"`
	// Queue Number of tasks in queue for the thread pool.
	Queue *int64 `json:"queue,omitempty"`
	// Rejected Number of tasks rejected by the thread pool executor.
	Rejected *int64 `json:"rejected,omitempty"`
	// Threads Number of threads in the thread pool.
	Threads *int64 `json:"threads,omitempty"`
}

func (s *ThreadCount) UnmarshalJSON(data []byte) error {

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

		case "active":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Active", err)
				}
				s.Active = &value
			case float64:
				f := int64(v)
				s.Active = &f
			}

		case "completed":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Completed", err)
				}
				s.Completed = &value
			case float64:
				f := int64(v)
				s.Completed = &f
			}

		case "largest":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Largest", err)
				}
				s.Largest = &value
			case float64:
				f := int64(v)
				s.Largest = &f
			}

		case "queue":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Queue", err)
				}
				s.Queue = &value
			case float64:
				f := int64(v)
				s.Queue = &f
			}

		case "rejected":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Rejected", err)
				}
				s.Rejected = &value
			case float64:
				f := int64(v)
				s.Rejected = &f
			}

		case "threads":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Threads", err)
				}
				s.Threads = &value
			case float64:
				f := int64(v)
				s.Threads = &f
			}

		}
	}
	return nil
}

// NewThreadCount returns a ThreadCount.
func NewThreadCount() *ThreadCount {
	r := &ThreadCount{}

	return r
}
