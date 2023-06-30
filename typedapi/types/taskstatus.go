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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// TaskStatus type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/tasks/_types/TaskStatus.ts#L24-L42
type TaskStatus struct {
	Batches              int64    `json:"batches"`
	Canceled             *string  `json:"canceled,omitempty"`
	Created              int64    `json:"created"`
	Deleted              int64    `json:"deleted"`
	Failures             []string `json:"failures,omitempty"`
	Noops                int64    `json:"noops"`
	RequestsPerSecond    float32  `json:"requests_per_second"`
	Retries              Retries  `json:"retries"`
	Throttled            Duration `json:"throttled,omitempty"`
	ThrottledMillis      int64    `json:"throttled_millis"`
	ThrottledUntil       Duration `json:"throttled_until,omitempty"`
	ThrottledUntilMillis int64    `json:"throttled_until_millis"`
	TimedOut             *bool    `json:"timed_out,omitempty"`
	Took                 *int64   `json:"took,omitempty"`
	Total                int64    `json:"total"`
	Updated              int64    `json:"updated"`
	VersionConflicts     int64    `json:"version_conflicts"`
}

func (s *TaskStatus) UnmarshalJSON(data []byte) error {

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

		case "batches":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Batches = value
			case float64:
				f := int64(v)
				s.Batches = f
			}

		case "canceled":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Canceled = &o

		case "created":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Created = value
			case float64:
				f := int64(v)
				s.Created = f
			}

		case "deleted":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Deleted = value
			case float64:
				f := int64(v)
				s.Deleted = f
			}

		case "failures":
			if err := dec.Decode(&s.Failures); err != nil {
				return err
			}

		case "noops":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Noops = value
			case float64:
				f := int64(v)
				s.Noops = f
			}

		case "requests_per_second":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return err
				}
				f := float32(value)
				s.RequestsPerSecond = f
			case float64:
				f := float32(v)
				s.RequestsPerSecond = f
			}

		case "retries":
			if err := dec.Decode(&s.Retries); err != nil {
				return err
			}

		case "throttled":
			if err := dec.Decode(&s.Throttled); err != nil {
				return err
			}

		case "throttled_millis":
			if err := dec.Decode(&s.ThrottledMillis); err != nil {
				return err
			}

		case "throttled_until":
			if err := dec.Decode(&s.ThrottledUntil); err != nil {
				return err
			}

		case "throttled_until_millis":
			if err := dec.Decode(&s.ThrottledUntilMillis); err != nil {
				return err
			}

		case "timed_out":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.TimedOut = &value
			case bool:
				s.TimedOut = &v
			}

		case "took":
			if err := dec.Decode(&s.Took); err != nil {
				return err
			}

		case "total":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Total = value
			case float64:
				f := int64(v)
				s.Total = f
			}

		case "updated":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Updated = value
			case float64:
				f := int64(v)
				s.Updated = f
			}

		case "version_conflicts":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.VersionConflicts = value
			case float64:
				f := int64(v)
				s.VersionConflicts = f
			}

		}
	}
	return nil
}

// NewTaskStatus returns a TaskStatus.
func NewTaskStatus() *TaskStatus {
	r := &TaskStatus{}

	return r
}
