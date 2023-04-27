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

	"strconv"

	"encoding/json"
)

// CommonGramsTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/_types/analysis/token_filters.ts#L173-L179
type CommonGramsTokenFilter struct {
	CommonWords     []string `json:"common_words,omitempty"`
	CommonWordsPath *string  `json:"common_words_path,omitempty"`
	IgnoreCase      *bool    `json:"ignore_case,omitempty"`
	QueryMode       *bool    `json:"query_mode,omitempty"`
	Type            string   `json:"type,omitempty"`
	Version         *string  `json:"version,omitempty"`
}

func (s *CommonGramsTokenFilter) UnmarshalJSON(data []byte) error {

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

		case "common_words":
			if err := dec.Decode(&s.CommonWords); err != nil {
				return err
			}

		case "common_words_path":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.CommonWordsPath = &o

		case "ignore_case":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.IgnoreCase = &value
			case bool:
				s.IgnoreCase = &v
			}

		case "query_mode":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.QueryMode = &value
			case bool:
				s.QueryMode = &v
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

// NewCommonGramsTokenFilter returns a CommonGramsTokenFilter.
func NewCommonGramsTokenFilter() *CommonGramsTokenFilter {
	r := &CommonGramsTokenFilter{}

	r.Type = "common_grams"

	return r
}
