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
// https://github.com/elastic/elasticsearch-specification/tree/abf9c2c6bb21328339daa197aae15af2ecbc46f0

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// DataPathStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/abf9c2c6bb21328339daa197aae15af2ecbc46f0/specification/nodes/_types/Stats.ts#L715-L791
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
	// FloodStageFreeSpace The amount of free disk space that, once reached, triggers the flood stage
	// disk watermark.
	FloodStageFreeSpace *string `json:"flood_stage_free_space,omitempty"`
	// FloodStageFreeSpaceInBytes The amount of free disk space, in bytes, that, once reached, triggers the
	// flood stage disk watermark.
	FloodStageFreeSpaceInBytes *int64 `json:"flood_stage_free_space_in_bytes,omitempty"`
	// Free Total amount of unallocated disk space in the file store.
	Free *string `json:"free,omitempty"`
	// FreeInBytes Total number of unallocated bytes in the file store.
	FreeInBytes *int64 `json:"free_in_bytes,omitempty"`
	// FrozenFloodStageFreeSpace The amount of free disk space that, once reached, triggers the frozen flood
	// stage disk watermark.
	FrozenFloodStageFreeSpace *string `json:"frozen_flood_stage_free_space,omitempty"`
	// FrozenFloodStageFreeSpaceInBytes The amount of free disk space, in bytes, that, once reached, triggers the
	// frozen flood stage disk watermark.
	FrozenFloodStageFreeSpaceInBytes *int64 `json:"frozen_flood_stage_free_space_in_bytes,omitempty"`
	// HighWatermarkFreeSpace The amount of free disk space that, once reached, triggers the high disk
	// watermark.
	HighWatermarkFreeSpace *string `json:"high_watermark_free_space,omitempty"`
	// HighWatermarkFreeSpaceInBytes The amount of free disk space, in bytes, that, once reached, triggers the
	// high disk watermark.
	HighWatermarkFreeSpaceInBytes *int64 `json:"high_watermark_free_space_in_bytes,omitempty"`
	// LowWatermarkFreeSpace The amount of free disk space that, once reached, triggers the low disk
	// watermark.
	LowWatermarkFreeSpace *string `json:"low_watermark_free_space,omitempty"`
	// LowWatermarkFreeSpaceInBytes The amount of free disk space, in bytes, that, once reached, triggers the low
	// disk watermark.
	LowWatermarkFreeSpaceInBytes *int64 `json:"low_watermark_free_space_in_bytes,omitempty"`
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

		case "disk_queue":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "DiskQueue", err)
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
				return fmt.Errorf("%s | %w", "DiskReadSize", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DiskReadSize = &o

		case "disk_read_size_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "DiskReadSizeInBytes", err)
				}
				s.DiskReadSizeInBytes = &value
			case float64:
				f := int64(v)
				s.DiskReadSizeInBytes = &f
			}

		case "disk_reads":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "DiskReads", err)
				}
				s.DiskReads = &value
			case float64:
				f := int64(v)
				s.DiskReads = &f
			}

		case "disk_write_size":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "DiskWriteSize", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DiskWriteSize = &o

		case "disk_write_size_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "DiskWriteSizeInBytes", err)
				}
				s.DiskWriteSizeInBytes = &value
			case float64:
				f := int64(v)
				s.DiskWriteSizeInBytes = &f
			}

		case "disk_writes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "DiskWrites", err)
				}
				s.DiskWrites = &value
			case float64:
				f := int64(v)
				s.DiskWrites = &f
			}

		case "flood_stage_free_space":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "FloodStageFreeSpace", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.FloodStageFreeSpace = &o

		case "flood_stage_free_space_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "FloodStageFreeSpaceInBytes", err)
				}
				s.FloodStageFreeSpaceInBytes = &value
			case float64:
				f := int64(v)
				s.FloodStageFreeSpaceInBytes = &f
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

		case "frozen_flood_stage_free_space":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "FrozenFloodStageFreeSpace", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.FrozenFloodStageFreeSpace = &o

		case "frozen_flood_stage_free_space_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "FrozenFloodStageFreeSpaceInBytes", err)
				}
				s.FrozenFloodStageFreeSpaceInBytes = &value
			case float64:
				f := int64(v)
				s.FrozenFloodStageFreeSpaceInBytes = &f
			}

		case "high_watermark_free_space":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "HighWatermarkFreeSpace", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.HighWatermarkFreeSpace = &o

		case "high_watermark_free_space_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "HighWatermarkFreeSpaceInBytes", err)
				}
				s.HighWatermarkFreeSpaceInBytes = &value
			case float64:
				f := int64(v)
				s.HighWatermarkFreeSpaceInBytes = &f
			}

		case "low_watermark_free_space":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "LowWatermarkFreeSpace", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.LowWatermarkFreeSpace = &o

		case "low_watermark_free_space_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "LowWatermarkFreeSpaceInBytes", err)
				}
				s.LowWatermarkFreeSpaceInBytes = &value
			case float64:
				f := int64(v)
				s.LowWatermarkFreeSpaceInBytes = &f
			}

		case "mount":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Mount", err)
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
				return fmt.Errorf("%s | %w", "Path", err)
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

		case "type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
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
