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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// DataPathStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/nodes/_types/Stats.ts#L550-L594
type DataPathStats struct {
	// Available Total amount of disk space available to this Java virtual machine on this
	// file store.
	Available *string `json:"available,omitempty"`
	// AvailableInBytes Total number of bytes available to this Java virtual machine on this file
	// store.
	AvailableInBytes     *int64  `json:"available_in_bytes,omitempty"`
	DiskQueue            *string `json:"disk_queue,omitempty"`
	DiskReadSize         *string `json:"disk_read_size,omitempty"`
	DiskReadSizeInBytes  *int64  `json:"disk_read_size_in_bytes,omitempty"`
	DiskReads            *int64  `json:"disk_reads,omitempty"`
	DiskWriteSize        *string `json:"disk_write_size,omitempty"`
	DiskWriteSizeInBytes *int64  `json:"disk_write_size_in_bytes,omitempty"`
	DiskWrites           *int64  `json:"disk_writes,omitempty"`
	// Free Total amount of unallocated disk space in the file store.
	Free *string `json:"free,omitempty"`
	// FreeInBytes Total number of unallocated bytes in the file store.
	FreeInBytes *int64 `json:"free_in_bytes,omitempty"`
	// Mount Mount point of the file store (for example: `/dev/sda2`).
	Mount *string `json:"mount,omitempty"`
	// Path Path to the file store.
	Path *string `json:"path,omitempty"`
	// Total Total size of the file store.
	Total *string `json:"total,omitempty"`
	// TotalInBytes Total size of the file store in bytes.
	TotalInBytes *int64 `json:"total_in_bytes,omitempty"`
	// Type Type of the file store (ex: ext4).
	Type *string `json:"type,omitempty"`
}

func (s *DataPathStats) UnmarshalJSON(data []byte) error {

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
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Available = &o

		case "available_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.AvailableInBytes = &value
			case float64:
				f := int64(v)
				s.AvailableInBytes = &f
			}

		case "disk_queue":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DiskQueue = &o

		case "disk_read_size":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DiskReadSize = &o

		case "disk_read_size_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.DiskReadSizeInBytes = &value
			case float64:
				f := int64(v)
				s.DiskReadSizeInBytes = &f
			}

		case "disk_reads":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.DiskReads = &value
			case float64:
				f := int64(v)
				s.DiskReads = &f
			}

		case "disk_write_size":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DiskWriteSize = &o

		case "disk_write_size_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.DiskWriteSizeInBytes = &value
			case float64:
				f := int64(v)
				s.DiskWriteSizeInBytes = &f
			}

		case "disk_writes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.DiskWrites = &value
			case float64:
				f := int64(v)
				s.DiskWrites = &f
			}

		case "free":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Free = &o

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

		case "mount":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Mount = &o

		case "path":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Path = &o

		case "total":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Total = &o

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

		case "type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Type = &o

		}
	}
	return nil
}

// NewDataPathStats returns a DataPathStats.
func NewDataPathStats() *DataPathStats {
	r := &DataPathStats{}

	return r
}
