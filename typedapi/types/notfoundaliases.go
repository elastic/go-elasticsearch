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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// NotFoundAliases type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/indices/get_alias/_types/response.ts#L28-L36
type NotFoundAliases struct {
	Error           string                  `json:"error"`
	NotFoundAliases map[string]IndexAliases `json:"-"`
	Status          int                     `json:"status"`
}

func (s *NotFoundAliases) UnmarshalJSON(data []byte) error {

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

		case "error":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Error", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Error = o

		case "status":
			if err := dec.Decode(&s.Status); err != nil {
				return fmt.Errorf("%s | %w", "Status", err)
			}

		default:

			if key, ok := t.(string); ok {
				if s.NotFoundAliases == nil {
					s.NotFoundAliases = make(map[string]IndexAliases, 0)
				}
				raw := NewIndexAliases()
				if err := dec.Decode(&raw); err != nil {
					return fmt.Errorf("%s | %w", "NotFoundAliases", err)
				}
				s.NotFoundAliases[key] = *raw
			}

		}
	}
	return nil
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s NotFoundAliases) MarshalJSON() ([]byte, error) {
	type opt NotFoundAliases
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]any, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.NotFoundAliases {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "NotFoundAliases")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewNotFoundAliases returns a NotFoundAliases.
func NewNotFoundAliases() *NotFoundAliases {
	r := &NotFoundAliases{
		NotFoundAliases: make(map[string]IndexAliases),
	}

	return r
}
