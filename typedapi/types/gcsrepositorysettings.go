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
// https://github.com/elastic/elasticsearch-specification/tree/cdb84fa39f1401846dab6e1c76781fb3090527ed

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// GcsRepositorySettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/cdb84fa39f1401846dab6e1c76781fb3090527ed/specification/snapshot/_types/SnapshotRepository.ts#L85-L91
type GcsRepositorySettings struct {
	ApplicationName        *string  `json:"application_name,omitempty"`
	BasePath               *string  `json:"base_path,omitempty"`
	Bucket                 string   `json:"bucket"`
	ChunkSize              ByteSize `json:"chunk_size,omitempty"`
	Client                 *string  `json:"client,omitempty"`
	Compress               *bool    `json:"compress,omitempty"`
	MaxRestoreBytesPerSec  ByteSize `json:"max_restore_bytes_per_sec,omitempty"`
	MaxSnapshotBytesPerSec ByteSize `json:"max_snapshot_bytes_per_sec,omitempty"`
	Readonly               *bool    `json:"readonly,omitempty"`
}

func (s *GcsRepositorySettings) UnmarshalJSON(data []byte) error {

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

		case "application_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ApplicationName", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ApplicationName = &o

		case "base_path":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "BasePath", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.BasePath = &o

		case "bucket":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Bucket", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Bucket = o

		case "chunk_size":
			if err := dec.Decode(&s.ChunkSize); err != nil {
				return fmt.Errorf("%s | %w", "ChunkSize", err)
			}

		case "client":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Client", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Client = &o

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

		case "max_restore_bytes_per_sec":
			if err := dec.Decode(&s.MaxRestoreBytesPerSec); err != nil {
				return fmt.Errorf("%s | %w", "MaxRestoreBytesPerSec", err)
			}

		case "max_snapshot_bytes_per_sec":
			if err := dec.Decode(&s.MaxSnapshotBytesPerSec); err != nil {
				return fmt.Errorf("%s | %w", "MaxSnapshotBytesPerSec", err)
			}

		case "readonly":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Readonly", err)
				}
				s.Readonly = &value
			case bool:
				s.Readonly = &v
			}

		}
	}
	return nil
}

// NewGcsRepositorySettings returns a GcsRepositorySettings.
func NewGcsRepositorySettings() *GcsRepositorySettings {
	r := &GcsRepositorySettings{}

	return r
}
