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
)

// FieldLookup type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/query_dsl/abstractions.ts#L436-L453
type FieldLookup struct {
	// Id `id` of the document.
	Id string `json:"id"`
	// Index Index from which to retrieve the document.
	Index *string `json:"index,omitempty"`
	// Path Name of the field.
	Path *string `json:"path,omitempty"`
	// Routing Custom routing value.
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
				return fmt.Errorf("%s | %w", "Id", err)
			}

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return fmt.Errorf("%s | %w", "Index", err)
			}

		case "path":
			if err := dec.Decode(&s.Path); err != nil {
				return fmt.Errorf("%s | %w", "Path", err)
			}

		case "routing":
			if err := dec.Decode(&s.Routing); err != nil {
				return fmt.Errorf("%s | %w", "Routing", err)
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

type FieldLookupVariant interface {
	FieldLookupCaster() *FieldLookup
}

func (s *FieldLookup) FieldLookupCaster() *FieldLookup {
	return s
}
