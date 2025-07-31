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

// TruncateTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/analysis/token_filters.ts#L435-L439
type TruncateTokenFilter struct {
	// Length Character limit for each token. Tokens exceeding this limit are truncated.
	// Defaults to `10`.
	Length  *int    `json:"length,omitempty"`
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *TruncateTokenFilter) UnmarshalJSON(data []byte) error {

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

		case "length":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Length", err)
				}
				s.Length = &value
			case float64:
				f := int(v)
				s.Length = &f
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
func (s TruncateTokenFilter) MarshalJSON() ([]byte, error) {
	type innerTruncateTokenFilter TruncateTokenFilter
	tmp := innerTruncateTokenFilter{
		Length:  s.Length,
		Type:    s.Type,
		Version: s.Version,
	}

	tmp.Type = "truncate"

	return json.Marshal(tmp)
}

// NewTruncateTokenFilter returns a TruncateTokenFilter.
func NewTruncateTokenFilter() *TruncateTokenFilter {
	r := &TruncateTokenFilter{}

	return r
}
