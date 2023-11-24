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
)

// RecoveryIndexStatus type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/indices/recovery/types.ts#L64-L74
type RecoveryIndexStatus struct {
	Bytes                      *RecoveryBytes `json:"bytes,omitempty"`
	Files                      RecoveryFiles  `json:"files"`
	Size                       RecoveryBytes  `json:"size"`
	SourceThrottleTime         Duration       `json:"source_throttle_time,omitempty"`
	SourceThrottleTimeInMillis int64          `json:"source_throttle_time_in_millis"`
	TargetThrottleTime         Duration       `json:"target_throttle_time,omitempty"`
	TargetThrottleTimeInMillis int64          `json:"target_throttle_time_in_millis"`
	TotalTime                  Duration       `json:"total_time,omitempty"`
	TotalTimeInMillis          int64          `json:"total_time_in_millis"`
}

func (s *RecoveryIndexStatus) UnmarshalJSON(data []byte) error {

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

		case "bytes":
			if err := dec.Decode(&s.Bytes); err != nil {
				return err
			}

		case "files":
			if err := dec.Decode(&s.Files); err != nil {
				return err
			}

		case "size":
			if err := dec.Decode(&s.Size); err != nil {
				return err
			}

		case "source_throttle_time":
			if err := dec.Decode(&s.SourceThrottleTime); err != nil {
				return err
			}

		case "source_throttle_time_in_millis":
			if err := dec.Decode(&s.SourceThrottleTimeInMillis); err != nil {
				return err
			}

		case "target_throttle_time":
			if err := dec.Decode(&s.TargetThrottleTime); err != nil {
				return err
			}

		case "target_throttle_time_in_millis":
			if err := dec.Decode(&s.TargetThrottleTimeInMillis); err != nil {
				return err
			}

		case "total_time":
			if err := dec.Decode(&s.TotalTime); err != nil {
				return err
			}

		case "total_time_in_millis":
			if err := dec.Decode(&s.TotalTimeInMillis); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewRecoveryIndexStatus returns a RecoveryIndexStatus.
func NewRecoveryIndexStatus() *RecoveryIndexStatus {
	r := &RecoveryIndexStatus{}

	return r
}
