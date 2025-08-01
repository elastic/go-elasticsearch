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

// SnapshotShardsStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/snapshot/_types/SnapshotShardsStats.ts#L22-L47
type SnapshotShardsStats struct {
	// Done The number of shards that initialized, started, and finalized successfully.
	Done int64 `json:"done"`
	// Failed The number of shards that failed to be included in the snapshot.
	Failed int64 `json:"failed"`
	// Finalizing The number of shards that are finalizing but are not done.
	Finalizing int64 `json:"finalizing"`
	// Initializing The number of shards that are still initializing.
	Initializing int64 `json:"initializing"`
	// Started The number of shards that have started but are not finalized.
	Started int64 `json:"started"`
	// Total The total number of shards included in the snapshot.
	Total int64 `json:"total"`
}

func (s *SnapshotShardsStats) UnmarshalJSON(data []byte) error {

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

		case "done":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Done", err)
				}
				s.Done = value
			case float64:
				f := int64(v)
				s.Done = f
			}

		case "failed":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Failed", err)
				}
				s.Failed = value
			case float64:
				f := int64(v)
				s.Failed = f
			}

		case "finalizing":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Finalizing", err)
				}
				s.Finalizing = value
			case float64:
				f := int64(v)
				s.Finalizing = f
			}

		case "initializing":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Initializing", err)
				}
				s.Initializing = value
			case float64:
				f := int64(v)
				s.Initializing = f
			}

		case "started":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Started", err)
				}
				s.Started = value
			case float64:
				f := int64(v)
				s.Started = f
			}

		case "total":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Total", err)
				}
				s.Total = value
			case float64:
				f := int64(v)
				s.Total = f
			}

		}
	}
	return nil
}

// NewSnapshotShardsStats returns a SnapshotShardsStats.
func NewSnapshotShardsStats() *SnapshotShardsStats {
	r := &SnapshotShardsStats{}

	return r
}
