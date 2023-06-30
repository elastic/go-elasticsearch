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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// StoreStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_types/Stats.ts#L233-L240
type StoreStats struct {
	Reserved                ByteSize `json:"reserved,omitempty"`
	ReservedInBytes         int64    `json:"reserved_in_bytes"`
	Size                    ByteSize `json:"size,omitempty"`
	SizeInBytes             int64    `json:"size_in_bytes"`
	TotalDataSetSize        ByteSize `json:"total_data_set_size,omitempty"`
	TotalDataSetSizeInBytes *int64   `json:"total_data_set_size_in_bytes,omitempty"`
}

func (s *StoreStats) UnmarshalJSON(data []byte) error {

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

		case "reserved":
			if err := dec.Decode(&s.Reserved); err != nil {
				return err
			}

		case "reserved_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.ReservedInBytes = value
			case float64:
				f := int64(v)
				s.ReservedInBytes = f
			}

		case "size":
			if err := dec.Decode(&s.Size); err != nil {
				return err
			}

		case "size_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.SizeInBytes = value
			case float64:
				f := int64(v)
				s.SizeInBytes = f
			}

		case "total_data_set_size":
			if err := dec.Decode(&s.TotalDataSetSize); err != nil {
				return err
			}

		case "total_data_set_size_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.TotalDataSetSizeInBytes = &value
			case float64:
				f := int64(v)
				s.TotalDataSetSizeInBytes = &f
			}

		}
	}
	return nil
}

// NewStoreStats returns a StoreStats.
func NewStoreStats() *StoreStats {
	r := &StoreStats{}

	return r
}
