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

// MergesStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/Stats.ts#L161-L178
type MergesStats struct {
	Current                    int64    `json:"current"`
	CurrentDocs                int64    `json:"current_docs"`
	CurrentSize                *string  `json:"current_size,omitempty"`
	CurrentSizeInBytes         int64    `json:"current_size_in_bytes"`
	Total                      int64    `json:"total"`
	TotalAutoThrottle          *string  `json:"total_auto_throttle,omitempty"`
	TotalAutoThrottleInBytes   int64    `json:"total_auto_throttle_in_bytes"`
	TotalDocs                  int64    `json:"total_docs"`
	TotalSize                  *string  `json:"total_size,omitempty"`
	TotalSizeInBytes           int64    `json:"total_size_in_bytes"`
	TotalStoppedTime           Duration `json:"total_stopped_time,omitempty"`
	TotalStoppedTimeInMillis   int64    `json:"total_stopped_time_in_millis"`
	TotalThrottledTime         Duration `json:"total_throttled_time,omitempty"`
	TotalThrottledTimeInMillis int64    `json:"total_throttled_time_in_millis"`
	TotalTime                  Duration `json:"total_time,omitempty"`
	TotalTimeInMillis          int64    `json:"total_time_in_millis"`
}

func (s *MergesStats) UnmarshalJSON(data []byte) error {

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

		case "current":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Current = value
			case float64:
				f := int64(v)
				s.Current = f
			}

		case "current_docs":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.CurrentDocs = value
			case float64:
				f := int64(v)
				s.CurrentDocs = f
			}

		case "current_size":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.CurrentSize = &o

		case "current_size_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.CurrentSizeInBytes = value
			case float64:
				f := int64(v)
				s.CurrentSizeInBytes = f
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

		case "total_auto_throttle":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.TotalAutoThrottle = &o

		case "total_auto_throttle_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.TotalAutoThrottleInBytes = value
			case float64:
				f := int64(v)
				s.TotalAutoThrottleInBytes = f
			}

		case "total_docs":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.TotalDocs = value
			case float64:
				f := int64(v)
				s.TotalDocs = f
			}

		case "total_size":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.TotalSize = &o

		case "total_size_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.TotalSizeInBytes = value
			case float64:
				f := int64(v)
				s.TotalSizeInBytes = f
			}

		case "total_stopped_time":
			if err := dec.Decode(&s.TotalStoppedTime); err != nil {
				return err
			}

		case "total_stopped_time_in_millis":
			if err := dec.Decode(&s.TotalStoppedTimeInMillis); err != nil {
				return err
			}

		case "total_throttled_time":
			if err := dec.Decode(&s.TotalThrottledTime); err != nil {
				return err
			}

		case "total_throttled_time_in_millis":
			if err := dec.Decode(&s.TotalThrottledTimeInMillis); err != nil {
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

// NewMergesStats returns a MergesStats.
func NewMergesStats() *MergesStats {
	r := &MergesStats{}

	return r
}
