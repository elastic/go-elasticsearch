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

// Hint type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/security/suggest_user_profiles/types.ts#L23-L34
type Hint struct {
	// Labels A single key-value pair to match against the labels section
	// of a profile. A profile is considered matching if it matches
	// at least one of the strings.
	Labels map[string][]string `json:"labels,omitempty"`
	// Uids A list of Profile UIDs to match against.
	Uids []string `json:"uids,omitempty"`
}

func (s *Hint) UnmarshalJSON(data []byte) error {

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

		case "labels":
			if s.Labels == nil {
				s.Labels = make(map[string][]string, 0)
			}
			rawMsg := make(map[string]json.RawMessage, 0)
			dec.Decode(&rawMsg)
			for key, value := range rawMsg {
				switch {
				case bytes.HasPrefix(value, []byte("\"")), bytes.HasPrefix(value, []byte("{")):
					o := new(string)
					err := json.NewDecoder(bytes.NewReader(value)).Decode(&o)
					if err != nil {
						return err
					}
					s.Labels[key] = append(s.Labels[key], *o)
				default:
					o := []string{}
					err := json.NewDecoder(bytes.NewReader(value)).Decode(&o)
					if err != nil {
						return err
					}
					s.Labels[key] = o
				}
			}

		case "uids":
			if err := dec.Decode(&s.Uids); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewHint returns a Hint.
func NewHint() *Hint {
	r := &Hint{
		Labels: make(map[string][]string, 0),
	}

	return r
}
