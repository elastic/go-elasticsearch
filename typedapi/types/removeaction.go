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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// RemoveAction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/indices/update_aliases/types.ts#L97-L122
type RemoveAction struct {
	// Alias Alias for the action.
	// Index alias names support date math.
	Alias *string `json:"alias,omitempty"`
	// Aliases Aliases for the action.
	// Index alias names support date math.
	Aliases []string `json:"aliases,omitempty"`
	// Index Data stream or index for the action.
	// Supports wildcards (`*`).
	Index *string `json:"index,omitempty"`
	// Indices Data streams or indices for the action.
	// Supports wildcards (`*`).
	Indices []string `json:"indices,omitempty"`
	// MustExist If `true`, the alias must exist to perform the action.
	MustExist *bool `json:"must_exist,omitempty"`
}

func (s *RemoveAction) UnmarshalJSON(data []byte) error {

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

		case "alias":
			if err := dec.Decode(&s.Alias); err != nil {
				return err
			}

		case "aliases":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.Aliases = append(s.Aliases, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Aliases); err != nil {
					return err
				}
			}

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return err
			}

		case "indices":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.Indices = append(s.Indices, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Indices); err != nil {
					return err
				}
			}

		case "must_exist":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.MustExist = &value
			case bool:
				s.MustExist = &v
			}

		}
	}
	return nil
}

// NewRemoveAction returns a RemoveAction.
func NewRemoveAction() *RemoveAction {
	r := &RemoveAction{}

	return r
}
