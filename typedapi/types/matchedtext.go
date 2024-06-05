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
// https://github.com/elastic/elasticsearch-specification/tree/07bf82537a186562d8699685e3704ea338b268ef

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// MatchedText type.
//
// https://github.com/elastic/elasticsearch-specification/blob/07bf82537a186562d8699685e3704ea338b268ef/specification/text_structure/test_grok_pattern/types.ts#L29-L32
type MatchedText struct {
	Fields  map[string][]MatchedField `json:"fields,omitempty"`
	Matched bool                      `json:"matched"`
}

func (s *MatchedText) UnmarshalJSON(data []byte) error {

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

		case "fields":
			if s.Fields == nil {
				s.Fields = make(map[string][]MatchedField, 0)
			}
			if err := dec.Decode(&s.Fields); err != nil {
				return fmt.Errorf("%s | %w", "Fields", err)
			}

		case "matched":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Matched", err)
				}
				s.Matched = value
			case bool:
				s.Matched = v
			}

		}
	}
	return nil
}

// NewMatchedText returns a MatchedText.
func NewMatchedText() *MatchedText {
	r := &MatchedText{
		Fields: make(map[string][]MatchedField, 0),
	}

	return r
}
