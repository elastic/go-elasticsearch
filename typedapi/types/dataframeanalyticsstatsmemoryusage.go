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

// DataframeAnalyticsStatsMemoryUsage type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/ml/_types/DataframeAnalytics.ts#L350-L359
type DataframeAnalyticsStatsMemoryUsage struct {
	// MemoryReestimateBytes This value is present when the status is hard_limit and it is a new estimate
	// of how much memory the job needs.
	MemoryReestimateBytes *int64 `json:"memory_reestimate_bytes,omitempty"`
	// PeakUsageBytes The number of bytes used at the highest peak of memory usage.
	PeakUsageBytes int64 `json:"peak_usage_bytes"`
	// Status The memory usage status.
	Status string `json:"status"`
	// Timestamp The timestamp when memory usage was calculated.
	Timestamp *int64 `json:"timestamp,omitempty"`
}

func (s *DataframeAnalyticsStatsMemoryUsage) UnmarshalJSON(data []byte) error {

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

		case "memory_reestimate_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.MemoryReestimateBytes = &value
			case float64:
				f := int64(v)
				s.MemoryReestimateBytes = &f
			}

		case "peak_usage_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.PeakUsageBytes = value
			case float64:
				f := int64(v)
				s.PeakUsageBytes = f
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

		case "timestamp":
			if err := dec.Decode(&s.Timestamp); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewDataframeAnalyticsStatsMemoryUsage returns a DataframeAnalyticsStatsMemoryUsage.
func NewDataframeAnalyticsStatsMemoryUsage() *DataframeAnalyticsStatsMemoryUsage {
	r := &DataframeAnalyticsStatsMemoryUsage{}

	return r
}
