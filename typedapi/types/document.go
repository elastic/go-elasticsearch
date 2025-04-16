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
// https://github.com/elastic/elasticsearch-specification/tree/52c473efb1fb5320a5bac12572d0b285882862fb

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// Document type.
//
// https://github.com/elastic/elasticsearch-specification/blob/52c473efb1fb5320a5bac12572d0b285882862fb/specification/ingest/_types/Simulation.ts#L62-L76
type Document struct {
	// Id_ Unique identifier for the document.
	// This ID must be unique within the `_index`.
	Id_ *string `json:"_id,omitempty"`
	// Index_ Name of the index containing the document.
	Index_ *string `json:"_index,omitempty"`
	// Source_ JSON body for the document.
	Source_ json.RawMessage `json:"_source,omitempty"`
}

func (s *Document) UnmarshalJSON(data []byte) error {

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

		case "_id":
			if err := dec.Decode(&s.Id_); err != nil {
				return fmt.Errorf("%s | %w", "Id_", err)
			}

		case "_index":
			if err := dec.Decode(&s.Index_); err != nil {
				return fmt.Errorf("%s | %w", "Index_", err)
			}

		case "_source":
			if err := dec.Decode(&s.Source_); err != nil {
				return fmt.Errorf("%s | %w", "Source_", err)
			}

		}
	}
	return nil
}

// NewDocument returns a Document.
func NewDocument() *Document {
	r := &Document{}

	return r
}

// true

type DocumentVariant interface {
	DocumentCaster() *Document
}

func (s *Document) DocumentCaster() *Document {
	return s
}
