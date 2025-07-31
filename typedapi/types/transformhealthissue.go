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

// TransformHealthIssue type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/transform/get_transform_stats/types.ts#L51-L63
type TransformHealthIssue struct {
	// Count Number of times this issue has occurred since it started
	Count int `json:"count"`
	// Details Details about the issue
	Details              *string  `json:"details,omitempty"`
	FirstOccurenceString DateTime `json:"first_occurence_string,omitempty"`
	// FirstOccurrence The timestamp this issue occurred for for the first time
	FirstOccurrence *int64 `json:"first_occurrence,omitempty"`
	// Issue A description of the issue
	Issue string `json:"issue"`
	// Type The type of the issue
	Type string `json:"type"`
}

func (s *TransformHealthIssue) UnmarshalJSON(data []byte) error {

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

		case "count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Count", err)
				}
				s.Count = value
			case float64:
				f := int(v)
				s.Count = f
			}

		case "details":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Details", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Details = &o

		case "first_occurence_string":
			if err := dec.Decode(&s.FirstOccurenceString); err != nil {
				return fmt.Errorf("%s | %w", "FirstOccurenceString", err)
			}

		case "first_occurrence":
			if err := dec.Decode(&s.FirstOccurrence); err != nil {
				return fmt.Errorf("%s | %w", "FirstOccurrence", err)
			}

		case "issue":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Issue", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Issue = o

		case "type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Type = o

		}
	}
	return nil
}

// NewTransformHealthIssue returns a TransformHealthIssue.
func NewTransformHealthIssue() *TransformHealthIssue {
	r := &TransformHealthIssue{}

	return r
}
