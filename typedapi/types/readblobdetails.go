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

// ReadBlobDetails type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/snapshot/repository_analyze/SnapshotAnalyzeRepositoryResponse.ts#L204-L248
type ReadBlobDetails struct {
	// BeforeWriteComplete Indicates whether the read operation may have started before the write
	// operation was complete.
	BeforeWriteComplete *bool `json:"before_write_complete,omitempty"`
	// Elapsed The length of time spent reading the blob.
	// If the blob was not found, this detail is omitted.
	Elapsed Duration `json:"elapsed,omitempty"`
	// ElapsedNanos The length of time spent reading the blob, in nanoseconds.
	// If the blob was not found, this detail is omitted.
	ElapsedNanos *int64 `json:"elapsed_nanos,omitempty"`
	// FirstByteTime The length of time waiting for the first byte of the read operation to be
	// received.
	// If the blob was not found, this detail is omitted.
	FirstByteTime Duration `json:"first_byte_time,omitempty"`
	// FirstByteTimeNanos The length of time waiting for the first byte of the read operation to be
	// received, in nanoseconds.
	// If the blob was not found, this detail is omitted.
	FirstByteTimeNanos int64 `json:"first_byte_time_nanos"`
	// Found Indicates whether the blob was found by the read operation.
	// If the read was started before the write completed or the write was ended
	// before completion, it might be false.
	Found bool `json:"found"`
	// Node The node that performed the read operation.
	Node SnapshotNodeInfo `json:"node"`
	// Throttled The length of time spent waiting due to the `max_restore_bytes_per_sec` or
	// `indices.recovery.max_bytes_per_sec` throttles during the read of the blob.
	// If the blob was not found, this detail is omitted.
	Throttled Duration `json:"throttled,omitempty"`
	// ThrottledNanos The length of time spent waiting due to the `max_restore_bytes_per_sec` or
	// `indices.recovery.max_bytes_per_sec` throttles during the read of the blob,
	// in nanoseconds.
	// If the blob was not found, this detail is omitted.
	ThrottledNanos *int64 `json:"throttled_nanos,omitempty"`
}

func (s *ReadBlobDetails) UnmarshalJSON(data []byte) error {

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

		case "before_write_complete":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "BeforeWriteComplete", err)
				}
				s.BeforeWriteComplete = &value
			case bool:
				s.BeforeWriteComplete = &v
			}

		case "elapsed":
			if err := dec.Decode(&s.Elapsed); err != nil {
				return fmt.Errorf("%s | %w", "Elapsed", err)
			}

		case "elapsed_nanos":
			if err := dec.Decode(&s.ElapsedNanos); err != nil {
				return fmt.Errorf("%s | %w", "ElapsedNanos", err)
			}

		case "first_byte_time":
			if err := dec.Decode(&s.FirstByteTime); err != nil {
				return fmt.Errorf("%s | %w", "FirstByteTime", err)
			}

		case "first_byte_time_nanos":
			if err := dec.Decode(&s.FirstByteTimeNanos); err != nil {
				return fmt.Errorf("%s | %w", "FirstByteTimeNanos", err)
			}

		case "found":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Found", err)
				}
				s.Found = value
			case bool:
				s.Found = v
			}

		case "node":
			if err := dec.Decode(&s.Node); err != nil {
				return fmt.Errorf("%s | %w", "Node", err)
			}

		case "throttled":
			if err := dec.Decode(&s.Throttled); err != nil {
				return fmt.Errorf("%s | %w", "Throttled", err)
			}

		case "throttled_nanos":
			if err := dec.Decode(&s.ThrottledNanos); err != nil {
				return fmt.Errorf("%s | %w", "ThrottledNanos", err)
			}

		}
	}
	return nil
}

// NewReadBlobDetails returns a ReadBlobDetails.
func NewReadBlobDetails() *ReadBlobDetails {
	r := &ReadBlobDetails{}

	return r
}
