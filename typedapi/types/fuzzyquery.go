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
// https://github.com/elastic/elasticsearch-specification/tree/26d0e2015b6bb2b1e0c549a4f1abeca6da16e89c

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// FuzzyQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/26d0e2015b6bb2b1e0c549a4f1abeca6da16e89c/specification/_types/query_dsl/term.ts#L40-L51
type FuzzyQuery struct {
	Boost          *float32  `json:"boost,omitempty"`
	Fuzziness      Fuzziness `json:"fuzziness,omitempty"`
	MaxExpansions  *int      `json:"max_expansions,omitempty"`
	PrefixLength   *int      `json:"prefix_length,omitempty"`
	QueryName_     *string   `json:"_name,omitempty"`
	Rewrite        *string   `json:"rewrite,omitempty"`
	Transpositions *bool     `json:"transpositions,omitempty"`
	Value          string    `json:"value"`
}

func (s *FuzzyQuery) UnmarshalJSON(data []byte) error {

	if !bytes.HasPrefix(data, []byte(`{`)) {
		err := json.NewDecoder(bytes.NewReader(data)).Decode(&s.Value)
		return err
	}

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

		case "boost":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return err
				}
				f := float32(value)
				s.Boost = &f
			case float64:
				f := float32(v)
				s.Boost = &f
			}

		case "fuzziness":
			if err := dec.Decode(&s.Fuzziness); err != nil {
				return err
			}

		case "max_expansions":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxExpansions = &value
			case float64:
				f := int(v)
				s.MaxExpansions = &f
			}

		case "prefix_length":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.PrefixLength = &value
			case float64:
				f := int(v)
				s.PrefixLength = &f
			}

		case "_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.QueryName_ = &o

		case "rewrite":
			if err := dec.Decode(&s.Rewrite); err != nil {
				return err
			}

		case "transpositions":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Transpositions = &value
			case bool:
				s.Transpositions = &v
			}

		case "value":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Value = o

		}
	}
	return nil
}

// NewFuzzyQuery returns a FuzzyQuery.
func NewFuzzyQuery() *FuzzyQuery {
	r := &FuzzyQuery{}

	return r
}
