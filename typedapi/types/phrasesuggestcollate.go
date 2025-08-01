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

// PhraseSuggestCollate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_global/search/_types/suggester.ts#L334-L347
type PhraseSuggestCollate struct {
	// Params Parameters to use if the query is templated.
	Params map[string]json.RawMessage `json:"params,omitempty"`
	// Prune Returns all suggestions with an extra `collate_match` option indicating
	// whether the generated phrase matched any document.
	Prune *bool `json:"prune,omitempty"`
	// Query A collate query that is run once for every suggestion.
	Query PhraseSuggestCollateQuery `json:"query"`
}

func (s *PhraseSuggestCollate) UnmarshalJSON(data []byte) error {

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

		case "params":
			if s.Params == nil {
				s.Params = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Params); err != nil {
				return fmt.Errorf("%s | %w", "Params", err)
			}

		case "prune":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Prune", err)
				}
				s.Prune = &value
			case bool:
				s.Prune = &v
			}

		case "query":
			if err := dec.Decode(&s.Query); err != nil {
				return fmt.Errorf("%s | %w", "Query", err)
			}

		}
	}
	return nil
}

// NewPhraseSuggestCollate returns a PhraseSuggestCollate.
func NewPhraseSuggestCollate() *PhraseSuggestCollate {
	r := &PhraseSuggestCollate{
		Params: make(map[string]json.RawMessage),
	}

	return r
}

type PhraseSuggestCollateVariant interface {
	PhraseSuggestCollateCaster() *PhraseSuggestCollate
}

func (s *PhraseSuggestCollate) PhraseSuggestCollateCaster() *PhraseSuggestCollate {
	return s
}
