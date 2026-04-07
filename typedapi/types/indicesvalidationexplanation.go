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
// https://github.com/elastic/elasticsearch-specification/tree/6ee016a765be615b0205fc209d3d3c515044689d

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// IndicesValidationExplanation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/6ee016a765be615b0205fc209d3d3c515044689d/specification/indices/validate_query/IndicesValidateQueryResponse.ts#L33-L39
type IndicesValidationExplanation struct {
	Error       *string `json:"error,omitempty"`
	Explanation *string `json:"explanation,omitempty"`
	Index       *string `json:"index,omitempty"`
	Shard       *int    `json:"shard,omitempty"`
	Valid       bool    `json:"valid"`
}

func (s *IndicesValidationExplanation) UnmarshalJSON(data []byte) error {

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

		case "error":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Error", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Error = &o

		case "explanation":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Explanation", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Explanation = &o

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return fmt.Errorf("%s | %w", "Index", err)
			}

		case "shard":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Shard", err)
				}
				s.Shard = &value
			case float64:
				f := int(v)
				s.Shard = &f
			}

		case "valid":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Valid", err)
				}
				s.Valid = value
			case bool:
				s.Valid = v
			}

		}
	}
	return nil
}

// NewIndicesValidationExplanation returns a IndicesValidationExplanation.
func NewIndicesValidationExplanation() *IndicesValidationExplanation {
	r := &IndicesValidationExplanation{}

	return r
}
