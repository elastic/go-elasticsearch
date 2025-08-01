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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// ListTypeValidation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/connector/_types/Connector.ts#L68-L71
type ListTypeValidation struct {
	Constraint string `json:"constraint"`
	Type       string `json:"type,omitempty"`
}

func (s *ListTypeValidation) UnmarshalJSON(data []byte) error {

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

		case "constraint":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Constraint", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Constraint = o

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s ListTypeValidation) MarshalJSON() ([]byte, error) {
	type innerListTypeValidation ListTypeValidation
	tmp := innerListTypeValidation{
		Constraint: s.Constraint,
		Type:       s.Type,
	}

	tmp.Type = "list_type"

	return json.Marshal(tmp)
}

// NewListTypeValidation returns a ListTypeValidation.
func NewListTypeValidation() *ListTypeValidation {
	r := &ListTypeValidation{}

	return r
}

type ListTypeValidationVariant interface {
	ListTypeValidationCaster() *ListTypeValidation
}

func (s *ListTypeValidation) ListTypeValidationCaster() *ListTypeValidation {
	return s
}

func (s *ListTypeValidation) ValidationCaster() *Validation {
	o := Validation(s)
	return &o
}
