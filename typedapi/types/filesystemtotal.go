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

// FileSystemTotal type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/nodes/_types/Stats.ts#L830-L859
type FileSystemTotal struct {
	// Available Total disk space available to this Java virtual machine on all file stores.
	// Depending on OS or process level restrictions, this might appear less than
	// `free`.
	// This is the actual amount of free disk space the Elasticsearch node can
	// utilise.
	Available *string `json:"available,omitempty"`
	// AvailableInBytes Total number of bytes available to this Java virtual machine on all file
	// stores.
	// Depending on OS or process level restrictions, this might appear less than
	// `free_in_bytes`.
	// This is the actual amount of free disk space the Elasticsearch node can
	// utilise.
	AvailableInBytes *int64 `json:"available_in_bytes,omitempty"`
	// Free Total unallocated disk space in all file stores.
	Free *string `json:"free,omitempty"`
	// FreeInBytes Total number of unallocated bytes in all file stores.
	FreeInBytes *int64 `json:"free_in_bytes,omitempty"`
	// Total Total size of all file stores.
	Total *string `json:"total,omitempty"`
	// TotalInBytes Total size of all file stores in bytes.
	TotalInBytes *int64 `json:"total_in_bytes,omitempty"`
}

func (s *FileSystemTotal) UnmarshalJSON(data []byte) error {

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

		case "available":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Available", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Available = &o

		case "available_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "AvailableInBytes", err)
				}
				s.AvailableInBytes = &value
			case float64:
				f := int64(v)
				s.AvailableInBytes = &f
			}

		case "free":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Free", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Free = &o

		case "free_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "FreeInBytes", err)
				}
				s.FreeInBytes = &value
			case float64:
				f := int64(v)
				s.FreeInBytes = &f
			}

		case "total":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Total", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Total = &o

		case "total_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "TotalInBytes", err)
				}
				s.TotalInBytes = &value
			case float64:
				f := int64(v)
				s.TotalInBytes = &f
			}

		}
	}
	return nil
}

// NewFileSystemTotal returns a FileSystemTotal.
func NewFileSystemTotal() *FileSystemTotal {
	r := &FileSystemTotal{}

	return r
}
