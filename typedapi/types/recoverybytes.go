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
)

// RecoveryBytes type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/indices/recovery/types.ts#L38-L48
type RecoveryBytes struct {
	Percent                      Percentage `json:"percent"`
	Recovered                    ByteSize   `json:"recovered,omitempty"`
	RecoveredFromSnapshot        ByteSize   `json:"recovered_from_snapshot,omitempty"`
	RecoveredFromSnapshotInBytes ByteSize   `json:"recovered_from_snapshot_in_bytes,omitempty"`
	RecoveredInBytes             ByteSize   `json:"recovered_in_bytes"`
	Reused                       ByteSize   `json:"reused,omitempty"`
	ReusedInBytes                ByteSize   `json:"reused_in_bytes"`
	Total                        ByteSize   `json:"total,omitempty"`
	TotalInBytes                 ByteSize   `json:"total_in_bytes"`
}

func (s *RecoveryBytes) UnmarshalJSON(data []byte) error {

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

		case "percent":
			if err := dec.Decode(&s.Percent); err != nil {
				return fmt.Errorf("%s | %w", "Percent", err)
			}

		case "recovered":
			if err := dec.Decode(&s.Recovered); err != nil {
				return fmt.Errorf("%s | %w", "Recovered", err)
			}

		case "recovered_from_snapshot":
			if err := dec.Decode(&s.RecoveredFromSnapshot); err != nil {
				return fmt.Errorf("%s | %w", "RecoveredFromSnapshot", err)
			}

		case "recovered_from_snapshot_in_bytes":
			if err := dec.Decode(&s.RecoveredFromSnapshotInBytes); err != nil {
				return fmt.Errorf("%s | %w", "RecoveredFromSnapshotInBytes", err)
			}

		case "recovered_in_bytes":
			if err := dec.Decode(&s.RecoveredInBytes); err != nil {
				return fmt.Errorf("%s | %w", "RecoveredInBytes", err)
			}

		case "reused":
			if err := dec.Decode(&s.Reused); err != nil {
				return fmt.Errorf("%s | %w", "Reused", err)
			}

		case "reused_in_bytes":
			if err := dec.Decode(&s.ReusedInBytes); err != nil {
				return fmt.Errorf("%s | %w", "ReusedInBytes", err)
			}

		case "total":
			if err := dec.Decode(&s.Total); err != nil {
				return fmt.Errorf("%s | %w", "Total", err)
			}

		case "total_in_bytes":
			if err := dec.Decode(&s.TotalInBytes); err != nil {
				return fmt.Errorf("%s | %w", "TotalInBytes", err)
			}

		}
	}
	return nil
}

// NewRecoveryBytes returns a RecoveryBytes.
func NewRecoveryBytes() *RecoveryBytes {
	r := &RecoveryBytes{}

	return r
}
