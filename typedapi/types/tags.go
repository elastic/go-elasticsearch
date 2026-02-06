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
// https://github.com/elastic/elasticsearch-specification/tree/d520d9e8cf14cad487de5e0654007686c395b494

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Tags type.
//
// https://github.com/elastic/elasticsearch-specification/blob/d520d9e8cf14cad487de5e0654007686c395b494/specification/project/tags/TagsResponse.ts#L23-L31
type Tags struct {
	Alias_        string            `json:"_alias"`
	Id_           string            `json:"_id"`
	Organisation_ string            `json:"_organisation"`
	Tags          map[string]string `json:"-"`
	Type_         string            `json:"_type"`
}

func (s *Tags) UnmarshalJSON(data []byte) error {

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

		case "_alias":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Alias_", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Alias_ = o

		case "_id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Id_", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Id_ = o

		case "_organisation":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Organisation_", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Organisation_ = o

		case "_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Type_", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Type_ = o

		default:

			if key, ok := t.(string); ok {
				if s.Tags == nil {
					s.Tags = make(map[string]string, 0)
				}
				raw := new(string)
				if err := dec.Decode(&raw); err != nil {
					return fmt.Errorf("%s | %w", "Tags", err)
				}
				if raw != nil {
					s.Tags[key] = *raw
				}
			}

		}
	}
	return nil
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s Tags) MarshalJSON() ([]byte, error) {
	type opt Tags
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]json.RawMessage, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.Tags {
		marshaled, err := json.Marshal(value)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal additional property %q: %w", key, err)
		}
		tmp[fmt.Sprintf("%s", key)] = marshaled
	}
	delete(tmp, "Tags")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewTags returns a Tags.
func NewTags() *Tags {
	r := &Tags{
		Tags: make(map[string]string),
	}

	return r
}
