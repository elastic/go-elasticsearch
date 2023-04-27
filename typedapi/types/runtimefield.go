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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/runtimefieldtype"

	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// RuntimeField type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/_types/mapping/RuntimeFields.ts#L26-L38
type RuntimeField struct {
	// FetchFields For type `lookup`
	FetchFields []RuntimeFieldFetchFields `json:"fetch_fields,omitempty"`
	Format      *string                   `json:"format,omitempty"`
	// InputField For type `lookup`
	InputField *string `json:"input_field,omitempty"`
	Script     Script  `json:"script,omitempty"`
	// TargetField For type `lookup`
	TargetField *string `json:"target_field,omitempty"`
	// TargetIndex For type `lookup`
	TargetIndex *string                           `json:"target_index,omitempty"`
	Type        runtimefieldtype.RuntimeFieldType `json:"type"`
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
				return err
			}

		case "format":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.Format = &o

		case "input_field":
			if err := dec.Decode(&s.InputField); err != nil {
				return err
			}

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return err
			}

		case "target_field":
			if err := dec.Decode(&s.TargetField); err != nil {
				return err
			}

		case "target_index":
			if err := dec.Decode(&s.TargetIndex); err != nil {
				return err
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewRuntimeField returns a RuntimeField.
func NewRuntimeField() *RuntimeField {
	r := &RuntimeField{}

	return r
}
