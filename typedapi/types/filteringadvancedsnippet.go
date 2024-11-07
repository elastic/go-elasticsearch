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
// https://github.com/elastic/elasticsearch-specification/tree/4fcf747dfafc951e1dcf3077327e3dcee9107db3

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// FilteringAdvancedSnippet type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4fcf747dfafc951e1dcf3077327e3dcee9107db3/specification/connector/_types/Connector.ts#L192-L196
type FilteringAdvancedSnippet struct {
	CreatedAt DateTime        `json:"created_at,omitempty"`
	UpdatedAt DateTime        `json:"updated_at,omitempty"`
	Value     json.RawMessage `json:"value,omitempty"`
}

func (s *FilteringAdvancedSnippet) UnmarshalJSON(data []byte) error {

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

		case "created_at":
			if err := dec.Decode(&s.CreatedAt); err != nil {
				return fmt.Errorf("%s | %w", "CreatedAt", err)
			}

		case "updated_at":
			if err := dec.Decode(&s.UpdatedAt); err != nil {
				return fmt.Errorf("%s | %w", "UpdatedAt", err)
			}

		case "value":
			if err := dec.Decode(&s.Value); err != nil {
				return fmt.Errorf("%s | %w", "Value", err)
			}

		}
	}
	return nil
}

// NewFilteringAdvancedSnippet returns a FilteringAdvancedSnippet.
func NewFilteringAdvancedSnippet() *FilteringAdvancedSnippet {
	r := &FilteringAdvancedSnippet{}

	return r
}
