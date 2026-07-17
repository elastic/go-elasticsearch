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
// https://github.com/elastic/elasticsearch-specification/tree/8076b1c4ff3b8bd4eb5372bc75372577a21d1b0c

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// A named ES|QL query parameter supplied in its classified form. Exactly one of
// `value`, `identifier`, or `pattern` must be set.
//
// https://github.com/elastic/elasticsearch-specification/blob/8076b1c4ff3b8bd4eb5372bc75372577a21d1b0c/specification/esql/_types/types.ts#L41-L59
type ClassifiedNamedParameter struct {
	// Identifier Interpret the parameter as an identifier, such as a field or function name.
	Identifier *string `json:"identifier,omitempty"`
	// Pattern Interpret the parameter as a pattern, such as an index or field name pattern.
	Pattern *string `json:"pattern,omitempty"`
	// Value Interpret the parameter as a literal value.
	Value []FieldValue `json:"value,omitempty"`
}

func (s *ClassifiedNamedParameter) UnmarshalJSON(data []byte) error {

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

		case "identifier":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Identifier", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Identifier = &o

		case "pattern":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Pattern", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Pattern = &o

		case "value":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(FieldValue)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Value", err)
				}

				s.Value = append(s.Value, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Value); err != nil {
					return fmt.Errorf("%s | %w", "Value", err)
				}
			}

		}
	}
	return nil
}

// NewClassifiedNamedParameter returns a ClassifiedNamedParameter.
func NewClassifiedNamedParameter() *ClassifiedNamedParameter {
	r := &ClassifiedNamedParameter{}

	return r
}

type ClassifiedNamedParameterVariant interface {
	ClassifiedNamedParameterCaster() *ClassifiedNamedParameter
}

func (s *ClassifiedNamedParameter) ClassifiedNamedParameterCaster() *ClassifiedNamedParameter {
	return s
}

func (s *ClassifiedNamedParameter) NamedParameterValueCaster() *NamedParameterValue {
	if s == nil {
		return nil
	}
	o := NamedParameterValue(s)
	return &o
}
