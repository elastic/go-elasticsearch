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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/noridecompoundmode"

	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// NoriAnalyzer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/_types/analysis/analyzers.ts#L66-L72
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
				return err
			}

		case "stoptags":
			if err := dec.Decode(&s.Stoptags); err != nil {
				return err
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return err
			}

		case "user_dictionary":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.UserDictionary = &o

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewNoriAnalyzer returns a NoriAnalyzer.
func NewNoriAnalyzer() *NoriAnalyzer {
	r := &NoriAnalyzer{}

	r.Type = "nori"

	return r
}
