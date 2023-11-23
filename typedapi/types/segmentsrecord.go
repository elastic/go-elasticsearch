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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// SegmentsRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/cat/segments/types.ts#L22-L107
type SegmentsRecord struct {
	// Committed If `true`, the segment is synced to disk.
	// Segments that are synced can survive a hard reboot.
	// If `false`, the data from uncommitted segments is also stored in the
	// transaction log so that Elasticsearch is able to replay changes on the next
	// start.
	Committed *string `json:"committed,omitempty"`
	// Compound If `true`, the segment is stored in a compound file.
	// This means Lucene merged all files from the segment in a single file to save
	// file descriptors.
	Compound *string `json:"compound,omitempty"`
	// DocsCount The number of documents in the segment.
	// This excludes deleted documents and counts any nested documents separately
	// from their parents.
	// It also excludes documents which were indexed recently and do not yet belong
	// to a segment.
	DocsCount *string `json:"docs.count,omitempty"`
	// DocsDeleted The number of deleted documents in the segment, which might be higher or
	// lower than the number of delete operations you have performed.
	// This number excludes deletes that were performed recently and do not yet
	// belong to a segment.
	// Deleted documents are cleaned up by the automatic merge process if it makes
	// sense to do so.
	// Also, Elasticsearch creates extra deleted documents to internally track the
	// recent history of operations on a shard.
	DocsDeleted *string `json:"docs.deleted,omitempty"`
	// Generation The segment generation number.
	// Elasticsearch increments this generation number for each segment written then
	// uses this number to derive the segment name.
	Generation *string `json:"generation,omitempty"`
	// Id The unique identifier of the node where it lives.
	Id *string `json:"id,omitempty"`
	// Index The index name.
	Index *string `json:"index,omitempty"`
	// Ip The IP address of the node where it lives.
	Ip *string `json:"ip,omitempty"`
	// Prirep The shard type: `primary` or `replica`.
	Prirep *string `json:"prirep,omitempty"`
	// Searchable If `true`, the segment is searchable.
	// If `false`, the segment has most likely been written to disk but needs a
	// refresh to be searchable.
	Searchable *string `json:"searchable,omitempty"`
	// Segment The segment name, which is derived from the segment generation and used
	// internally to create file names in the directory of the shard.
	Segment *string `json:"segment,omitempty"`
	// Shard The shard name.
	Shard *string `json:"shard,omitempty"`
	// Size The segment size in bytes.
	Size ByteSize `json:"size,omitempty"`
	// SizeMemory The segment memory in bytes.
	// A value of `-1` indicates Elasticsearch was unable to compute this number.
	SizeMemory ByteSize `json:"size.memory,omitempty"`
	// Version The version of Lucene used to write the segment.
	Version *string `json:"version,omitempty"`
}

func (s *SegmentsRecord) UnmarshalJSON(data []byte) error {

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

		case "committed", "ic", "isCommitted":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Committed = &o

		case "compound", "ico", "isCompound":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Compound = &o

		case "docs.count", "dc", "docsCount":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DocsCount = &o

		case "docs.deleted", "dd", "docsDeleted":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DocsDeleted = &o

		case "generation", "g", "gen":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Generation = &o

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return err
			}

		case "index", "i", "idx":
			if err := dec.Decode(&s.Index); err != nil {
				return err
			}

		case "ip":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Ip = &o

		case "prirep", "p", "pr", "primaryOrReplica":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Prirep = &o

		case "searchable", "is", "isSearchable":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Searchable = &o

		case "segment", "seg":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Segment = &o

		case "shard", "s", "sh":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Shard = &o

		case "size", "si":
			if err := dec.Decode(&s.Size); err != nil {
				return err
			}

		case "size.memory", "sm", "sizeMemory":
			if err := dec.Decode(&s.SizeMemory); err != nil {
				return err
			}

		case "version", "v":
			if err := dec.Decode(&s.Version); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewSegmentsRecord returns a SegmentsRecord.
func NewSegmentsRecord() *SegmentsRecord {
	r := &SegmentsRecord{}

	return r
}
