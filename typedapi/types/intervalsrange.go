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

// IntervalsRange type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/query_dsl/fulltext.ts#L237-L264
type IntervalsRange struct {
	// Analyzer Analyzer used to analyze the `prefix`.
	Analyzer *string `json:"analyzer,omitempty"`
	// Gt Lower term, either gte or gt must be provided.
	Gt *string `json:"gt,omitempty"`
	// Gte Lower term, either gte or gt must be provided.
	Gte *string `json:"gte,omitempty"`
	// Lt Upper term, either lte or lt must be provided.
	Lt *string `json:"lt,omitempty"`
	// Lte Upper term, either lte or lt must be provided.
	Lte *string `json:"lte,omitempty"`
	// UseField If specified, match intervals from this field rather than the top-level
	// field.
	// The `prefix` is normalized using the search analyzer from this field, unless
	// `analyzer` is specified separately.
	UseField *string `json:"use_field,omitempty"`
}

func (s *IntervalsRange) UnmarshalJSON(data []byte) error {

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

		case "analyzer":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Analyzer", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Analyzer = &o

		case "gt":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Gt", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Gt = &o

		case "gte":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Gte", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Gte = &o

		case "lt":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Lt", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Lt = &o

		case "lte":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Lte", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Lte = &o

		case "use_field":
			if err := dec.Decode(&s.UseField); err != nil {
				return fmt.Errorf("%s | %w", "UseField", err)
			}

		}
	}
	return nil
}

// NewIntervalsRange returns a IntervalsRange.
func NewIntervalsRange() *IntervalsRange {
	r := &IntervalsRange{}

	return r
}
