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

// CgroupCpuStat type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/nodes/_types/Stats.ts#L206-L210
type CgroupCpuStat struct {
	NumberOfElapsedPeriods *int64 `json:"number_of_elapsed_periods,omitempty"`
	NumberOfTimesThrottled *int64 `json:"number_of_times_throttled,omitempty"`
	TimeThrottledNanos     *int64 `json:"time_throttled_nanos,omitempty"`
}

func (s *CgroupCpuStat) UnmarshalJSON(data []byte) error {

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

		case "number_of_elapsed_periods":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.NumberOfElapsedPeriods = &value
			case float64:
				f := int64(v)
				s.NumberOfElapsedPeriods = &f
			}

		case "number_of_times_throttled":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.NumberOfTimesThrottled = &value
			case float64:
				f := int64(v)
				s.NumberOfTimesThrottled = &f
			}

		case "time_throttled_nanos":
			if err := dec.Decode(&s.TimeThrottledNanos); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewCgroupCpuStat returns a CgroupCpuStat.
func NewCgroupCpuStat() *CgroupCpuStat {
	r := &CgroupCpuStat{}

	return r
}
