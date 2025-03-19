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
// https://github.com/elastic/elasticsearch-specification/tree/3ea9ce260df22d3244bff5bace485dd97ff4046d

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Alias type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3ea9ce260df22d3244bff5bace485dd97ff4046d/specification/indices/_types/Alias.ts#L23-L53
type Alias struct {
	// Filter Query used to limit documents the alias can access.
	Filter *Query `json:"filter,omitempty"`
	// IndexRouting Value used to route indexing operations to a specific shard.
	// If specified, this overwrites the `routing` value for indexing operations.
	IndexRouting *string `json:"index_routing,omitempty"`
	// IsHidden If `true`, the alias is hidden.
	// All indices for the alias must have the same `is_hidden` value.
	IsHidden *bool `json:"is_hidden,omitempty"`
	// IsWriteIndex If `true`, the index is the write index for the alias.
	IsWriteIndex *bool `json:"is_write_index,omitempty"`
	// Routing Value used to route indexing and search operations to a specific shard.
	Routing *string `json:"routing,omitempty"`
	// SearchRouting Value used to route search operations to a specific shard.
	// If specified, this overwrites the `routing` value for search operations.
	SearchRouting *string `json:"search_routing,omitempty"`
}

func (s *Alias) UnmarshalJSON(data []byte) error {

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
				return fmt.Errorf("%s | %w", "Filter", err)
			}

		case "index_routing":
			if err := dec.Decode(&s.IndexRouting); err != nil {
				return fmt.Errorf("%s | %w", "IndexRouting", err)
			}

		case "is_hidden":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IsHidden", err)
				}
				s.IsHidden = &value
			case bool:
				s.IsHidden = &v
			}

		case "is_write_index":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IsWriteIndex", err)
				}
				s.IsWriteIndex = &value
			case bool:
				s.IsWriteIndex = &v
			}

		case "routing":
			if err := dec.Decode(&s.Routing); err != nil {
				return fmt.Errorf("%s | %w", "Routing", err)
			}

		case "search_routing":
			if err := dec.Decode(&s.SearchRouting); err != nil {
				return fmt.Errorf("%s | %w", "SearchRouting", err)
			}

		}
	}
	return nil
}

// NewAlias returns a Alias.
func NewAlias() *Alias {
	r := &Alias{}

	return r
}

// true

type AliasVariant interface {
	AliasCaster() *Alias
}

func (s *Alias) AliasCaster() *Alias {
	return s
}
