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

// AddAction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/indices/update_aliases/types.ts#L41-L95
type AddAction struct {
	// Alias Alias for the action.
	// Index alias names support date math.
	Alias *string `json:"alias,omitempty"`
	// Aliases Aliases for the action.
	// Index alias names support date math.
	Aliases []string `json:"aliases,omitempty"`
	// Filter Query used to limit documents the alias can access.
	Filter *Query `json:"filter,omitempty"`
	// Index Data stream or index for the action.
	// Supports wildcards (`*`).
	Index *string `json:"index,omitempty"`
	// IndexRouting Value used to route indexing operations to a specific shard.
	// If specified, this overwrites the `routing` value for indexing operations.
	// Data stream aliases don’t support this parameter.
	IndexRouting *string `json:"index_routing,omitempty"`
	// Indices Data streams or indices for the action.
	// Supports wildcards (`*`).
	Indices []string `json:"indices,omitempty"`
	// IsHidden If `true`, the alias is hidden.
	IsHidden *bool `json:"is_hidden,omitempty"`
	// IsWriteIndex If `true`, sets the write index or data stream for the alias.
	IsWriteIndex *bool `json:"is_write_index,omitempty"`
	// MustExist If `true`, the alias must exist to perform the action.
	MustExist *bool `json:"must_exist,omitempty"`
	// Routing Value used to route indexing and search operations to a specific shard.
	// Data stream aliases don’t support this parameter.
	Routing *string `json:"routing,omitempty"`
	// SearchRouting Value used to route search operations to a specific shard.
	// If specified, this overwrites the `routing` value for search operations.
	// Data stream aliases don’t support this parameter.
	SearchRouting *string `json:"search_routing,omitempty"`
}

func (s *AddAction) UnmarshalJSON(data []byte) error {

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

		case "alias":
			if err := dec.Decode(&s.Alias); err != nil {
				return err
			}

		case "aliases":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.Aliases = append(s.Aliases, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Aliases); err != nil {
					return err
				}
			}

		case "filter":
			if err := dec.Decode(&s.Filter); err != nil {
				return err
			}

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return err
			}

		case "index_routing":
			if err := dec.Decode(&s.IndexRouting); err != nil {
				return err
			}

		case "indices":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.Indices = append(s.Indices, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Indices); err != nil {
					return err
				}
			}

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

		case "must_exist":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.MustExist = &value
			case bool:
				s.MustExist = &v
			}

		case "routing":
			if err := dec.Decode(&s.Routing); err != nil {
				return err
			}

		case "search_routing":
			if err := dec.Decode(&s.SearchRouting); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewAddAction returns a AddAction.
func NewAddAction() *AddAction {
	r := &AddAction{}

	return r
}
