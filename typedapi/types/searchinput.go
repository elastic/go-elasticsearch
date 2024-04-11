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

// SearchInput type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1/specification/watcher/_types/Input.ts#L112-L116
type SearchInput struct {
	Extract []string                     `json:"extract,omitempty"`
	Request SearchInputRequestDefinition `json:"request"`
	Timeout Duration                     `json:"timeout,omitempty"`
}

func (s *SearchInput) UnmarshalJSON(data []byte) error {

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

		case "extract":
			if err := dec.Decode(&s.Extract); err != nil {
				return fmt.Errorf("%s | %w", "Extract", err)
			}

		case "request":
			if err := dec.Decode(&s.Request); err != nil {
				return fmt.Errorf("%s | %w", "Request", err)
			}

		case "timeout":
			if err := dec.Decode(&s.Timeout); err != nil {
				return fmt.Errorf("%s | %w", "Timeout", err)
			}

		}
	}
	return nil
}

// NewSearchInput returns a SearchInput.
func NewSearchInput() *SearchInput {
	r := &SearchInput{}

	return r
}
