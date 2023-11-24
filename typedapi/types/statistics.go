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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// Statistics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/slm/_types/SnapshotLifecycle.ts#L51-L74
type Statistics struct {
	Policy                        *string  `json:"policy,omitempty"`
	RetentionDeletionTime         Duration `json:"retention_deletion_time,omitempty"`
	RetentionDeletionTimeMillis   *int64   `json:"retention_deletion_time_millis,omitempty"`
	RetentionFailed               *int64   `json:"retention_failed,omitempty"`
	RetentionRuns                 *int64   `json:"retention_runs,omitempty"`
	RetentionTimedOut             *int64   `json:"retention_timed_out,omitempty"`
	TotalSnapshotDeletionFailures *int64   `json:"total_snapshot_deletion_failures,omitempty"`
	TotalSnapshotsDeleted         *int64   `json:"total_snapshots_deleted,omitempty"`
	TotalSnapshotsFailed          *int64   `json:"total_snapshots_failed,omitempty"`
	TotalSnapshotsTaken           *int64   `json:"total_snapshots_taken,omitempty"`
}

func (s *Statistics) UnmarshalJSON(data []byte) error {

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

		case "policy":
			if err := dec.Decode(&s.Policy); err != nil {
				return err
			}

		case "retention_deletion_time":
			if err := dec.Decode(&s.RetentionDeletionTime); err != nil {
				return err
			}

		case "retention_deletion_time_millis":
			if err := dec.Decode(&s.RetentionDeletionTimeMillis); err != nil {
				return err
			}

		case "retention_failed":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.RetentionFailed = &value
			case float64:
				f := int64(v)
				s.RetentionFailed = &f
			}

		case "retention_runs":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.RetentionRuns = &value
			case float64:
				f := int64(v)
				s.RetentionRuns = &f
			}

		case "retention_timed_out":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.RetentionTimedOut = &value
			case float64:
				f := int64(v)
				s.RetentionTimedOut = &f
			}

		case "total_snapshot_deletion_failures", "snapshot_deletion_failures":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.TotalSnapshotDeletionFailures = &value
			case float64:
				f := int64(v)
				s.TotalSnapshotDeletionFailures = &f
			}

		case "total_snapshots_deleted", "snapshots_deleted":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.TotalSnapshotsDeleted = &value
			case float64:
				f := int64(v)
				s.TotalSnapshotsDeleted = &f
			}

		case "total_snapshots_failed", "snapshots_failed":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.TotalSnapshotsFailed = &value
			case float64:
				f := int64(v)
				s.TotalSnapshotsFailed = &f
			}

		case "total_snapshots_taken", "snapshots_taken":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.TotalSnapshotsTaken = &value
			case float64:
				f := int64(v)
				s.TotalSnapshotsTaken = &f
			}

		}
	}
	return nil
}

// NewStatistics returns a Statistics.
func NewStatistics() *Statistics {
	r := &Statistics{}

	return r
}
