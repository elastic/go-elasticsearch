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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// IndexingStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/Stats.ts#L168-L186
type IndexingStats struct {
	DeleteCurrent        int64                    `json:"delete_current"`
	DeleteTime           Duration                 `json:"delete_time,omitempty"`
	DeleteTimeInMillis   int64                    `json:"delete_time_in_millis"`
	DeleteTotal          int64                    `json:"delete_total"`
	IndexCurrent         int64                    `json:"index_current"`
	IndexFailed          int64                    `json:"index_failed"`
	IndexTime            Duration                 `json:"index_time,omitempty"`
	IndexTimeInMillis    int64                    `json:"index_time_in_millis"`
	IndexTotal           int64                    `json:"index_total"`
	IsThrottled          bool                     `json:"is_throttled"`
	NoopUpdateTotal      int64                    `json:"noop_update_total"`
	PeakWriteLoad        *Float64                 `json:"peak_write_load,omitempty"`
	RecentWriteLoad      *Float64                 `json:"recent_write_load,omitempty"`
	ThrottleTime         Duration                 `json:"throttle_time,omitempty"`
	ThrottleTimeInMillis int64                    `json:"throttle_time_in_millis"`
	Types                map[string]IndexingStats `json:"types,omitempty"`
	WriteLoad            *Float64                 `json:"write_load,omitempty"`
}

func (s *IndexingStats) UnmarshalJSON(data []byte) error {

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

		case "delete_current":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "DeleteCurrent", err)
				}
				s.DeleteCurrent = value
			case float64:
				f := int64(v)
				s.DeleteCurrent = f
			}

		case "delete_time":
			if err := dec.Decode(&s.DeleteTime); err != nil {
				return fmt.Errorf("%s | %w", "DeleteTime", err)
			}

		case "delete_time_in_millis":
			if err := dec.Decode(&s.DeleteTimeInMillis); err != nil {
				return fmt.Errorf("%s | %w", "DeleteTimeInMillis", err)
			}

		case "delete_total":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "DeleteTotal", err)
				}
				s.DeleteTotal = value
			case float64:
				f := int64(v)
				s.DeleteTotal = f
			}

		case "index_current":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "IndexCurrent", err)
				}
				s.IndexCurrent = value
			case float64:
				f := int64(v)
				s.IndexCurrent = f
			}

		case "index_failed":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "IndexFailed", err)
				}
				s.IndexFailed = value
			case float64:
				f := int64(v)
				s.IndexFailed = f
			}

		case "index_time":
			if err := dec.Decode(&s.IndexTime); err != nil {
				return fmt.Errorf("%s | %w", "IndexTime", err)
			}

		case "index_time_in_millis":
			if err := dec.Decode(&s.IndexTimeInMillis); err != nil {
				return fmt.Errorf("%s | %w", "IndexTimeInMillis", err)
			}

		case "index_total":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "IndexTotal", err)
				}
				s.IndexTotal = value
			case float64:
				f := int64(v)
				s.IndexTotal = f
			}

		case "is_throttled":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IsThrottled", err)
				}
				s.IsThrottled = value
			case bool:
				s.IsThrottled = v
			}

		case "noop_update_total":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "NoopUpdateTotal", err)
				}
				s.NoopUpdateTotal = value
			case float64:
				f := int64(v)
				s.NoopUpdateTotal = f
			}

		case "peak_write_load":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "PeakWriteLoad", err)
				}
				f := Float64(value)
				s.PeakWriteLoad = &f
			case float64:
				f := Float64(v)
				s.PeakWriteLoad = &f
			}

		case "recent_write_load":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "RecentWriteLoad", err)
				}
				f := Float64(value)
				s.RecentWriteLoad = &f
			case float64:
				f := Float64(v)
				s.RecentWriteLoad = &f
			}

		case "throttle_time":
			if err := dec.Decode(&s.ThrottleTime); err != nil {
				return fmt.Errorf("%s | %w", "ThrottleTime", err)
			}

		case "throttle_time_in_millis":
			if err := dec.Decode(&s.ThrottleTimeInMillis); err != nil {
				return fmt.Errorf("%s | %w", "ThrottleTimeInMillis", err)
			}

		case "types":
			if s.Types == nil {
				s.Types = make(map[string]IndexingStats, 0)
			}
			if err := dec.Decode(&s.Types); err != nil {
				return fmt.Errorf("%s | %w", "Types", err)
			}

		case "write_load":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "WriteLoad", err)
				}
				f := Float64(value)
				s.WriteLoad = &f
			case float64:
				f := Float64(v)
				s.WriteLoad = &f
			}

		}
	}
	return nil
}

// NewIndexingStats returns a IndexingStats.
func NewIndexingStats() *IndexingStats {
	r := &IndexingStats{
		Types: make(map[string]IndexingStats),
	}

	return r
}
