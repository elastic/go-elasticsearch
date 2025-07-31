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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/runtimefieldtype"
)

// RuntimeField type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/mapping/RuntimeFields.ts#L26-L50
type RuntimeField struct {
	// FetchFields For type `lookup`
	FetchFields []RuntimeFieldFetchFields `json:"fetch_fields,omitempty"`
	// Fields For type `composite`
	Fields map[string]CompositeSubField `json:"fields,omitempty"`
	// Format A custom format for `date` type runtime fields.
	Format *string `json:"format,omitempty"`
	// InputField For type `lookup`
	InputField *string `json:"input_field,omitempty"`
	// Script Painless script executed at query time.
	Script *Script `json:"script,omitempty"`
	// TargetField For type `lookup`
	TargetField *string `json:"target_field,omitempty"`
	// TargetIndex For type `lookup`
	TargetIndex *string `json:"target_index,omitempty"`
	// Type Field type, which can be: `boolean`, `composite`, `date`, `double`,
	// `geo_point`, `ip`,`keyword`, `long`, or `lookup`.
	Type runtimefieldtype.RuntimeFieldType `json:"type"`
}

func (s *RuntimeField) UnmarshalJSON(data []byte) error {

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

		case "fetch_fields":
			if err := dec.Decode(&s.FetchFields); err != nil {
				return fmt.Errorf("%s | %w", "FetchFields", err)
			}

		case "fields":
			if s.Fields == nil {
				s.Fields = make(map[string]CompositeSubField, 0)
			}
			if err := dec.Decode(&s.Fields); err != nil {
				return fmt.Errorf("%s | %w", "Fields", err)
			}

		case "format":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Format", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Format = &o

		case "input_field":
			if err := dec.Decode(&s.InputField); err != nil {
				return fmt.Errorf("%s | %w", "InputField", err)
			}

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return fmt.Errorf("%s | %w", "Script", err)
			}

		case "target_field":
			if err := dec.Decode(&s.TargetField); err != nil {
				return fmt.Errorf("%s | %w", "TargetField", err)
			}

		case "target_index":
			if err := dec.Decode(&s.TargetIndex); err != nil {
				return fmt.Errorf("%s | %w", "TargetIndex", err)
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		}
	}
	return nil
}

// NewRuntimeField returns a RuntimeField.
func NewRuntimeField() *RuntimeField {
	r := &RuntimeField{
		Fields: make(map[string]CompositeSubField),
	}

	return r
}
