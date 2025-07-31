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

// ClusterFileSystem type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/cluster/stats/types.ts#L35-L75
type ClusterFileSystem struct {
	// Available Total number of bytes available to JVM in file stores across all selected
	// nodes.
	// Depending on operating system or process-level restrictions, this number may
	// be less than `nodes.fs.free_in_byes`.
	// This is the actual amount of free disk space the selected Elasticsearch nodes
	// can use.
	Available ByteSize `json:"available,omitempty"`
	// AvailableInBytes Total number of bytes available to JVM in file stores across all selected
	// nodes.
	// Depending on operating system or process-level restrictions, this number may
	// be less than `nodes.fs.free_in_byes`.
	// This is the actual amount of free disk space the selected Elasticsearch nodes
	// can use.
	AvailableInBytes           *int64   `json:"available_in_bytes,omitempty"`
	FloodStageFreeSpace        ByteSize `json:"flood_stage_free_space,omitempty"`
	FloodStageFreeSpaceInBytes *int64   `json:"flood_stage_free_space_in_bytes,omitempty"`
	// Free Total number of unallocated bytes in file stores across all selected nodes.
	Free ByteSize `json:"free,omitempty"`
	// FreeInBytes Total number, in bytes, of unallocated bytes in file stores across all
	// selected nodes.
	FreeInBytes                      *int64   `json:"free_in_bytes,omitempty"`
	FrozenFloodStageFreeSpace        ByteSize `json:"frozen_flood_stage_free_space,omitempty"`
	FrozenFloodStageFreeSpaceInBytes *int64   `json:"frozen_flood_stage_free_space_in_bytes,omitempty"`
	HighWatermarkFreeSpace           ByteSize `json:"high_watermark_free_space,omitempty"`
	HighWatermarkFreeSpaceInBytes    *int64   `json:"high_watermark_free_space_in_bytes,omitempty"`
	LowWatermarkFreeSpace            ByteSize `json:"low_watermark_free_space,omitempty"`
	LowWatermarkFreeSpaceInBytes     *int64   `json:"low_watermark_free_space_in_bytes,omitempty"`
	Mount                            *string  `json:"mount,omitempty"`
	Path                             *string  `json:"path,omitempty"`
	// Total Total size of all file stores across all selected nodes.
	Total ByteSize `json:"total,omitempty"`
	// TotalInBytes Total size, in bytes, of all file stores across all selected nodes.
	TotalInBytes *int64  `json:"total_in_bytes,omitempty"`
	Type         *string `json:"type,omitempty"`
}

func (s *ClusterFileSystem) UnmarshalJSON(data []byte) error {

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
			if err := dec.Decode(&s.Available); err != nil {
				return fmt.Errorf("%s | %w", "Available", err)
			}

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

		case "flood_stage_free_space":
			if err := dec.Decode(&s.FloodStageFreeSpace); err != nil {
				return fmt.Errorf("%s | %w", "FloodStageFreeSpace", err)
			}

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
			if err := dec.Decode(&s.Free); err != nil {
				return fmt.Errorf("%s | %w", "Free", err)
			}

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
			if err := dec.Decode(&s.FrozenFloodStageFreeSpace); err != nil {
				return fmt.Errorf("%s | %w", "FrozenFloodStageFreeSpace", err)
			}

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
			if err := dec.Decode(&s.HighWatermarkFreeSpace); err != nil {
				return fmt.Errorf("%s | %w", "HighWatermarkFreeSpace", err)
			}

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
			if err := dec.Decode(&s.LowWatermarkFreeSpace); err != nil {
				return fmt.Errorf("%s | %w", "LowWatermarkFreeSpace", err)
			}

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
			if err := dec.Decode(&s.Total); err != nil {
				return fmt.Errorf("%s | %w", "Total", err)
			}

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

// NewClusterFileSystem returns a ClusterFileSystem.
func NewClusterFileSystem() *ClusterFileSystem {
	r := &ClusterFileSystem{}

	return r
}
