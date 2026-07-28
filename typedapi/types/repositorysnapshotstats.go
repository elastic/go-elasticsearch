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
// https://github.com/elastic/elasticsearch-specification/tree/7560c979602e6941815872bdaec801200bc7ec4e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// RepositorySnapshotStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7560c979602e6941815872bdaec801200bc7ec4e/specification/nodes/_types/Stats.ts#L124-L181
type RepositorySnapshotStats struct {
	// ShardSnapshotsCompleted The number of shard snapshots completed for this repository.
	ShardSnapshotsCompleted int64 `json:"shard_snapshots_completed"`
	// ShardSnapshotsInProgress The number of shard snapshots currently in progress for this repository.
	ShardSnapshotsInProgress int64 `json:"shard_snapshots_in_progress"`
	// ShardSnapshotsStarted The number of shard snapshots started for this repository.
	ShardSnapshotsStarted int64 `json:"shard_snapshots_started"`
	// TotalReadThrottledTime The cumulative time spent throttling read operations for this repository.
	TotalReadThrottledTime Duration `json:"total_read_throttled_time,omitempty"`
	// TotalReadThrottledTimeNanos The cumulative time, in nanoseconds, spent throttling read operations for
	// this repository.
	TotalReadThrottledTimeNanos int64 `json:"total_read_throttled_time_nanos"`
	// TotalReadTime The cumulative time spent reading blobs while uploading to this repository.
	TotalReadTime Duration `json:"total_read_time,omitempty"`
	// TotalReadTimeInMillis The cumulative time, in milliseconds, spent reading blobs while uploading to
	// this repository.
	TotalReadTimeInMillis int64 `json:"total_read_time_in_millis"`
	// TotalUploadTime The cumulative time spent uploading blobs to this repository.
	TotalUploadTime Duration `json:"total_upload_time,omitempty"`
	// TotalUploadTimeInMillis The cumulative time, in milliseconds, spent uploading blobs to this
	// repository.
	TotalUploadTimeInMillis int64 `json:"total_upload_time_in_millis"`
	// TotalWriteThrottledTime The cumulative time spent throttling write operations for this repository.
	TotalWriteThrottledTime Duration `json:"total_write_throttled_time,omitempty"`
	// TotalWriteThrottledTimeNanos The cumulative time, in nanoseconds, spent throttling write operations for
	// this repository.
	TotalWriteThrottledTimeNanos int64 `json:"total_write_throttled_time_nanos"`
	// UploadedBlobs The number of blobs uploaded to this repository.
	UploadedBlobs int64 `json:"uploaded_blobs"`
	// UploadedSize The cumulative size of the blobs uploaded to this repository.
	UploadedSize ByteSize `json:"uploaded_size,omitempty"`
	// UploadedSizeInBytes The cumulative size, in bytes, of the blobs uploaded to this repository.
	UploadedSizeInBytes int64 `json:"uploaded_size_in_bytes"`
}

func (s *RepositorySnapshotStats) UnmarshalJSON(data []byte) error {

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

		case "shard_snapshots_completed":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "ShardSnapshotsCompleted", err)
				}
				s.ShardSnapshotsCompleted = value
			case float64:
				f := int64(v)
				s.ShardSnapshotsCompleted = f
			}

		case "shard_snapshots_in_progress":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "ShardSnapshotsInProgress", err)
				}
				s.ShardSnapshotsInProgress = value
			case float64:
				f := int64(v)
				s.ShardSnapshotsInProgress = f
			}

		case "shard_snapshots_started":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "ShardSnapshotsStarted", err)
				}
				s.ShardSnapshotsStarted = value
			case float64:
				f := int64(v)
				s.ShardSnapshotsStarted = f
			}

		case "total_read_throttled_time":
			if err := dec.Decode(&s.TotalReadThrottledTime); err != nil {
				return fmt.Errorf("%s | %w", "TotalReadThrottledTime", err)
			}

		case "total_read_throttled_time_nanos":
			if err := dec.Decode(&s.TotalReadThrottledTimeNanos); err != nil {
				return fmt.Errorf("%s | %w", "TotalReadThrottledTimeNanos", err)
			}

		case "total_read_time":
			if err := dec.Decode(&s.TotalReadTime); err != nil {
				return fmt.Errorf("%s | %w", "TotalReadTime", err)
			}

		case "total_read_time_in_millis":
			if err := dec.Decode(&s.TotalReadTimeInMillis); err != nil {
				return fmt.Errorf("%s | %w", "TotalReadTimeInMillis", err)
			}

		case "total_upload_time":
			if err := dec.Decode(&s.TotalUploadTime); err != nil {
				return fmt.Errorf("%s | %w", "TotalUploadTime", err)
			}

		case "total_upload_time_in_millis":
			if err := dec.Decode(&s.TotalUploadTimeInMillis); err != nil {
				return fmt.Errorf("%s | %w", "TotalUploadTimeInMillis", err)
			}

		case "total_write_throttled_time":
			if err := dec.Decode(&s.TotalWriteThrottledTime); err != nil {
				return fmt.Errorf("%s | %w", "TotalWriteThrottledTime", err)
			}

		case "total_write_throttled_time_nanos":
			if err := dec.Decode(&s.TotalWriteThrottledTimeNanos); err != nil {
				return fmt.Errorf("%s | %w", "TotalWriteThrottledTimeNanos", err)
			}

		case "uploaded_blobs":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "UploadedBlobs", err)
				}
				s.UploadedBlobs = value
			case float64:
				f := int64(v)
				s.UploadedBlobs = f
			}

		case "uploaded_size":
			if err := dec.Decode(&s.UploadedSize); err != nil {
				return fmt.Errorf("%s | %w", "UploadedSize", err)
			}

		case "uploaded_size_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "UploadedSizeInBytes", err)
				}
				s.UploadedSizeInBytes = value
			case float64:
				f := int64(v)
				s.UploadedSizeInBytes = f
			}

		}
	}
	return nil
}

// NewRepositorySnapshotStats returns a RepositorySnapshotStats.
func NewRepositorySnapshotStats() *RepositorySnapshotStats {
	r := &RepositorySnapshotStats{}

	return r
}
