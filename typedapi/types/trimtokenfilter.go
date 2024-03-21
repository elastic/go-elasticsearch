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
// https://github.com/elastic/elasticsearch-specification/tree/accc26662ab4c58f4f6fb0fc1d9fc5249d0de339

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// TrimTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/accc26662ab4c58f4f6fb0fc1d9fc5249d0de339/specification/_types/analysis/token_filters.ts#L326-L328
type TrimTokenFilter struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *TrimTokenFilter) UnmarshalJSON(data []byte) error {

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
func (s TrimTokenFilter) MarshalJSON() ([]byte, error) {
	type innerTrimTokenFilter TrimTokenFilter
	tmp := innerTrimTokenFilter{
		Type:    s.Type,
		Version: s.Version,
	}

	tmp.Type = "trim"

	return json.Marshal(tmp)
}

// NewTrimTokenFilter returns a TrimTokenFilter.
func NewTrimTokenFilter() *TrimTokenFilter {
	r := &TrimTokenFilter{}

	return r
}
