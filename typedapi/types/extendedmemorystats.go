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

// ExtendedMemoryStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/nodes/_types/Stats.ts#L261-L264
type ExtendedMemoryStats struct {
	AdjustedTotalInBytes *int64  `json:"adjusted_total_in_bytes,omitempty"`
	FreeInBytes          *int64  `json:"free_in_bytes,omitempty"`
	FreePercent          *int    `json:"free_percent,omitempty"`
	Resident             *string `json:"resident,omitempty"`
	ResidentInBytes      *int64  `json:"resident_in_bytes,omitempty"`
	Share                *string `json:"share,omitempty"`
	ShareInBytes         *int64  `json:"share_in_bytes,omitempty"`
	TotalInBytes         *int64  `json:"total_in_bytes,omitempty"`
	TotalVirtual         *string `json:"total_virtual,omitempty"`
	TotalVirtualInBytes  *int64  `json:"total_virtual_in_bytes,omitempty"`
	UsedInBytes          *int64  `json:"used_in_bytes,omitempty"`
	UsedPercent          *int    `json:"used_percent,omitempty"`
}

func (s *ExtendedMemoryStats) UnmarshalJSON(data []byte) error {

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

		case "adjusted_total_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.AdjustedTotalInBytes = &value
			case float64:
				f := int64(v)
				s.AdjustedTotalInBytes = &f
			}

		case "free_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.FreeInBytes = &value
			case float64:
				f := int64(v)
				s.FreeInBytes = &f
			}

		case "free_percent":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.FreePercent = &value
			case float64:
				f := int(v)
				s.FreePercent = &f
			}

		case "resident":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Resident = &o

		case "resident_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.ResidentInBytes = &value
			case float64:
				f := int64(v)
				s.ResidentInBytes = &f
			}

		case "share":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Share = &o

		case "share_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.ShareInBytes = &value
			case float64:
				f := int64(v)
				s.ShareInBytes = &f
			}

		case "total_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.TotalInBytes = &value
			case float64:
				f := int64(v)
				s.TotalInBytes = &f
			}

		case "total_virtual":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.TotalVirtual = &o

		case "total_virtual_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.TotalVirtualInBytes = &value
			case float64:
				f := int64(v)
				s.TotalVirtualInBytes = &f
			}

		case "used_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.UsedInBytes = &value
			case float64:
				f := int64(v)
				s.UsedInBytes = &f
			}

		case "used_percent":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.UsedPercent = &value
			case float64:
				f := int(v)
				s.UsedPercent = &f
			}

		}
	}
	return nil
}

// NewExtendedMemoryStats returns a ExtendedMemoryStats.
func NewExtendedMemoryStats() *ExtendedMemoryStats {
	r := &ExtendedMemoryStats{}

	return r
}
