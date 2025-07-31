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

// ReadSummaryInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/snapshot/repository_analyze/SnapshotAnalyzeRepositoryResponse.ts#L115-L160
type ReadSummaryInfo struct {
	// Count The number of read operations performed in the test.
	Count int `json:"count"`
	// MaxWait The maximum time spent waiting for the first byte of any read request to be
	// received.
	MaxWait Duration `json:"max_wait"`
	// MaxWaitNanos The maximum time spent waiting for the first byte of any read request to be
	// received, in nanoseconds.
	MaxWaitNanos int64 `json:"max_wait_nanos"`
	// TotalElapsed The total elapsed time spent on reading blobs in the test.
	TotalElapsed Duration `json:"total_elapsed"`
	// TotalElapsedNanos The total elapsed time spent on reading blobs in the test, in nanoseconds.
	TotalElapsedNanos int64 `json:"total_elapsed_nanos"`
	// TotalSize The total size of all the blobs or partial blobs read in the test.
	TotalSize ByteSize `json:"total_size"`
	// TotalSizeBytes The total size of all the blobs or partial blobs read in the test, in bytes.
	TotalSizeBytes int64 `json:"total_size_bytes"`
	// TotalThrottled The total time spent waiting due to the `max_restore_bytes_per_sec` or
	// `indices.recovery.max_bytes_per_sec` throttles.
	TotalThrottled Duration `json:"total_throttled"`
	// TotalThrottledNanos The total time spent waiting due to the `max_restore_bytes_per_sec` or
	// `indices.recovery.max_bytes_per_sec` throttles, in nanoseconds.
	TotalThrottledNanos int64 `json:"total_throttled_nanos"`
	// TotalWait The total time spent waiting for the first byte of each read request to be
	// received.
	TotalWait Duration `json:"total_wait"`
	// TotalWaitNanos The total time spent waiting for the first byte of each read request to be
	// received, in nanoseconds.
	TotalWaitNanos int64 `json:"total_wait_nanos"`
}

func (s *ReadSummaryInfo) UnmarshalJSON(data []byte) error {

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

		case "max_wait":
			if err := dec.Decode(&s.MaxWait); err != nil {
				return fmt.Errorf("%s | %w", "MaxWait", err)
			}

		case "max_wait_nanos":
			if err := dec.Decode(&s.MaxWaitNanos); err != nil {
				return fmt.Errorf("%s | %w", "MaxWaitNanos", err)
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
			if err := dec.Decode(&s.TotalThrottledNanos); err != nil {
				return fmt.Errorf("%s | %w", "TotalThrottledNanos", err)
			}

		case "total_wait":
			if err := dec.Decode(&s.TotalWait); err != nil {
				return fmt.Errorf("%s | %w", "TotalWait", err)
			}

		case "total_wait_nanos":
			if err := dec.Decode(&s.TotalWaitNanos); err != nil {
				return fmt.Errorf("%s | %w", "TotalWaitNanos", err)
			}

		}
	}
	return nil
}

// NewReadSummaryInfo returns a ReadSummaryInfo.
func NewReadSummaryInfo() *ReadSummaryInfo {
	r := &ReadSummaryInfo{}

	return r
}
