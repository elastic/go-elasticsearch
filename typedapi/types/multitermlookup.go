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
// https://github.com/elastic/elasticsearch-specification/tree/55f8d05b44cea956ae5ceddfcb02770ea2a24ff6

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// MultiTermLookup type.
//
// https://github.com/elastic/elasticsearch-specification/blob/55f8d05b44cea956ae5ceddfcb02770ea2a24ff6/specification/_types/aggregations/bucket.ts#L646-L666
type MultiTermLookup struct {
	// Field A field from which to retrieve terms.
	// It is required if `script` is not provided.
	Field *string `json:"field,omitempty"`
	// Missing The value to apply to documents that do not have a value.
	// By default, documents without a value are ignored.
	Missing Missing `json:"missing,omitempty"`
	// Script A script to calculate terms to aggregate on.
	// It is required if `field` is not provided.
	Script *Script `json:"script,omitempty"`
}

func (s *MultiTermLookup) UnmarshalJSON(data []byte) error {

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

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return fmt.Errorf("%s | %w", "Field", err)
			}

		case "missing":
			if err := dec.Decode(&s.Missing); err != nil {
				return fmt.Errorf("%s | %w", "Missing", err)
			}

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return fmt.Errorf("%s | %w", "Script", err)
			}

		}
	}
	return nil
}

// NewMultiTermLookup returns a MultiTermLookup.
func NewMultiTermLookup() *MultiTermLookup {
	r := &MultiTermLookup{}

	return r
}

type MultiTermLookupVariant interface {
	MultiTermLookupCaster() *MultiTermLookup
}

func (s *MultiTermLookup) MultiTermLookupCaster() *MultiTermLookup {
	return s
}
