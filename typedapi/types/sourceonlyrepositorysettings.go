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

// SourceOnlyRepositorySettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/snapshot/_types/SnapshotRepository.ts#L117-L124
type SourceOnlyRepositorySettings struct {
	ChunkSize              ByteSize `json:"chunk_size,omitempty"`
	Compress               *bool    `json:"compress,omitempty"`
	DelegateType           *string  `json:"delegate_type,omitempty"`
	MaxNumberOfSnapshots   *int     `json:"max_number_of_snapshots,omitempty"`
	MaxRestoreBytesPerSec  ByteSize `json:"max_restore_bytes_per_sec,omitempty"`
	MaxSnapshotBytesPerSec ByteSize `json:"max_snapshot_bytes_per_sec,omitempty"`
	ReadOnly               *bool    `json:"read_only,omitempty"`
}

func (s *SourceOnlyRepositorySettings) UnmarshalJSON(data []byte) error {

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

		case "chunk_size":
			if err := dec.Decode(&s.ChunkSize); err != nil {
				return fmt.Errorf("%s | %w", "ChunkSize", err)
			}

		case "compress":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Compress", err)
				}
				s.Compress = &value
			case bool:
				s.Compress = &v
			}

		case "delegate_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "DelegateType", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DelegateType = &o

		case "max_number_of_snapshots":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxNumberOfSnapshots", err)
				}
				s.MaxNumberOfSnapshots = &value
			case float64:
				f := int(v)
				s.MaxNumberOfSnapshots = &f
			}

		case "max_restore_bytes_per_sec":
			if err := dec.Decode(&s.MaxRestoreBytesPerSec); err != nil {
				return fmt.Errorf("%s | %w", "MaxRestoreBytesPerSec", err)
			}

		case "max_snapshot_bytes_per_sec":
			if err := dec.Decode(&s.MaxSnapshotBytesPerSec); err != nil {
				return fmt.Errorf("%s | %w", "MaxSnapshotBytesPerSec", err)
			}

		case "read_only", "readonly":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ReadOnly", err)
				}
				s.ReadOnly = &value
			case bool:
				s.ReadOnly = &v
			}

		}
	}
	return nil
}

// NewSourceOnlyRepositorySettings returns a SourceOnlyRepositorySettings.
func NewSourceOnlyRepositorySettings() *SourceOnlyRepositorySettings {
	r := &SourceOnlyRepositorySettings{}

	return r
}
