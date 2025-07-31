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

// WriteSummaryInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/snapshot/repository_analyze/SnapshotAnalyzeRepositoryResponse.ts#L162-L191
type WriteSummaryInfo struct {
	// Count The number of write operations performed in the test.
	Count int `json:"count"`
	// TotalElapsed The total elapsed time spent on writing blobs in the test.
	TotalElapsed Duration `json:"total_elapsed"`
	// TotalElapsedNanos The total elapsed time spent on writing blobs in the test, in nanoseconds.
	TotalElapsedNanos int64 `json:"total_elapsed_nanos"`
	// TotalSize The total size of all the blobs written in the test.
	TotalSize ByteSize `json:"total_size"`
	// TotalSizeBytes The total size of all the blobs written in the test, in bytes.
	TotalSizeBytes int64 `json:"total_size_bytes"`
	// TotalThrottled The total time spent waiting due to the `max_snapshot_bytes_per_sec`
	// throttle.
	TotalThrottled Duration `json:"total_throttled"`
	// TotalThrottledNanos The total time spent waiting due to the `max_snapshot_bytes_per_sec`
	// throttle, in nanoseconds.
	TotalThrottledNanos int64 `json:"total_throttled_nanos"`
}

func (s *WriteSummaryInfo) UnmarshalJSON(data []byte) error {

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

		case "count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Count", err)
				}
				s.Count = value
			case float64:
				f := int(v)
				s.Count = f
			}

		case "total_elapsed":
			if err := dec.Decode(&s.TotalElapsed); err != nil {
				return fmt.Errorf("%s | %w", "TotalElapsed", err)
			}

		case "total_elapsed_nanos":
			if err := dec.Decode(&s.TotalElapsedNanos); err != nil {
				return fmt.Errorf("%s | %w", "TotalElapsedNanos", err)
			}

		case "total_size":
			if err := dec.Decode(&s.TotalSize); err != nil {
				return fmt.Errorf("%s | %w", "TotalSize", err)
			}

		case "total_size_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "TotalSizeBytes", err)
				}
				s.TotalSizeBytes = value
			case float64:
				f := int64(v)
				s.TotalSizeBytes = f
			}

		case "total_throttled":
			if err := dec.Decode(&s.TotalThrottled); err != nil {
				return fmt.Errorf("%s | %w", "TotalThrottled", err)
			}

		case "total_throttled_nanos":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "TotalThrottledNanos", err)
				}
				s.TotalThrottledNanos = value
			case float64:
				f := int64(v)
				s.TotalThrottledNanos = f
			}

		}
	}
	return nil
}

// NewWriteSummaryInfo returns a WriteSummaryInfo.
func NewWriteSummaryInfo() *WriteSummaryInfo {
	r := &WriteSummaryInfo{}

	return r
}
