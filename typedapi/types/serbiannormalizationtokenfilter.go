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

// SerbianNormalizationTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/analysis/token_filters.ts#L539-L541
type SerbianNormalizationTokenFilter struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *SerbianNormalizationTokenFilter) UnmarshalJSON(data []byte) error {

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
func (s SerbianNormalizationTokenFilter) MarshalJSON() ([]byte, error) {
	type innerSerbianNormalizationTokenFilter SerbianNormalizationTokenFilter
	tmp := innerSerbianNormalizationTokenFilter{
		Type:    s.Type,
		Version: s.Version,
	}

	tmp.Type = "serbian_normalization"

	return json.Marshal(tmp)
}

// NewSerbianNormalizationTokenFilter returns a SerbianNormalizationTokenFilter.
func NewSerbianNormalizationTokenFilter() *SerbianNormalizationTokenFilter {
	r := &SerbianNormalizationTokenFilter{}

	return r
}

type SerbianNormalizationTokenFilterVariant interface {
	SerbianNormalizationTokenFilterCaster() *SerbianNormalizationTokenFilter
}

func (s *SerbianNormalizationTokenFilter) SerbianNormalizationTokenFilterCaster() *SerbianNormalizationTokenFilter {
	return s
}

func (s *SerbianNormalizationTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	o := TokenFilterDefinition(s)
	return &o
}
