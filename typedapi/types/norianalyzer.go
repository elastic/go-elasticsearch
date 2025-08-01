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

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/noridecompoundmode"
)

// NoriAnalyzer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/analysis/analyzers.ts#L323-L330
type NoriAnalyzer struct {
	DecompoundMode *noridecompoundmode.NoriDecompoundMode `json:"decompound_mode,omitempty"`
	Stoptags       []string                               `json:"stoptags,omitempty"`
	Type           string                                 `json:"type,omitempty"`
	UserDictionary *string                                `json:"user_dictionary,omitempty"`
	Version        *string                                `json:"version,omitempty"`
}

func (s *NoriAnalyzer) UnmarshalJSON(data []byte) error {

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

		case "decompound_mode":
			if err := dec.Decode(&s.DecompoundMode); err != nil {
				return fmt.Errorf("%s | %w", "DecompoundMode", err)
			}

		case "stoptags":
			if err := dec.Decode(&s.Stoptags); err != nil {
				return fmt.Errorf("%s | %w", "Stoptags", err)
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		case "user_dictionary":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "UserDictionary", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.UserDictionary = &o

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s NoriAnalyzer) MarshalJSON() ([]byte, error) {
	type innerNoriAnalyzer NoriAnalyzer
	tmp := innerNoriAnalyzer{
		DecompoundMode: s.DecompoundMode,
		Stoptags:       s.Stoptags,
		Type:           s.Type,
		UserDictionary: s.UserDictionary,
		Version:        s.Version,
	}

	tmp.Type = "nori"

	return json.Marshal(tmp)
}

// NewNoriAnalyzer returns a NoriAnalyzer.
func NewNoriAnalyzer() *NoriAnalyzer {
	r := &NoriAnalyzer{}

	return r
}

type NoriAnalyzerVariant interface {
	NoriAnalyzerCaster() *NoriAnalyzer
}

func (s *NoriAnalyzer) NoriAnalyzerCaster() *NoriAnalyzer {
	return s
}

func (s *NoriAnalyzer) AnalyzerCaster() *Analyzer {
	o := Analyzer(s)
	return &o
}
