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

// SegmentsRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/cat/segments/types.ts#L22-L96
type SegmentsRecord struct {
	// Committed is segment committed
	Committed *string `json:"committed,omitempty"`
	// Compound is segment compound
	Compound *string `json:"compound,omitempty"`
	// DocsCount number of docs in segment
	DocsCount *string `json:"docs.count,omitempty"`
	// DocsDeleted number of deleted docs in segment
	DocsDeleted *string `json:"docs.deleted,omitempty"`
	// Generation segment generation
	Generation *string `json:"generation,omitempty"`
	// Id unique id of node where it lives
	Id *string `json:"id,omitempty"`
	// Index index name
	Index *string `json:"index,omitempty"`
	// Ip ip of node where it lives
	Ip *string `json:"ip,omitempty"`
	// Prirep primary or replica
	Prirep *string `json:"prirep,omitempty"`
	// Searchable is segment searched
	Searchable *string `json:"searchable,omitempty"`
	// Segment segment name
	Segment *string `json:"segment,omitempty"`
	// Shard shard name
	Shard *string `json:"shard,omitempty"`
	// Size segment size in bytes
	Size ByteSize `json:"size,omitempty"`
	// SizeMemory segment memory in bytes
	SizeMemory ByteSize `json:"size.memory,omitempty"`
	// Version version
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
