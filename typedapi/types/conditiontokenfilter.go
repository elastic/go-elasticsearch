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
// https://github.com/elastic/elasticsearch-specification/tree/b7d4fb5356784b8bcde8d3a2d62a1fd5621ffd67

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// ConditionTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/b7d4fb5356784b8bcde8d3a2d62a1fd5621ffd67/specification/_types/analysis/token_filters.ts#L181-L185
type ConditionTokenFilter struct {
	Filter  []string `json:"filter"`
	Script  Script   `json:"script"`
	Type    string   `json:"type,omitempty"`
	Version *string  `json:"version,omitempty"`
}

func (s *ConditionTokenFilter) UnmarshalJSON(data []byte) error {

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

		case "filter":
			if err := dec.Decode(&s.Filter); err != nil {
				return err
			}

		case "script":
			message := json.RawMessage{}
			if err := dec.Decode(&message); err != nil {
				return err
			}
			keyDec := json.NewDecoder(bytes.NewReader(message))
			for {
				t, err := keyDec.Token()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return err
				}

				switch t {

				case "lang", "options", "source":
					o := NewInlineScript()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return err
					}
					s.Script = o

				case "id":
					o := NewStoredScriptId()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return err
					}
					s.Script = o

				}
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return err
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return err
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s ConditionTokenFilter) MarshalJSON() ([]byte, error) {
	type innerConditionTokenFilter ConditionTokenFilter
	tmp := innerConditionTokenFilter{
		Filter:  s.Filter,
		Script:  s.Script,
		Type:    s.Type,
		Version: s.Version,
	}

	tmp.Type = "condition"

	return json.Marshal(tmp)
}

// NewConditionTokenFilter returns a ConditionTokenFilter.
func NewConditionTokenFilter() *ConditionTokenFilter {
	r := &ConditionTokenFilter{}

	return r
}
