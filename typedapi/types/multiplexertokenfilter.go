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
)

// MultiplexerTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/analysis/token_filters.ts#L356-L362
type MultiplexerTokenFilter struct {
	// Filters A list of token filters to apply to incoming tokens.
	Filters []string `json:"filters"`
	// PreserveOriginal If `true` (the default) then emit the original token in addition to the
	// filtered tokens.
	PreserveOriginal Stringifiedboolean `json:"preserve_original,omitempty"`
	Type             string             `json:"type,omitempty"`
	Version          *string            `json:"version,omitempty"`
}

func (s *MultiplexerTokenFilter) UnmarshalJSON(data []byte) error {

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

		case "filters":
			if err := dec.Decode(&s.Filters); err != nil {
				return fmt.Errorf("%s | %w", "Filters", err)
			}

		case "preserve_original":
			if err := dec.Decode(&s.PreserveOriginal); err != nil {
				return fmt.Errorf("%s | %w", "PreserveOriginal", err)
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s MultiplexerTokenFilter) MarshalJSON() ([]byte, error) {
	type innerMultiplexerTokenFilter MultiplexerTokenFilter
	tmp := innerMultiplexerTokenFilter{
		Filters:          s.Filters,
		PreserveOriginal: s.PreserveOriginal,
		Type:             s.Type,
		Version:          s.Version,
	}

	tmp.Type = "multiplexer"

	return json.Marshal(tmp)
}

// NewMultiplexerTokenFilter returns a MultiplexerTokenFilter.
func NewMultiplexerTokenFilter() *MultiplexerTokenFilter {
	r := &MultiplexerTokenFilter{}

	return r
}
