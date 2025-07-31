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

// FrenchStemTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/analysis/token_filters.ts#L555-L557
type FrenchStemTokenFilter struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *FrenchStemTokenFilter) UnmarshalJSON(data []byte) error {

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

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s FrenchStemTokenFilter) MarshalJSON() ([]byte, error) {
	type innerFrenchStemTokenFilter FrenchStemTokenFilter
	tmp := innerFrenchStemTokenFilter{
		Type:    s.Type,
		Version: s.Version,
	}

	tmp.Type = "french_stem"

	return json.Marshal(tmp)
}

// NewFrenchStemTokenFilter returns a FrenchStemTokenFilter.
func NewFrenchStemTokenFilter() *FrenchStemTokenFilter {
	r := &FrenchStemTokenFilter{}

	return r
}

type FrenchStemTokenFilterVariant interface {
	FrenchStemTokenFilterCaster() *FrenchStemTokenFilter
}

func (s *FrenchStemTokenFilter) FrenchStemTokenFilterCaster() *FrenchStemTokenFilter {
	return s
}

func (s *FrenchStemTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	o := TokenFilterDefinition(s)
	return &o
}
