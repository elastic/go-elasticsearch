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

// AliasesRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/cat/aliases/types.ts#L22-L53
type AliasesRecord struct {
	// Alias alias name
	Alias *string `json:"alias,omitempty"`
	// Filter filter
	Filter *string `json:"filter,omitempty"`
	// Index index alias points to
	Index *string `json:"index,omitempty"`
	// IsWriteIndex write index
	IsWriteIndex *string `json:"is_write_index,omitempty"`
	// RoutingIndex index routing
	RoutingIndex *string `json:"routing.index,omitempty"`
	// RoutingSearch search routing
	RoutingSearch *string `json:"routing.search,omitempty"`
}

func (s *AliasesRecord) UnmarshalJSON(data []byte) error {

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

		case "alias", "a":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Alias = &o

		case "filter", "f", "fi":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Filter = &o

		case "index", "i", "idx":
			if err := dec.Decode(&s.Index); err != nil {
				return err
			}

		case "is_write_index", "w", "isWriteIndex":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.IsWriteIndex = &o

		case "routing.index", "ri", "routingIndex":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.RoutingIndex = &o

		case "routing.search", "rs", "routingSearch":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.RoutingSearch = &o

		}
	}
	return nil
}

// NewAliasesRecord returns a AliasesRecord.
func NewAliasesRecord() *AliasesRecord {
	r := &AliasesRecord{}

	return r
}
