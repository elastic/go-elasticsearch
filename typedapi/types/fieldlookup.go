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
	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// FieldLookup type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/_types/query_dsl/abstractions.ts#L164-L169
type FieldLookup struct {
	Id      string  `json:"id"`
	Index   *string `json:"index,omitempty"`
	Path    *string `json:"path,omitempty"`
	Routing *string `json:"routing,omitempty"`
}

func (s *FieldLookup) UnmarshalJSON(data []byte) error {

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

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return err
			}

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return err
			}

		case "path":
			if err := dec.Decode(&s.Path); err != nil {
				return err
			}

		case "routing":
			if err := dec.Decode(&s.Routing); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewFieldLookup returns a FieldLookup.
func NewFieldLookup() *FieldLookup {
	r := &FieldLookup{}

	return r
}
