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

// SnapshotPolicyStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/slm/_types/SnapshotLifecycle.ts#L153-L159
type SnapshotPolicyStats struct {
	Policy                   string `json:"policy"`
	SnapshotDeletionFailures int64  `json:"snapshot_deletion_failures"`
	SnapshotsDeleted         int64  `json:"snapshots_deleted"`
	SnapshotsFailed          int64  `json:"snapshots_failed"`
	SnapshotsTaken           int64  `json:"snapshots_taken"`
}

func (s *SnapshotPolicyStats) UnmarshalJSON(data []byte) error {

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
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Policy", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Policy = o

		case "snapshot_deletion_failures":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "SnapshotDeletionFailures", err)
				}
				s.SnapshotDeletionFailures = value
			case float64:
				f := int64(v)
				s.SnapshotDeletionFailures = f
			}

		case "snapshots_deleted":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "SnapshotsDeleted", err)
				}
				s.SnapshotsDeleted = value
			case float64:
				f := int64(v)
				s.SnapshotsDeleted = f
			}

		case "snapshots_failed":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "SnapshotsFailed", err)
				}
				s.SnapshotsFailed = value
			case float64:
				f := int64(v)
				s.SnapshotsFailed = f
			}

		case "snapshots_taken":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "SnapshotsTaken", err)
				}
				s.SnapshotsTaken = value
			case float64:
				f := int64(v)
				s.SnapshotsTaken = f
			}

		}
	}
	return nil
}

// NewSnapshotPolicyStats returns a SnapshotPolicyStats.
func NewSnapshotPolicyStats() *SnapshotPolicyStats {
	r := &SnapshotPolicyStats{}

	return r
}
