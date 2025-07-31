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

// SnapshotCurrentCounts type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/cluster/stats/types.ts#L642-L663
type SnapshotCurrentCounts struct {
	// Cleanups Cleanups in progress, not counted in concurrent_operations as they are not
	// concurrent
	Cleanups int `json:"cleanups"`
	// ConcurrentOperations Sum of snapshots and snapshot_deletions
	ConcurrentOperations int `json:"concurrent_operations"`
	// ShardSnapshots Incomplete shard snapshots
	ShardSnapshots int `json:"shard_snapshots"`
	// SnapshotDeletions Snapshots deletions in progress
	SnapshotDeletions int `json:"snapshot_deletions"`
	// Snapshots Snapshots currently in progress
	Snapshots int `json:"snapshots"`
}

func (s *SnapshotCurrentCounts) UnmarshalJSON(data []byte) error {

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

		case "cleanups":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Cleanups", err)
				}
				s.Cleanups = value
			case float64:
				f := int(v)
				s.Cleanups = f
			}

		case "concurrent_operations":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ConcurrentOperations", err)
				}
				s.ConcurrentOperations = value
			case float64:
				f := int(v)
				s.ConcurrentOperations = f
			}

		case "shard_snapshots":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ShardSnapshots", err)
				}
				s.ShardSnapshots = value
			case float64:
				f := int(v)
				s.ShardSnapshots = f
			}

		case "snapshot_deletions":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "SnapshotDeletions", err)
				}
				s.SnapshotDeletions = value
			case float64:
				f := int(v)
				s.SnapshotDeletions = f
			}

		case "snapshots":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Snapshots", err)
				}
				s.Snapshots = value
			case float64:
				f := int(v)
				s.Snapshots = f
			}

		}
	}
	return nil
}

// NewSnapshotCurrentCounts returns a SnapshotCurrentCounts.
func NewSnapshotCurrentCounts() *SnapshotCurrentCounts {
	r := &SnapshotCurrentCounts{}

	return r
}
