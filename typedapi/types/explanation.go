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
// https://github.com/elastic/elasticsearch-specification/tree/5260ec5b7c899ab1a7939f752218cae07ef07dd7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// Explanation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5260ec5b7c899ab1a7939f752218cae07ef07dd7/specification/_global/explain/types.ts#L22-L26
type Explanation struct {
	Description string              `json:"description"`
	Details     []ExplanationDetail `json:"details"`
	Value       float32             `json:"value"`
}

func (s *Explanation) UnmarshalJSON(data []byte) error {

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

		case "description":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Description = o

		case "details":
			if err := dec.Decode(&s.Details); err != nil {
				return err
			}

		case "value":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return err
				}
				f := float32(value)
				s.Value = f
			case float64:
				f := float32(v)
				s.Value = f
			}

		}
	}
	return nil
}

// NewExplanation returns a Explanation.
func NewExplanation() *Explanation {
	r := &Explanation{}

	return r
}
