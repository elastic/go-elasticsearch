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
// https://github.com/elastic/elasticsearch-specification/tree/abf9c2c6bb21328339daa197aae15af2ecbc46f0

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// FlushStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/abf9c2c6bb21328339daa197aae15af2ecbc46f0/specification/_types/Stats.ts#L148-L155
type FlushStats struct {
	Periodic                                int64    `json:"periodic"`
	Total                                   int64    `json:"total"`
	TotalTime                               Duration `json:"total_time,omitempty"`
	TotalTimeExcludingWaiting               Duration `json:"total_time_excluding_waiting,omitempty"`
	TotalTimeExcludingWaitingOnLockInMillis int64    `json:"total_time_excluding_waiting_on_lock_in_millis"`
	TotalTimeInMillis                       int64    `json:"total_time_in_millis"`
}

func (s *FlushStats) UnmarshalJSON(data []byte) error {

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

		case "periodic":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Periodic", err)
				}
				s.Periodic = value
			case float64:
				f := int64(v)
				s.Periodic = f
			}

		case "total":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Total", err)
				}
				s.Total = value
			case float64:
				f := int64(v)
				s.Total = f
			}

		case "total_time":
			if err := dec.Decode(&s.TotalTime); err != nil {
				return fmt.Errorf("%s | %w", "TotalTime", err)
			}

		case "total_time_excluding_waiting":
			if err := dec.Decode(&s.TotalTimeExcludingWaiting); err != nil {
				return fmt.Errorf("%s | %w", "TotalTimeExcludingWaiting", err)
			}

		case "total_time_excluding_waiting_on_lock_in_millis":
			if err := dec.Decode(&s.TotalTimeExcludingWaitingOnLockInMillis); err != nil {
				return fmt.Errorf("%s | %w", "TotalTimeExcludingWaitingOnLockInMillis", err)
			}

		case "total_time_in_millis":
			if err := dec.Decode(&s.TotalTimeInMillis); err != nil {
				return fmt.Errorf("%s | %w", "TotalTimeInMillis", err)
			}

		}
	}
	return nil
}

// NewFlushStats returns a FlushStats.
func NewFlushStats() *FlushStats {
	r := &FlushStats{}

	return r
}
