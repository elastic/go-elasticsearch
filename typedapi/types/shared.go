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

// Shared type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/searchable_snapshots/cache_stats/Response.ts#L34-L43
type Shared struct {
	BytesReadInBytes    ByteSize `json:"bytes_read_in_bytes"`
	BytesWrittenInBytes ByteSize `json:"bytes_written_in_bytes"`
	Evictions           int64    `json:"evictions"`
	NumRegions          int      `json:"num_regions"`
	Reads               int64    `json:"reads"`
	RegionSizeInBytes   ByteSize `json:"region_size_in_bytes"`
	SizeInBytes         ByteSize `json:"size_in_bytes"`
	Writes              int64    `json:"writes"`
}

func (s *Shared) UnmarshalJSON(data []byte) error {

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

		case "bytes_read_in_bytes":
			if err := dec.Decode(&s.BytesReadInBytes); err != nil {
				return fmt.Errorf("%s | %w", "BytesReadInBytes", err)
			}

		case "bytes_written_in_bytes":
			if err := dec.Decode(&s.BytesWrittenInBytes); err != nil {
				return fmt.Errorf("%s | %w", "BytesWrittenInBytes", err)
			}

		case "evictions":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Evictions", err)
				}
				s.Evictions = value
			case float64:
				f := int64(v)
				s.Evictions = f
			}

		case "num_regions":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumRegions", err)
				}
				s.NumRegions = value
			case float64:
				f := int(v)
				s.NumRegions = f
			}

		case "reads":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Reads", err)
				}
				s.Reads = value
			case float64:
				f := int64(v)
				s.Reads = f
			}

		case "region_size_in_bytes":
			if err := dec.Decode(&s.RegionSizeInBytes); err != nil {
				return fmt.Errorf("%s | %w", "RegionSizeInBytes", err)
			}

		case "size_in_bytes":
			if err := dec.Decode(&s.SizeInBytes); err != nil {
				return fmt.Errorf("%s | %w", "SizeInBytes", err)
			}

		case "writes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Writes", err)
				}
				s.Writes = value
			case float64:
				f := int64(v)
				s.Writes = f
			}

		}
	}
	return nil
}

// NewShared returns a Shared.
func NewShared() *Shared {
	r := &Shared{}

	return r
}
