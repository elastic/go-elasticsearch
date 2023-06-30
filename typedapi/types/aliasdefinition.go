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

// AliasDefinition type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/indices/_types/AliasDefinition.ts#L22-L30
type AliasDefinition struct {
	Filter        *Query  `json:"filter,omitempty"`
	IndexRouting  *string `json:"index_routing,omitempty"`
	IsHidden      *bool   `json:"is_hidden,omitempty"`
	IsWriteIndex  *bool   `json:"is_write_index,omitempty"`
	Routing       *string `json:"routing,omitempty"`
	SearchRouting *string `json:"search_routing,omitempty"`
}

func (s *AliasDefinition) UnmarshalJSON(data []byte) error {

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

		case "filter":
			if err := dec.Decode(&s.Filter); err != nil {
				return err
			}

		case "index_routing":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.IndexRouting = &o

		case "is_hidden":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.IsHidden = &value
			case bool:
				s.IsHidden = &v
			}

		case "is_write_index":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.IsWriteIndex = &value
			case bool:
				s.IsWriteIndex = &v
			}

		case "routing":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Routing = &o

		case "search_routing":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SearchRouting = &o

		}
	}
	return nil
}

// NewAliasDefinition returns a AliasDefinition.
func NewAliasDefinition() *AliasDefinition {
	r := &AliasDefinition{}

	return r
}
