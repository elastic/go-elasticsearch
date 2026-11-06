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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Segment type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/indices/segments/types.ts#L28-L75
type Segment struct {
	// Attributes Contains information about whether high compression was enabled and per-field
	// vector formats.
	Attributes map[string]string `json:"attributes,omitempty"`
	// AutoCalibration The auto-calibration information of the segment, per `dense_vector` field
	// that supports auto-calibration. Only returned when `include_auto_calibration`
	// is `true`.
	AutoCalibration map[string]SegmentAutoCalibration `json:"auto_calibration,omitempty"`
	// Committed If `true`, the segment is synced to disk. Segments that are synced can
	// survive a hard reboot. If `false`, the data from uncommitted segments is also
	// stored in the transaction log so that Elasticsearch is able to replay changes
	// on the next start.
	Committed bool `json:"committed"`
	// Compound If `true`, Lucene merged all files from the segment into a single file to
	// save file descriptors.
	Compound *bool `json:"compound,omitempty"`
	// DeletedDocs The number of deleted documents as reported by Lucene, which may be higher or
	// lower than the number of delete operations you have performed. This number
	// excludes deletes that were performed recently and do not yet belong to a
	// segment. Deleted documents are cleaned up by the automatic merge process if
	// it makes sense to do so. Also, Elasticsearch creates extra deleted documents
	// to internally track the recent history of operations on a shard.
	DeletedDocs int64 `json:"deleted_docs"`
	// Generation Generation number, such as `0`. Elasticsearch increments this generation
	// number for each segment written then uses this number to derive the segment
	// name.
	Generation int `json:"generation"`
	// MergeId The ID of the merge the segment is part of. Only returned while the segment
	// is being merged.
	MergeId *string `json:"merge_id,omitempty"`
	// NumDocs The number of documents as reported by Lucene. This excludes deleted
	// documents and counts any nested documents separately from their parents. It
	// also excludes documents which were indexed recently and do not yet belong to
	// a segment.
	NumDocs int64 `json:"num_docs"`
	// Search If `true`, the segment is searchable. If `false`, the segment has most likely
	// been written to disk but needs a refresh to be searchable.
	Search bool `json:"search"`
	// Size Disk space used by the segment.
	Size ByteSize `json:"size,omitempty"`
	// SizeInBytes Disk space used by the segment, in bytes.
	SizeInBytes Float64 `json:"size_in_bytes"`
	// Sort The index sort of the segment. Only returned when the index is sorted.
	Sort []SegmentSort `json:"sort,omitempty"`
	// Version Version of Lucene used to write the segment.
	Version *string `json:"version,omitempty"`
}

func (s *Segment) UnmarshalJSON(data []byte) error {

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

		case "attributes":
			if s.Attributes == nil {
				s.Attributes = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.Attributes); err != nil {
				return fmt.Errorf("%s | %w", "Attributes", err)
			}

		case "auto_calibration":
			if s.AutoCalibration == nil {
				s.AutoCalibration = make(map[string]SegmentAutoCalibration, 0)
			}
			refs := make(map[string]json.RawMessage, 0)
			dec.Decode(&refs)
			for key, message := range refs {
				kind := make(map[string]any)
				buf := bytes.NewReader(message)
				localDec := json.NewDecoder(buf)
				localDec.Decode(&kind)
				buf.Seek(0, io.SeekStart)

				switch kind["type"] {
				case "osq":
					oo := NewOsqSegmentAutoCalibration()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("AutoCalibration | %w", err)
					}
					s.AutoCalibration[key] = oo
				default:
					oo := new(SegmentAutoCalibration)
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("new(SegmentAutoCalibration) | %w", err)
					}
					s.AutoCalibration[key] = oo
				}
			}

		case "committed":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Committed", err)
				}
				s.Committed = value
			case bool:
				s.Committed = v
			}

		case "compound":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Compound", err)
				}
				s.Compound = &value
			case bool:
				s.Compound = &v
			}

		case "deleted_docs":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "DeletedDocs", err)
				}
				s.DeletedDocs = value
			case float64:
				f := int64(v)
				s.DeletedDocs = f
			}

		case "generation":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Generation", err)
				}
				s.Generation = value
			case float64:
				f := int(v)
				s.Generation = f
			}

		case "merge_id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "MergeId", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MergeId = &o

		case "num_docs":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumDocs", err)
				}
				s.NumDocs = value
			case float64:
				f := int64(v)
				s.NumDocs = f
			}

		case "search":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Search", err)
				}
				s.Search = value
			case bool:
				s.Search = v
			}

		case "size":
			if err := dec.Decode(&s.Size); err != nil {
				return fmt.Errorf("%s | %w", "Size", err)
			}

		case "size_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "SizeInBytes", err)
				}
				f := Float64(value)
				s.SizeInBytes = f
			case float64:
				f := Float64(v)
				s.SizeInBytes = f
			}

		case "sort":
			if err := dec.Decode(&s.Sort); err != nil {
				return fmt.Errorf("%s | %w", "Sort", err)
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		}
	}
	return nil
}

// NewSegment returns a Segment.
func NewSegment() *Segment {
	r := &Segment{
		Attributes:      make(map[string]string),
		AutoCalibration: make(map[string]SegmentAutoCalibration),
	}

	return r
}
