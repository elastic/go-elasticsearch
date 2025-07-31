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
)

// DetailsInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/snapshot/repository_analyze/SnapshotAnalyzeRepositoryResponse.ts#L286-L321
type DetailsInfo struct {
	// Blob A description of the blob that was written and read.
	Blob BlobDetails `json:"blob"`
	// OverwriteElapsed The elapsed time spent overwriting the blob.
	// If the blob was not overwritten, this information is omitted.
	OverwriteElapsed Duration `json:"overwrite_elapsed,omitempty"`
	// OverwriteElapsedNanos The elapsed time spent overwriting the blob, in nanoseconds.
	// If the blob was not overwritten, this information is omitted.
	OverwriteElapsedNanos *int64 `json:"overwrite_elapsed_nanos,omitempty"`
	// WriteElapsed The elapsed time spent writing the blob.
	WriteElapsed Duration `json:"write_elapsed"`
	// WriteElapsedNanos The elapsed time spent writing the blob, in nanoseconds.
	WriteElapsedNanos int64 `json:"write_elapsed_nanos"`
	// WriteThrottled The length of time spent waiting for the `max_snapshot_bytes_per_sec` (or
	// `indices.recovery.max_bytes_per_sec` if the recovery settings for managed
	// services are set) throttle while writing the blob.
	WriteThrottled Duration `json:"write_throttled"`
	// WriteThrottledNanos The length of time spent waiting for the `max_snapshot_bytes_per_sec` (or
	// `indices.recovery.max_bytes_per_sec` if the recovery settings for managed
	// services are set) throttle while writing the blob, in nanoseconds.
	WriteThrottledNanos int64 `json:"write_throttled_nanos"`
	// WriterNode The node which wrote the blob and coordinated the read operations.
	WriterNode SnapshotNodeInfo `json:"writer_node"`
}

func (s *DetailsInfo) UnmarshalJSON(data []byte) error {

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

		case "blob":
			if err := dec.Decode(&s.Blob); err != nil {
				return fmt.Errorf("%s | %w", "Blob", err)
			}

		case "overwrite_elapsed":
			if err := dec.Decode(&s.OverwriteElapsed); err != nil {
				return fmt.Errorf("%s | %w", "OverwriteElapsed", err)
			}

		case "overwrite_elapsed_nanos":
			if err := dec.Decode(&s.OverwriteElapsedNanos); err != nil {
				return fmt.Errorf("%s | %w", "OverwriteElapsedNanos", err)
			}

		case "write_elapsed":
			if err := dec.Decode(&s.WriteElapsed); err != nil {
				return fmt.Errorf("%s | %w", "WriteElapsed", err)
			}

		case "write_elapsed_nanos":
			if err := dec.Decode(&s.WriteElapsedNanos); err != nil {
				return fmt.Errorf("%s | %w", "WriteElapsedNanos", err)
			}

		case "write_throttled":
			if err := dec.Decode(&s.WriteThrottled); err != nil {
				return fmt.Errorf("%s | %w", "WriteThrottled", err)
			}

		case "write_throttled_nanos":
			if err := dec.Decode(&s.WriteThrottledNanos); err != nil {
				return fmt.Errorf("%s | %w", "WriteThrottledNanos", err)
			}

		case "writer_node":
			if err := dec.Decode(&s.WriterNode); err != nil {
				return fmt.Errorf("%s | %w", "WriterNode", err)
			}

		}
	}
	return nil
}

// NewDetailsInfo returns a DetailsInfo.
func NewDetailsInfo() *DetailsInfo {
	r := &DetailsInfo{}

	return r
}
