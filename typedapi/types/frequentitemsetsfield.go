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
// https://github.com/elastic/elasticsearch-specification/tree/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// FrequentItemSetsField type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1/specification/_types/aggregations/bucket.ts#L1144-L1156
type FrequentItemSetsField struct {
	// Exclude Values to exclude.
	// Can be regular expression strings or arrays of strings of exact terms.
	Exclude []string `json:"exclude,omitempty"`
	Field   string   `json:"field"`
	// Include Values to include.
	// Can be regular expression strings or arrays of strings of exact terms.
	Include TermsInclude `json:"include,omitempty"`
}

func (s *FrequentItemSetsField) UnmarshalJSON(data []byte) error {

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

		case "exclude":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Exclude", err)
				}

				s.Exclude = append(s.Exclude, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Exclude); err != nil {
					return fmt.Errorf("%s | %w", "Exclude", err)
				}
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return fmt.Errorf("%s | %w", "Field", err)
			}

		case "include":
			if err := dec.Decode(&s.Include); err != nil {
				return fmt.Errorf("%s | %w", "Include", err)
			}

		}
	}
	return nil
}

// NewFrequentItemSetsField returns a FrequentItemSetsField.
func NewFrequentItemSetsField() *FrequentItemSetsField {
	r := &FrequentItemSetsField{}

	return r
}
