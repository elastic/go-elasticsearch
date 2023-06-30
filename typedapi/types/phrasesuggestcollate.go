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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// PhraseSuggestCollate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_global/search/_types/suggester.ts#L184-L188
type PhraseSuggestCollate struct {
	Params map[string]json.RawMessage `json:"params,omitempty"`
	Prune  *bool                      `json:"prune,omitempty"`
	Query  PhraseSuggestCollateQuery  `json:"query"`
}

func (s *PhraseSuggestCollate) UnmarshalJSON(data []byte) error {

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

		case "params":
			if s.Params == nil {
				s.Params = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Params); err != nil {
				return err
			}

		case "prune":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Prune = &value
			case bool:
				s.Prune = &v
			}

		case "query":
			if err := dec.Decode(&s.Query); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewPhraseSuggestCollate returns a PhraseSuggestCollate.
func NewPhraseSuggestCollate() *PhraseSuggestCollate {
	r := &PhraseSuggestCollate{
		Params: make(map[string]json.RawMessage, 0),
	}

	return r
}
