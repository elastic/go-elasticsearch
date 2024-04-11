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
// https://github.com/elastic/elasticsearch-specification/tree/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// SettingsSimilarityScripted type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1/specification/indices/_types/IndexSettings.ts#L222-L226
type SettingsSimilarityScripted struct {
	Script       Script `json:"script"`
	Type         string `json:"type,omitempty"`
	WeightScript Script `json:"weight_script,omitempty"`
}

func (s *SettingsSimilarityScripted) UnmarshalJSON(data []byte) error {

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

		case "script":
			message := json.RawMessage{}
			if err := dec.Decode(&message); err != nil {
				return fmt.Errorf("%s | %w", "Script", err)
			}
			keyDec := json.NewDecoder(bytes.NewReader(message))
			for {
				t, err := keyDec.Token()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return fmt.Errorf("%s | %w", "Script", err)
				}

				switch t {

				case "lang", "options", "source":
					o := NewInlineScript()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "Script", err)
					}
					s.Script = o

				case "id":
					o := NewStoredScriptId()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "Script", err)
					}
					s.Script = o

				}
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		case "weight_script":
			message := json.RawMessage{}
			if err := dec.Decode(&message); err != nil {
				return fmt.Errorf("%s | %w", "WeightScript", err)
			}
			keyDec := json.NewDecoder(bytes.NewReader(message))
			for {
				t, err := keyDec.Token()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return fmt.Errorf("%s | %w", "WeightScript", err)
				}

				switch t {

				case "lang", "options", "source":
					o := NewInlineScript()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "WeightScript", err)
					}
					s.WeightScript = o

				case "id":
					o := NewStoredScriptId()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "WeightScript", err)
					}
					s.WeightScript = o

				}
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s SettingsSimilarityScripted) MarshalJSON() ([]byte, error) {
	type innerSettingsSimilarityScripted SettingsSimilarityScripted
	tmp := innerSettingsSimilarityScripted{
		Script:       s.Script,
		Type:         s.Type,
		WeightScript: s.WeightScript,
	}

	tmp.Type = "scripted"

	return json.Marshal(tmp)
}

// NewSettingsSimilarityScripted returns a SettingsSimilarityScripted.
func NewSettingsSimilarityScripted() *SettingsSimilarityScripted {
	r := &SettingsSimilarityScripted{}

	return r
}
